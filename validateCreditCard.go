package main

/*
A program to check whther a given credit card number is valid or not.
A valid credit card has the following characteristics .
> It must start with 4,5 or 6.
> It must contain exactly 16  digits.
> It must only consist of digits (0-9).
> It may have digits in groups of 4, separated by one hyphen "-".
> It must NOT use any other separator like ' ' , '_', etc.
> It must NOT have 4 or more consecutive repeated digits.

*/

// Imports necessary pacakges
import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gijsbers/go-pcre"
)

/*
function checks whether the given credit card number has the reuired characteristics

Parameters: str<string> - credit card number
Return Value: returns boolean
*/
func validateNumber(str string) bool {

	regex := regexp.MustCompile(`^[4-6]\d{15}|^[4-6]\d{3}-(?:\d{4}-){2}\d{4}?$`)
	format_check := regex.MatchString(str)

	trimHyphen := strings.Replace(str, "-", "", -1)

	consecutive_occurence := pcre.MustCompile(`^.*(\d)\1{3}`, 0)
	consecutive_occurence_check := consecutive_occurence.MatcherString(trimHyphen, 0)

	if (format_check == true) && (!consecutive_occurence_check.Matches()) {
		return true
	} else {
		return false
	}
}

/*
function checks whether the first input N is 0<N>100.

Parameters: firstInput<integer> - users first input
Return Value: returns boolean
*/
func validInput(firstInput int) bool {

	if (firstInput > 0) && (firstInput < 100) {
		return true
	} else {
		return false
	}
}

// main function
func main() {

	length := 0

	inputformat := "Enter the inputs in the following format \n" +
		"The first line of input contains an integer . \n" +
		"The next  lines contain credit card numbers. "

	fmt.Println(inputformat)

	fmt.Scanln(&length)

	if validInput(length) {
		creditCardNumbersList := make([]string, length)
		for i := 0; i < length; i++ {
			fmt.Scanln(&creditCardNumbersList[i])
		}

		for i := 0; i < length; i++ {
			if validateNumber(creditCardNumbersList[i]) == true {
				fmt.Println("Valid")
			} else {
				fmt.Println("Invalid")
			}
		}
	} else {
		fmt.Println("Invalid input")
	}

}
