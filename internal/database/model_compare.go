package database

import (
	"fmt"
	"github.com/Akkadius/spire/internal/structs"
	"reflect"
	"strings"
)

type ModelDifference struct {
	Field string
	Old   interface{}
	New   interface{}
}

// ResultDifference will return a map[string]interface{} based on differences between
// two of the same models to be used in database updates
func ResultDifference(v1 interface{}, v2 interface{}) map[string]interface{} {
	results := make(map[string]interface{}, 0)
	for _, d := range CompareModels(v1, v2) {
		results[d.Field] = d.New
	}

	return results
}

// CompareModels will compare two of the same data models and determine
// what fields are different
// Example
//
// []database.ModelDifference{
//  database.ModelDifference{
//    Field: "type",
//    Old:   2.000000,
//    New:   1.000000,
//  },
//}
func CompareModels(v1 interface{}, v2 interface{}) []ModelDifference {
	var v1Model = structs.Map(v1)
	var v2Model = structs.Map(v2)

	var differences []ModelDifference
	for v1Field := range v1Model {
		for v2Field := range v2Model {
			fType := fmt.Sprintf("%v", reflect.TypeOf(v1Model[v1Field]))

			// struct mapper will export interfaces
			isValidField := !strings.Contains(fType, "interface")

			if v2Field == v1Field && isValidField && v1Model[v1Field] != v2Model[v1Field] {
				differences = append(differences, ModelDifference{
					Field: v1Field,
					Old:   v1Model[v1Field],
					New:   v2Model[v1Field],
				})
			}
		}
	}

	return differences
}
