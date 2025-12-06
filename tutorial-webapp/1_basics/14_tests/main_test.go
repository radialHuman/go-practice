package main

import "testing"

// need to create go mod init for test to work

// 1.manual way

func TestDivide(t *testing.T) { // t is the usual way
	// place all possible scenarios here and if else for success and failed situations
	_, err := divide(10, 2)
	if err != nil {
		t.Error("Got an error even when it was not supposed to")
	}

}

func TestZeorDivide(t *testing.T) { // t is the usual way
	// place all possible scenarios here and if else for success and failed situations
	_, err := divide(10, 0)
	if err == nil {
		t.Error("Doesn catch divide by zero error")
	}

}

// 2. automated table test
// cant write all the test fucntions 1/1
// using this we can declare what situations we want to test and it will do ti
var tests = []struct {
	testName string
	param1   float32
	param2   float32
	expected float32
	isError  bool
}{
	{"testNormal", 100, 10, 10, false},
	{"testZero", 100, 0, 0, true},
}

func TestDivision(t *testing.T) { // has to be capital for it to run
	for _, tt := range tests {
		result, err := divide(tt.param1, tt.param2)
		if tt.isError {
			if err == nil {
				t.Error("Expected error but got nothing")
			}
		} else {
			if err != nil {
				t.Error("Expected no error but got one")
			}
		}
		if result != tt.expected {
			t.Errorf("Expected %f but got %f", tt.expected, result)

		}

	}
}

// to run this : go test -v
// to test how much is of the main file is tested in % : go test -cover
// to see this in browser : go test -coverprofile=coverage.out && go tool cover -html=coverage.out
