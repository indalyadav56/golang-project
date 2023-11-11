package utils

import (
	"bufio"
	"golang-project/models"
	"log"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

// load environment
func LoadEnv() map[string]string {
  f, err := os.Open(".env")
  if err != nil {
    log.Fatal(err)
  }
  defer f.Close()

  scanner := bufio.NewScanner(f)
  vars := make(map[string]string)
  
  for scanner.Scan() {
    line := scanner.Text()
    if equal := strings.Index(line, "="); equal >= 0 {
      key := line[:equal]
      value := line[equal+1:]
      vars[key] = value
    }
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

  return vars
}

func GenerateToken(user models.User) (string, error) {

  token := jwt.New(jwt.SigningMethodHS256)

  claims := token.Claims.(jwt.MapClaims)
  claims["authorized"] = true
  claims["user_id"] = user.ID
  claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix() // 1 week expiration

  return token.SignedString([]byte("secret"))
}