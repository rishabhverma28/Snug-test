package main

import (
	"testing"
)

func TestCheckLeapYear(t *testing.T) {
	var tests = []struct {
		input    int
		expected bool
	}{
		{2016, true},
		{2012, true},
		{2009, false},
		{2010, false},
	}

	for _, test := range tests {
		if output := checkLeapYear(test.input); output != test.expected {
			t.Error("Test Failed. Input: {}, Expected:{}, Received: {}", test.input, test.expected, output)
		}
	}
}

func TestValidateStringLength(t *testing.T) {
	var tests = []struct {
		input    []string
		expected bool
	}{
		{[]string{"05", "05", "2016"}, true},
		{[]string{"01", "02", "2012"}, true},
		{[]string{"1", "2", "2009"}, false},
		{[]string{"04", "06", "10"}, false},
	}

	for _, test := range tests {
		if output := validateStringLength(test.input); output != test.expected {
			t.Error("Test Failed. Input: {}, Expected:{}, Received: {}", test.input, test.expected, output)
		}
	}
}

func TestConvertToInteger(t *testing.T) {
	var tests = []struct {
		input    []string
		expected [3]int
	}{
		{[]string{"05", "05", "2016"}, [3]int{5, 5, 2016}},
		{[]string{"01", "02", "2012"}, [3]int{1, 2, 2012}},
		{[]string{"12", "12", "2009"}, [3]int{12, 12, 2009}},
		{[]string{"04", "06", "2010"}, [3]int{4, 6, 2010}},
	}

	for _, test := range tests {
		if output := convertToInteger(test.input); output != test.expected {
			t.Error("Test Failed. Input: {}, Expected:{}, Received: {}", test.input, test.expected, output)
		}
	}
}

func TestCheckValidMonthYear(t *testing.T) {
	var tests = []struct {
		input    dateType
		expected bool
	}{
		{dateType{5, 5, 2016, true, 0}, true},
		{dateType{1, 2, 2012, true, 0}, true},
		{dateType{45, 56, 2016, false, 0}, false},
		{dateType{5, 5, 201612, true, 0}, false},
		{dateType{5, 5, 1888, true, 0}, false},
		{dateType{5, -5, 2016, true, 0}, false},
		{dateType{225, -5, 2016, true, 0}, false},
	}
	for _, test := range tests {
		if output := (*dateType).checkValidMonthYear(&test.input); output != test.expected {
			t.Error("Test Failed. Input {}, Expected: {}, Received: {}", test.input, test.expected, output)
		}
	}
}

func TestCheckValidDay(t *testing.T) {
	var tests = []struct {
		input    dateType
		expected bool
	}{
		{dateType{5, 5, 2016, true, 0}, true},
		{dateType{1, 2, 2012, true, 0}, true},
		{dateType{45, 5, 2016, false, 0}, false},
		{dateType{65, 5, 2004, true, 0}, false},
	}
	for _, test := range tests {
		if output := (*dateType).checkValidDay(&test.input); output != test.expected {
			t.Error("Test Failed. Input {}, Expected: {}, Received: {}", test.input, test.expected, output)
		}
	}
}

func TestDaysSinceOrigin(t *testing.T) {
	var tests = []struct {
		input    dateType
		expected int
	}{
		{dateType{2, 6, 1983, false, 724062}, 724062},
		{dateType{21, 6, 1983, false, 724081}, 724081},
		{dateType{21, 6, 2016, true, 736135}, 736135},
		{dateType{1, 1, 1, false, 0}, 0},
	}
	for _, test := range tests {
		if output := (*dateType).daysSinceOrigin(&test.input); output != test.expected {
			t.Error("Test Failed. Input {}, Expected: {}, Received: {}", test.input, test.expected, output)
		}
	}
}

func TestMinMax(t *testing.T) {
	got1, got2 := minMax(20, 10)
	want1, want2 := 20, 10
	if got1 != want1 || got2 != want2 {
		t.Error("Test Failed")
	}

	got1test2, got2test2 := minMax(10, 20)
	want1test2, want2test2 := 20, 10
	if got1test2 != want1test2 || got2test2 != want2test2 {
		t.Error("Test Failed")
	}

}
