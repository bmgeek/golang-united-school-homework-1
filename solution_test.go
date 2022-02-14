package solution

import (
	"testing"
)

func TestGetMessage(t *testing.T) {
	var s interface{} = GetMessage()
	if _, ok := s.(string); !ok {
		t.Fatal("Variable is not a String")
	}
	t.Log("Test is submit")
}
