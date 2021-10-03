package generators

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gertd/go-pluralize"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/iancoleman/strcase"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
)

type GenerateVueFormContext struct {
	TablesToGenerate      []string
	UseDatabaseSchemaDocs bool
}

type GenerateVueForm struct {
	options GenerateVueFormContext
	logger  *logrus.Logger
	pluralize *pluralize.Client
}

func NewGenerateVueForm(options GenerateVueFormContext, logger *logrus.Logger) *GenerateVueForm {
	return &GenerateVueForm{
		options:   options,
		logger:    logger,
		pluralize: pluralize.NewClient(),
	}
}

type VueFormField struct {
	Disabled     bool   `json:"disabled,omitempty"`
	Featured     bool   `json:"featured"`
	InputType    string `json:"inputType"`
	Label        string `json:"label"`
	Model        string `json:"model"`
	Readonly     bool   `json:"readonly,omitempty"`
	Rows         string `json:"rows,omitempty"` // used in textArea
	StyleClasses string `json:"styleClasses,omitempty"`
	Type         string `json:"type,omitempty"`
	Help         string `json:"help,omitempty"`
}

func (g *GenerateVueForm) Generate() error {
	if len(g.options.TablesToGenerate) == 0 {
		g.options.TablesToGenerate = GetDatabaseTables()
	}

	for _, table := range g.options.TablesToGenerate {

		// template
		tpl, err := template.ParseFiles("./internal/generators/templates/vue_form.tmpl")

		type templateData struct {
			FormFields            string
			EntityNameSnake       string
			EntityNameSnakePlural string
		}

		// Create form fields from database table columns
		vueFormFields := GetVueFormConfigTable(table)

		// Accepts an empty interface, prefix and indent
		formFieldsJson, err := json.MarshalIndent(vueFormFields, "          ", "  ")
		if err != nil {
			return err
		}

		var out bytes.Buffer
		err = tpl.ExecuteTemplate(
			&out, "vue_form.tmpl", templateData{
				FormFields:            string(formFieldsJson),
				EntityNameSnake:       strcase.ToSnake(table),
				EntityNameSnakePlural: g.pluralize.Plural(strcase.ToSnake(table)),
			},
		)
		if err != nil {
			fmt.Println(err)
		}

		// write file
		fileName := fmt.Sprintf("./frontend/src/components/forms/%vForm.vue", strcase.ToCamel(table))

		g.logger.Info(fmt.Sprintf("Wrote [%v]", fileName))

		// create file
		file, err := os.Create(fileName)
		if err != nil {
			return err
		}

		defer file.Close()

		// write contents
		_, err = file.WriteString(out.String())
		if err != nil {
			return err
		}
	}

	return nil
}

const (
	// database schema doc .yaml reference http url
	dbSchemaRawUrl = "https://raw.githubusercontent.com/eqemu/docs-db-schema/master/database-schema-reference.yml"
	// database schema doc .yaml local reference
	dbSchemaLocalCachePath = "./internal/generators/config/database-schema-reference.yml"
)

var dbSchemaYaml map[interface{}]interface{}

func (g *GenerateVueForm) getSchemaYaml() (map[interface{}]interface{}, error) {
	if len(dbSchemaYaml) > 0 {
		return dbSchemaYaml, nil
	}

	nullMap := make(map[interface{}]interface{})

	if _, err := os.Stat(dbSchemaLocalCachePath); os.IsNotExist(err) {

		fmt.Printf("Fetching manifest again\n")

		// get
		resp, err := http.Get(dbSchemaRawUrl)
		if err != nil {
			return nullMap, err
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nullMap, err
		}

		// open file
		file, err := os.Create(dbSchemaLocalCachePath)
		if err != nil {
			return nullMap, nil
		}
		defer file.Close()

		// write
		_, err = file.WriteString(string(body))
		if err != nil {
			return nullMap, nil
		}
	}

	schemaYaml, err := ioutil.ReadFile(dbSchemaLocalCachePath)
	if err != nil {
		return nullMap, err
	}

	// load yaml
	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal(schemaYaml, &m)
	if err != nil {
		return nullMap, err
	}

	dbSchemaYaml = m

	return m, nil
}

