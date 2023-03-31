package rtdata

import (
	"fmt"
	"testing"
)

func TestLocalVars(t *testing.T) {
	frameLocalVars := newFrame(100, 100).localVars

	var tests = []struct {
		input uint
		want  interface{}
	}{
		{0, int32(100)},
		{1, int32(-100)},
		{2, int64(2997924580)},
		{4, int64(-2997924580)},
		{6, float32(3.1415926)},
		{7, float64(2.71828182845)},
		{9, nil},
	}

	for _, test := range tests {
		switch test.want.(type) {
		case int32:
			frameLocalVars.SetInt(test.input, test.want.(int32))
		case int64:
			frameLocalVars.SetLong(test.input, test.want.(int64))
		case float32:
			frameLocalVars.SetFloat(test.input, test.want.(float32))
		case float64:
			frameLocalVars.SetDouble(test.input, test.want.(float64))
		default:
			if test.want == nil {
				frameLocalVars.SetRef(test.input, nil)
			}
		}
	}

	for _, test := range tests {
		switch test.want.(type) {
		case int32:
			if got := frameLocalVars.GetInt(test.input); got != test.want {
				t.Errorf("wrong at get int")
			} else {
				fmt.Printf("put %v as Int get %v as Int\n", test.want, got)
			}
		case int64:
			if got := frameLocalVars.GetLong(test.input); got != test.want {
				t.Errorf("wrong at get long")
			} else {
				fmt.Printf("put %v as Long get %v as Long\n", test.want, got)
			}
		case float32:
			if got := frameLocalVars.GetFloat(test.input); got != test.want {
				t.Errorf("wrong at get float")
			} else {
				fmt.Printf("put %v as Float get %v as Float\n", test.want, got)
			}
		case float64:
			if got := frameLocalVars.GetDouble(test.input); got != test.want {
				t.Errorf("wrong at get double")
			} else {
				fmt.Printf("put %v as Double get %v as Double\n", test.want, got)
			}
		default:
			if frameLocalVars.GetRef(test.input) != nil {
				t.Errorf("wrong at get ref")
			}
		}
	}
}
