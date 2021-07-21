package writer_test

import (
	"fmt"
	"testing"

	"figoxu.me/snail_girl/pkg/biz/writer"
)

func TestScaffoldDomain(t *testing.T) {
	d := &writer.ScaffoldDomain{
		Parser:         &writer.ScaffoldDomainParser{},
		ServiceWriter:  &writer.ScaffoldDomainServiceWriter{},
		RepoWriter:     &writer.ScaffoldDomainRepoWriter{},
		TestDataWriter: &writer.ScaffoldDomainTestDataWriter{},
	}
	v := d.GenerateResult("/Users/xujianhui/xxbmm/projects/workspace_go/meishi/m2/pkg/domain/search_key.go", "")
	fmt.Println(v)
}
