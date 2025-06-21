package main

func main() {
	// fmt.Println(ReverseString("indal"))
	// fmt.Println(ReverseArray([6]int{2, 4, 6, 8, 10, 12}))
	// fmt.Println(isPelindrome(""))

	// for i := range 10 {
	// 	fmt.Println(fibonnachi(i))
	// }
	// fmt.Println(EvenOdd(123))

	//

	// fmt.Println([]byte(strconv.Itoa(123)))

	// for i, val := range "123" {
	// 	fmt.Println(i, val)
	// }

	// utf - extened version of ascii

	// strings to bytes

	// byte := 0- 255(not more then  only positive)

	// byteData, _ := json.Marshal(map[string]interface{}{
	// 	"key": "value",
	// })

	// fmt.Println(string(byteData))

	// strings.Builder()

	// fmt.Printf("Bytes [49 50] as string: %s\n", []byte{49, 50})

	// m := "indal"

	// b := []byte(m)

	// b[0] = 97

	// m = string(b)

	// fmt.Println(string(b))

	// m := "indal"
	// r := []rune(m)
	// r[0] = 'H'
	// m = string(r)
	// fmt.Println(m) // Output: Hndal

	// nameChange := strings.Map(func(r rune) rune {
	// 	if r == 97 {
	// 		r = 65
	// 	}
	// 	return r
	// }, "indal")

	// fmt.Println(nameChange)

	// 10. Using bytes.Buffer for efficient byte manipulation
	// var buffer bytes.Buffer
	// buffer.WriteString("Hello, ")
	// buffer.Write([]byte("Buffer!"))
	// buffer.Write([]byte("Buffer!"))

	// fmt.Println(buffer.String())

	// os.WriteFile("te", buffer.Bytes(),)
	// fmt.Println(buffer.Len())

	// reader := strings.NewReader("Streamed data aewra wera ewr awer awe rawe rawe r awer awer awer ")
	// chunk := make([]byte, 5)
	// fmt.Println("\nReading from reader:")
	// for {
	// 	n, err := reader.Read(chunk)
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		fmt.Println("Read error:", err)
	// 		break
	// 	}
	// 	fmt.Printf("Read %d bytes: %q\n", n, string(chunk[:n]))
	// }

	// fmt.Println(strings.Fields("indal kumar"))
	// NewMethod(&NewStruct{})

	// fallthrough
	// x := 2
	// switch x {
	// case 1:
	// 	fmt.Println("One")
	// case 2:
	// 	fmt.Println("Two")
	// 	fallthrough
	// case 3:
	// 	fmt.Println("Three") // Will be printed because of fallthrough
	// 	fallthrough
	// case 4:
	// 	fmt.Println("Four")
	// 	fallthrough
	// default:
	// 	fmt.Println("Other")
	// }

	// i := 0

	// LoopStart:
	// 	for i < 5 {
	// 		if i == 3 {
	// 			fmt.Println("Going to skip 3 using goto")
	// 			i++
	// 			goto LoopStart // jumps back to the start of the loop
	// 		}
	// 		fmt.Println("i =", i)
	// 		i++
	// 	}

	// 	for i := 0; i < 3; i++ {
	// 		for j := 0; j < 3; j++ {
	// 			if i == 1 && j == 1 {
	// 				fmt.Println("Breaking out of both loops")
	// 				goto Done
	// 			}
	// 			fmt.Printf("i=%d, j=%d\n", i, j)
	// 		}
	// 	}

	// Done:
	// 	fmt.Println("Exited loops")

	// data := make([]byte, 3)

	// data[0] = 97
	// data[1] = 98
	// data[2] = 99

	// fmt.Println(bytes.Index(data, []byte("97")))

	// TestBytes(data)

	// fmt.Println(data)

	// name := "indal"

	// bytesData := []byte(name)

	// // bytesData = append(bytesData, []byte("kumar")...)
	// bytesData = append(bytesData, 12)
	// bytesData = append(bytesData, 120)
	// bytesData = append(bytesData, 124)
	// bytesData = append(bytesData, 125)

	// fmt.Println(string(bytesData), len(bytesData), cap(bytesData))

	// var buf bytes.Buffer

	// // Encode JSON into buffer
	// json.NewEncoder(&buf).Encode(map[string]interface{}{"name": "indal"})

	// // Write another raw JSON string into buffer
	// buf.Write([]byte(`{"name":"indal"}`))
	// buf.Write([]byte(`{"name":"indal"}`))
	// buf.Write([]byte(`{"name":"indal"}`))
	// buf.Write([]byte(`{"name":"indal"}`))
	// buf.Write([]byte(`{"name":"indal"}`))

	// // Print buffer contents
	// fmt.Println(buf.String())

	// a := []int{1, 2, 3}
	// b := a
	// b[0] = 100

	// b = append(b, 10)

	// fmt.Println(a)

	// fmt.Println(b)

	//heap :- *a+ 10

	// s1 := []string{"James", "Wagner", "Christene", "Mike"}
	// s2 := []string{"Paul", "Haaland", "Patrick"}
	// s3 := []string{"Peter", "Mark", "Luke"}
	// s3 = append(append(s1, s2...), s3...)
	// fmt.Println(s3)

	// a := "strings"

	// for i, runeChar := range a {
	// 	fmt.Println(i, string(runeChar))
	// }

	// fmt.Println([]rune(strconv.Itoa(123)))

	// b := []byte(strconv.Itoa(123)) // 0 -2555
	// fmt.Println(string(b))

	// byteData, _ := json.Marshal(map[string]interface{}{})

	// r := rune('😀')
	// b := []byte(string(r))

	// fmt.Println(b)
}

// func TestBytes(s []byte) {
// 	s[0] = 100
// 	s = append(s, 100)
// }

// result := strings.Map(func(r rune) rune {
// 	if r >= 'A' && r <= 'Z' {
// 		return r + 32 // convert to lowercase
// 	}
// 	return r
// }, "HeLLo Go!")

// fmt.Println(result) // Output: hello go!
