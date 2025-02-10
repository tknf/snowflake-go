# Snowflake ID for Go

This is the simplest Snowflake ID generator library.

For production use, I recommend [bwmarrin/snowflake](https://github.com/bwmarrin/snowflake).

## Installation

```bash
go get github.com/tknf/snowflake-go
```

## Usage

```go
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

```

## License

MIT License
