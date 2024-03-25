package stringx

import (
	"bufio"
	"io"
)

// ReadLines 逐行读取字符串
func ReadLines(input io.Reader, cb func(s string)) error {
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		cb(scanner.Text())
		if scanner.Err() != nil {
			return scanner.Err()
		}
	}
	return nil
}
