package mapping

import (
	"time"

	"github.com/omniscale/imposm3/element"
	"github.com/omniscale/imposm3/geom"
)

func init() {

	registerFieldTypes(
		FieldType{
			Name:     "current_dateutc",
			GoType:   "date",
			Func:     getField_CurrentDateUTC,
			MakeFunc: nil,
		},
		FieldTypeDoc{
			Label: "Current Date -  example ",
			Doc: "\n For demonstration purpose. This mapping create a DATE field contains current date value  " +
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

func getField_CurrentDateUTC(val string, elem *element.OSMElem, geom *geom.Geometry, match Match) interface{} {
	return time.Now().UTC().Format(time.RFC3339)[:10]
}
