package main

import (
	"fmt"

	"github.com/tknf/snowflake-go"
)

func main() {
	s := snowflake.NewSnowflake(1)
	id := s.Generate()
	fmt.Println(id)
}
