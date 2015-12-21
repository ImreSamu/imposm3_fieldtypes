package mapping

import (
	"fmt"

	"github.com/omniscale/imposm3/element"
	"github.com/omniscale/imposm3/geom"
)

func init() {

	registerFieldTypes(
		FieldType{
			Name:     "jsonb_example",
			GoType:   "jsonb_string",
			Func:     getField_jsonb_example,
			MakeFunc: nil,
		},
		FieldTypeDoc{
			Label: "JSONB mapping example ( only for PostgreSQL>=9.4 ) ",
			Doc: "\n For demonstration purpose. This mapping create a JSONB field : {\"value\":  value}` " +
				"\n .... " +
				"\n Example ....     ", // markdown text ?
			Parameters: "",
			Status:     "example", // draft,notest,tested,notsupported,production,depricated
			Category:   "tags",    // tags,geometry,osm_id,system
			Repository: "https://github.com/..../.....",
			Lastupdate: "2015-12-21",
			Author:     "@ImreSamu", // e-mail or other nick name
			Tags:       nil,         //  any other key value ... map[string][]string
		},
	)

}

func getField_jsonb_example(val string, elem *element.OSMElem, geom *geom.Geometry, match Match) interface{} {
	return fmt.Sprintf(` {"value": "%s"}`, val)
}
