package client

import (
	"context"
	"reflect"
	"strings"
	"time"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/thoas/go-funk"

	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func TimestampTypeTransformer(field reflect.StructField) (arrow.DataType, error) {
	if !strings.HasSuffix(field.Name, "At") {
		return nil, nil // fallback
	}

	switch field.Type {
	case reflect.TypeOf(""), reflect.TypeOf(new(string)):
		return arrow.FixedWidthTypes.Timestamp_us, nil
	default:
		return nil, nil // fallback
	}
}

func TimestampResolverTransformer(field reflect.StructField, path string) schema.ColumnResolver {
	if t, _ := TimestampTypeTransformer(field); t != arrow.FixedWidthTypes.Timestamp_us {
		return transformers.DefaultResolverTransformer(field, path)
	}

	if field.Type.Kind() == reflect.Pointer {
		return func(_ context.Context, _ schema.ClientMeta, r *schema.Resource, c schema.Column) error {
			val := funk.Get(r.Item, path)
			if val == nil {
				return r.Set(c.Name, nil)
			}

			t, err := time.Parse(time.RFC3339, *(val.(*string)))
			if err != nil {
				return err
			}
			return r.Set(c.Name, t)
		}
	}

	return func(_ context.Context, _ schema.ClientMeta, r *schema.Resource, c schema.Column) error {
		val := funk.Get(r.Item, path)
		if val == nil {
			return r.Set(c.Name, nil)
		}

		t, err := time.Parse(time.RFC3339, val.(string))
		if err != nil {
			return err
		}
		return r.Set(c.Name, t)
	}
}

// TagsResolverTransformer adds possibility to override the column resolver for "Tags" field
func TagsResolverTransformer(field reflect.StructField, path string) schema.ColumnResolver {
	if path == "Tags" {
		return ResolveTags
	}
	return transformers.DefaultResolverTransformer(field, path)
}
