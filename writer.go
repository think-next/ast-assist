package ast_assist

import (
	"bytes"
	"fmt"
	"go/printer"
)

type CodeWriter struct {
	buf *bytes.Buffer
}

// WriteLn a newline is appended to buf
func (c *CodeWriter) WriteLn(v ...interface{}) {
	for _, x := range v {
		fmt.Fprint(c.buf, x)
	}

	fmt.Fprintln(c.buf)
}

func (c *CodeWriter) Format(mode printer.Mode, width int) (*bytes.Buffer, error) {
	codeParser, err := NewCodeParser("", c.buf)
	if err != nil {
		return nil, err
	}

	bufReformat := new(bytes.Buffer)
	if err := (&printer.Config{Mode: mode, Tabwidth: width}).Fprint(bufReformat, codeParser.GetFileSet(), codeParser.GetFileNode()); err != nil {
		return nil, err
	}

	return bufReformat, nil
}
