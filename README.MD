# ESEP Grade Calculator

This project is a simple grade calculator for ESEP at the University of Florida.
It computes final grades for a student based on assignments, exams, and essays, using weighted averages.

## Tech Stack
- Go
- Go testing framework

## Project Structure
- **grade_calculator.go** – main implementation of the GradeCalculator
- **grade_calculator_test.go** – automated unit tests for the calculator
- **go.mod** – Go module definition
- **README.md** – project documentation

## How to Clone
1. Clone the repository:
   ```bash
   git clone https://github.com/Mo-Sanchez/esep-grade-calculator
   cd esep-grade-calculator/grade-calculator

## How to run tests
Run unit tests:
```
cd grade_calculator/
go test ./...
```

Run unit tests with code coverage:
```
cd grade_calculator/
go test -coverprofile=coverage.txt ./...
```
