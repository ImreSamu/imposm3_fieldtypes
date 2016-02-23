package mapping

import (
	"time"

	"github.com/omniscale/imposm3/element"
	"github.com/omniscale/imposm3/geom"
)

func init() {

	RegisterFieldTypes(
		FieldType{
			Name:     "current_dateutc",
			GoType:   "date",
			Func:     getField_CurrentDateUTC,
			MakeFunc: nil,
		},
	)

}

func getField_CurrentDateUTC(val string, elem *element.OSMElem, geom *geom.Geometry, match Match) interface{} {
	return time.Now().UTC().Format(time.RFC3339)[:10]
}
