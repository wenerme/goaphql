package gqlp

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/wenerme/goaphql/gqll"
)

func ParseContent(content string) (doc *gqll.Document, err error) {
	input := antlr.NewInputStream(content)
	lexer := NewGraphQLLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)

	parser := NewGraphQLParser(stream)
	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	parser.BuildParseTrees = true

	ctx := parser.Document()
	doc = ctx.Accept(NewGraphQLLangVisitor()).(*gqll.Document)
	return
}
