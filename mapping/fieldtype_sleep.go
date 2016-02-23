package mapping

import (
	"time"

	"github.com/omniscale/imposm3/element"
	"github.com/omniscale/imposm3/geom"
)

func init() {

	RegisterFieldTypes(
		FieldType{
			Name:     "sleep",
			GoType:   "string",
			Func:     getField_sleep,
			MakeFunc: nil,
		},
	)

}

func getField_sleep(val string, elem *element.OSMElem, geom *geom.Geometry, match Match) interface{} {

	time.Sleep(100 * time.Millisecond)

	// log.Print("IMPOSM3 too fast .. sleep  ...  Id :", elem.Id)

	return "IMPOSM3 too fast - sleep 100 millisecond"
}
