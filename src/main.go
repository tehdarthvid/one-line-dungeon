package main

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"

	"./encounter"
)

func main() {
	hasher := sha512.New()
	hasher.Write([]byte("string"))

	enc := encounter.Create()

	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	//fmt.Println("世界[" + getWhereAmI() + "] = " + string(generateFloor()))
	fmt.Println("世界[" + enc.WorldName + "] = " + enc.Floor + " " + sha)
}
