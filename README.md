# imposm3_fieldtypes
imposm3 user fieldtypes for https://github.com/omniscale/imposm3 


Working with:
    
https://github.com/ImreSamu/imposm3/tree/fieldtypes_addon branch! 
or https://github.com/omniscale/imposm3/pull/86 


Experimental fieldtypes:
* mapping/fieldtype_coalesce.go
* mapping/fieldtype_current_datetimeutc.go
* mapping/fieldtype_current_dateutc.go
* mapping/fieldtype_echo_parameters.go
* mapping/fieldtype_enumerate_tiny.go
* mapping/fieldtype_hstore_tags_drop.go
* mapping/fieldtype_hstore_tags_keep.go
* mapping/fieldtype_integer_long.go
* mapping/fieldtype_integer_tiny.go
* mapping/fieldtype_jobstart_datetimeutc.go
* mapping/fieldtype_jsonb_example.go   ( PG >= 9.4 )
* mapping/fieldtype_json_example.go    ( PG >= 9.2 )
* mapping/fieldtype_mapping_key_value.go
* mapping/fieldtype_sleep.go
* mapping/fieldtype_string_char1.go

Example: 
* experimental_mapping.json   -   be careful, and remove the `sleep` mapping, before testing with big osm files!


