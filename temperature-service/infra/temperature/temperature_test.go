package temperature

import "testing"

func TestCelsiusToFahrenheit(t *testing.T) {
	tests := []struct {
		celsius    float64
		fahrenheit float64
	}{
		{9, 48.2},
		{16, 60.8},
		{30, 86},
		{39, 102.2},
	}

	for _, test := range tests {
		result := CelsiusToFahrenheit(test.celsius)
		if result != test.fahrenheit {
			t.Errorf("CelsiusToFahrenheit(%f) = %f, expected %f", test.celsius, result, test.fahrenheit)
		}
	}
}

func TestCelsiusToKelvin(t *testing.T) {
	tests := []struct {
		celsius float64
		kelvin  float64
	}{
		{9, 282},
		{16, 289},
		{30, 303},
		{39, 312},
	}

	for _, test := range tests {
		result := CelsiusToKelvin(test.celsius)
		if result != test.kelvin {
			t.Errorf("CelsiusToKelvin(%f) = %f, expected %f", test.celsius, result, test.kelvin)
		}
	}
}
