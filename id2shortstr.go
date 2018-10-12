package id2shortstr

import (
	"math"
)

type Id2ShortStr struct {
	Key    []byte
	KeyPos map[byte]byte
	Hex    byte
	Id2Str map[uint64]string
	Str2Id map[string]uint64
}

var (
	Key = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_")
)

func New() *Id2ShortStr {
	keyPos := map[byte]byte{}
	keyLen := len(Key)
	for i := 0; i < keyLen; i++ {
		keyPos[Key[i]] = byte(i)
	}
	return &Id2ShortStr{Key, keyPos, byte(64), map[uint64]string{}, map[string]uint64{}}
}

func (s *Id2ShortStr) GetStr(id uint64) string {
	if s.Id2Str[id] != "" {
		return s.Id2Str[id]
	}
	var hexArr []uint64
	var ret = []byte{}
	_, hexArr = s.encode(id, hexArr)
	len := len(hexArr)
	for i := len - 1; i >= 0; i-- {
		ret = append(ret, s.Key[int(hexArr[i])])
	}
	str := string(ret)
	s.Save(id, str)
	return str
}

func (s *Id2ShortStr) GetId(str string) uint64 {
	if s.Str2Id[str] != 0 {
		return s.Str2Id[str]
	}
	strlen := len(str)
	var hexArr []byte
	for i := 0; i < strlen; i++ {
		hexArr = append(hexArr, s.KeyPos[byte(str[i])])
	}
	ret := s.decode(hexArr)
	s.Save(ret, str)
	return ret
}

func (s *Id2ShortStr) decode(arr []byte) uint64 {
	arrLen := len(arr)
	var ret uint64
	var i = 1
	for _, b := range arr {
		ret += uint64(b) * uint64(math.Pow(float64(s.Hex), float64(arrLen-i)))
		i++
	}
	return ret
}

func (s *Id2ShortStr) encode(id uint64, arr []uint64) (uint64, []uint64) {
	quotient := id / uint64(s.Hex)
	remainder := id % uint64(s.Hex)
	arr = append(arr, remainder)
	if quotient < uint64(s.Hex) {
		if quotient != 0 {
			arr = append(arr, quotient)
		}
		return quotient, arr
	} else {
		return s.encode(quotient, arr)
	}
}

func (s *Id2ShortStr) Save(id uint64, str string) {
	s.Id2Str[id] = str
	s.Str2Id[str] = id
}
