package helper

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Contains(list []string, target string) bool {
	for _, v := range list {
		if v == target {
			return true
		}
	}
	return false
}

func IndexOf[T comparable](slice []T, val T) int {
	for i, v := range slice {
		if v == val {
			return i
		}
	}
	return -1
}

func Reverse[T any](slice []T) []T {
	out := make([]T, len(slice))
	copy(out, slice)
	for i := 0; i < len(out)/2; i++ {
		j := len(out) - 1 - i
		out[i], out[j] = out[j], out[i]
	}
	return out
}

func Map[T any, R any](slice []T, f func(T) R) []R {
	out := make([]R, len(slice))
	for i, v := range slice {
		out[i] = f(v)
	}
	return out
}

func Filter[T any](slice []T, f func(T) bool) []T {
	out := make([]T, 0)
	for _, v := range slice {
		if f(v) {
			out = append(out, v)
		}
	}
	return out
}

func Reduce[T any, R any](slice []T, init R, f func(R, T) R) R {
	acc := init
	for _, v := range slice {
		acc = f(acc, v)
	}
	return acc
}

func ReadLines(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")

	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	return lines, nil
}

func ReadLinesAsInts(path string) ([]int, error) {
	lines, err := ReadLines(path)
	if err != nil {
		return nil, err
	}

	out := make([]int, len(lines))
	for i, l := range lines {
		v, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}
		out[i] = v
	}

	return out, nil
}
func ReadLinesAsIntArrays(path string) ([][]int, error) {
	lines, err := ReadLines(path)
	if err != nil {
		return nil, err
	}

	out := make([][]int, len(lines))
	for i, l := range lines {
		if err != nil {
			return nil, err
		}
		out[i], err = StringToIntSlice(l)
	}

	return out, nil
}

func ReadLinesAsStringArrays(path string) ([][]string, error) {
	lines, err := ReadLines(path)
	if err != nil {
		return nil, err
	}

	out := make([][]string, len(lines))
	for i, l := range lines {
		if err != nil {
			return nil, err
		}
		out[i] = StringToSlice(l)
	}

	return out, nil
}

func StringToIntSlice(s string) ([]int, error) {
	result := make([]int, len(s))
	for i, ch := range s {
		if ch < '0' || ch > '9' {
			return nil, fmt.Errorf("ungültiges Zeichen: %q", ch)
		}
		result[i] = int(ch - '0')
	}
	return result, nil
}

func StringToSlice(s string) []string {
	result := make([]string, len(s))

	for i, ch := range s {
		result[i] = string(ch)
	}

	return result
}

func SplitByChars(s string, delims string) []string {
	return strings.FieldsFunc(s, func(r rune) bool {
		return strings.ContainsRune(delims, r)
	})
}

func DeepCopy(field [][]string) [][]string {
	copyField := make([][]string, len(field))
	for i := range field {
		copyField[i] = make([]string, len(field[i]))
		copy(copyField[i], field[i])
	}
	return copyField
}
