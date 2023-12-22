package main

import (
	"go/ast"
	"go/types"
)

//Package
type Package struct {
	name  string
	defs  map[*ast.Ident]types.Object
	files []*File
}
