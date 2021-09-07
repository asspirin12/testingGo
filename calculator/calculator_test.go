package calculator_test

import (
	"testing"
	"testingGo/calculator"
)

type TestCase struct {
	value    int
	expected bool
	actual   bool
}

func TestCalculateArmstrong(t *testing.T) {
	testCase := TestCase{
		value:    371,
		expected: true,
	}

	testCase.actual = calculator.CalculateArmstrong(testCase.value)
	if testCase.actual != testCase.expected {
		t.Fail()
	}
}

func TestNegativeCalculateArmstrong(t *testing.T) {
	testCase := TestCase{
		value:    3371,
		expected: false,
	}

	testCase.actual = calculator.CalculateArmstrong(testCase.value)
	if testCase.actual != testCase.expected {
		t.Fail()
	}
}
