package mapping

import (
	"fmt"

	"github.com/omniscale/imposm3/element"
	"github.com/omniscale/imposm3/geom"
)

func init() {

	RegisterFieldTypes(
		FieldType{
			Name:     "jsonb_example",
			GoType:   "jsonb_string",
			Func:     getField_jsonb_example,
			MakeFunc: nil,
		},
	)

}

func getField_jsonb_example(val string, elem *element.OSMElem, geom *geom.Geometry, match Match) interface{} {
	return fmt.Sprintf(` {"value": "%s"}`, val)
}
