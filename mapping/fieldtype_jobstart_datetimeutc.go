package mapping

import (
	"time"

	"github.com/omniscale/imposm3/element"
	"github.com/omniscale/imposm3/geom"
)

var getField_JobStartDatetimeUTCformatted string

func init() {

	RegisterFieldTypes(
		FieldType{
			Name:     "jobstart_datetimeutc",
			GoType:   "timestamp",
			Func:     getField_JobStartDatetimeUTC,
			MakeFunc: nil,
		},
	)

	getField_JobStartDatetimeUTCformatted = time.Now().UTC().Format(time.RFC3339)
}

func getField_JobStartDatetimeUTC(val string, elem *element.OSMElem, geom *geom.Geometry, match Match) interface{} {
	return getField_JobStartDatetimeUTCformatted
}
