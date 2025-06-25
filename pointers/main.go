package main

import "fmt"

func main() {

	// data := 9

	// fmt.Printf("%d, %b \n", data, data)

	// // pointer to struct
	// ptrStruct := PointerStuct{
	// 	Name:  "INDAL",
	// 	Email: "indal@gmail.com",
	// }

	// fmt.Println(ptrStruct)

	// ptr := &ptrStruct

	// ptr.Name = "newName"
	// ptr.Email = "newEmail"

	// fmt.Println(*ptr)

	// // pointer to slices
	// slices := []int{1, 2, 3, 4, 5}

	// sp := &slices

	// slices = append(*sp, 6, 7, 8)

	// fmt.Println(slices)
	// fmt.Println(*sp)

	// var p *int

	// * opeartion (defaine o) and dereferenceing

	// pointer type(base type should be only int)

	// str := "indal"
	// x := 10

	// p = &x

	// *p = 1000

	// fmt.Println(*p)

	// x := 10
	// p = &x

	// fmt.Println(*p)

	// newInt := int64(1999999999999999999)
	// fmt.Printf("%d, %b", newInt, newInt)

	// obj1 := PointerStuct{
	// 	Name:  "indal",
	// 	Email: "indal@gmail.com",
	// }
	// fmt.Printf("%p\n", &obj1)
	// fmt.Println(&obj1)

	// bytes, _ := json.Marshal(obj1)
	// fmt.Printf("%v, %b \n", obj1, bytes)

	// myData := []byte("indal-kumar-yadav")
	// fmt.Printf("%b \n", myData)
	// fmt.Printf("%p", &myData)

	// there is another way also

	// var ptr *[]byte = &myData
	// fmt.Println(ptr)

	// var intData *int
	// var strData *string
	// var boolData *bool

	// str := "indal"
	// intD := 10
	// boolD := true

	// intData = &intD
	// strData = &str
	// boolData = &boolD

	// fmt.Println(intData)
	// fmt.Println(strData)
	// fmt.Println(boolData)

	// obj1 := PointerStuct{
	// 	Name:  "indal",
	// 	Email: "indal@gmail.com",
	// }

	// josnBytes, _ := json.Marshal(obj1)
	// fmt.Printf("%p  %#v %[2]T %x", &obj1, obj1, josnBytes)

	// arr := [...]int{1, 2, 3, 4, 5}

	// fmt.Println("current Array:=>", arr)
	// PointerSlice(&arr)
	// fmt.Println("after Array:=>", arr)

	// arr := [5]int{10, 20, 30, 40, 50}

	// // Get pointer to first element
	// ptr := (*int)(unsafe.Pointer(&arr[0]))

	// fmt.Println(ptr)

	// size := unsafe.Sizeof(arr[0])

	// // Access second element using pointer arithmetic
	// second := *(*int)(unsafe.Pointer(uintptr(ptr) + size))

	// fmt.Println(second)

	// Size of int
	// size := unsafe.Sizeof(arr[0])

	// var ptr *func() string

	// f := func() string {
	// 	fmt.Println("Hello")
	// 	return "INDAL"
	// }
	// ptr = &f
	// fmt.Println((*ptr)())

	// vals := []any{true, "off", "hello"}
	// fmt.Println(vals)

	// out := new(PointerStuct)
	// fmt.Println(&out)

	// var pd *int

	// slices := []int{1, 2, 3}

	// sp := &slices

	// fmt.Println(sp)

	// ptr := unsafe.Pointer(&slices[0])

	sl := []int{1, 2, 3}

	fmt.Printf("Before modification slice Addr %p \n", &sl)
	SlicePointer(&sl)
	fmt.Printf("After modification slice Addr %p \n", &sl)

}

func SlicePointer(slPtr *[]int) {
	fmt.Printf("SlicePointer Before mmodification, inner func pointer addr  %p \n", slPtr)
	fmt.Printf("SlicePointer Before mmodification, inner func pointer addr  %p \n", &slPtr)
	*slPtr = append(*slPtr, 100, 200, 300)
	fmt.Printf("SlicePointer After modification, inner func pointer addr  %p \n", &slPtr)
}

type PointerStuct struct {
	Name  string
	Email string
}

func PointerSlice(s *[5]int) {
	s[3] = 100
}
