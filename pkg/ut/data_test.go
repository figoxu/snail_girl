package ut

import (
	"fmt"
	"testing"
)

func TestData(t *testing.T) {
	// seq:=IntSeq{
	// }
	// for i:=0;i<100;i++ {
	// 	fmt.Println("a"+seq.Next()+"a")
	// }
	ss := NewStringSeq()
	for i := 0; i < 100; i++ {
		fmt.Println("a " + ss.Next() + " a")
	}
	fmt.Print(int('a'), int('z'), int('A'), int('Z'))
}
