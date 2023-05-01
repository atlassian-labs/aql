package aql

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestQueryBuilder(t *testing.T) {

	for _, test := range []struct {
		name     string
		query    *QueryBuilder
		expected string
	}{
		{
			"should generate a basic statement",
			NewQueryBuilder().Equal("attribute", "value"),
			`attribute = "value"`,
		},
		{
			"should generate a statement with a int",
			NewQueryBuilder().Equal("number", 135),
			`number = 135`,
		},
		{
			"should group query together",
			NewQueryBuilder().Group(NewQueryBuilder().Equal("value", "one").Or().Equal("value", "three")),
			`( value = "one" OR value = "three" )`,
		},
		{
			"should support more complex queries",
			NewQueryBuilder().Equal("attribute", "value").
				And().GtrEqualTo("len", "0").
				And().LessEqualTo("len", "0").
				And().In("thinger", []any{"one", "two"}).
				And().Contains("comment", "mchandler"),

			`attribute = "value" AND len >= "0" AND len <= "0" AND thinger in ('one', 'two') AND comment ~ "mchandler"`,
		},
		{
			"should support more or queries",
			NewQueryBuilder().Equal("attribute", "value").
				And().GtrEqualTo("len", "0").
				Or().LessEqualTo("len", "0"),
			`attribute = "value" AND len >= "0" OR len <= "0"`,
		},
		{
			"should support literal strings like null",
			NewQueryBuilder().Equal("attribute", "value").
				And().NotEqual("recommended", Literal("null")),
			`attribute = "value" AND recommended != null`,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.query.Encode())
		})
	}

}

func TestParse(t *testing.T) {
	for _, test := range []struct {
		in       string
		err      bool
		rawQuery []string
	}{
		{
			`something = null`,
			false,
			[]string{"something", "=", `null`},
		},
		{
			`something = "contains thing"`,
			false,
			[]string{"something", "=", `"contains thing"`},
		},
		{
			`something != "contains thing"`,
			false,
			[]string{"something", "!=", `"contains thing"`},
		},
		{
			`something ~ "contains thing"`,
			false,
			[]string{"something", "~", `"contains thing"`},
		},
		{
			`something > "2011-01-01"`,
			false,
			[]string{"something", ">", `"2011-01-01"`},
		},
		{
			`something >= "2011-01-01"`,
			false,
			[]string{"something", ">=", `"2011-01-01"`},
		},
		{
			`something in ("one", "two")`,
			false,
			[]string{"something", "in", `('one', 'two')`},
		},
		{
			`something in ("one", "two", "three")`,
			false,
			[]string{"something", "in", `('one', 'two', 'three')`},
		},
		{
			`something < "2011-01-01"`,
			false,
			[]string{"something", "<", `"2011-01-01"`},
		},
		{
			`something <= "2011-01-01"`,
			false,
			[]string{"something", "<=", `"2011-01-01"`},
		},
		{
			`small query`,
			true,
			nil,
		},
		{
			`something !! failure`,
			true,
			nil,
		},
	} {
		t.Run(test.in, func(t *testing.T) {
			q, err := Parse(test.in)
			if test.err {
				assert.Error(t, err)
				assert.Nil(t, q)
				return
			}

			require.NoError(t, err)
			assert.NotNil(t, q)
			assert.Equal(t, test.rawQuery, q.q)
		})
	}
}
