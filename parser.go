package ast_assist

import (
	"github.com/think-next/ast-assist/internal"
	"go/ast"
)

type DeclPair struct {
	Name string
	Type string
}

// A FuncDecl represents a function declaration
type FuncDecl struct {
	Name     string
	Receiver DeclPair
	Param    []DeclPair
	Result   []DeclPair

	root *ast.FuncDecl
}

func getFuncDecl(name string, file *ast.File) FuncDecl {
	funcDecl := FuncDecl{}

	ast.Inspect(file, func(node ast.Node) bool {
		switch t := node.(type) {
		case *ast.FuncDecl:
			funcDecl.Name = t.Name.Name
			funcDecl.root = t

			if t.Recv != nil {
				field := t.Recv.List[0]
				for _, v := range field.Names {
					funcDecl.Receiver.Name = v.Name
					break
				}
				funcDecl.Receiver.Type = internal.TypeExprStr(field.Type)
			}

			if t.Type.Params != nil {
				for _, v := range t.Type.Params.List {
					for _, iv := range v.Names {
						pair := DeclPair{}
						pair.Name = iv.Name
						pair.Type = internal.TypeExprStr(v.Type)
						funcDecl.Param = append(funcDecl.Param, pair)
					}
				}
			}

			if t.Type.Results != nil {
				for _, v := range t.Type.Results.List {
					for _, iv := range v.Names {
						pair := DeclPair{}
						pair.Name = iv.Name
						pair.Type = internal.TypeExprStr(v.Type)
						funcDecl.Result = append(funcDecl.Result, pair)
					}
				}
			}
		}
		return true
	})

	return funcDecl
}
