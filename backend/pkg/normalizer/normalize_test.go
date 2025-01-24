package normalizer_test

import (
	"backend/pkg/normalizer"
	"testing"
)

func TestNormalizer(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{" wladimir putin", "wladimir_putin"},
		{"rok1stock$", "rok1stock"},
		{" to   test    1", "to_test_1"},
		{"to_test_203", "to_test_203"},
		{"  hello    world!  ", "hello_world"},
		{"special@chars#remove!", "special_chars_remove"},
		{"multiple__underscores__here", "multiple__underscores__here"},
		{"", ""},
		{"1234 !@#$%^&*()", "1234"},
		{"hello   world   123", "hello_world_123"},
		{"sp@cial!Chars", "sp_cial_Chars"},
		{"   clean    up", "clean_up"},
		{"testing-with-dashes", "testing_with_dashes"},
		{"underscores___should___stay", "underscores___should___stay"},
		{"spaces    between__words", "spaces_between__words"},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			res := normalizer.Normalizer([]string{test.input})[0]
			if res != test.expected {
				t.Errorf("For input %q, expected %q, but got %q", test.input, test.expected, res)
			}
		})
	}
}
