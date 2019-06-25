# neo-rpc

A JSON-RPC client for neo-cli.

## Installation
```
go get github.com/o3labs/neo-rpc
```

## Importing
```go
import (
    "github.com/o3labs/neo-rpc/neorpc/neorpc"
)
```


## Usage

```go
client := neorpc.Client("https://us-east.o3node.org")
response, err := client.GetBlock(3623800)
if err != nil {
	log.Printf("err %v", err)
	return
}
log.Printf("%v", response)
```

