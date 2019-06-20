package id2shortstr

import (
	"log"
	"testing"
)

func TestId2ShortStr(t *testing.T) {
	var i = uint64(12312323423)
	log.Println("Id", i)

	id2ShortStr := New()

	hexStr := id2ShortStr.GetStr(i)
	log.Println("encode", hexStr)

	hexInt := id2ShortStr.GetId(hexStr)
	log.Println("decode", hexInt)
}
