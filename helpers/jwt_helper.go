package helpers

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)
func JwtToken(){
	secretKey := []byte("your-secret-key")

    // Create a JWT
    token := jwt.New(jwt.SigningMethodHS256)
    claims := token.Claims.(jwt.MapClaims)
    claims["user_id"] = 123
    claims["username"] = "john.doe"
    claims["exp"] = time.Now().Add(time.Hour * 1).Unix() 

    // Sign the token with the secret key
    tokenString, err := token.SignedString(secretKey)
    if err != nil {
        fmt.Println("Error creating token:", err)
        return
    }

    fmt.Println("Token:", tokenString)

    // Verify the token
    parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        // Make sure the signing method is what you expect
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
        }
        return secretKey, nil
    })

    if err != nil {
        fmt.Println("Token validation failed:", err)
        return
    }

    if parsedToken.Valid {
        fmt.Println("Token is valid")
        // Access claims
        if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok {
            fmt.Println("User ID:", claims["user_id"])
            fmt.Println("Username:", claims["username"])
        }
    } else {
        fmt.Println("Token is invalid")
    }
}


func VerifyToken(tokenString string) error {
    const (
        SecretKey = "your-secret-key"
    )

    // Parse the token
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
      // Validate method
      if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
        return nil, fmt.Errorf("Unexpected signing method") 
      }
      
      // Return key for validation
      return []byte(SecretKey), nil
    })
  
    // Check parse error
    if err != nil {
        fmt.Println("Error:", err)
      return err
    }
  
    // Check token validity
    if !token.Valid {
      return errors.New("Invalid token")
    }
  
    // Token is valid, return nil error
    return nil
  }