package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

// GetPasswordHash generates password hash.
func GetPasswordHash(salt string, password string) string {
	result := password + salt

	for i := 0; i < 1000; i++ {
		out, err := exec.Command("sh", "-c", "echo -n '"+result+"' | openssl sha256").Output()
		if err != nil {
			panic("Failed to generate the hash.")
		}

		outputRaw := string(out)
		output := strings.Split(outputRaw, " ")
		log.Printf("out: %#v", output)
		result = strings.TrimRight(output[1], "\n")

	}
	fmt.Printf("GetPasswordHash() salt: %v password: %v, result:%v", salt, password, result)
	return result
}

func main() {

	salt := "51d1ddd410c2ad3529bc8a2d4819c98e809bd48206b7ff9b34c0f498c3ea50796f96a94c6b5fcf37d9545cb8ee16d36fea72252ec7f4bd89bb3d839f220d9a79"
	password := "uq2kiIFuK"
	result := "5d1bbfddbb7ec6a022f0139cb28f9d228308053ed919b3e6551a0ff7c98a5797"

	res := GetPasswordHash(salt, password)
	log.Printf("res:%v, %v", res, result)
}
