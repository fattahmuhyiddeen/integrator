package model

import (
	"fmt"
	"testing"
)

func TestGetDriver(t *testing.T) {
	type testObject struct {
		Input     string
		Expeceted string
	}
	testObjects := []testObject{
		testObject{"pgsql", "postgres"},
		testObject{"mssql", "sqlserver"},
		testObject{"postgres", "postgres"},
	}

	for index, item := range testObjects {
		if getDriver(item.Input) != item.Expeceted {
			t.Errorf(fmt.Sprintf("Fail at test case #: %d", index+1))
		}
	}
}
