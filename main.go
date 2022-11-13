package main

import (
	"github.com/acheong08/crystals-go/dilithium"
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
}
