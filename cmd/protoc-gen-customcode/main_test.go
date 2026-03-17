package main

import (
	"testing"

	"google.golang.org/protobuf/compiler/protogen"
)

func TestExtractStartValue(t *testing.T) {
	tests := []struct {
		name     string
		comments protogen.Comments
		expected int
	}{
		{
			name:     "valid start with space",
			comments: "//start 101\n",
			expected: 101,
		},
		{
			name:     "valid start with colon",
			comments: "//start:202\n",
			expected: 202,
		},
		{
			name:     "valid start with equals",
			comments: "//start=303\n",
			expected: 303,
		},
		{
			name:     "valid with comment",
			comments: "//start 404 - some comment\n",
			expected: 404,
		},
		{
			name:     "invalid format",
			comments: "// other comment\n",
			expected: 0,
		},
		{
			name:     "empty comments",
			comments: "",
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := extractStartValue(tt.comments)
			if result != tt.expected {
				t.Errorf("extractStartValue() = %d, want %d", result, tt.expected)
			}
		})
	}
}

func TestToCamelCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"basic_success", "BasicSuccess"},
		{"conf_not_exist", "ConfNotExist"},
		{"param_invalid", "ParamInvalid"},
		{"team_member_info_empty", "TeamMemberInfoEmpty"},
		{"gem", "Gem"},
		{"gem_success", "GemSuccess"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := toCamelCase(tt.input)
			if result != tt.expected {
				t.Errorf("toCamelCase(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}
