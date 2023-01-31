package services

import (
	"bytes"
	"embed"
	"fmt"
	"go/format"
	"io"
	"os"
	"path"
	"reflect"
	"runtime"
	"strings"
	"text/template"

	caser "github.com/cloudquery/plugin-sdk/caser"
)

//go:embed templates/*.go.tpl
var templatesFS embed.FS

// these method name prefixes will be part of the generated client interface
var acceptedPrefixes = []string{
	"List", "Get", "Describe", "Search", "BatchGet",
}

// these methods will be included despite not starting with an accepted prefix
var exceptions = []string{
	"QuerySchemaVersionMetadata",
	"GenerateCredentialReport",
}

// Adapted from https://stackoverflow.com/a/54129236
func signature(name string, f interface{}) string {
	t := reflect.TypeOf(f)
	if t.Kind() != reflect.Func {
		return "<not a function>"
	}

	buf := strings.Builder{}
	buf.WriteString(name + "(")
	for i := 0; i < t.NumIn(); i++ {
		if i > 0 {
			buf.WriteString(", ")
		}
		if t.IsVariadic() && i == t.NumIn()-1 {
			buf.WriteString("..." + strings.TrimPrefix(t.In(i).String(), "[]"))
		} else {
			buf.WriteString(t.In(i).String())
		}
	}
	buf.WriteString(")")
	if numOut := t.NumOut(); numOut > 0 {
		if numOut > 1 {
			buf.WriteString(" (")
		} else {
			buf.WriteString(" ")
		}
		for i := 0; i < t.NumOut(); i++ {
			if i > 0 {
				buf.WriteString(", ")
			}
			buf.WriteString(t.Out(i).String())
		}
		if numOut > 1 {
			buf.WriteString(")")
		}
	}

	return buf.String()
}

func shouldInclude(name string) bool {
	for _, t := range acceptedPrefixes {
		if strings.HasPrefix(name, t) {
			return true
		}
	}
	for _, e := range exceptions {
		if name == e {
			return true
		}
	}
	return false
}

type serviceInfo struct {
	Import      string
	Name        string
	PackageName string
	ClientName  string
	Signatures  []string
}

func getServiceInfo(client interface{}) serviceInfo {
	v := reflect.ValueOf(client)
	t := v.Type()
	pkgPath := t.Elem().PkgPath()
	parts := strings.Split(pkgPath, "/")
	pkgName := parts[len(parts)-1]
	csr := caser.New()
	name := csr.ToPascal(pkgName)
	clientName := name + "Client"
	signatures := make([]string, 0)
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		if shouldInclude(method.Name) {
			sig := signature(method.Name, v.Method(i).Interface())
			signatures = append(signatures, sig)
		}
	}
	return serviceInfo{
		Import:      pkgPath,
		Name:        name,
		PackageName: pkgName,
		ClientName:  clientName,
		Signatures:  signatures,
	}
}

// Generate generates a services.go file and individual service files from the clients defined in clients.go
func Generate() error {
	services := make([]serviceInfo, 0)
	for _, client := range clients {
		services = append(services, getServiceInfo(client))
	}

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return fmt.Errorf("failed to get caller information")
	}

	// write services.go file
	servicesTpl, err := template.New("services.go.tpl").ParseFS(templatesFS, "templates/services.go.tpl")
	if err != nil {
		return err
	}

	var buff bytes.Buffer
	all := append(services, customClients...)
	if err := servicesTpl.Execute(&buff, all); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}
	filePath := path.Join(path.Dir(filename), "../../client/services.go")
	formatAndWriteFile(filePath, buff)

	// write individual service files
	serviceTpl, err := template.New("service.go.tpl").ParseFS(templatesFS, "templates/service.go.tpl")
	if err != nil {
		return err
	}

	for _, service := range services {
		buff = bytes.Buffer{}
		if err := serviceTpl.Execute(&buff, service); err != nil {
			return fmt.Errorf("failed to execute template: %w", err)
		}
		filePath := path.Join(path.Dir(filename), fmt.Sprintf("../../client/services/%s.go", service.PackageName))
		formatAndWriteFile(filePath, buff)
	}

	for _, custom := range customClients {
		src := path.Join(path.Dir(filename), "custom", custom.PackageName+".go")
		dst := path.Join(path.Dir(filename), "../../client/services/"+custom.PackageName+".go")
		err := copyCustomFile(src, dst)
		if err != nil {
			return err
		}
	}

	return nil
}

func copyCustomFile(src, dst string) error {
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	return err
}

func formatAndWriteFile(filePath string, buff bytes.Buffer) error {
	content := buff.Bytes()
	formattedContent, err := format.Source(buff.Bytes())
	if err != nil {
		fmt.Printf("failed to format source: %s: %v\n", filePath, err)
	} else {
		content = formattedContent
	}
	if err := os.WriteFile(filePath, content, 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", filePath, err)
	}
	return nil
}
