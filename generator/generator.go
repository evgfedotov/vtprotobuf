package generator

import (
	"runtime/debug"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type helper struct {
	path   protogen.GoImportPath
	plugin string
}

type Extensions struct {
	Poolable map[protogen.GoIdent]bool
}

type Generator struct {
	seen map[helper]bool
	ext  *Extensions
}

func NewGenerator(ext *Extensions) *Generator {
	return &Generator{
		seen: make(map[helper]bool),
		ext:  ext,
	}
}

func (gen *Generator) GenerateFile(gf *protogen.GeneratedFile, file *protogen.File) bool {
	if file.Desc.Syntax() != protoreflect.Proto3 {
		return false
	}

	p := &GeneratedFile{
		GeneratedFile: gf,
		Ext:           gen.ext,
	}

	p.P("// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.")
	if bi, ok := debug.ReadBuildInfo(); ok {
		p.P("// protoc-gen-go-vtproto version: ", bi.Main.Version)
	}
	p.P("// source: ", file.Desc.Path())
	p.P()
	p.P("package ", file.GoPackageName)
	p.P()

	var generated bool
	for _, plugin := range pluginsForFile(p) {
		if plugin.GenerateFile(file) {
			generated = true

			key := helper{
				path:   file.GoImportPath,
				plugin: plugin.Name(),
			}
			if !gen.seen[key] {
				plugin.GenerateHelpers()
				gen.seen[key] = true
			}
		}
	}

	return generated
}
