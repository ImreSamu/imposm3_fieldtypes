package mapping

import (
	"strconv"

	"github.com/omniscale/imposm3/element"
	"github.com/omniscale/imposm3/geom"
)

func init() {

	registerFieldTypes(
		FieldType{
			Name:     "integer_long",
			GoType:   "int64",
			Func:     getField_IntegerLong,
			MakeFunc: nil,
		},
		FieldTypeDoc{
			Label: "Convert to long integer -  example ",
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

func getField_IntegerLong(val string, elem *element.OSMElem, geom *geom.Geometry, match Match) interface{} {
	v, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return nil
	}
	return v
}
