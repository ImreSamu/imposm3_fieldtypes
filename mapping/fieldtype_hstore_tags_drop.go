package mapping

import (
	"errors"
	"path"
	"strings"

	"github.com/omniscale/imposm3/element"
	"github.com/omniscale/imposm3/geom"
)

func init() {

	registerFieldTypes(
		FieldType{
			Name:     "hstore_tags_drop",
			GoType:   "hstore_string",
			Func:     nil,
			MakeFunc: getField_MakeHstoreDrop,
		},
		FieldTypeDoc{
			Label: "hstore tags drop -  example ",
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

func getField_MakeHstoreDrop(fieldName string, fieldType FieldType, field Field) (MakeValue, error) {
	_keysList, ok := field.Args["drop_keys"]
	if !ok {
		return nil, errors.New("missing keys in 'drop_keys' for 'hstore_tags_drop' ")
	}

	keysList, ok := _keysList.([]interface{})
	if !ok {
		return nil, errors.New("keys in 'drop_keys' for 'hstore_tags_drop'  not a string list")
	}

	checkedKeysList := make([]Key, len(keysList), len(keysList))
	for i, key := range keysList {
		skey, ok := key.(string)
		if !ok {
			return nil, errors.New("value in 'drop_keys' not a string  ( 'hstore_tags_keep' )")
		} else {
			checkedKeysList[i] = Key(skey)
		}
	}
	dropFilter := newExcludeFilter(checkedKeysList)

	hstoredrop := func(val string, elem *element.OSMElem, geom *geom.Geometry, match Match) interface{} {

		var drop bool
		tags := make([]string, 0, len(elem.Tags))
		for k, v := range elem.Tags {

			drop = false

			if _, ok := dropFilter.keys[Key(k)]; ok {
				drop = true
			} else if dropFilter.matches != nil {
				for _, exkey := range dropFilter.matches {
					if ok, _ := path.Match(exkey, k); ok {
						drop = true
						break
					}
				}
			}

			if !drop {
				tags = append(tags, `"`+hstoreReplacer.Replace(k)+`"=>"`+hstoreReplacer.Replace(v)+`"`)
			}
		}
		return strings.Join(tags, ", ")
	}

	return hstoredrop, nil
}
