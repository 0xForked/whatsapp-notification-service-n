package utils_test

import (
	"github.com/aasumitro/gowa/internal/utils"
	"testing"
)

func TestInArray(t *testing.T) {
	var array = []string{"a", "b", "c"}
	var value = "a"
	var result = utils.InArray(value, array)
	if !result {
		t.Errorf("Expected %s to be in array %v", value, array)
	}
}

func TestExplode(t *testing.T) {
	var tests = []struct {
		input  string
		delim  string
		output []string
	}{
		{"a:b:c", ":", []string{"a", "b", "c"}},
		{"a:b:c", ",", []string{"a:b:c"}},
		{"a:b:c", "x", []string{"a:b:c"}},
	}

	for _, test := range tests {
		result := utils.Explode(test.input, test.delim)
		if len(result) != len(test.output) {
			t.Errorf("Expected %v, got %v", test.output, result)
		}
		for i := range result {
			if result[i] != test.output[i] {
				t.Errorf("Expected %v, got %v", test.output, result)
			}
		}
	}
}
