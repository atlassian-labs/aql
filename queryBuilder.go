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

// Literal can be used to tell the JqlQuery builder to not wrap the value
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

func (qb *QueryBuilder) sprintf(value any, quote string) string {
	switch t := value.(type) {
	case Literal:
		return string(t)
	case int:
		return strconv.Itoa(t)
	default:
		return fmt.Sprintf(`%s%s%s`, quote, value, quote)
	}
}

func (qb *QueryBuilder) append(vals ...string) *QueryBuilder {
	qb.q = append(qb.q, vals...)
	return qb
}

func (qb *QueryBuilder) Group(qb2 *QueryBuilder) *QueryBuilder {
	return qb.append(fmt.Sprintf("( %s )", strings.Join(qb2.q, " ")))
}

func (qb *QueryBuilder) And() *QueryBuilder {
	return qb.append("AND")
}

func (qb *QueryBuilder) Contains(attr string, value any) *QueryBuilder {
	return qb.append(attr, "~", qb.sprintf(value, quote))
}

func (qb *QueryBuilder) Or() *QueryBuilder {
	return qb.append("OR")
}

func (qb *QueryBuilder) Gtr(attr string, value any) *QueryBuilder {
	return qb.append(attr, ">", qb.sprintf(value, quote))
}

func (qb *QueryBuilder) Less(attr string, value any) *QueryBuilder {
	return qb.append(attr, "<", qb.sprintf(value, quote))
}

func (qb *QueryBuilder) GtrEqualTo(attr string, value any) *QueryBuilder {
	return qb.append(attr, ">=", qb.sprintf(value, quote))
}

func (qb *QueryBuilder) LessEqualTo(attr string, value any) *QueryBuilder {
	return qb.append(attr, "<=", qb.sprintf(value, quote))
}

func (qb *QueryBuilder) Equal(attr string, value any) *QueryBuilder {
	return qb.append(attr, "=", qb.sprintf(value, quote))
}

func (qb *QueryBuilder) NotEqual(attr string, value any) *QueryBuilder {
	return qb.append(attr, "!=", qb.sprintf(value, quote))
}

func (qb *QueryBuilder) In(attr string, values []any) *QueryBuilder {

	var tupleStr string
	for i, v := range values {
		if i == 0 {
			tupleStr += "("
		}

		tupleStr += qb.sprintf(v, singleQuote)

		if i == len(values)-1 {
			tupleStr += ")"
		} else {
			tupleStr += ", "
		}
	}

	return qb.append(attr, "in", tupleStr)
}

func (qb *QueryBuilder) Encode() (JqlQuery, error) {
	return NewJqlQuery(strings.Join(qb.q, " "))
}

func NewQueryBuilder() *QueryBuilder {
	return &QueryBuilder{
		q: make([]string, 0),
	}
}
