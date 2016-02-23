package mapping

import (
	"fmt"

	"github.com/omniscale/imposm3/element"
	"github.com/omniscale/imposm3/geom"
)

func init() {

	RegisterFieldTypes(
		FieldType{
			Name:     "json_example",
			GoType:   "json_string",
			Func:     getField_json_example,
			MakeFunc: nil,
		},
	)

}

func getField_json_example(val string, elem *element.OSMElem, geom *geom.Geometry, match Match) interface{} {
	return fmt.Sprintf(` {"value": "%s"}`, val)
}
