package dynamo

import (
	"fmt"
	"testing"
)

func TestSubExpr(t *testing.T) {
	s := subber{}

	subbed, err := s.subExpr("$ > ? AND begins_with (Title, ?)", []interface{}{"Count", "1", "foo"})
	if err != nil {
		t.Error(err)
	}

	const format = "%s > :v0 AND begins_with (Title, :v1)"
	// you should be able to sub the same name twice and get the same result
	expect := fmt.Sprintf(format, s.subName("Count"))
	if subbed != expect {
		t.Errorf("bad subbed expr: %v ≠ %v", subbed, expect)
	}
}
