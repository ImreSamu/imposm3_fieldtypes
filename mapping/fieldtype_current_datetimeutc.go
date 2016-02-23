package mapping

import (
	"time"

	"github.com/omniscale/imposm3/element"
	"github.com/omniscale/imposm3/geom"
)

func init() {

	RegisterFieldTypes(
		FieldType{
			Name:     "current_datetimeutc",
			GoType:   "timestamp",
			Func:     getField_CurrentDatetimeUTC,
			MakeFunc: nil,
		},
	)

}

func getField_CurrentDatetimeUTC(val string, elem *element.OSMElem, geom *geom.Geometry, match Match) interface{} {
	return time.Now().UTC().Format(time.RFC3339)
}
