package slice

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Pet struct {
	name string
}

func CopySlice() {
	s1 := make([]*Pet, 0, 3)
	s1 = append(s1, &Pet{name: "dog"})

	// 深拷贝
	//s2 := make([]*Pet, len(s1))
	//copy(s2, s1) 
	
	// 浅拷贝
	s2 := s1 
	fmt.Println(cap(s2)) // 未超容量，应该不会触发扩容，还是指向原来的底层数据
	//s2[0] = &Pet{name: "cat"} // 为什么s2[0]对底层数组的修改能够同步到s1
	s2 = append(s2, &Pet{name: "cat"}) // 但是使用append的话s2修改不影响s1
	// 回答：append操作修改了切片的len，所以s1还是原来的len，s2是新的len
	// 也就是说，底层数组还是被修改了，但是s1因为len不足看不到修改

	// 所以如果s1再去修改底层数组的第二个位置，那么s2刚加的cat就会被改为tiger
	s1 = append(s1, &Pet{name: "tiger"})
	s2[1] = &Pet{name: "lion"}
	fmt.Println("s1[1]:",s1[1].name)
	fmt.Println("s2[1]:",s2[1].name)
	fmt.Println("s1的底层数组地址",(*reflect.SliceHeader)(unsafe.Pointer(&s1)).Data)
	fmt.Println("s2的底层数组地址",(*reflect.SliceHeader)(unsafe.Pointer(&s2)).Data)

	// 除非通过扩容使得其中一个切片扩容，开辟一个新的底层数组
	s1 = append(s1, &Pet{name: "elephant"})
	s1 = append(s1, &Pet{name: "mouse"})
	// 此时s1和s2已经指向不同的底层数组了，所以s2对底层数组的修改不再影响s1
	s2[1] = &Pet{name: "bird"}

	fmt.Println("s1[1]:",s1[1].name)
	fmt.Println("s2[1]:",s2[1].name)
	fmt.Println("s1的底层数组地址",(*reflect.SliceHeader)(unsafe.Pointer(&s1)).Data)
	fmt.Println("s2的底层数组地址",(*reflect.SliceHeader)(unsafe.Pointer(&s2)).Data)

}