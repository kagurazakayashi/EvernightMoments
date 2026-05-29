package main

import (
	"testing"
)

func TestMatchPattern_BasenameOnly(t *testing.T) {
	patterns := []string{"*.cos", "*.dop", "*.xmp"}
	testCases := []struct {
		path     string
		expected bool
	}{
		{"D:\\Photos\\photo.cos", true},
		{"D:\\Photos\\photo.dop", true},
		{"D:\\Photos\\photo.jpg", false},
		{"photo.cos", true},
	}
	for _, tc := range testCases {
		if got := matchesAnyPattern(tc.path, patterns); got != tc.expected {
			t.Errorf("matchesAnyPattern(%q, %v) = %v, want %v", tc.path, patterns, got, tc.expected)
		}
	}
}

func TestMatchPattern_PathPattern_Relative(t *testing.T) {
	patterns := []string{"CaptureOne/Settings153/*.cos", "SubDir/*.dop"}
	testCases := []struct {
		path     string
		expected bool
	}{
		{"D:\\Photos\\Album\\CaptureOne\\Settings153\\photo.cos", true},
		{"D:\\Photos\\Album\\CaptureOne\\Settings153\\readme.txt", false},
		{"D:\\Photos\\SubDir\\file.dop", true},
		{"D:\\Photos\\SubDir\\Other\\file.dop", false},
	}
	for _, tc := range testCases {
		if got := matchesAnyPattern(tc.path, patterns); got != tc.expected {
			t.Errorf("matchesAnyPattern(%q, %v) = %v, want %v", tc.path, patterns, got, tc.expected)
		}
	}
}

func TestMatchPattern_PathPattern_Absolute(t *testing.T) {
	patterns := []string{"D:\\Photos\\*.arw"}
	testCases := []struct {
		path     string
		expected bool
	}{
		{"D:\\Photos\\photo.arw", true},
		{"D:\\Photos\\SubDir\\photo.arw", false},
		{"C:\\Other\\photo.arw", false},
	}
	for _, tc := range testCases {
		if got := matchesAnyPattern(tc.path, patterns); got != tc.expected {
			t.Errorf("matchesAnyPattern(%q, %v) = %v, want %v", tc.path, patterns, got, tc.expected)
		}
	}
}

func TestMatchPattern_MixedPatterns(t *testing.T) {
	patterns := []string{"*.dop", "CaptureOne\\Settings153\\*.cos"}
	testCases := []struct {
		path     string
		expected bool
	}{
		{"D:\\Photos\\photo.dop", true},
		{"D:\\Photos\\Album\\CaptureOne\\Settings153\\photo.cos", true},
		{"D:\\Photos\\photo.jpg", false},
	}
	for _, tc := range testCases {
		if got := matchesAnyPattern(tc.path, patterns); got != tc.expected {
			t.Errorf("matchesAnyPattern(%q, %v) = %v, want %v", tc.path, patterns, got, tc.expected)
		}
	}
}

func TestMatchPattern_WindowsBackslash(t *testing.T) {
	patterns := []string{`CaptureOne\Settings153\*.cos`}
	testCases := []struct {
		path     string
		expected bool
	}{
		{`D:\Photos\Album\CaptureOne\Settings153\photo.cos`, true},
		{`D:\Photos\Album\CaptureOne\Settings153\readme.txt`, false},
	}
	for _, tc := range testCases {
		if got := matchesAnyPattern(tc.path, patterns); got != tc.expected {
			t.Errorf("matchesAnyPattern(%q, %v) = %v, want %v", tc.path, patterns, got, tc.expected)
		}
	}
}

func TestMatchPattern_EmptyPatterns(t *testing.T) {
	if matchesAnyPattern("D:\\Photos\\photo.jpg", nil) {
		t.Error("matchesAnyPattern with nil patterns should return false")
	}
	if matchesAnyPattern("D:\\Photos\\photo.jpg", []string{}) {
		t.Error("matchesAnyPattern with empty patterns should return false")
	}
}

func TestParsePatterns(t *testing.T) {
	testCases := []struct {
		input    string
		expected []string
	}{
		{"", nil},
		{"*.cos", []string{"*.cos"}},
		{"*.cos, *.dop", []string{"*.cos", "*.dop"}},
		{"  *.cos ,  *.dop  ", []string{"*.cos", "*.dop"}},
		{"CaptureOne\\Settings153\\*.cos, *.dop", []string{"CaptureOne\\Settings153\\*.cos", "*.dop"}},
	}
	for _, tc := range testCases {
		got := parsePatterns(tc.input)
		if len(got) != len(tc.expected) {
			t.Errorf("parsePatterns(%q) len = %d, want %d", tc.input, len(got), len(tc.expected))
			continue
		}
		for i := range got {
			if got[i] != tc.expected[i] {
				t.Errorf("parsePatterns(%q)[%d] = %q, want %q", tc.input, i, got[i], tc.expected[i])
			}
		}
	}
}
