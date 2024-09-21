package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"log"
	"math/rand"
)

func padToBlockSize(d []byte) []byte {
	n := len(d)
	blockSize := aes.BlockSize
	//确保 n 是 blockSize 的整数倍
	if n % blockSize != 0{
		n += blockSize - (n % blockSize)
		tmp := make([]byte, n)
		//最后会用 0 来填充
		copy(tmp, d)
		d = tmp 
	}
	return d
}

func test(key, data[]byte) []byte{
	key = padToBlockSize(key)
	block, err := aes.NewCipher(key)
	if err !=nil {
		log.Fatalln(err)
	}
	//enc 
	src := padToBlockSize(data)
	enc := make([]byte, len(src))
	mode := cipher.NewCBCEncrypter(block, key[:aes.BlockSize])
	mode.CryptBlocks(enc, src)

	//dec
	dec := make([]byte, len(enc))
	mode = cipher.NewCBCDecrypter(block, key[:aes.BlockSize])
	mode.CryptBlocks(dec, enc)

	return dec
}


func main(){
	for i:=0; i<10000;i++{
		key := make([]byte, 3)
		data := make([]byte, i)
		rand.Read(key)
		rand.Read(data)

		r := test(key, data)
		if !bytes.Equal(r[:i], data){
			log.Fatalln(i)
		}
	}
}