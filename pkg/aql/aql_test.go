package aql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringQuery(t *testing.T) {

	for _, test := range []struct {
		name     string
		query    Query
		expected string
	}{
		{
			"should generate a basic statement",
			NewStringQuery().Equal("attribute", "value"),
			`attribute = "value"`,
		},
		{
			"should generate a statement with a int",
			NewStringQuery().Equal("number", 135),
			`number = 135`,
		},
		{
			"should group query together",
			NewStringQuery().Group(NewStringQuery().Equal("value", "one").Or().Equal("value", "three")),
			`( value = "one" OR value = "three" )`,
		},
		{
			"should support more complex queries",
			NewStringQuery().Equal("attribute", "value").
				And().GtrEqualTo("len", "0").
				And().LessEqualTo("len", "0").
				And().In("thinger", []any{"one", "two"}),
			`attribute = "value" AND len >= "0" AND len <= "0" AND thinger in ('one', 'two')`,
		},
		{
			"should support more or queries",
			NewStringQuery().Equal("attribute", "value").
				And().GtrEqualTo("len", "0").
				Or().LessEqualTo("len", "0"),
			`attribute = "value" AND len >= "0" OR len <= "0"`,
		},
		{
			"should support literal strings like null",
			NewStringQuery().Equal("attribute", "value").
				And().NotEqual("recommended", Literal("null")),
			`attribute = "value" AND recommended != null`,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.query.Encode())
		})
	}

}
