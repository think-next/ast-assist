package ast_assist

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"io/ioutil"
)

type CodeParser struct {
	fileSet  *token.FileSet
	fileNode *ast.File
	buf      *bytes.Buffer
	file     string
}

func (c *CodeParser) GetFileSet() *token.FileSet {
	return c.fileSet
}

func (c *CodeParser) GetFileNode() *ast.File {
	return c.fileNode
}

func (c *CodeParser) GetBuf() *bytes.Buffer {
	return c.buf
}

func NewCodeParser(file string, src interface{}) (*CodeParser, error) {
	parser := &CodeParser{
		buf: bytes.NewBuffer([]byte{}),
	}

	if file != "" {
		content, err := ioutil.ReadFile(file)
		if err != nil {
			return nil, err
		}
		parser.buf = bytes.NewBuffer(content)
		parser.file = file
	} else {

		var srcBytes []byte
		if src != nil {
			switch s := src.(type) {
			case string:
				srcBytes = []byte(s)
			case []byte:
				srcBytes = s
			case *bytes.Buffer:
				// is io.Reader, but src is already available in []byte form
				if s != nil {
					srcBytes = s.Bytes()
				}
			case io.Reader:
				srcBytes, _ = io.ReadAll(s)
			default:
				return nil, errors.New(fmt.Sprintf("debug info: invalid source type(%T)", src))
			}

			parser.buf = bytes.NewBuffer(srcBytes)
		}
	}

	if err := parser.init(file, src); err != nil {
		return nil, err
	}

	return parser, nil
}

func (c *CodeParser) init(filename string, src interface{}) error {
	fSet := token.NewFileSet()
	fNode, err := parser.ParseFile(fSet, filename, src, parser.ParseComments)
	if err != nil {
		// Print out the bad code with line numbers.
		// This should never happen in practice, but it can while changing generated code
		// so consider this a debugging aid.
		var src bytes.Buffer

		s := bufio.NewScanner(bytes.NewReader(c.buf.Bytes()))
		for line := 1; s.Scan(); line++ {
			fmt.Fprintf(&src, "%5d\t%s\n", line, s.Bytes())
		}
		return fmt.Errorf("debug info: %v unparsable code source: %v\n%v", filename, err, c.buf.String())
	}

	c.fileSet = fSet
	c.fileNode = fNode
	return nil
}
