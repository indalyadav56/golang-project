package main

func main() {
	// emoji := []byte{240, 159, 152, 132}
	// fmt.Println(string(emoji))
	// var b byte = 65         // decimal
	// fmt.Printf("%08b\n", b) // prints: 01000001 → binary
	// fmt.Println(string(b))  // "A"

	// fmt.Println(&b)

	// var b byte = 170 // 10101010
	// fmt.Printf("Decimal: %d\n", b)
	// fmt.Printf("Binary:  %08b\n", b)
	// fmt.Printf("Hex:     %02X\n", b)

	// str := "indal"
	// for i, runeChar := range str {
	// 	fmt.Println(i, runeChar)
	// }

	// data := "123aa"

	// runes := []rune(data)

	// fmt.Println([]rune(strconv.Itoa(123)))

	// fmt.Println(runes)

	// fmt.Println(strings.Map(func(r rune) rune {
	// 	if r == 'a' {
	// 		r = 'A'
	// 	}
	// 	return r
	// }, "testa"))

	// fmt.Println(byte(255))

	// byte is is binary representation
	//  which can hold the binary

	// 2^8  := 0-255
	// 9090909090

	// byteData, _ := json.Marshal(map[int]interface{}{1: "test"})
	// fmt.Println(byteData)

	// fmt.Println(unsafe.Sizeof([]byte{255, 155, 255, 255}))

	// 0 1

	// lowes urerese of binary bit

	// !!!!bytes

	// r := 'a' // 01100001
	// fmt.Println(byte(r))
	// fmt.Printf("%08b", r)

	// fmt.Print(unsafe.Sizeof(int16(1.00)))
	// fmt.Println(unsafe.Sizeof(byte('a'))) // Outputs: 1 (size of byte)

	// fmt.Println(unsafe.Sizeof([2]byte{97, 98}))

	// fmt.Println(unsafe.Sizeof([2]byte{97, 98}))

	// name := "indal"
	// fmt.Println(name[1:4])

	//
	// fmt.Println(unsafe.Sizeof([]byte("indal")))
	// fmt.Println(unsafe.Sizeof([5]rune{97, 97, 97, 97, 97}))
	// var b [5]byte
	// copy(b[:], "indal")           // Copy the string's bytes into the array
	// fmt.Println(unsafe.Sizeof(b)) // Outputs: 5

	// type SliceHeader struct {
	// 	Data uintptr // [97, 97].  8
	// 	Len  int 8
	// 	Cap  int 8
	// }

	// s := []byte("ing")

	// h := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	// // fmt.Println(h)
	// // fmt.Println(s)
	// fmt.Println(unsafe.Sizeof(s))
	// fmt.Println(unsafe.Sizeof(&h.Data))

	// // Size of underlying data
	// fmt.Println("Size of backing array (approx):", len(s)*int(unsafe.Sizeof(8))) // 10MB
	// // fmt.Println(len(s))
	// fmt.Println(cap(s))
	// fmt.Println(reflect.SliceHeader)
	// fmt.Println(s)

	// slice := make([]int, 0, 10)

	// fmt.Println(cap(slice))
	// fmt.Println(len(slice))

	// m := map[string]int{"x": 1}
	// h := *(*reflect.StringHeader)(unsafe.Pointer(&m)) // Not even allowed
	// h := *(*reflect.StringHeader)(unsafe.Pointer(&m)) // Not even allowed

	// fmt.Println(unsafe.Map)

	// chanage := "change"
	// sh.data = &chanage

	//DEMO =============================

	// r := rune('😀')
	// fmt.Println(r, string(r))
	// // fmt.Println([]byte(string(r)))
	// fmt.Println(string([]byte(string(r))))

	// //

	// str := "indal😀"

	// runes := []rune(str)

	// fmt.Println([]byte(string(runes)))

	// "strings are immutable"

	// name := "indal"
	// // name[0] = 'H'
	// // fmt.Println(name[1:3])
	// r := []rune(name)
	// r[0] = 'H'
	// fmt.Println(name)
	// fmt.Println(string(r))

	// "string manipulation "

	// for loop string

	// strings.ToUpper()

	// fmt.Println(strings.Map(func(r rune) rune {
	// 	if r >= 'a' && r <= 'z' {
	// 		r = r - 32
	// 	}
	// 	return r
	// }, "indal"))

	// 65 - 'A'
	// 97 - 'a'

	// DEMO BYTES =====================

	// b := byte('a')
	// fmt.Println(b)

	// str := "inda;"

	// b := []byte()

	// str := "indal"

	// ru := []rune(str)

	// b := []byte(string(ru))

	// fmt.Println(b)

	// fmt.Println([]byte(strconv.Itoa(123)))

}

// A slice is a 3-field struct:
// 	A pointer to the array

// 	A length

// 	A capacity

// runes
//bytes

// 1 byte = 8 bits

// ascii := 97
// architechture :=
