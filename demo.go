package main

import (
	"github.com/mathew-bowersox/graphql"
	"github.com/mathew-bowersox/jflect"
	"io"
	"log"
	"os"
	"time"
	"context"

)

func main() {

	// create a client (safe to share across requests)
	client := graphql.NewClient("https://countries.trevorblades.com/")
	client.Log = func(s string) { log.Println(s) }
	client.IndentLoggedJson = true

	client.ProcessResult = func (r io.Reader) error {
		strNme := "Results"
		err := generate.Generate(r, os.Stdout, &strNme)
		return err
	}
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	var responseData map[string]interface{}
	req := graphql.NewRequest(`
    query  {
		country(code: "BR") {
			name
            native
            emoji
            currency
            languages {
                code
                name
            }
        }
	}
`)

	err := client.Run(ctx, req, &responseData)
	log.Println(err)


}
