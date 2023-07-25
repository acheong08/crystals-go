package main

import (
	"github.com/acheong08/crystals-go/dilithium"
	"github.com/acheong08/crystals-go/kyber"
)

// Just testing a few things to make sure the code is working
func main() {
	d := dilithium.NewDilithium3()
	pk, sk := d.KeyGen([]byte("test"))
	msg := []byte("hello world")
	sig := d.Sign(sk, msg)
	if !d.Verify(pk, msg, sig) {
		panic("failed")
	} else {
		println("success")
	}
	k := kyber.NewKyber512()
	pk, sk = k.KeyGen([]byte("test"))
	msg = []byte("hello world")
	c, ss := k.Encaps(pk, msg)
	newss := k.Decaps(sk, c)
	// Convert to string and compare
	if string(ss) != string(newss) {
		panic("failed")
	} else {
		println("success")
	}
}