// pulls down database schema .yaml from docs and pulls the description
// keyed by table and column name
// we keep a local copy to keep from pulling it down repeatedly
func (g *GenerateVueForm) getSchemaDocDescriptionForColumn(table string, column string) (string, error) {
	m, err := g.getSchemaYaml()
	if err != nil {
		return "", err
	}

	// table -> column
	tableSchemaRef, ok := m[table].(map[string]interface{})
	if !ok {
		return "", errors.New("Invalid index reference for table -> column")
	}

	// column -> description
	columnData, ok := tableSchemaRef[column].(map[string]interface{})
	if !ok {
		return "", errors.New("Invalid index reference for table -> column")
	}

	space := regexp.MustCompile(`\s+`)
	columnDescription := space.ReplaceAllString(columnData["description"].(string), " ")

	return columnDescription, nil
}

func (g *GenerateVueForm) GenerateConfig() error {
	for _, table := range GetDatabaseTables() {

		vueFormConfig := []VueFormField{}

		htmlFlags := html.CommonFlags | html.HrefTargetBlank | html.SkipHTML
		opts := html.RendererOptions{Flags: htmlFlags}
		renderer := html.NewRenderer(opts)

		// Create form fields from database table columns
		tableColumns := GetDbSchemaConfigTable(table)
		for _, column := range tableColumns {
			columnDescription := ""
			if g.options.UseDatabaseSchemaDocs {
				desc, err := g.getSchemaDocDescriptionForColumn(table, column.Column)
				if err != nil {
					columnDescription = ""
					g.logger.Warn(
						fmt.Sprintf(
							"Could not find schema description for [%v] [%v]\n",
							table,
							column.Column,
						),
					)
				}
				columnDescription = string(
					markdown.ToHTML(
						[]byte(desc),
						nil,
						renderer,
					),
				)

				columnDescription = strings.TrimRight(columnDescription, "\n")
			}
			vueFormConfig = append(
				vueFormConfig, VueFormField{
					Disabled:     false,
					Featured:     false,
					InputType:    "text",
					Label:        strings.Title(strings.Replace(column.Column, "_", " ", -1)),
					Model:        column.Column,
					Readonly:     false,
					StyleClasses: "col-12",
					Type:         "input",
					Help:         columnDescription,
				},
			)
		}

		// yaml
		var b bytes.Buffer
		config := yaml.NewEncoder(&b)
		config.SetIndent(2)
		if err := config.Encode(&vueFormConfig); err != nil {
			return err
		}

		vueFieldsConfigFile := fmt.Sprintf("./internal/generators/config/vue-forms/%v.yml", table)

		dir, err := filepath.Abs(filepath.Dir(vueFieldsConfigFile))
		if err != nil {
			g.logger.Fatal(err)
		}

		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			g.logger.Fatal(err)
		}

		// create config
		file, err := os.Create(vueFieldsConfigFile)
		if err != nil {
			g.logger.Fatal(err)
		}

		defer file.Close()

		// write config
		_, err = file.Write(b.Bytes())
		if err != nil {
			g.logger.Fatal(err)
		}

		g.logger.Infof("Wrote configuration [%v]", vueFieldsConfigFile)
	}

	return nil
}

// get vue form config table
func GetVueFormConfigTable(form string) []VueFormField {
	var m []VueFormField

	vueFieldsConfigFile := fmt.Sprintf("./internal/generators/config/vue-forms/%v.yml", form)

	config, err := ioutil.ReadFile(vueFieldsConfigFile)
	if err != nil {
		log.Fatal(err)
	}

	// load yaml
	err = yaml.Unmarshal(config, &m)
	if err != nil {
		log.Fatal(err)
	}

	return m
}
