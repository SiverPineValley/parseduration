package parseduration

import (
	"fmt"
	"regexp"
	"time"
)

var (
	ErrInvalidTimeUnit    = fmt.Errorf("Invalid time unit")
	ErrDuplicatedTimeUnit = fmt.Errorf("Duplicated time unit")

	unit = map[string]int64{
		"ns": int64(time.Nanosecond),
		"us": int64(time.Microsecond),
		"ms": int64(time.Millisecond),
		"s":  int64(time.Second),
		"m":  int64(time.Minute),
		"h":  int64(time.Hour),
		"d":  int64(time.Hour) * 24,   // 1 Days
		"w":  int64(time.Hour) * 168,  // 7 Days
		"M":  int64(time.Hour) * 720,  // 30 Days
		"y":  int64(time.Hour) * 8760, // 365 Days
	}
)

func ParseDuration(s string) (d time.Duration, e error) {
	r1, _ := regexp.Compile("[-+]?[0-9]*(ns|us|ms|s|m|h|d|w|M|y)+")
	parsed := r1.FindAllString(s, -1)

	if len(parsed) == 0 {
		return 0, ErrInvalidTimeUnit
	}

	dup := map[string]struct{}{}
	var y int64
	for _, v := range parsed {
		var m int64
		var p string

		m, p, e = parseUnit(v)
		if e != nil {
			return
		}

		if _, exists := dup[p]; exists {
			return 0, ErrDuplicatedTimeUnit
		}

		dup[p] = struct{}{}
		y += m * unit[p]
	}

	d = time.Duration(y)
	return
}

func parseUnit(s string) (int64, string, error) {
	if s == "" {
		return 0, "", ErrInvalidTimeUnit
	}

	i := 0
	isMinor := false
	var p string
	var m int64
	for ; i < len(s); i++ {
		c := s[i]
		switch true {
		case i == 0 && c == '-':
			isMinor = true
		case i == 0 && c == '+':
			continue
		case c >= '0' && c <= '9':
			m = m*10 + (int64(c) - '0')
		case c >= 'a' && c <= 'z':
			p = s[i:]
		case c >= 'A' && c <= 'Z':
			p = s[i:]
		default:
			return 0, "", ErrInvalidTimeUnit
		}

		if p != "" {
			if _, exists := unit[p]; !exists {
				return 0, "", ErrInvalidTimeUnit
			}

			break
		}
	}

	if p == "" {
		return 0, "", ErrInvalidTimeUnit
	}

	if isMinor {
		m = -1 * m
	}

	return m, p, nil
}
