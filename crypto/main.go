package main

import (
	"crypto/cipher"
	"crypto/des"
	"fmt"
	"math/rand"
)

func makeRand(size int) []byte {
	r := make([]byte, size)
	for n := 0; n < len(r); n++ {
		r[n] = uint8(rand.Int() % 256)
	}
	return r
}

func main() {
	in := []byte("helloworld")
	key := []byte("love_go!")

	ci, _ := des.NewCipher(key[0:8])
	iv := makeRand(8)

	enc := cipher.NewCBCEncrypter(ci, iv)
	cin := make([]byte, int(len(in)/des.BlockSize)*des.BlockSize+des.BlockSize)
	out := make([]byte, len(cin))
	enc.CryptBlocks(out, cin)
	fmt.Printf("%X", iv)
}
