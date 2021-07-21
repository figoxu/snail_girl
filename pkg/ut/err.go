package ut

import "strings"

func Chk(err error) {
	if err != nil {
		panic(err)
	}
}

func ErrNotExist(err error) bool {
	if err == nil {
		return false
	}
	features := []string{
		"no such file or directory",
		"file does not exist",
	}
	msg := err.Error()
	for _, feature := range features {
		if strings.Contains(msg, feature) {
			return true
		}
	}
	return false
}
