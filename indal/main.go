package main

import (
	"errors"
	"fmt"
)

type CustomError struct {
	ErrorCode string
	ErrorMsg  string
	Cause     error `json:"-"` // Not serialized in JSON
}

func (c *CustomError) Error() string {
	return c.ErrorMsg
}

// Unwrap allows unwrapping to get the underlying cause
func (e *CustomError) Unwrap() error {
	return e.Cause
}

type User struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	// resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	// if err != nil {
	// 	panic(err)
	// }
	// defer resp.Body.Close()

	// if resp.StatusCode != http.StatusOK {
	// 	panic("Failed to fetch data: " + resp.Status)
	// }

	// byteData, err := io.ReadAll(resp.Body)

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(string(byteData))

	// var userObj User
	// if err := json.Unmarshal(byteData, &userObj); err != nil {
	// 	panic(err)
	// }

	// fmt.Println(userObj.UserID)
	// fmt.Println(userObj.ID)
	// fmt.Println(userObj.Title)
	// fmt.Println(userObj.Completed)

	fmt.Println(GetData().Unwrap())

}

func GetData() *CustomError {
	return &CustomError{
		ErrorCode: "20",
		ErrorMsg:  "hello this is custom error",
		Cause:     errors.New("this is error cause for cusotm use"),
	}
}
