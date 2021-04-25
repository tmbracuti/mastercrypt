package main

//import "C"

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"
)
/*
func ExampleNewCFBDecrypter(cipher_in string) string {
	
	key := []byte{0xaa, 0xde, 0xff, 0xfa, 0xaa, 0xde, 0xff, 0xfa,0xa0, 0xd0, 0xff, 0xc1,0xca, 0xde, 0x07, 0x01}  	
	ciphertext, err := hex.DecodeString(cipher_in)
	for _, b := range ciphertext {
		fmt.Printf("%x ", b)
	}
	fmt.Printf("\n")

	if err != nil {
		fmt.Println(err);
	}
  	block, err := aes.NewCipher(key)
  	if err != nil {
  		panic(err)
  	}

  	if len(ciphertext) < aes.BlockSize {
  		panic("ciphertext too short")
  	}
  	
  	iv := ciphertext[:aes.BlockSize]
	for _, b := range iv {
		fmt.Printf("%x ", b)
	}
	fmt.Printf("\n")
	
	ciphertext = ciphertext[aes.BlockSize:]
	for _, b := range ciphertext {
		fmt.Printf("%x ", b)
	}
	fmt.Printf("\n")
  
  	stream := cipher.NewCFBDecrypter(block, iv)
    	
  	stream.XORKeyStream(ciphertext, ciphertext)
  	ret := fmt.Sprintf("%s", ciphertext)
  	return ret
}
*/

func StarfishCFBEncrypter(plainTxt string) string {
  	
	key:= []byte{0xaa, 0xde, 0xff, 0xfa, 0xaa, 0xde, 0xff, 0xfa,0xa0, 0xd0, 0xff, 0xc1,0xca, 0xde, 0x07, 0x01}	
  	plaintext := []byte(plainTxt)
  
  	block, err := aes.NewCipher(key)
  	if err != nil {
  		panic(err)
  	}
    	
  	ciphertext := make([]byte, aes.BlockSize+len(plaintext))	
  	iv := ciphertext[:aes.BlockSize]
	
  	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
  		panic(err)
  	}
  
  	stream := cipher.NewCFBEncrypter(block, iv)
  	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
	txt := fmt.Sprintf("%x", ciphertext)	
	//fmt.Printf("ciphertext bytes: %x\n", ciphertext)
	return txt
  
}

func main() { 
	//fmt.Printf("aes block size: %d\n", aes.BlockSize)
	if len(os.Args) < 2 {
		fmt.Println("usage: ./mastercrypt <password>")
		return
	}
	
	pwd := os.Args[1]
	fmt.Printf("encrypting password %s...\n", pwd)

	hexciphertxt := StarfishCFBEncrypter(pwd)
	fmt.Printf("use this cipher text: %s\n", hexciphertxt)

} 