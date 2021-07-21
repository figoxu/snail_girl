package sniffer

import (
	"strings"

	"figoxu.me/snail_girl/pkg/ut"
)

type GoPkgSniffer struct {
}

func (p *GoPkgSniffer) AsPkg(filePath string) (string, error) {
	const pkgFlagStr = "/pkg"
	pkgIndex := strings.Index(filePath, pkgFlagStr)
	if pkgIndex == -1 {
		return "", nil
	}
	prefixPath := filePath[:pkgIndex]
	modPath := prefixPath + "/go.mod"
	ft := ut.FileUt{}
	vs, err := ft.ReadLinesSlice(modPath)
	if err != nil {
		return "", err
	}
	for _, v := range vs {
		prefix := "module"
		prefixFlag := strings.HasPrefix(v, prefix)
		if prefixFlag {
			content := strings.Replace(v, prefix, "", -1)
			content = strings.TrimSpace(content)
			suffix := filePath[pkgIndex:]
			return content + suffix, nil
		}
	}
	return modPath, nil
}
