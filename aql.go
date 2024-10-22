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

type StringQuery struct {
	q []string
}

func (s *StringQuery) sprintf(value any, quote string) string {
	switch t := value.(type) {
	case Literal:
		return string(t)
	case int:
		return strconv.Itoa(t)
	default:
		return fmt.Sprintf(`%s%s%s`, quote, value, quote)
	}
}

func (s *StringQuery) append(vals ...string) *StringQuery {
	s.q = append(s.q, vals...)
	return s
}

func (s *StringQuery) Group(q *StringQuery) *StringQuery {
	return s.append(fmt.Sprintf("( %s )", q.Encode()))
}

func (s *StringQuery) And() *StringQuery {
	return s.append("AND")
}

func (s *StringQuery) Contains(attr string, value any) *StringQuery {
	return s.append(attr, "~", s.sprintf(value, quote))
}

func (s *StringQuery) Or() *StringQuery {
	return s.append("OR")
}

func (s *StringQuery) Gtr(attr string, value any) *StringQuery {
	return s.append(attr, ">", s.sprintf(value, quote))
}

func (s *StringQuery) Less(attr string, value any) *StringQuery {
	return s.append(attr, "<", s.sprintf(value, quote))
}

func (s *StringQuery) GtrEqualTo(attr string, value any) *StringQuery {
	return s.append(attr, ">=", s.sprintf(value, quote))
}

func (s *StringQuery) LessEqualTo(attr string, value any) *StringQuery {
	return s.append(attr, "<=", s.sprintf(value, quote))
}

func (s *StringQuery) Equal(attr string, value any) *StringQuery {
	return s.append(attr, "=", s.sprintf(value, quote))
}

func (s *StringQuery) NotEqual(attr string, value any) *StringQuery {
	return s.append(attr, "!=", s.sprintf(value, quote))
}

func (s *StringQuery) In(attr string, values []any) *StringQuery {

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

func (s *StringQuery) Encode() string {
	return strings.Join(s.q, " ")
}

func NewStringQuery() *StringQuery {
	return &StringQuery{
		q: make([]string, 0),
	}
}
