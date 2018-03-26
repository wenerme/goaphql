package gqlp

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/wenerme/goaphql/gqll"
)

type ParseOption struct {
	Source  string // source of the content
	Content string
}

func ParseContent(content string) (doc *gqll.Document, err error) {
	return ParseWithOption(ParseOption{
		Source:  "inline",
		Content: content,
	})
}
func ParseWithOption(opt ParseOption) (doc *gqll.Document, err error) {
	input := antlr.NewInputStream(opt.Content)
	lexer := NewGraphQLLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)

	parser := NewGraphQLParser(stream)
	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	parser.BuildParseTrees = true

	ctx := parser.Document()
	visitor := NewGraphQLLangVisitor()
	visitor.Source = opt.Source
	doc = ctx.Accept(visitor).(*gqll.Document)
	return
}
