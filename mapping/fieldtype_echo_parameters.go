package mapping

import (
	"github.com/omniscale/imposm3/element"
	"github.com/omniscale/imposm3/geom"
)

func init() {

	registerFieldTypes(
		FieldType{
			Name:     "echo_hello_world",
			GoType:   "string",
			Func:     getField_Echo_parameters,
			MakeFunc: nil,
		},
		FieldTypeDoc{
			Label: "echo 'Hello world' value -  example ",
			Doc: "\n For demonstration purpose. ....  " +
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

func getField_Echo_parameters(val string, elem *element.OSMElem, geom *geom.Geometry, match Match) interface{} {

	return "Hello World!"
}
