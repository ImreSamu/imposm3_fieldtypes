package mapping

import (
	"github.com/omniscale/imposm3/element"
	"github.com/omniscale/imposm3/geom"
)

func init() {

	RegisterFieldTypes(
		FieldType{
			Name:     "string_char1",
			GoType:   "char1",
			Func:     getField_String1,
			MakeFunc: nil,
		},
	)

}

func getField_String1(val string, elem *element.OSMElem, geom *geom.Geometry, match Match) interface{} {
	return (val + " ")[:1]
}
