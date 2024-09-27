package utils

import (
	"go/parser"
	"go/token"
)

// CheckSyntax checks if the provided code has any syntax errors
func CheckSyntax(code string) error {
	fset := token.NewFileSet()
	_, err := parser.ParseFile(fset, "", code, parser.ParseComments)
	return err
}
