package main

import (
	"crypto/sha512"
	"fmt"
	"time"
)

// An Encounter and relevant details
type Encounter struct {
	Arena     string
	Health    uint8
	Monesters []byte
	Items     []byte
	Encounter string
	Checksum  []byte
}

func getTime() int64 {
	return time.Now().UnixNano()
}

func getWhereAmI() string {
	return "dev"
}

func getMonster(seed int64) string {
	return "k"
}

func generateEncounter() string {
	return "<.\\" + ".@!.[..." + getMonster(getTime()) + "..\u2640|"
}

func main() {
	hasher := sha512.New()
	hasher.Write([]byte("string"))

	//sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	fmt.Println("世界[" + getWhereAmI() + "] = " + generateEncounter())
}
