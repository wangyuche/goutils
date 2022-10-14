package random

import (
	"strings"
	"testing"
)

func TestRandString(t *testing.T) {
	testCases := []struct {
		mask           RandStringMode
		expectedString string
	}{
		{
			mask:           EngCapMode,
			expectedString: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		},
		{
			mask:           EngLowMode,
			expectedString: "abcdefghijklmnopqrstuvwxyz",
		},
		{
			mask:           SpecialsMode,
			expectedString: "~=+%^*/()[]{}/!@#$?|",
		},
		{
			mask:           DigitsMode,
			expectedString: "0123456789",
		},
		{
			mask:           DigitsMode | EngCapMode | SpecialsMode,
			expectedString: "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
		},
		{
			mask:           AllMode,
			expectedString: "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789~=+%^*/()[]{}/!@#$?|",
		},
	}
	for _, tc := range testCases {
		r := RandString(5, tc.mask)
		if strings.ContainsAny(tc.expectedString, r) {
			t.Logf("RandString Result: %s", r)
		} else {
			t.Errorf("RandString Error Result: %s", r)
		}
	}
}

func TestRandInt(t *testing.T) {
	testCases := []struct {
		max      int
		expected int
	}{
		{
			max: 50,
		},
		{
			max: 5,
		},
		{
			max: 150,
		},
	}
	for _, tc := range testCases {
		r := RandInt(5, tc.max)
		for _, i := range r {
			if i >= tc.max {
				t.Errorf("RandString Error Result: %d", r)
			}
		}
	}
}
