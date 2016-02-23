package mapping

import (
	_ "strconv"

	"github.com/omniscale/imposm3/element"
	"github.com/omniscale/imposm3/geom"
)

func init() {

	RegisterFieldTypes(
		FieldType{
			Name:     "mapping_key_value",
			GoType:   "string",
			Func:     getField_KeyValueName,
			MakeFunc: nil,
		},
	)
}

func getField_KeyValueName(val string, elem *element.OSMElem, geom *geom.Geometry, match Match) interface{} {
	return match.Key + "_" + match.Value
}
