This repository contains a Go program which validates whether given set of a
credit cards have the following characteristics:

1) It must start with a 4, 5 or 6.
2) It must contain exactly 16 digits.
3) It must only consist of digits (0-9).
4) It may have digits in groups of 4, separated by one hyphen "-".
5) It must NOT use any other separator like ' ' , '_', etc.
6) It must NOT have 4 or more consecutive repeated digits.

## Prerequistes

* Install go

	```
	brew install go

	```

* Install PCRE library

	```
	brew install pcre

	```

---
**NOTE**

Golang regex package regexp uses the [re2 engine](https://github.com/google/re2/wiki/Syntax)
which doesn’t support backreferences.

So we are using another golang package [go-pcre](https://github.com/gijsbers/go-pcre)
which uses `PCRE`(Perl Compatible Regular Expressions).

---

### Steps to run the program

 1) Clone the github repository

	 ```
	 git clone https://github.com/Bincy1234/validateCreditCardNumber

	 ```

 2) Build the Project

		```
		go build

		```

 3) Run

		```
		go run validateCreditCard.go

		```

### Steps for Unit Testing

	```
	go test -v

	```
