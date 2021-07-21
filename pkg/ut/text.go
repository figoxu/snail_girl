package ut

import (
	"regexp"
	"strings"
)

type TextUt struct {
}

func (p *TextUt) CleanUZero(content string) string {
	content = strings.ReplaceAll(content, "\u0000", ``)
	return content
}

func (p *TextUt) MergeMultiBlankLine(content string) string {
	return regexp.MustCompile("\n{2,}").ReplaceAllString(content, "\n")
}

func (p *TextUt) Match(reg string, contents ...string) []string {
	var result []string
	for _, content := range contents {
		rs := regexp.MustCompile("(?i)"+reg).FindAllString(content, -1)
		result = append(result, rs...)
	}
	return result
}

func (p *TextUt) Clean(reg string, contents ...string) []string {
	var result []string
	for _, content := range contents {
		rs := regexp.MustCompile(reg).ReplaceAllString(content, "")
		result = append(result, rs)
	}
	return result
}
