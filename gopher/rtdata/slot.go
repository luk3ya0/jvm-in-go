package rtdata

import "gopher/rtdata/heap"

type Slot struct {
	num int32
	ref *heap.Object
}
