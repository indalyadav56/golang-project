package services

import (
	"context"
	"testing"
)

func TestUserServiceCreateUser(t *testing.T) {
	userSrv := NewUserService()
	userSrv.CreateUser(context.TODO())
}
