package main

import (
	"crypto/md5"
	"fmt"
)

//Funcionalidad para pasar de MD5 a Int a Nro de avatar
func main() {
	user := "Juan Carlos Villagran"
	userMD5 := md5.Sum([]byte(user))
	fmt.Printf("MD5 Hash: %x \n", userMD5)

	sum := 0
	for i := 0; i < 16; i++ {
		byte := userMD5[i]
		byteToInt := int(byte)
		sum += byteToInt
		fmt.Printf("Byte %d: %x ", i, byte)
		fmt.Println(byteToInt)
		fmt.Println(sum)
	}
	fmt.Println(sum)
}
