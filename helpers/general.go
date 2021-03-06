package helpers

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	schema "github.com/lestrrat/go-jsschema"
)

var (
	rspace = regexp.MustCompile(`\s+`)
	rsnake = regexp.MustCompile(`_`)
)

func NotLast(i, len int) bool {
	return i != len-1
}

func GetKeyFromJSONPath(url string) string {
	s := strings.Split(url, "/")
	return UpperCamelCase(s[len(s)-1])
}

func LinkTitle(title, method, suffix string) string {
	return method + UpperCamelCase(title) + suffix
}

func JoinTypes(ts schema.PrimitiveTypes, sep string) string {
	var strs []string
	for _, t := range ts {
		strs = append(strs, t.String())
	}
	return strings.Join(strs, sep)
}

func Serialize(v interface{}) string {
	j, err := json.Marshal(v)
	if err != nil {
		return fmt.Sprint(v)
	}
	return string(j)
}

func ToStringLiteral(v interface{}) string {
	s, ok := v.(string)
	if ok {
		return fmt.Sprintf("\"%s\"", s)
	}
	return fmt.Sprint(v)
}

func Slice(s ...interface{}) []interface{} {
	return s
}

func In(e interface{}, s []interface{}) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}
