package main

import (
	"fmt"
	"testing"
)

func TestContains(t *testing.T) {
	type testObject struct {
		Data      []string
		Target    string
		Expeceted bool
	}
	testObjects := []testObject{
		testObject{[]string{"a", "aa", "b"}, "b", true},
		testObject{[]string{"a", "aa", "b"}, "a", true},
		testObject{[]string{"a", "aa", "b"}, "bbb", false},
	}

	for index, item := range testObjects {
		if contains(item.Data, item.Target) != item.Expeceted {
			t.Errorf(fmt.Sprintf("Fail at test case #: %d", index+1))
		}
	}
}
