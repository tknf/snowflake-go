package main

import (
	"fmt"

	"github.com/tknf/snowflake"
)

func main() {
	s := snowflake.NewSnowflake(1)
	id := s.Generate()
	fmt.Println(id)
}
