package aql

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	singleQuote = `'`
	quote       = `"`
)

// Literal can be used to tell the Query builder to not wrap the value
// with quotes.  An example would be something like
// q.NotEqual("subject", aql.Literal("null"))
// this would become subject != null
type Literal string

const (
	Null Literal = "null"
)

type QueryBuilder struct {
	q []string
}

func (s *QueryBuilder) sprintf(value any, quote string) string {
	switch t := value.(type) {
	case Literal:
		return string(t)
	case int:
		return strconv.Itoa(t)
	default:
		return fmt.Sprintf(`%s%s%s`, quote, value, quote)
	}
}

func (s *QueryBuilder) append(vals ...string) *QueryBuilder {
	s.q = append(s.q, vals...)
	return s
}

func (s *QueryBuilder) Group(q *QueryBuilder) *QueryBuilder {
	return s.append(fmt.Sprintf("( %s )", q.Encode()))
}

func (s *QueryBuilder) And() *QueryBuilder {
	return s.append("AND")
}

func (s *QueryBuilder) Contains(attr string, value any) *QueryBuilder {
	return s.append(attr, "~", s.sprintf(value, quote))
}

func (s *QueryBuilder) Or() *QueryBuilder {
	return s.append("OR")
}

func (s *QueryBuilder) Gtr(attr string, value any) *QueryBuilder {
	return s.append(attr, ">", s.sprintf(value, quote))
}

func (s *QueryBuilder) Less(attr string, value any) *QueryBuilder {
	return s.append(attr, "<", s.sprintf(value, quote))
}

func (s *QueryBuilder) GtrEqualTo(attr string, value any) *QueryBuilder {
	return s.append(attr, ">=", s.sprintf(value, quote))
}

func (s *QueryBuilder) LessEqualTo(attr string, value any) *QueryBuilder {
	return s.append(attr, "<=", s.sprintf(value, quote))
}

func (s *QueryBuilder) Equal(attr string, value any) *QueryBuilder {
	return s.append(attr, "=", s.sprintf(value, quote))
}

func (s *QueryBuilder) NotEqual(attr string, value any) *QueryBuilder {
	return s.append(attr, "!=", s.sprintf(value, quote))
}

func (s *QueryBuilder) In(attr string, values []any) *QueryBuilder {

	var tupleStr string
	for i, v := range values {
		if i == 0 {
			tupleStr += "("
		}

		tupleStr += s.sprintf(v, singleQuote)

		if i == len(values)-1 {
			tupleStr += ")"
		} else {
			tupleStr += ", "
		}
	}

	return s.append(attr, "in", tupleStr)
}

func (s *QueryBuilder) Encode() string {
	return strings.Join(s.q, " ")
}

func NewQueryBuilder() *QueryBuilder {
	return &QueryBuilder{
		q: make([]string, 0),
	}
}
