package aql

import (
	"fmt"
	"strings"
)

type Query interface {
	Equal(attr, value string) Query
	NotEqual(attr, value string) Query
	In(attr string, values []string) Query
	Gtr(attr string, value string) Query
	Less(attr string, value string) Query
	GtrEqualTo(attr string, value string) Query
	LessEqualTo(attr string, value string) Query

	And() Query
	Or() Query

	Encode() string
}

type StringQuery struct {
	q []string
}

func (s *StringQuery) append(vals ...string) *StringQuery {
	s.q = append(s.q, vals...)
	return s
}

func (s *StringQuery) And() Query {
	return s.append("AND")
}

func (s *StringQuery) Or() Query {
	return s.append("OR")
}

func (s *StringQuery) Gtr(attr, value string) Query {
	return s.append(attr, ">", fmt.Sprintf(`"%s"`, value))
}

func (s *StringQuery) Less(attr, value string) Query {
	return s.append(attr, "<", fmt.Sprintf(`"%s"`, value))
}

func (s *StringQuery) GtrEqualTo(attr, value string) Query {
	return s.append(attr, ">=", fmt.Sprintf(`"%s"`, value))
}

func (s *StringQuery) LessEqualTo(attr, value string) Query {
	return s.append(attr, "<=", fmt.Sprintf(`"%s"`, value))
}

func (s *StringQuery) Equal(attr, value string) Query {
	return s.append(attr, "=", fmt.Sprintf(`"%s"`, value))
}

func (s *StringQuery) NotEqual(attr, value string) Query {
	return s.append(attr, "!=", fmt.Sprintf(`"%s"`, value))
}

func (s *StringQuery) In(attr string, values []string) Query {

	var tupleStr string
	for i, v := range values {
		if i == 0 {
			tupleStr += "("
		}

		tupleStr += fmt.Sprintf(`'%s'`, v)

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
