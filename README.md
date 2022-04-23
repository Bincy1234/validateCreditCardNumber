This repository contains a Go program which validates whether given set of a credit card has the following characteristics:

► It must start with a 4, 5 or 6.
► It must contain exactly 16 digits.
► It must only consist of digits (0-9).
► It may have digits in groups of 4, separated by one hyphen "-".
► It must NOT use any other separator like ' ' , '_', etc.
► It must NOT have 4 or more consecutive repeated digits.

Installing
----------------
brew install go


Clone repository

git clone  

To build and run
----------------
cd <dir>

go build
go get <>
go run main.go

Run unit tests
--------------
go test
