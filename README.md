# id2shortstr
format int to short string

### 测试用例
``` golang
go test -v *.go
```

### 接口调用
``` golang

	package main

	import (
		"github.com/irebit/id2shortstr"
		"log"
	)

	func main() {
		var i = uint64(12312323423)
		log.Println("Id", i)
		id2ShortStr := id2shortstr.New()
		hexStr := id2ShortStr.GetStr(i)
		log.Println("encode", hexStr)
		hexInt := id2ShortStr.GetId(hexStr)
		log.Println("decode", hexInt)
	}
```