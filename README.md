This repository contains a Go program which validates whether given set of a
credit cards have the following characteristics:

► It must start with a 4, 5 or 6.
► It must contain exactly 16 digits.
► It must only consist of digits (0-9).
► It may have digits in groups of 4, separated by one hyphen "-".
► It must NOT use any other separator like ' ' , '_', etc.
► It must NOT have 4 or more consecutive repeated digits.

## Prerequistes

* Install go

	```
	brew install go

	```

* Install PCRE library

	```
	brew install pcre


	```

### Steps to run the program

 1) Clone the github repository

	 ```
	 	git clone  

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
--------------

	```
	go test -v

	```
