package ut

import (
	"regexp"
	"strings"
)

type Parser struct {
	PrepareReg []string
	ProcessReg []string
}

func (p *Parser) Exe(content string) []string {
	prep := func(reg string, contents ...string) []string {
		var result []string
		for _, content := range contents {
			rs := regexp.MustCompile("(?i)"+reg).FindAllString(content, -1)
			result = append(result, rs...)
		}
		return result
	}

	proc := func(reg string, contents ...string) []string {
		var result []string
		for _, content := range contents {
			rs := regexp.MustCompile("(?i)"+reg).ReplaceAllString(content, "")
			result = append(result, rs)
		}
		return result
	}

	trimAndClear := func(strs ...string) []string {
		result := []string{}
		for _, v := range strs {
			v = strings.TrimSpace(v)
			if v != "" {
				result = append(result, v)
			}
		}
		return result
	}

	result := []string{content}
	for _, reg := range p.PrepareReg {
		result = prep(reg, result...)
	}
	for _, reg := range p.ProcessReg {
		result = proc(reg, result...)
	}
	return trimAndClear(result...)
}
