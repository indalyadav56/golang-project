package services

import (
	"context"
	"fmt"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (u *UserService) CreateUser(ctx context.Context) error {
	fmt.Println("user created!!!!!!")
	HeavyTask()
	return nil
}

func HeavyTask() {
	for i := range 20 {
		fibonaci(i)
	}
}

func fibonaci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonaci(n-1) + fibonaci(n-2)
}
