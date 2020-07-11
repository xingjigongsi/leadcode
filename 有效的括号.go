package leadcode

import (
	"strings"
)

func isValid1(s string) bool {
	rifo := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			rifo = append(rifo, s[i])
		} else {
			if len(rifo) == 0 {
				return false
			}
			temp := rifo[len(rifo)-1]
			rifo = rifo[0 : len(rifo)-1]
			if temp == '(' && s[i] != ')' {
				return false
			}
			if temp == '[' && s[i] != ']' {
				return false
			}
			if temp == '{' && s[i] != '}' {
				return false
			}
		}
	}
	if len(rifo) != 0 {
		return false
	}
	return true
}

func isValid2(s string) bool {
	rifo := []byte{}
	t := map[byte]byte{'{': '}', '[': ']', '(': ')', '?': '?'}
	for i := 0; i < len(s); i++ {
		v, ok := t[s[i]]
		if ok {
			rifo = append(rifo, v)
		} else {
			if len(rifo) == 0 {
				return false
			}
			temp := rifo[len(rifo)-1]
			rifo = rifo[:len(rifo)-1]
			if temp != s[i] {
				return false
			}
		}
	}
	if len(rifo) != 0 {
		return false
	}
	return true
}

func isValid3(s string) bool {
	for {
		if len(s) == 0 {
			break
		}
		if !strings.Contains(s, "{}") && !strings.Contains(s, "[]") &&
			!strings.Contains(s, "()") {
			break
		}
		s = strings.Replace(s, "{}", "", -1)
		s = strings.Replace(s, "[]", "", -1)
		s = strings.Replace(s, "()", "", -1)
	}
	if len(s) != 0 {
		return false
	}
	return true
}
