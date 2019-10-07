package jrpc

import (
	"bytes"
	"fmt"
)

type Service struct {
	Methods []*Method
}

type Method struct {
	URL          string
	HTTPMethod   string
	RequestType  string
	ResponseType string
}

type FileDescriptor struct {
	PackageName    string
	PackageComment []string
	Imports        []*Import
	Service        *Service

	out *bytes.Buffer
}

func New() *FileDescriptor {
	return &FileDescriptor{
		Imports: []*Import{
			{
				Path: "github.com/jakewright/home-automation/libraries/go/rpc",
			},
		},
		out: bytes.NewBuffer(nil),
	}
}

type Import struct {
	Alias string
	Path  string
}

func (f *FileDescriptor) Generate() string {

	f.pn("// Code generated by protoc-gen-jrpc. DO NOT EDIT.")
	f.pn()
	for _, c := range f.PackageComment {
		f.pn("// ", c)
	}
	f.pn("package ", f.PackageName)
	f.pn()

	f.p(renderImports(f.Imports))

	for _, m := range f.Service.Methods {
		f.pn("func (r *", m.RequestType, ") Do() (*", m.ResponseType, ", error) {")
		f.pn("\treq := rpc.Request{")
		f.pn("\t\tMethod: \"", m.HTTPMethod, "\",")
		f.pn("\t\tURL:    \"", m.URL, "\",")
		f.pn("\t\tBody:   r,")
		f.pn("\t}")
		f.pn()
		f.pn("\tvar rsp *", m.ResponseType)
		f.pn()
		f.pn("\t_, err := rpc.Do(req, &rsp)")
		f.pn("\tif err != nil {")
		f.pn("\t\treturn nil, err")
		f.pn("\t}")
		f.pn()
		f.pn("\treturn rsp, nil")
		f.pn("}")
	}

	return f.out.String()
}

func renderImports(imports []*Import) string {
	if len(imports) == 0 {
		return ""
	}

	if len(imports) == 1 {
		s := "import "
		if imports[0].Alias != "" {
			s += imports[0].Alias + " "
		}
		return fmt.Sprintf("%s \"%s\"\n\n", s, imports[0].Path)
	}

	s := "import (\n"
	for _, i := range imports {
		s += "\t"
		if i.Alias != "" {
			s += i.Alias + " "
		}
		s += "\"" + i.Path + "\"\n"
	}
	s += ")\n\n"

	return s
}

func (f *FileDescriptor) p(args ...string) {
	for _, v := range args {
		f.out.WriteString(v)
	}
}

func (f *FileDescriptor) pn(args ...string) {
	for _, v := range args {
		f.out.WriteString(v)
	}
	f.out.WriteByte('\n')
}