package solution

import (
	"testing"
)

func TestGetMessage(t *testing.T) {
	var s interface{} = GetMessage()
	if _, ok := s.(string); !ok {
		t.Error("Variable is not a String")
	}
}
