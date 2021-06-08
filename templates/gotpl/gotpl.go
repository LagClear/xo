// Package gotpl provides a Go template for xo.
package gotpl

import (
	"context"
	"embed"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/xo/xo/templates"
	"golang.org/x/tools/imports"
	"mvdan.cc/gofumpt/format"
)

func init() {
	first := true
	knownTypes := map[string]bool{
		"bool":        true,
		"string":      true,
		"byte":        true,
		"rune":        true,
		"int":         true,
		"int16":       true,
		"int32":       true,
		"int64":       true,
		"uint":        true,
		"uint8":       true,
		"uint16":      true,
		"uint32":      true,
		"uint64":      true,
		"float32":     true,
		"float64":     true,
		"Slice":       true,
		"StringSlice": true,
	}
	shortNames := map[string]string{
		"bool":        "b",
		"string":      "s",
		"byte":        "b",
		"rune":        "r",
		"int":         "i",
		"int16":       "i",
		"int32":       "i",
		"int64":       "i",
		"uint":        "u",
		"uint8":       "u",
		"uint16":      "u",
		"uint32":      "u",
		"uint64":      "u",
		"float32":     "f",
		"float64":     "f",
		"Slice":       "s",
		"StringSlice": "ss",
	}
	templates.Register("go", &templates.TemplateSet{
		Files:   Files,
		FileExt: ".xo.go",
		AddType: func(typ string) {
			knownTypes[typ] = true
		},
		Flags: []templates.Flag{
			{
				ContextKey: NotFirstKey,
				Desc:       "disable package comment (ie, not first generated file)",
				Short:      '2',
				Default:    "false",
				Value:      false,
			},
			{
				ContextKey:  Int32Key,
				Desc:        "int32 type (default: int)",
				PlaceHolder: "int",
				Default:     "int",
				Value:       "",
			},
			{
				ContextKey:  Uint32Key,
				Desc:        "uint32 type (default: uint)",
				PlaceHolder: "uint",
				Default:     "uint",
				Value:       "",
			},
			{
				ContextKey:  PkgKey,
				Desc:        "package name",
				PlaceHolder: "<name>",
				Value:       "",
			},
			{
				ContextKey:  TagsKey,
				Desc:        "build tags",
				PlaceHolder: "<tags>",
				Value:       "",
			},
			{
				ContextKey:  CustomKey,
				Desc:        "package name for custom types",
				PlaceHolder: "<name>",
				Value:       "",
			},
			{
				ContextKey:  ConflictKey,
				Desc:        "name conflict suffix (default: Val)",
				PlaceHolder: "Val",
				Default:     "Val",
				Value:       "",
			},
			{
				ContextKey:  EscKey,
				Desc:        "escape fields (none, schema, table, column, all; default: none)",
				PlaceHolder: "none",
				Default:     "none",
				Value:       []string{},
				Enums:       []string{"none", "schema", "table", "column", "all"},
			},
		},
		Funcs: func(ctx context.Context) template.FuncMap {
			return NewFuncs(ctx, knownTypes, shortNames, &first).FuncMap()
		},
		HeaderTemplate: func(ctx context.Context) *templates.Template {
			return &templates.Template{
				Template: "hdr",
			}
		},
		PackageTemplates: func(ctx context.Context) []*templates.Template {
			if NotFirst(ctx) {
				return nil
			}
			return []*templates.Template{
				{
					Template: "db",
					Name:     "db",
				},
			}
		},
		FileName: func(ctx context.Context, tpl *templates.Template) string {
			if templates.GenType(ctx) == "schema" {
				switch tpl.Template {
				case "typedef", "enum", "index", "foreignkey", "proc":
					return strings.ToLower(tpl.Type)
				}
			}
			file := tpl.Name
			if file == "" {
				file = tpl.Type
			}
			return strings.ToLower(file)
		},
		Post: func(ctx context.Context, buf []byte) ([]byte, error) {
			// imports processing
			buf, err := imports.Process("", buf, nil)
			if err != nil {
				return nil, err
			}
			// format
			return format.Source(buf, format.Options{
				ExtraRules: true,
			})
		},
	})
}

// Context keys.
const (
	NotFirstKey templates.ContextKey = "not-first"
	Int32Key    templates.ContextKey = "int32"
	Uint32Key   templates.ContextKey = "uint32"
	PkgKey      templates.ContextKey = "pkg"
	TagsKey     templates.ContextKey = "tags"
	CustomKey   templates.ContextKey = "custom"
	ConflictKey templates.ContextKey = "conflict"
	EscKey      templates.ContextKey = "esc"
)

// NotFirst returns not-first from the context.
func NotFirst(ctx context.Context) bool {
	b, _ := ctx.Value(NotFirstKey).(bool)
	return b
}

// Int32 returns int32 from the context.
func Int32(ctx context.Context) string {
	s, _ := ctx.Value(Int32Key).(string)
	return s
}

// Uint32 returns uint32 from the context.
func Uint32(ctx context.Context) string {
	s, _ := ctx.Value(Uint32Key).(string)
	return s
}

// Pkg returns pkg from the context.
func Pkg(ctx context.Context) string {
	s, _ := ctx.Value(PkgKey).(string)
	if s == "" {
		s = filepath.Base(templates.Out(ctx))
	}
	return strings.ToLower(s)
}

// Tags returns tags from the context.
func Tags(ctx context.Context) string {
	s, _ := ctx.Value(TagsKey).(string)
	return s
}

// Custom returns custom-pkg from the context.
func Custom(ctx context.Context) string {
	s, _ := ctx.Value(CustomKey).(string)
	return s
}

// Conflict returns conflict from the context.
func Conflict(ctx context.Context) string {
	s, _ := ctx.Value(ConflictKey).(string)
	return s
}

// Esc returns esc from the context.
func Esc(ctx context.Context) []string {
	v, _ := ctx.Value(EscKey).([]string)
	return v
}

// contains returns true when s is in v.
func contains(v []string, s string) bool {
	for _, z := range v {
		if z == s {
			return true
		}
	}
	return false
}

// Files are the embedded Go templates.
//
//go:embed *.tpl
//go:embed */*.tpl
var Files embed.FS