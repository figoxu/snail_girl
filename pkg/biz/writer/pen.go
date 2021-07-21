package writer

import (
	"sync"

	"figoxu.me/snail_girl/pkg/ut"
	"github.com/quexer/utee"
)

type Pen struct {
	delayWriter *utee.TimerCache
	lock        sync.Mutex
}

func NewPen() *Pen {
	pen := &Pen{
		lock: sync.Mutex{},
	}
	pen.delayWriter = utee.NewTimerCache(1, pen.delayWriting)
	return pen
}

func (p *Pen) Write(fileName, content string) {
	p.lock.Lock()
	defer p.lock.Unlock()
	v := p.delayWriter.Get(fileName)
	if v != nil {
		preData := v.(string)
		content = preData + content
	}
	p.delayWriter.Put(fileName, content)
}

func (p *Pen) delayWriting(key, value interface{}) {
	fileName := key.(string)
	content := value.(string)
	ut.File.FlushWrite(fileName, content)
}
