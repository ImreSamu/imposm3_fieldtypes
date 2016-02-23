package mapping

import (
	"strconv"

	"github.com/omniscale/imposm3/element"
	"github.com/omniscale/imposm3/geom"
)

func init() {

	RegisterFieldTypes(
		FieldType{
			Name:     "integer_tiny",
			GoType:   "int8",
			Func:     getField_IntegerTiny,
			MakeFunc: nil,
		},
	)

}

func getField_IntegerTiny(val string, elem *element.OSMElem, geom *geom.Geometry, match Match) interface{} {
	v, err := strconv.ParseInt(val, 10, 8)
	if err != nil {
		return nil
	}
	return v
}
