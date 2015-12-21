package mapping

import (
	_ "strconv"

	"github.com/omniscale/imposm3/element"
	"github.com/omniscale/imposm3/geom"
)

func init() {

	registerFieldTypes(
		FieldType{
			Name:     "mapping_key_value",
			GoType:   "string",
			Func:     getField_KeyValueName,
			MakeFunc: nil,
		},
		FieldTypeDoc{
			Label: "Mapping 'key' AND 'value' separated by '_'  - example ",
			Doc: "\n For demonstration purpose. " +
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

func getField_KeyValueName(val string, elem *element.OSMElem, geom *geom.Geometry, match Match) interface{} {
	return match.Key + "_" + match.Value
}
