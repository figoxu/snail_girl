package ut

import (
	"encoding/hex"
	"fmt"
)

type DataSeqTp int

const (
	DataSeqTpUnKnown DataSeqTp = iota
	DataSeqTpInt
	DataSeqTpBool
	DataSeqTpFloat
	DataSeqTpString
)

type DataSeq struct {
	Int    ISeq
	Bool   ISeq
	Float  ISeq
	String ISeq
}

func NewDataSeq() *DataSeq {
	return &DataSeq{
		Int:    &IntSeq{},
		Bool:   &BoolSeq{},
		Float:  &FloatSeq{},
		String: NewStringSeq(),
	}
}

func (p *DataSeq) Next(tp DataSeqTp) string {
	switch tp {
	case DataSeqTpInt:
		return p.Int.Next()
	case DataSeqTpBool:
		return p.Bool.Next()
	case DataSeqTpFloat:
		return p.Float.Next()
	case DataSeqTpString:
		return p.String.Next()
	}
	return "undefined"
}

type ISeq interface {
	Next() string
}

type IntSeq struct {
	val int
}

func (p *IntSeq) Next() string {
	if p.val <= 0 {
		p.val = 0
	}
	p.val = p.val + 1
	return fmt.Sprintf(" %d ", p.val)
}

type BoolSeq struct {
	val bool
}

func (p *BoolSeq) Next() string {
	p.val = !p.val
	return fmt.Sprint(p.val)
}

type FloatSeq struct {
	val float64
}

func (p *FloatSeq) Next() string {
	p.val = p.val + 0.3
	return fmt.Sprint(p.val)
}

type StringSeq struct {
	val rune
}

func NewStringSeq() *StringSeq {
	var c rune = 'A'
	return &StringSeq{
		val: c - 1,
	}
}

func (p *StringSeq) Next() string {
	wrap := func(x string) string {
		return `"` + x + `"`
	}
	p.val = p.val + 1
	var end01 rune = 'Z'
	if p.val == end01+1 {
		p.val = 'a'
		return wrap(string(p.val))
	}
	if p.val > 122 {
		return wrap("h" + hex.EncodeToString([]byte(string(p.val))))
	}
	return wrap(string(p.val))
}
