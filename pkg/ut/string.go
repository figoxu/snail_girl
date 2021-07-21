package ut

import (
	"encoding/json"
	"strings"

	"github.com/jinzhu/inflection"
)

// snake string, XxYy to xx_yy , XxYY to xx_yy
func SnakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:]))
}

func SnakeStrings(strs ...string) (result []string) {
	for _, str := range strs {
		s := SnakeString(str)
		result = append(result, s)
	}
	return result
}

// camel string, xx_yy to XxYy
func CamelString(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}

func SingularName(v string) string {
	vs := strings.Split(v, "_")
	idx := len(vs) - 1
	w := inflection.Singular(vs[idx])
	vs[idx] = w
	v = strings.Join(vs, "_")
	return CamelString(v)
}

func JsonString(obj interface{}) string {
	b, err := json.Marshal(obj)
	if err != nil {
		return ""
	}
	return string(b)
}
