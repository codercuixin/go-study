package main

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"log"
)


func marshal()(bsk, bpk[]byte){
	sk, _ := rsa.GenerateKey(rand.Reader, 1024)
	pk := &sk.PublicKey

	//序列化
	bsk = x509.MarshalPKCS1PrivateKey(sk)
	bpk = x509.MarshalPKCS1PublicKey(pk)

	//反序列化
	sk2, _ := x509.ParsePKCS1PrivateKey(bsk)
	pk2, _ := x509.ParsePKCS1PublicKey(bpk)

	if !sk.Equal(sk2) || !pk.Equal(pk2){
		log.Fatal()
	}
	return
}


func encode(sk, pk[]byte)(string, string){
	msk := pem.EncodeToMemory(&pem.Block{
		Type: "RSA PRIVATE KEY",
		Bytes: sk,
	})

	mpk := pem.EncodeToMemory(&pem.Block{
		Type: "PUBLIC KEY",
		Bytes: pk ,
	})

	bsk, _ := pem.Decode(msk)
	bpk, _ := pem.Decode(mpk)
	if !bytes.Equal(sk, bsk.Bytes) ||!bytes.Equal(pk, bpk.Bytes){
			log.Fatal()
	}

	return string(msk), string(mpk)
}

func main(){
	bsk, bpk := marshal()
	psk, ppk := encode(bsk, bpk)
	println(psk)
	println(ppk)
}