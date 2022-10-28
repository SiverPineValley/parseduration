# Golang parseduration

<a href="https://github.com/SiverPineValley/parseduration"><img src="https://goreportcard.com/badge/github.com/SiverPineValley/parseduration" alt="Go Report Card"></a>

The built-in package time offer very useful functions.

However, when you use time.ParseDuration, you cannot use time unit that is bigger than hours. This package also offer day, week, month, year time units for parsing string to duration.


## Installation
```shell
go get https://github.com/SiverPineValley/parseduration
```

## Format
The standard for parsing is same as time.ParseDuration. You can use `ParseDuration` sending parameters such as `2w3d20h`.


Negative numbers are also provided, and if you enter only letters without entering a number, it will be recognized as 1. `d == 1d`


## Standard
Time units are as provided in the table below. A time unit that is not provided or a time unit that is duplicated will cause an error.


In addition, the maximum value for each unit was not separately filtered.


|Time Unit|Duration|
|------|---|
|ns|1 Nano seconds|
|us|1 Micro Seconds|
|ms|1 Milli Seconds|
|s|1 Seconds|
|m|1 Minutes|
|h|1 Hours|
|d|1 Days|
|w|1 Weeks (7 Days)|
|M|1 Months (30 Days)|
|y|1 Years (365 Days)|

## Examples

```go
package main

import (
	"fmt"
    pd "github.com/SiverPineValley/parseduration"
)

// It returns 4 weeks + 2 days + 3 hours Duration
func main() {
	fmt.Println(pd.ParseDuration("4w2d3h"))
}
```