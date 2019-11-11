# github.com/gotoolkit/log Go項目日誌配置


## Install

```console
go get -u github.com/gotoolkit/log
```

## Example


```go
package main

import (
    "github.com/gotoolkit/log"
)

func main() {
    err := log.Setup()
	if err != nil {
		log.Fatalln(err)
    }
    
    log.Println("hello, world")
}
```