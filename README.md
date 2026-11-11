# Math Skills

A command-line Go program that reads a dataset of integers from a file and calculates its average, median, population variance, and standard deviation.

## Features

- Reads integer values from a text file.
- Accepts values separated by spaces, tabs, or newlines.
- Calculates:
  - **Average** — arithmetic mean.
  - **Median** — middle value after sorting, or the average of the two middle values.
  - **Variance** — population variance (divides by the number of values, `N`).
  - **Standard deviation** — square root of the population variance.
- Rounds the displayed results to the nearest integer.
- Reports an error when the file cannot be read or contains an invalid value.
- Handles an empty input file.

## Requirements

- Go installed on your machine.

The project uses Go's standard library; no third-party packages are required.

## Usage

Place your integer dataset in a text file, for example `data.txt`:

```text
1
2
3
4
5
```

Run the program from the project directory:

```bash
go run . data.txt
```

Example output:

```text
Average: 3
Median: 3
Variance: 2
Standard Deviation: 1
```

You can also build an executable:

```bash
go build -o math-skills .
```

Then run it with your input file:

```bash
./math-skills data.txt
```

On Windows, run:

```powershell
.\math-skills.exe data.txt
```

## Input format

The file must contain valid base-10 integers. Values can be separated by any whitespace, including newlines and spaces.

Valid example:

```text
-10  -5  0
5
10
```

Decimal values such as `1.5` are not accepted. An empty file is reported as `Empty file`.

## Calculations

Let the input contain `N` values, with mean `μ`:

**Average**

```text
sum(values) / N
```

**Median**

The middle value of the sorted dataset. For an even number of values, it is the average of the two middle values.

**Population variance**

```text
sum((x - μ)²) / N
```

**Standard deviation**

```text
sqrt(variance)
```

The program uses `math/big` (`big.Int` and `big.Rat`) to avoid integer overflow and preserve exact rational arithmetic in the average, median, and variance calculations. The input is still parsed into Go's native `int` type, so the accepted integer range depends on the platform. Standard deviation is calculated using `big.Float` and then converted to `float64` for the final square-root result, so extreme values may still lose some precision.

Displayed values are rounded to the nearest integer. Exact halfway cases are rounded away from zero.

## Tests

Run the unit tests with:

```bash
go test -v
```

## Project structure

A typical project layout:

```text
math-skills/
├── main.go
├── main_test.go
├── data.txt          # Your input file (optional)
└── README.md
```

## Notes

- Variance is **population variance**, not sample variance.
- The input must contain at least one integer to perform calculations.
- The file path is provided as the single command-line argument.
