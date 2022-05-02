package main

import "testing"

func TestValidateCreditCard_postive1(t *testing.T) {

	testCardNumber := "4123456789123456"
	actual := validateNumber(testCardNumber)
	expected := true
	if actual != expected {
		t.Errorf("Expected Valid")

	}
}
func TestValidateCreditCard_positive2(t *testing.T) {

	testCardNumber := "5122-2368-7954-3214"
	actual := validateNumber(testCardNumber)
	expected := true
	if actual != expected {
		t.Errorf("Expected Valid")

	}
}

func TestValidateCreditCard_negative1(t *testing.T) {

	testCardNumber := "61234--8912-3456"
	actual := validateNumber(testCardNumber)
	expected := false
	if actual != expected {
		t.Errorf("Expected Valid")

	}
}

func TestValidateCreditCard_negative2(t *testing.T) {

	testCardNumber := "51-67-8912-3456"
	actual := validateNumber(testCardNumber)
	expected := false
	if actual != expected {
		t.Errorf("Expected Valid")

	}
}

func TestValidateCreditCard_negative3(t *testing.T) {

	testCardNumber := "5123 - 4567 - 8912 - 3456"
	actual := validateNumber(testCardNumber)
	expected := false
	if actual != expected {
		t.Errorf("Expected Valid")

	}
}

func TestValidateCreditCard_negative4(t *testing.T) {

	testCardNumber := "4567"
	actual := validateNumber(testCardNumber)
	expected := false
	if actual != expected {
		t.Errorf("Expected Valid")

	}
}

func TestValidInput(t *testing.T) {
	firstInput := 2
	actual := validInput(firstInput)
	expected := true

	if actual != expected {
		t.Errorf("Input condition not met")

	}
}

func TestValidInput_edge1(t *testing.T) {
	firstInput := 0
	actual := validInput(firstInput)
	expected := false

	if actual != expected {
		t.Errorf("Input condition not met")

	}
}

func TestValidInput_edge2(t *testing.T) {
	firstInput := 100
	actual := validInput(firstInput)
	expected := false

	if actual != expected {
		t.Errorf("Input condition not met")

	}

}
