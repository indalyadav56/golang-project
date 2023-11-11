package graphql

import (
	"log"
	"net/http"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

var schema *graphql.Schema

func main() {

  // Define simple schema
  schema = graphql.MustParseSchema(`
    type Query {
      hello: String
    }
  `, &Query{})

  http.Handle("/query", &relay.Handler{Schema: schema})

  log.Fatal(http.ListenAndServe(":8080", nil))
}

// Define Query resolver
type Query struct {}

func (_ *Query) Hello() string {
  return "Hello, world!" 
}