package stringx

import (
	"bufio"
	"io"
	"os"
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

func FileExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		return false
	}
	return true
}
