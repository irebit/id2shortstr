package main

import (
	"log"

	"github.com/irebit/id2shortstr"
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
