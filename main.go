package main

import (
	"fmt"
	"github.com/AlmirKadric/redash-client-go/redash"
)

func main() {
	c, err := redash.NewClient(&redash.Config{
		RedashURI: "https://redash.data.internal.atlassian.com/",
		APIKey:    "something",
	})
	if err != nil {
		panic(err)
	}

	queries, err := c.GetQueries()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(queries)
}
