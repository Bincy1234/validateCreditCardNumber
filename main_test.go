package main

import "testing"

func TestValidateCreditCard(t *testing.T) {

	testCardNumber := "4123456789123456"
	actual := validateCreditCard(testCardNumber)
	expected := true
	if actual != expected {
		t.Errorf("Expected Valid")

	}
}

func TestValidInput(t *testing.T) {
	firstInput := 101
	actual := validInput(firstInput)
	expected := false

	if actual != expected {
		t.Errorf("Input condition not met")

	}
}
