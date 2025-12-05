# PropelAuth Go Client Library

## Installation

```bash
go get github.com/propelauth/byo-go
```

## Usage

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/propelauth/byo-go"
)

func main() {
    // Create client
    client, err := byo.NewClient(
        "https://api.propelauth.com",
        "your_api_key_here",
    )
    if err != nil {
        log.Fatal(err)
    }

    ctx := context.Background()

    // Use the client
    resp, err := client.Ping(ctx)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Ping response:", resp)
}
```
