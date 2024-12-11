package mathutils

import "testing"

func TestFindMin(t *testing.T) {
	tests := []struct {
		name          string
		a, b, c, want float64
	}{
		{"all positive", 10.5, 2.3, 4.7, 2.3},
		{"with negative", -1.0, 2.0, 3.0, -1.0},
		{"all negative", -1.0, -2.0, -3.0, -3.0},
		{"equal numbers", 1.0, 1.0, 1.0, 1.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindMin(tt.a, tt.b, tt.c)
			if got != tt.want {
				t.Errorf("FindMin(%v, %v, %v) = %v, want %v",
					tt.a, tt.b, tt.c, got, tt.want)
			}
		})
	}
}

func TestCalculateAverage(t *testing.T) {
	tests := []struct {
		name          string
		a, b, c, want float64
	}{
		{"positive numbers", 1.0, 2.0, 3.0, 2.0},
		{"with negative", -1.0, 0.0, 1.0, 0.0},
		{"all same", 2.0, 2.0, 2.0, 2.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateAverage(tt.a, tt.b, tt.c)
			if got != tt.want {
				t.Errorf("CalculateAverage(%v, %v, %v) = %v, want %v",
					tt.a, tt.b, tt.c, got, tt.want)
			}
		})
	}
}

func TestSolveLinearEquation(t *testing.T) {
	tests := []struct {
		name       string
		a, b, want float64
	}{
		{"positive coefficients", 2.0, -4.0, 2.0},
		{"negative coefficients", -2.0, -4.0, -2.0},
		{"fractional result", 2.0, -1.0, 0.5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SolveLinearEquation(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("SolveLinearEquation(%v, %v) = %v, want %v",
					tt.a, tt.b, got, tt.want)
			}
		})
	}
}
