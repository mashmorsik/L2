package main

import (
	"os"
	"testing"
)

func TestDefaultSort(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		expected string
	}{
		{name: "test_default_sort_numbers", filePath: "nums.txt", expected: "1 2 4 5"},
		{name: "test_default_sort_words", filePath: "words.txt", expected: "apple banana melon watermelon"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpFile, err := os.CreateTemp("", "test_default.txt")
			if err != nil {
				t.Fatalf("Failed to create temporary file: %v", err)
			}
			defer tmpFile.Close()
			defer os.Remove(tmpFile.Name())

			err = DefaultSort(tmpFile, tt.filePath)

			content, err := os.ReadFile(tmpFile.Name())
			if err != nil {
				t.Fatalf("Failed to read temporary file: %v", err)
			}

			if string(content) != tt.expected {
				t.Errorf("Test failed - results do not match.\nGot:\n%v\nExpected:\n%v", string(content), tt.expected)
			}

		})
	}
}

func TestReverseSort(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		expected string
	}{
		{name: "test_reverse_sort_numbers", filePath: "nums.txt", expected: "5 4 2 1"},
		{name: "test_reverse_sort_words", filePath: "words.txt", expected: "watermelon melon banana apple"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpFile, err := os.CreateTemp("", "test_default.txt")
			if err != nil {
				t.Fatalf("Failed to create temporary file: %v", err)
			}
			defer tmpFile.Close()
			defer os.Remove(tmpFile.Name())

			err = ReverseSort(tmpFile, tt.filePath)

			content, err := os.ReadFile(tmpFile.Name())
			if err != nil {
				t.Fatalf("Failed to read temporary file: %v", err)
			}

			if string(content) != tt.expected {
				t.Errorf("Test failed - results do not match.\nGot:\n%v\nExpected:\n%v", string(content), tt.expected)
			}

		})
	}
}

func TestColumnSort(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		column   int
		expected string
	}{
		{name: "test_columns_sort_column_3",
			filePath: "columns.txt",
			column:   3,
			expected: "5 7 0 4 3 1 1 2 3"},
		{name: "test_columns_sort_column_2",
			filePath: "columns.txt",
			column:   2,
			expected: "1 2 3 4 3 1 5 7 0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpFile, err := os.CreateTemp("", "test_default.txt")
			if err != nil {
				t.Fatalf("Failed to create temporary file: %v", err)
			}
			defer tmpFile.Close()
			defer os.Remove(tmpFile.Name())

			err = ColumnsSort(tmpFile, tt.filePath, tt.column)

			content, err := os.ReadFile(tmpFile.Name())
			if err != nil {
				t.Fatalf("Failed to read temporary file: %v", err)
			}

			if string(content) != tt.expected {
				t.Errorf("Test failed - results do not match.\nGot:\n%v\nExpected:\n%v", string(content), tt.expected)
			}

		})
	}
}

func TestNumericSort(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		expected string
		expError string
	}{
		{name: "test_numeric_sort_numbers", filePath: "nums.txt", expected: "1 2 4 5"},
		{name: "test_numeric_sort_words", filePath: "words.txt", expected: "", expError: "file needs to contain integers"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpFile, err := os.CreateTemp("", "test_default.txt")
			if err != nil {
				t.Fatalf("Failed to create temporary file: %v", err)
			}
			defer tmpFile.Close()
			defer os.Remove(tmpFile.Name())

			err = NumericSort(tmpFile, tt.filePath)
			if err != nil && err.Error() != tt.expError {
				t.Errorf("Test failed - results do not match.\nGot:\n%v\nExpected:\n%v", err.Error(), tt.expError)
			}

			content, err := os.ReadFile(tmpFile.Name())
			if err != nil {
				t.Fatalf("Failed to read temporary file: %v", err)
			}

			if string(content) != tt.expected {
				t.Errorf("Test failed - results do not match.\nGot:\n%v\nExpected:\n%v", string(content), tt.expected)
			}

		})
	}
}

func TestUniqueSort(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		expected string
	}{
		{name: "test_unique_sort_numbers", filePath: "doubles.txt", expected: "3 5 1 7 2 4 "},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpFile, err := os.CreateTemp("", "test_default.txt")
			if err != nil {
				t.Fatalf("Failed to create temporary file: %v", err)
			}
			defer tmpFile.Close()
			defer os.Remove(tmpFile.Name())

			err = UniqueSort(tmpFile, tt.filePath)

			content, err := os.ReadFile(tmpFile.Name())
			if err != nil {
				t.Fatalf("Failed to read temporary file: %v", err)
			}

			if string(content) != tt.expected {
				t.Errorf("Test failed - results do not match.\nGot:\n%v\nExpected:\n%v", string(content), tt.expected)
			}

		})
	}
}
