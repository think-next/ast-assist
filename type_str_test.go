package ast_assist

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"testing"
)

var src = `
package main

var a string = live[0].shows.list[1]
`

func TestName(t *testing.T) {
	fSet := token.NewFileSet()
	fNode, err := parser.ParseFile(fSet, "", src, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	ast.Inspect(fNode, func(node ast.Node) bool {
		switch n := node.(type) {
		case *ast.IndexExpr:
			x := n.X
			i := n.Index
		}
		return true
	})
}
