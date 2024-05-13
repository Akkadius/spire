package generators

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/iancoleman/strcase"
	"log"
	"os"
	"strings"
	"text/template"
)

type SyncControllersToInjector struct {
}

func NewSyncControllersToInjector() *SyncControllersToInjector {
	return &SyncControllersToInjector{}
}

const crudInjectFile = "./boot/inject_http_crud_controllers.go"

type injectHttpCrudControllerTmpl struct {
	NewControllers      string
	ControllersParam    string
	ControllersRegister string
}

// loop through existing controllers
// inject crud controllers into ./boot/inject_http_crud_controllers.go
func (s *SyncControllersToInjector) Sync() {
	files, err := os.ReadDir(crudControllerPath)
	if err != nil {
		log.Fatal(err)
	}

	var controllerNames []string

	// fetch struct names
	for _, f := range files {
		//fmt.Println(f.Name())
		file, err := os.Open(fmt.Sprintf("%v%v", crudControllerPath, f.Name()))
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			l := scanner.Text()
			if strings.Contains(l, "type") && strings.Contains(l, "struct") && strings.Contains(l, "Controller") {
				l = strings.Replace(l, "type ", "", -1)
				l = strings.Replace(l, "struct {", "", -1)
				l = strings.TrimSpace(l)

				controllerNames = append(controllerNames, l)
			}
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}

	newControllersString := ""
	controllersParamString := ""
	controllersRegisterString := ""
	for _, name := range controllerNames {
		fmt.Println(name)
		newControllersString += fmt.Sprintf("\tcrudcontrollers.New%v,\n", name)
		controllersParamString += fmt.Sprintf("\t%v *crudcontrollers.%v,\n", strcase.ToLowerCamel(name), name)
		controllersRegisterString += fmt.Sprintf("\t\t\t%v,\n", strcase.ToLowerCamel(name))
	}

	tpl, err := template.ParseFiles("./internal/generators/templates/inject_http_crud_controller.tmpl")
	t := injectHttpCrudControllerTmpl{
		NewControllers:      strings.TrimSuffix(newControllersString, "\n"),
		ControllersParam:    strings.TrimSuffix(controllersParamString, "\n"),
		ControllersRegister: strings.TrimSuffix(controllersRegisterString, "\n"),
	}

	var out bytes.Buffer
	err = tpl.ExecuteTemplate(&out, "inject_http_crud_controller.tmpl", t)
	if err != nil {
		fmt.Println(err)
	}

	// write file
	fileName := fmt.Sprintf("%v", crudInjectFile)

	// create file
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}

	// write contents
	_, err = file.WriteString(out.String())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fmt.Sprintf("Synced controllers to [%v]\n", crudInjectFile))

	defer file.Close()

}
