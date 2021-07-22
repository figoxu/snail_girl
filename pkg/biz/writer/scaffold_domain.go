package writer

import (
	"fmt"
	"strings"

	"figoxu.me/snail_girl/pkg/ut"
	"github.com/ahmetb/go-linq/v3"
	"github.com/flosch/pongo2"
	"github.com/pkg/errors"
)

type ScaffoldDomain struct {
	Parser         *ScaffoldDomainParser
	ServiceWriter  *ScaffoldDomainServiceWriter
	RepoWriter     *ScaffoldDomainRepoWriter
	TestDataWriter *ScaffoldDomainTestDataWriter
	VaildWriter    *ScaffoldDomainVaildWriter
}

func (p *ScaffoldDomain) MatchInput(content string) []string {
	var vs []string
	return vs
}

func (p *ScaffoldDomain) GenerateResult(fileSeed, clazzName string) string {
	writers := []ScaffoldDomainWriter{p.ServiceWriter, p.RepoWriter, p.TestDataWriter, p.VaildWriter}
	content, err := ut.File.ReadAll(fileSeed)
	ut.Chk(err)
	vs := p.Parser.Parse(content)
	if len(vs) <= 0 {
		return ""
	}
	fileName := ut.FilePath(fileSeed).FileName()
	if clazzName == "" {
		clazzName = ut.CamelString(strings.ReplaceAll(fileName, ".go", ""))
	}
	for _, v := range vs {
		matchFlag := strings.ToLower(v.Name) == strings.ToLower(clazzName)
		if matchFlag {
			for _, writer := range writers {
				err = writer.WriteFile(fileSeed, v)
				ut.Chk(err)
			}
			return "ok"
		}
	}
	return ""
}

type ScaffoldDomainParser struct {
}

type DomainTp string

type DomainTpList []DomainTp

func (p *DomainTp) GetSeqTp() ut.DataSeqTp {
	isIn := func(x DomainTp, vs DomainTpList) bool {
		foundFlag := linq.From(vs).WhereT(func(v DomainTp) bool {
			return v == x
		}).Count() > 0
		return foundFlag
	}
	v := *p
	if isIn(v, p.intTps()) {
		return ut.DataSeqTpInt
	} else if isIn(v, p.floatTps()) {
		return ut.DataSeqTpFloat
	} else if isIn(v, p.stringTps()) {
		return ut.DataSeqTpString
	} else if isIn(v, p.booleanTps()) {
		return ut.DataSeqTpBool
	}
	return ut.DataSeqTpInt // 不识别的类型都当做Int枚举
}

func (p *DomainTp) intTps() DomainTpList {
	return []DomainTp{
		"uint8", "uint16", "uint32", "uint64",
		"int8", "int16", "int32", "int64",
		"byte", "rune", "uint", "int", "uintptr",
		"utee.Tick"}
}

func (p *DomainTp) floatTps() DomainTpList {
	return []DomainTp{"float32", "float64", "complex64", "complex128"}
}

func (p *DomainTp) stringTps() DomainTpList {
	return []DomainTp{"string"}
}

func (p *DomainTp) booleanTps() DomainTpList {
	return []DomainTp{"bool"}
}

type DomainField struct {
	Name    string
	Type    DomainTp
	MockVal string
	GormTag string
}

type DomainStruct struct {
	Name   string
	Fields []*DomainField
}

func (p *ScaffoldDomainParser) Parse(content string) []*DomainStruct {
	var vs []*DomainStruct
	pr := ut.Parser{
		PrepareReg: []string{".*type.*struct.*\\{(.*\\n)+?.*\\}"},
	}
	structContents := pr.Exe(content)
	for _, structContent := range structContents {
		fields := p.parseFields(structContent)
		v := &DomainStruct{
			Name:   p.parseName(structContent),
			Fields: fields,
		}
		vs = append(vs, v)
	}
	return vs
}

func (p *ScaffoldDomainParser) parseName(structContent string) string {
	pr := ut.Parser{
		PrepareReg: []string{"type.*struct"},
		ProcessReg: []string{"type", "struct", " "},
	}
	vs := pr.Exe(structContent)
	if len(vs) <= 0 {
		return ""
	}
	return vs[0]
}

func (p *ScaffoldDomainParser) parseFields(structContent string) []*DomainField {
	var vs []*DomainField
	pr := ut.Parser{
		PrepareReg: []string{`.*\n`},
		ProcessReg: []string{"type.*struct.*", "//.*"},
	}
	contents := pr.Exe(structContent)
	for _, content := range contents {
		var subs []string
		linq.From(strings.Split(content, " ")).WhereT(func(x string) bool {
			return x != ""
		}).ToSlice(&subs)
		field := &DomainField{
			Name: subs[0],
			Type: DomainTp(subs[1]),
		}
		vs = append(vs, field)
	}
	return vs
}

type ScaffoldDomainWriter interface {
	WriteFile(seedFile string, domainStruct *DomainStruct) error
}

type ScaffoldDomainTestDataWriter struct {
}

