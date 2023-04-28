package aql

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
