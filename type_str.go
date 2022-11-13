package ast_assist

import (
	"go/ast"
)

func GetIndexExprStr(index *ast.IndexExpr) string {
	l := ""
	r := ""
	switch e := index.X.(type) {
	case *ast.IndexExpr:
		l += GetIndexExprStr(e)
	}

	switch e := index.Index.(type) {
	case *ast.BasicLit:
		r += GetBasicLitStr(e)
	}
	return l + "[" + r + "]"
}

func GetBasicLitStr(basic *ast.BasicLit) string {
	return basic.Value
}
