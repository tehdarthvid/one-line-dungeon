package main

import (
	"encoding/base64"
	"fmt"

	"./encounter"
)

func main() {
	//hasher := sha512.New()
	//hasher.Write([]byte("string"))

	//enc := encounter.CreateFromSeed(10)

	enc := encounter.Create()

	sha := base64.URLEncoding.EncodeToString(enc.Checksum)
	//fmt.Println("世界[" + getWhereAmI() + "] = " + string(generateFloor()))
	fmt.Println("世界[" + enc.WorldName + "] = " + enc.Floor + " " + sha)
}
