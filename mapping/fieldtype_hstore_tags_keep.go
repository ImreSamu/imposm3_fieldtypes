package mapping

import (
	"errors"
	"path"
	"strings"

	"github.com/omniscale/imposm3/element"
	"github.com/omniscale/imposm3/geom"
)

func init() {

	//registerFieldTypes(FieldType{"hstore_tags_keep", "hstore_string", nil, MakeHstoreKeep})

	RegisterFieldTypes(
		FieldType{
			Name:     "hstore_tags_keep",
			GoType:   "hstore_string",
			Func:     nil,
			MakeFunc: getField_MakeHstoreKeep,
		},
	)

}

func getField_MakeHstoreKeep(fieldName string, fieldType FieldType, field Field) (MakeValue, error) {
	_keysList, ok := field.Args["keep_keys"]
	if !ok {
		return nil, errors.New("missing keys in 'keep_keys' for 'hstore_tags_keep' ")
	}

	keysList, ok := _keysList.([]interface{})
	if !ok {
		return nil, errors.New("keys in 'keep_keys' for 'hstore_tags_keep'  not a string list")
	}

	checkedKeysList := make([]Key, len(keysList), len(keysList))
	for i, key := range keysList {
		skey, ok := key.(string)
		if !ok {
			return nil, errors.New("value in 'keep_keys' not a string  ( 'hstore_tags_keep' )")
		} else {
			checkedKeysList[i] = Key(skey)
		}
	}
	keepFilter := newExcludeFilter(checkedKeysList)

	hstorekeep := func(val string, elem *element.OSMElem, geom *geom.Geometry, match Match) interface{} {

		var keep bool
		tags := make([]string, 0, len(elem.Tags))
		for k, v := range elem.Tags {

			keep = false

			if _, ok := keepFilter.keys[Key(k)]; ok {
				keep = true
			} else if keepFilter.matches != nil {
				for _, exkey := range keepFilter.matches {
					if ok, _ := path.Match(exkey, k); ok {
						keep = true
						break
					}
				}
			}

			if keep {
				tags = append(tags, `"`+hstoreReplacer.Replace(k)+`"=>"`+hstoreReplacer.Replace(v)+`"`)
			}
		}
		return strings.Join(tags, ", ")
	}

	return hstorekeep, nil
}
