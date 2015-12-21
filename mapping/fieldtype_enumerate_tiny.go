package mapping

import (
	_ "github.com/omniscale/imposm3/element"
	_ "github.com/omniscale/imposm3/geom"
)

func init() {

	registerFieldTypes(
		FieldType{
			Name:     "enumerate_tiny",
			GoType:   "int8",
			Func:     nil,
			MakeFunc: MakeEnumerate,
		},
		FieldTypeDoc{
			Label: "Enumerate tiny mapping example ",
			Doc: "\n For demonstration purpose.   " +
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
