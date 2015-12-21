package mapping

import (
	"github.com/omniscale/imposm3/element"
	"github.com/omniscale/imposm3/geom"
)

func init() {

	registerFieldTypes(
		FieldType{
			Name:     "string_char1",
			GoType:   "char1",
			Func:     getField_String1,
			MakeFunc: nil,
		},
		FieldTypeDoc{
			Label: "1char field  -  example ",
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

}

func getField_String1(val string, elem *element.OSMElem, geom *geom.Geometry, match Match) interface{} {
	return (val + " ")[:1]
}
