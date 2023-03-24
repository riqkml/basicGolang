package internal

import (
	"container/list"
	"fmt"
)

func GenerateList() {
	data := list.New()

	data.PushFront("begin")
	data.PushBack("one")
	data.PushBack("two")

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
