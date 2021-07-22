package writer

import (
	"fmt"
	"strings"

	"figoxu.me/snail_girl/pkg/ut"
	"github.com/flosch/pongo2"
	"github.com/pkg/errors"
	"github.com/quexer/utee"
	"github.com/sirupsen/logrus"
)

var Tpl struct {
	declare                 string
	content                 string
	pushOrder               string
	curl                    string
	domainData              string
	domainRepo              string
	domainRepoImpl          string
	domainRepoInternalModel string
	domainRepoTest          string
	domainService           string
	domainServiceImpl       string
	domainServiceTest       string
	domainVaild             string
}

func init() {
	Tpl.domainData = ut.File.ReadAsset(`/assets/domain_data.tpl`)
	Tpl.domainRepo = ut.File.ReadAsset("/assets/domain_repo.tpl")
	Tpl.domainRepoImpl = ut.File.ReadAsset("/assets/domain_repo_impl.tpl")
	Tpl.domainRepoInternalModel = ut.File.ReadAsset("/assets/domain_repo_internal_model.tpl")
	Tpl.domainRepoTest = ut.File.ReadAsset("/assets/domain_repo_impl_test.tpl")
	Tpl.domainService = ut.File.ReadAsset("/assets/domain_service.tpl")
	Tpl.domainServiceImpl = ut.File.ReadAsset("/assets/domain_service_impl.tpl")
	Tpl.domainServiceTest = ut.File.ReadAsset("/assets/domain_service_impl_test.tpl")
	Tpl.curl = ut.File.ReadAsset(`/assets/domain_data.tpl`)
	Tpl.domainVaild = ut.File.ReadAsset(`/assets/domain_vaild.tpl`)
}

type Writer interface {
	HandleChange(string) error
	HandleDelete(string) error
}

type WordProcess interface {
	MatchInput(string) []string
	GenerateResult(string, string) string
}

type FileFeature struct {
	Watch string
	Write string
}

type CommonWriter struct {
	delayChange *utee.TimerCache
	feature     FileFeature
	processes   []WordProcess
	pen         *Pen
}

func (p *CommonWriter) HandleDelete(key string) error {
	if ut.File.Exist(key) {
		return nil
	}
	name := p.getFileName(key)
	if ut.File.Exist(name) {
		err := ut.File.Remove(name)
		return errors.WithStack(err)
	}
	return nil
}

func (p *CommonWriter) HandleChange(key string) error {
	p.delayChange.Put(key, nil)
	return nil
}

func (p *CommonWriter) excDelayChange(key, value interface{}) {
	inputFile := fmt.Sprint(key)
	if p.isGenFile(inputFile) {
		return
	}
	if !p.canProcess(inputFile) {
		return
	}
	content, err := ut.File.ReadAll(inputFile)
	if ut.ErrNotExist(err) {
		return
	}
	content = ut.Text.CleanUZero(content)
	contents := []string{}
	for _, process := range p.processes {
		vs := process.MatchInput(content)
		for _, v := range vs {
			logrus.Println(`process at ` + v)
			content := process.GenerateResult(v, content)
			contents = append(contents, content)
		}
	}

	pongoTpl, err := pongo2.FromString(Tpl.content)
	ut.Chk(err)
	out, err := pongoTpl.Execute(pongo2.Context{"contents": contents})
	ut.Chk(err)
	out = ut.Text.MergeMultiBlankLine(out)
	p.pen.Write(p.getFileName(inputFile), out)
}

func (p *CommonWriter) getFileName(key string) string {
	return strings.ReplaceAll(key, p.feature.Watch, p.feature.Write)
}

func (p *CommonWriter) isGenFile(key string) bool {
	return strings.Contains(key, p.feature.Write)
}

func (p *CommonWriter) canProcess(key string) bool {
	return strings.Contains(key, p.feature.Watch)
}
