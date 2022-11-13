package ast_assist

import (
	"go/ast"
)

func GetIndexExprStr(index *ast.IndexExpr) string {
	l := ""
	r := ""
	switch e := index.X.(type) {
	// case *ast.IndexExpr:
	// 	l += GetIndexExprStr(e)
	case *ast.SelectorExpr:
		l += GetSelectorExprStr(e)
	case *ast.Ident:
		l += GetIdentStr(e)
	}

	switch e := index.Index.(type) {
	case *ast.BasicLit:
		r += GetBasicLitStr(e)
	}
	return l + "[" + r + "]"
}

func GetSelectorExprStr(expr *ast.SelectorExpr) string {
	selStr := expr.Sel.Name
	xStr := ""

	x := expr.X
	switch ti := x.(type) {
	case *ast.Ident:
		xStr = GetIdentStr(ti)
	case *ast.IndexExpr:
		xStr = GetIndexExprStr(ti)
	case *ast.SelectorExpr:
		xStr = GetSelectorExprStr(ti)
	}

	return xStr + "." + selStr
}

func GetIdentStr(basic *ast.Ident) string {
	return basic.Name
}

func GetBasicLitStr(basic *ast.BasicLit) string {
	return basic.Value
}