func (p *ScaffoldDomainTestDataWriter) WriteFile(seedFile string, domainStruct *DomainStruct) error {
	genData := func() {
		seq := ut.NewDataSeq()
		linq.From(domainStruct.Fields).ForEachT(func(x *DomainField) {
			tp := x.Type.GetSeqTp()
			x.MockVal = seq.Next(tp)
		})
	}

	genData()
	content, err := genContent(seedFile, Tpl.domainData, domainStruct)
	if err != nil {
		return err
	}
	target := suggestPath(seedFile, "/../testdata/")
	if err := touch(target); err != nil {
		return err
	}
	ut.File.FlushWrite(target.String(), content)
	return nil
}

func suggestPath(seedFile, middlePath string) ut.FilePath {
	f := ut.FilePath(seedFile)
	fileName := f.FileName()
	folderName := f.FolderName()
	target := ut.FilePath(folderName + middlePath + fileName)
	if target.Exist() {
		target = ut.FilePath(folderName + middlePath + fileName + ".txt")
	}
	return target
}

func touch(target ut.FilePath) error {
	fut := ut.FileUt{}
	file, err := fut.MakeFile(target.FolderName(), target.FileName())
	if err != nil {
		return errors.WithStack(err)
	}
	err = file.Close()
	return errors.WithStack(err)
}

func genContent(fileSeed, tplContent string, domainStruct *DomainStruct) (string, error) {
	fileName := ut.FilePath(fileSeed).FileName()
	tpl, err := pongo2.FromString(tplContent)
	if err != nil {
		return "", errors.WithStack(err)
	}
	fmt.Println("----------->")
	fmt.Println(ut.JsonString(domainStruct))
	fmt.Println("<-----------")
	out, err := tpl.Execute(pongo2.Context{
		"content":  domainStruct,
		"fileName": fileName,
	})
	return out, errors.WithStack(err)
}

type ScaffoldDomainRepoWriter struct {
}

func (p *ScaffoldDomainRepoWriter) WriteFile(seedFile string, domainStruct *DomainStruct) error {
	genGormTag := func() {
		seq := ut.NewDataSeq()
		linq.From(domainStruct.Fields).ForEachT(func(x *DomainField) {
			if strings.ToLower(x.Name) == "id" {
				x.GormTag = `gorm:"primary_key"`
				return
			}
			tp := x.Type.GetSeqTp()
			switch tp {
			case ut.DataSeqTpInt:
				x.GormTag = `gorm:"not null"` //;default:0
				return
			case ut.DataSeqTpBool:
				x.GormTag = `gorm:"not null"` //;default:false
				return
			case ut.DataSeqTpFloat:
				x.GormTag = `gorm:"not null"` //;default:0.0
				return
			case ut.DataSeqTpString:
				x.GormTag = `gorm:"not null"` //;default:''
				return
			}
			x.MockVal = seq.Next(tp)
		})
	}

	genFile := func(tplContent, middlePath string) error {
		content, err := genContent(seedFile, tplContent, domainStruct)
		if err != nil {
			return err
		}
		target := suggestPath(seedFile, middlePath)
		if err := touch(target); err != nil {
			return err
		}
		ut.File.FlushWrite(target.String(), content)
		return nil
	}
	err := genFile(Tpl.domainRepo, "/../repo/")
	if err != nil {
		return err
	}
	err = genFile(Tpl.domainRepoImpl, "/../repo/impl/")
	if err != nil {
		return err
	}

	genGormTag()
	err = genFile(Tpl.domainRepoInternalModel, "/../repo/internal/model/")
	if err != nil {
		return err
	}

	seedFile = strings.ReplaceAll(seedFile, ".go", "_test.go")
	err = genFile(Tpl.domainRepoTest, "/../repo/impl/")
	return err
}

type ScaffoldDomainServiceWriter struct {
}

func (p *ScaffoldDomainServiceWriter) WriteFile(seedFile string, domainStruct *DomainStruct) error {
	genFile := func(tplContent, middlePath string) error {
		content, err := genContent(seedFile, tplContent, domainStruct)
		if err != nil {
			return err
		}
		target := suggestPath(seedFile, middlePath)
		if err := touch(target); err != nil {
			return err
		}
		ut.File.FlushWrite(target.String(), content)
		return nil
	}
	err := genFile(Tpl.domainService, "/../service/")
	if err != nil {
		return err
	}
	err = genFile(Tpl.domainServiceImpl, "/../service/impl/")
	if err != nil {
		return err
	}
	seedFile = strings.ReplaceAll(seedFile, ".go", "_test.go")
	err = genFile(Tpl.domainServiceTest, "/../service/impl/")
	return err
}

type ScaffoldDomainVaildWriter struct {
}

func (p *ScaffoldDomainVaildWriter) WriteFile(seedFile string, domainStruct *DomainStruct) error {
	genFile := func(tplContent, middlePath string) error {
		content, err := genContent(seedFile, tplContent, domainStruct)
		if err != nil {
			return err
		}
		target := suggestPath(seedFile, middlePath)
		if err := touch(target); err != nil {
			return err
		}
		ut.File.FlushWrite(target.String(), content)
		return nil
	}
	err := genFile(Tpl.domainVaild, "./")
	return err
}
