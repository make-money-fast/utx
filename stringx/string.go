package stringx

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Base64Encode 使用 Base64 编码字符串
func Base64Encode(input string) string {
	return base64.StdEncoding.EncodeToString([]byte(input))
}

// Base64Decode 使用 Base64 解码字符串
func Base64Decode(input string) (string, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return "", err
	}
	return string(decodedBytes), nil
}

// PrettyPrintJSON 格式化打印 JSON
func PrettyPrintJSON(input interface{}) error {
	prettyJSON, err := json.MarshalIndent(input, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(prettyJSON))
	return nil
}

// JSONToMap 将 JSON 字符串转换为 map[string]interface{}
func JSONToMap(input string) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := json.Unmarshal([]byte(input), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MapToJSON 将 map[string]interface{} 转换为 JSON 字符串
func MapToJSON(input map[string]interface{}) (string, error) {
	jsonBytes, err := json.Marshal(input)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

// SplitString 切割字符串
func SplitString(input, delimiter string) []string {
	parts := strings.Split(input, delimiter)
	return parts
}

// HiddenString 隐藏部分字符串
func HiddenString(s string, keepStart int, keepEnd int) string {
	if len(s) < keepStart+keepEnd {
		return s
	}
	return s[:keepStart] + "****" + s[len(s)-keepEnd:]
}

const (
	numSeeds  = "1234567890"
	charSeeds = "abcdefgABCDEFGhijklmnHIJKLMNopqOPQrstTSTuvwxyzUVWXYZ"
	fullSeeds = numSeeds + charSeeds
)

// RandString 随机字符串
func RandString(n int) string {
	var s string
	for i := 0; i < n; i++ {
		s += string(fullSeeds[rand.Intn(len(fullSeeds))])
	}
	return s
}

// RandNumber 随机数字字符串
func RandNumber(n int) string {
	var s string
	for i := 0; i < n; i++ {
		s += string(numSeeds[rand.Intn(len(numSeeds))])
	}
	return s
}

// RandChars 随机英文字符
func RandChars(n int) string {
	var s string
	for i := 0; i < n; i++ {
		s += string(charSeeds[rand.Intn(len(charSeeds))])
	}
	return s
}

// Uniq 去重.
func Uniq(s []string) []string {
	var x = make(map[string]struct{}, len(s))
	for _, i := range s {
		x[i] = struct{}{}
	}
	var res = make([]string, len(x))
	for k := range x {
		res = append(res, k)
	}
	return res
}
