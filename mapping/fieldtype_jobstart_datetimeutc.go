package mapping

import (
	"time"

	"github.com/omniscale/imposm3/element"
	"github.com/omniscale/imposm3/geom"
)

var getField_JobStartDatetimeUTCformatted string

func init() {

	registerFieldTypes(
		FieldType{
			Name:     "jobstart_datetimeutc",
			GoType:   "timestamp",
			Func:     getField_JobStartDatetimeUTC,
			MakeFunc: nil,
		},
		FieldTypeDoc{
			Label: "return: IMPOSM3 Start datetime  -  example ",
			Doc: "\n For demonstration purpose.  " +
				"\n .... " +
				"\n Example ....     ", // markdown text ?
			Parameters: "",
			Status:     "example", // draft,notest,tested,notsupported,production,depricated
			Category:   "system",  // tags,geometry,osm_id,system
			Repository: "https://github.com/..../.....",
			Lastupdate: "2015-12-21",
			Author:     "@ImreSamu", // e-mail or other nick name
			Tags:       nil,         //  any other key value ... map[string][]string
		},
	)

	getField_JobStartDatetimeUTCformatted = time.Now().UTC().Format(time.RFC3339)
}

func getField_JobStartDatetimeUTC(val string, elem *element.OSMElem, geom *geom.Geometry, match Match) interface{} {
	return getField_JobStartDatetimeUTCformatted
}
