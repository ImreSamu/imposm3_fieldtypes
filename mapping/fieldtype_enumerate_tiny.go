package mapping

import (
	_ "github.com/omniscale/imposm3/element"
	_ "github.com/omniscale/imposm3/geom"
)

func init() {

	RegisterFieldTypes(
		FieldType{
			Name:     "enumerate_tiny",
			GoType:   "int8",
			Func:     nil,
			MakeFunc: MakeEnumerate,
		},
	)
}
