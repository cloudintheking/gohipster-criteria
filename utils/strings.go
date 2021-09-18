package utils

import "unicode"

func NewString(data string) *string {
	s := new(string)
	*s = data
	return s
}
func NewInt(data int) *int {
	i := new(int)
	*i = data
	return i
}

/**
 * @Description: 首字母大写
 * @param str
 * @return string
 */
func Ucfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

/**
 * @Description: 首字母小写
 * @param str
 * @return string
 */
func Lcfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}
