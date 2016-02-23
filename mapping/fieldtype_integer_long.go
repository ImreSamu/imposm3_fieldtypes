package mapping

import (
	"strconv"

	"github.com/omniscale/imposm3/element"
	"github.com/omniscale/imposm3/geom"
)

func init() {

	RegisterFieldTypes(
		FieldType{
			Name:     "integer_long",
			GoType:   "int64",
			Func:     getField_IntegerLong,
			MakeFunc: nil,
		},
	)
}

func getField_IntegerLong(val string, elem *element.OSMElem, geom *geom.Geometry, match Match) interface{} {
	v, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return nil
	}
	return v
}
