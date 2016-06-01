package utils


//True if the string s is empty. Spaces are considered not empty.
func StrEmpty(s string) bool {
	return s == ""
}

//True if the string s is not empty. Spaces are considered not empty.
func StrNotEmpty(s string) bool {
	return s != ""
}

//StringDefault returns the string value or a default value if original value is empty
func StringDefault(str string, defaultStr string) string {
	if str == "" {
		return defaultStr
	} else {
		return str
	}
}