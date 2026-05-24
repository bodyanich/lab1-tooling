package internal

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{
			name:     "positive numbers",
			a:        2,
			b:        3,
			expected: 5,
		},
		{
			name:     "negative numbers",
			a:        -2,
			b:        -3,
			expected: -5,
		},
		{
			name:     "positive and negative number",
			a:        10,
			b:        -4,
			expected: 6,
		},
		{
			name:     "zeros",
			a:        0,
			b:        0,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name        string
		a           int
		b           int
		expected    int
		expectError bool
	}{
		{
			name:        "positive numbers",
			a:           10,
			b:           2,
			expected:    5,
			expectError: false,
		},
		{
			name:        "negative number",
			a:           -10,
			b:           2,
			expected:    -5,
			expectError: false,
		},
		{
			name:        "division by zero",
			a:           10,
			b:           0,
			expected:    0,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Divide(tt.a, tt.b)

			if tt.expectError && err == nil {
				t.Fatal("expected error, got nil")
			}

			if !tt.expectError && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("Divide(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestIsEven(t *testing.T) {
	tests := []struct {
		name     string
		number   int
		expected bool
	}{
		{
			name:     "even positive number",
			number:   4,
			expected: true,
		},
		{
			name:     "odd positive number",
			number:   5,
			expected: false,
		},
		{
			name:     "zero",
			number:   0,
			expected: true,
		},
		{
			name:     "negative even number",
			number:   -6,
			expected: true,
		},
		{
			name:     "negative odd number",
			number:   -7,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsEven(tt.number)
			if result != tt.expected {
				t.Errorf("IsEven(%d) = %v; want %v", tt.number, result, tt.expected)
			}
		})
	}
}
