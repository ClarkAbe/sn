package main
import (
	"fmt"
	"github.com/clarkabe/sn"
)
func main() {
	q := sn.Init()
	q.Append([]byte("1"))
	q.Append([]byte("2"))
	q.Append([]byte("3"))
	q.Append([]byte("5"))
	q.Append([]byte("6"))
	q.Insert(3, []byte("4"))
	q.Each(func (k uint64, v []byte) {
		fmt.Println(k, string(v))
	})
	fmt.Scanln()
} 
