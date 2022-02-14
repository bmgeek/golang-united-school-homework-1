package solution

import (
	"fmt"
	"testing"
)

func TestGetMessage(t *testing.T) {
	var s interface{} = GetMessage()
	if _, ok := s.(string); ok {
		fmt.Println("Variable is a string")
	} else {
		fmt.Println("Variable is not a string")
	}
}
