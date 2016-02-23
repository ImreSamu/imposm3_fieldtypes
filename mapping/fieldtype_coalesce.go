package mapping

import (
	"errors"

	"github.com/omniscale/imposm3/element"
	"github.com/omniscale/imposm3/geom"
)

func init() {

	RegisterFieldTypes(
		FieldType{
			Name:     "coalesce",
			GoType:   "string",
			Func:     nil,
			MakeFunc: getField_MakeCoalesce,
		},
	)
}

func getField_MakeCoalesce(fieldName string, fieldType FieldType, field Field) (MakeValue, error) {
	_keysList, ok := field.Args["keys"]
	if !ok {
		return nil, errors.New("missing keys in args for coalesce")
	}

	keysList, ok := _keysList.([]interface{})
	if !ok {
		return nil, errors.New("keys in args for coalesce not a list")
	}

	for _, key := range keysList {
		_, ok := key.(string)
		if !ok {
			return nil, errors.New("value in keys not a string  ( coalesce mapping )")
		}
	}

	coalesce := func(val string, elem *element.OSMElem, geom *geom.Geometry, match Match) interface{} {
		if field.Key == "" {
			for _, value := range keysList {
				if r, ok := elem.Tags[value.(string)]; ok {
					return r
				}
			}
			return ""
		}
		return "-error-"
	}
	return coalesce, nil
}
