package mapping

import (
	"time"

	"github.com/omniscale/imposm3/element"
	"github.com/omniscale/imposm3/geom"
)

func init() {

	registerFieldTypes(
		FieldType{
			Name:     "sleep",
			GoType:   "string",
			Func:     getField_sleep,
			MakeFunc: nil,
		},
		FieldTypeDoc{
			Label: "Sleep -  example ",
			Doc: "\n For demonstration purpose. ...  " +
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

func getField_sleep(val string, elem *element.OSMElem, geom *geom.Geometry, match Match) interface{} {

	time.Sleep(100 * time.Millisecond)

	// log.Print("IMPOSM3 too fast .. sleep  ...  Id :", elem.Id)

	return "IMPOSM3 too fast - sleep 100 millisecond"
}
