package main

import (
	"errors"
	"math"
	"strconv"
	"strings"
	"unicode"
)

// Bytes unit convert
const (
	B = 1 << (10 * iota)
	KB
	MB
	GB
	TB
	PB
	EB
)

// FormatByteSize returns human string of byte size
func FormatByteSize(n int64, precision int) string {
	value, unit := float64(n), "B"
	switch {
	case value >= EB:
		value, unit = value/EB, "EB"
	case value >= PB:
		value, unit = value/PB, "PB"
	case value >= TB:
		value, unit = value/TB, "TB"
	case value >= GB:
		value, unit = value/GB, "GB"
	case value >= MB:
		value, unit = value/MB, "MB"
	case value >= KB:
		value, unit = value/KB, "KB"
	}

	r := strconv.FormatFloat(value, 'f', precision, 64)
	r += unit

	return r
}

// ParseByteSize returns int size of string size
func ParseByteSize(s string) (int64, error) {
	s = strings.TrimSpace(strings.ToUpper(s))
	i := strings.IndexFunc(s, unicode.IsLetter)
	if i == -1 {
		i = len(s)
		s += "B"
	}

	value, unit := strings.TrimSpace(s[:i]), strings.TrimSpace(s[i:])
	bytes, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, err
	}

	if bytes < 0 {
		return 0, errors.New("byte size string invalid")
	}

	switch unit {
	case "E", "EB":
		return int64(bytes * EB), nil
	case "P", "PB":
		return int64(bytes * PB), nil
	case "T", "TB":
		return int64(bytes * TB), nil
	case "G", "GB":
		return int64(bytes * GB), nil
	case "M", "MB":
		return int64(bytes * MB), nil
	case "K", "KB":
		return int64(bytes * KB), nil
	case "B":
		return int64(bytes * B), nil
	default:
		return 0, errors.New("byte size string invalid")
	}
}

// Round returns round number with precision
func Round(n float64, precision int) (r float64) {
	pow := math.Pow(10, float64(precision))
	num := n * pow
	_, div := math.Modf(num)

	if n >= 0 && div >= 0.5 {
		r = math.Ceil(num)
	} else if n < 0 && div > -0.5 {
		r = math.Ceil(num)
	} else {
		r = math.Floor(num)
	}

	return r / pow
}

// Comma returns number string with comma
func Comma(n float64, precision int) string {
	s := strconv.FormatFloat(n, 'f', precision, 64)

	sc := ""
	if s[0] == '-' {
		sc = "-"
		s = s[1:]
	}

	si, sf := s, ""
	if strings.Contains(s, ".") {
		ss := strings.Split(s, ".")
		si, sf = ss[0], ss[1]
	}

	ss := []string{}
	for {
		if len(si) == 0 {
			break
		} else {
			start := len(si) - 3
			if start < 0 {
				start = 0
			}
			ss = append([]string{si[start:]}, ss...)
			si = si[:start]
		}
	}

	s = sc + strings.Join(ss, ",")
	if sf != "" {
		s += "." + sf
	}

	return s
}
