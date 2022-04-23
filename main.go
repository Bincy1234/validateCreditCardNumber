package main

//
import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gijsbers/go-pcre"
)

//
func validateCreditCard(str string) bool {

	regex := regexp.MustCompile(`^[4-6]\d{15}|^[4-6]\d{3}-(?:\d{4}-){2}\d{4}?$`)
	format_check := regex.MatchString(str)

	trimHyphen := strings.Replace(str, "-", "", -1)

	consecutive_occurence := pcre.MustCompile(`^.*(.)\1{3}`, 0)
	consecutive_occurence_check := consecutive_occurence.MatcherString(trimHyphen, 0)

	if (format_check == true) && (!consecutive_occurence_check.Matches()) {
		return true
	} else {
		return false
	}
}

//
func validInput(firstInput int) bool {

	if (firstInput >= 0) && (firstInput < 100) {
		return true
	} else {
		return false
	}
}

//
func main() {

	length := 0

	fmt.Println("Enter the inputs")
	fmt.Scanln(&length)

	creditCardNumbersList := make([]string, length)

	if validInput(length) {

		for i := 0; i < length; i++ {
			fmt.Scanln(&creditCardNumbersList[i])
		}

	} else {

		fmt.Println("Invalid input")
	}
	fmt.Println(creditCardNumbersList)

	for i := 0; i < length; i++ {
		if validateCreditCard(creditCardNumbersList[i]) == true {
			fmt.Println("Valid")
		} else {
			fmt.Println("Invalid")
		}
	}

}
