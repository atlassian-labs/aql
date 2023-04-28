package aql

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/atlassian-labs/aql/parser"
	"strconv"
	"strings"
)

type JqlQuery string
type SyntaxError struct {
	message string
}

func (e SyntaxError) Error() string {
	return e.message
}

type syntaxErrorListener struct {
	*antlr.DefaultErrorListener
	errors []string
	query  string
}

func (c *syntaxErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	c.errors = append(c.errors, fmt.Sprintln("line "+strconv.Itoa(line)+":"+strconv.Itoa(column)+" "+msg))
}

func (c *syntaxErrorListener) getSyntaxError() error {
	if len(c.errors) == 0 {
		return nil
	}
	return SyntaxError{message: strings.Join(c.errors, ";")}
}

func NewJqlQuery(q string) (JqlQuery, error) {
	syntaxErrorListener := &syntaxErrorListener{}

	is := antlr.NewInputStream(q)
	lexer := parser.NewJQLLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	jqlParser := parser.NewJQLParser(stream)
	jqlParser.RemoveErrorListeners()
	jqlParser.AddErrorListener(syntaxErrorListener)
	_ = jqlParser.JqlQuery()

	if len(syntaxErrorListener.errors) > 0 {
		return "", syntaxErrorListener.getSyntaxError()
	}
	return JqlQuery(q), nil
}
