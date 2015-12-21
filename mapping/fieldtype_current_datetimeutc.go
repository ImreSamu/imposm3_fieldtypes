package mapping

import (
	"time"

	"github.com/omniscale/imposm3/element"
	"github.com/omniscale/imposm3/geom"
)

func init() {

	registerFieldTypes(
		FieldType{
			Name:     "current_datetimeutc",
			GoType:   "timestamp",
			Func:     getField_CurrentDatetimeUTC,
			MakeFunc: nil,
		},
		FieldTypeDoc{
			Label: "JSON mapping example ",
			Doc: "\n For demonstration purpose. This mapping create a JSON field:  {\"value\":  value}` " +
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

}

func getField_CurrentDatetimeUTC(val string, elem *element.OSMElem, geom *geom.Geometry, match Match) interface{} {
	return time.Now().UTC().Format(time.RFC3339)
}
