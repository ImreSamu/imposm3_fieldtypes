package mapping

import (
	"github.com/omniscale/imposm3/element"
	"github.com/omniscale/imposm3/geom"
)

func init() {

	RegisterFieldTypes(
		FieldType{
			Name:     "echo_hello_world",
			GoType:   "string",
			Func:     getField_Echo_parameters,
			MakeFunc: nil,
		},
	)

}

func getField_Echo_parameters(val string, elem *element.OSMElem, geom *geom.Geometry, match Match) interface{} {

	return "Hello World!"
}
