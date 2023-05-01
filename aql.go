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

// Literal can be used to tell the *StringQuery builder to not wrap the value
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

func Parse(in string) (*QueryBuilder, error) {
	var q = NewQueryBuilder()

	var tokens = strings.Split(strings.TrimSpace(in), " ")
	if len(tokens) < 3 {
		return nil, fmt.Errorf("aql: query should be more then 3 characters")
	}

	var scanValue = func(start int, tokens []string) (int, any) {
		var value = tokens[start]

		if strings.HasPrefix(value, quote) || strings.HasPrefix(value, singleQuote) {
			value = value[1:]
		} else {
			return start, Literal(value)
		}

		var i int
		for i = start + 1; i <= len(tokens)-1; i++ {
			nv := tokens[i]
			value = value + " " + nv

			if strings.HasSuffix(value, quote) || strings.HasSuffix(value, singleQuote) {
				break
			}
		}

		value = value[:len(value)-1]

		return i, value
	}

	var scanGroup = func(start int, tokens []string) (int, []any) {
		var out = []any{tokens[start][2 : len(tokens[start])-2]}

		var i int
		for i = start + 1; i <= len(tokens); i++ {
			t := tokens[i]
			if strings.HasSuffix(t, ",") {
				out = append(out, t[1:len(t)-2])
			}
			if strings.HasSuffix(t, ")") {
				out = append(out, t[1:len(t)-2])
				break
			}
		}

		return i, out

	}

	for i := 1; i <= len(tokens)-1; i++ {
		var char = tokens[i]

		if char == " " {
			continue
		}

		var (
			attr   string
			value  any
			values []any
		)

		attr = tokens[i-1]

		switch char {
		case "<=":
			i, value = scanValue(i+1, tokens)
			q = q.LessEqualTo(attr, value)
		case "<":
			i, value = scanValue(i+1, tokens)
			q = q.Less(attr, value)
		case ">=":
			i, value = scanValue(i+1, tokens)
			q = q.GtrEqualTo(attr, value)
		case ">":
			i, value = scanValue(i+1, tokens)
			q = q.Gtr(attr, value)
		case "=":
			i, value = scanValue(i+1, tokens)
			q = q.Equal(attr, value)
		case "!=":
			i, value = scanValue(i+1, tokens)
			q = q.NotEqual(attr, value)
		case "~":
			i, value = scanValue(i+1, tokens)
			q = q.Contains(attr, value)
		case "in":
			i, values = scanGroup(i+1, tokens)
			q = q.In(attr, values)
		case "AND":
			q = q.And()
		case "OR":
			q = q.Or()
		default:
			return nil, fmt.Errorf("aql: unknown validation %s", char)
		}
	}

	return q, nil
}
