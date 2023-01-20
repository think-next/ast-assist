package internal

import (
	"fmt"
	"go/ast"
)

func TypeExprStr(expr ast.Expr) string {
	switch e := expr.(type) {
	case *ast.Ident:
		return e.String()
	case *ast.StarExpr:
		return fmt.Sprintf("*%s", TypeExprStr(e.X))
	case *ast.SelectorExpr:
		return fmt.Sprintf("%s.%s", TypeExprStr(e.X), TypeExprStr(e.Sel))
	case *ast.ArrayType:
		return fmt.Sprintf("[]%s", TypeExprStr(e.Elt))
	case *ast.MapType:
		return fmt.Sprintf("map[%s]%s", TypeExprStr(e.Key), TypeExprStr(e.Value))
	case *ast.ChanType:
		if e.Dir == 1 {
			return fmt.Sprintf("chan<- %s", TypeExprStr(e.Value))
		}
		if e.Dir == 2 {
			return fmt.Sprintf("<-chan %s", TypeExprStr(e.Value))
		}
		return fmt.Sprintf("chan %s", TypeExprStr(e.Value))
	case *ast.Ellipsis:
		return fmt.Sprintf("...%s", TypeExprStr(e.Elt))
	default:
		fmt.Printf("debug info: this type(%T) is unsupport \n", e)
	}

	return ""
}
