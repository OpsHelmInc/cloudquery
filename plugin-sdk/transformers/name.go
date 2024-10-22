package transformers

import (
	"reflect"
	"strings"

	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/caser"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
)

type NameTransformer func(reflect.StructField) (string, error)

var defaultCaser = caser.New()

func DefaultNameTransformer(field reflect.StructField) (string, error) {
	name := field.Name
	if jsonTag := strings.Split(field.Tag.Get("json"), ",")[0]; len(jsonTag) > 0 {
		// return empty string if the field is not related api response
		if jsonTag == "-" {
			return "", nil
		}
		if nameFromJSONTag := defaultCaser.ToSnake(jsonTag); schema.ValidColumnName(nameFromJSONTag) {
			return nameFromJSONTag, nil
		}
	}
	return defaultCaser.ToSnake(name), nil
}

var _ NameTransformer = DefaultNameTransformer
