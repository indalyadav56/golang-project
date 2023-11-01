package utils

import (
	"bufio"
	"log"
	"os"
	"strings"
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