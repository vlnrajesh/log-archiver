package utils

import (
	"testing"
)

//TestGreetHello performs Command line test arguments
func TestGreetHello(t *testing.T) {
	actualResult := GreetHello()
	expectedResult := "Hello"
	if actualResult != expectedResult {
		t.Fatalf("Expected %T Type but got %T", expectedResult, actualResult)
	}
}
