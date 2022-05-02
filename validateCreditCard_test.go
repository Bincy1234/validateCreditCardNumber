package main

import "testing"

/* This test case validates the following
1)credit card number should start with 4,5 or 6.
2)credit card number must contain 16 digits.
*/
func TestValidateCreditCard_postive1(t *testing.T) {

	testCardNumber := "4123456789123456"
	actual := validateNumber(testCardNumber)
	expected := true
	if actual != expected {
		t.Errorf("Expected Valid")

	}
}

/* This test case validates that the credit card number can have digits in
   groups of 4, separated by one hyphen "-".
*/
func TestValidateCreditCard_positive2(t *testing.T) {

	testCardNumber := "5122-2368-7954-3214"
	actual := validateNumber(testCardNumber)
	expected := true
	if actual != expected {
		t.Errorf("Expected Valid")

	}
}

/* This test case verifies that groups of 5 cannot be used and also two sepators
   (-) cannot be used `
*/
func TestValidateCreditCard_negative1(t *testing.T) {

	testCardNumber := "61234--8912-3456"
	actual := validateNumber(testCardNumber)
	expected := false
	if actual != expected {
		t.Errorf("Expected Valid")

	}
}

//This test case verifies that groups of 2 cannot be used.
func TestValidateCreditCard_negative2(t *testing.T) {

	testCardNumber := "51-67-8912-3456"
	actual := validateNumber(testCardNumber)
	expected := false
	if actual != expected {
		t.Errorf("Expected Valid")

	}
}

// This test case verifies that space ' ' and (-) cannot be used as separators.
func TestValidateCreditCard_negative3(t *testing.T) {

	testCardNumber := "5123 - 4567 - 8912 - 3456"
	actual := validateNumber(testCardNumber)
	expected := false
	if actual != expected {
		t.Errorf("Expected Valid")

	}
}

// This test verifies that credit card number must contain exactly 16 digits.
func TestValidateCreditCard_negative4(t *testing.T) {

	testCardNumber := "4567"
	actual := validateNumber(testCardNumber)
	expected := false
	if actual != expected {
		t.Errorf("Expected Valid")

	}
}

/* This test case validate that the first line of input contains an integer N and
   it follows the contraint 0 < N > 100.
*/

func TestValidInput(t *testing.T) {
	firstInput := 2
	actual := validInput(firstInput)
	expected := true

	if actual != expected {
		t.Errorf("Input condition not met")

	}
}

// This test case validates that first input 0 is not valid.
func TestValidInput_edge1(t *testing.T) {
	firstInput := 0
	actual := validInput(firstInput)
	expected := false

	if actual != expected {
		t.Errorf("Input condition not met")

	}
}

// This test case validates that first input 100 is not valid.
func TestValidInput_edge2(t *testing.T) {
	firstInput := 100
	actual := validInput(firstInput)
	expected := false

	if actual != expected {
		t.Errorf("Input condition not met")

	}
}
