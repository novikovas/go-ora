package converters

import (
	"math"
	"reflect"
	"testing"
)

// Some documentation:
//	https://gotodba.com/2015/03/24/how-are-numbers-saved-in-oracle/

func TestDecodeDouble(t *testing.T) {

	for _, tt := range testFloatValue {
		t.Run(tt.SelectText, func(t *testing.T) {
			got := DecodeDouble(tt.Binary)
			e := math.Abs((got - tt.Float) / tt.Float)
			if e > 1e-15 {
				t.Errorf("DecodeDouble() = %v, want %v, Error= %e", got, tt.Float, e)
			}
		})
	}
}

func TestDecodeInt(t *testing.T) {
	for _, tt := range testFloatValue {
		// Test only with interger values
		if tt.IsInteger {
			t.Run(tt.SelectText, func(t *testing.T) {
				got := DecodeInt(tt.Binary)
				if got != tt.Integer {
					t.Errorf("DecodeInt() = %v, want %v", got, tt.Integer)
				}
			})
		}
	}
}

func TestEncodeInt64(t *testing.T) {
	for _, tt := range testFloatValue {
		// Test only with interger values
		if tt.IsInteger {
			t.Run(tt.SelectText, func(t *testing.T) {
				got := EncodeInt64(tt.Integer)

				n2 := DecodeInt(got)
				if n2 != tt.Integer {
					t.Errorf("DecodeInt(EncodeInt64(%d)) = %v", tt.Integer, n2)
				}

				if !reflect.DeepEqual(got, tt.Binary) {
					t.Errorf("EncodeInt64() = %v, want %v", got, tt.Binary)
				}
			})
		}
	}
}

func TestEncodeInt(t *testing.T) {
	for _, tt := range testFloatValue {
		// Test only with interger values
		if tt.IsInteger && tt.Float >= math.MinInt64 && tt.Float <= math.MaxInt64 {
			t.Run(tt.SelectText, func(t *testing.T) {
				i := int(tt.Integer)
				got := EncodeInt(i)

				n2 := int(DecodeInt(got))
				if n2 != i {
					t.Errorf("DecodeInt(EncodeInt(%d)) = %v", i, n2)
				}

				if !reflect.DeepEqual(got, tt.Binary) {
					t.Errorf("EncodeInt() = %v, want %v", got, tt.Binary)
				}
			})
		}
	}
}

func TestEncodeDouble(t *testing.T) {

	for _, tt := range testFloatValue {
		t.Run(tt.SelectText, func(t *testing.T) {
			got, err := EncodeDouble(tt.Float)
			if err != nil {
				t.Errorf("Unexpected error: %s", err)
				return
			}

			f := DecodeDouble(got)

			if tt.Float != 0.0 {
				e := math.Abs((f - tt.Float) / tt.Float)
				if e > 1e-15 {
					t.Errorf("DecodeDouble(EncodeDouble(%g)) = %g,  Error= %e", tt.Float, f, e)
				}
			}

			if len(tt.Binary) < 10 {
				if !reflect.DeepEqual(tt.Binary, got) {
					t.Errorf("EncodeDouble(%g) = %v want %v", tt.Float, got, tt.Binary)
				}
			}
		})
	}
}
