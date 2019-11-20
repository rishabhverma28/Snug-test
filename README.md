# Snug Developer Test

## Description
Find the number of days between 2 given dates.

## Steps to run the Python Code
- Clone the repository.
- Run `python3 datediff.py [date1] - [date2]`, replacing `[date1]` and `[date2]` whith the start and end dates on the terminal.

## Steps to run the Golang code
- Clone the repository.
- To build the code run `go build datediff.go`
- Run `./datediff [date1] - [date2]`, replacing `[date1]` and `[date2]` whith the start and end dates on the terminal.
- To see the Test Coverage results, open `cover.html` in any browser.
- To regenerate the Test coverage report, run the following commands:

```go
go test -coverprofile cover.out   
go tool cover -html=cover.out -o cover.html
```