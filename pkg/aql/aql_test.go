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
			"should support more complex queries",
			NewStringQuery().Equal("attribute", "value").
				And().GtrEqualTo("len", "0").
				And().LessEqualTo("len", "0").
				And().In("thinger", []string{"one", "two"}),
			`attribute = "value" AND len >= "0" AND len <= "0" AND thinger in ('one', 'two')`,
		},
		{
			"should support more or queries",
			NewStringQuery().Equal("attribute", "value").
				And().GtrEqualTo("len", "0").
				Or().LessEqualTo("len", "0"),
			`attribute = "value" AND len >= "0" OR len <= "0"`,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.query.Encode())
		})
	}

}
