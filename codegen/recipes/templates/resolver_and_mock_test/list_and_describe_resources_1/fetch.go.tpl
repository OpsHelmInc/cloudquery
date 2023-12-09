// Code generated by codegen; DO NOT EDIT.

package {{.Service}}

import (
	"context"

    "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/{{.Service}}"
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/schema"
)

func {{.Table.Resolver}}(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
    var input {{.Service}}.{{.ListMethod.Method.Name}}Input{{ if .CustomListInput }} = {{.CustomListInput}}{{ end }}
	c := meta.(*client.Client)
	svc := c.Services().{{.CloudQueryServiceName}}
	for {
        response, err := svc.{{.ListMethod.Method.Name}}(ctx, &input)
        if err != nil {
            return err
        }
        {{- if .ListMethod.OutputFieldName }}
        res <- response.{{.ListMethod.OutputFieldName}}
        {{- else }}
        res <- response
        {{- end }}

        if aws.ToString(response.NextToken) == "" {
            break
        }
        input.NextToken = response.NextToken
    }
    return nil
}

func {{.Table.PreResourceResolver}}(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().{{.CloudQueryServiceName}}
	var input {{.Service}}.{{.DescribeMethod.Method.Name}}Input{{ if .CustomDescribeInput }} = {{.CustomDescribeInput}}{{ end }}
	output, err := svc.{{.DescribeMethod.Method.Name}}(ctx, &input)
	if err != nil {
		return err
	}
	{{- if .DescribeMethod.OutputFieldName }}
	resource.Item = output.{{.DescribeMethod.OutputFieldName}}
	{{- else }}
	resource.Item = output
	{{- end }}
	return nil
}
