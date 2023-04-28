package aql

import "testing"

func TestSyntaxError(t *testing.T) {
	s := "asdfdasf  asdfsdfsda  sdfasdfads"
	q, err := NewJqlQuery(s)
	if err == nil {
		t.Error("A query with syntax errors was accepted")
	}
	if q != "" {
		t.Error("A query was returned containing a syntax error")
	}
}

func TestValidSyntax(t *testing.T) {
	s := "project = 'XXXXX'"
	q, err := NewJqlQuery(s)
	if err != nil {
		t.Error("A query without syntax errors was rejected")
	}
	if q == "" {
		t.Error("A valid query was not returned")
	}
}
