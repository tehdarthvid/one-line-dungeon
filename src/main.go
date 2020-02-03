package main

import (
	"crypto/sha512"
	"fmt"
	"time"
)

var dungeonFeatures = []rune{'^', '<', '>'}
var dungeonItems = []rune{'!', ';', '~', ']', '↑', '♀', '*', '%', '♪'}
var dungeonCreatures = []rune{'k', 'g', 'D', 'w', 'b', 'B', 'z', 'j', 'J'}

// Scenario is a self contained event or encounter.
type Scenario struct {
	WorldName string
	Seed      int64
	Floor     []byte // for now, put everything on the floor
	Checksum  []byte // not really necessary, but I want to put in some math calculaitons
}

func getTime() int64 {
	return time.Now().UnixNano()
}

func getWhereAmI() string {
	//fmt.Println(os.Getenv("OLD_DEPLOYMENT_ENV"))
	//return os.Getenv("OLD_DEPLOYMENT_ENV")
	return "dev"
}

func getMonster(seed int64) string {
	return "k"
}

func generateEncounter() string {
	return "<.\\" + ".@!.[..." + getMonster(getTime()) + "..\u2640|"
}

func generateFloor() []byte {
	var floor []byte
	for i := 0; i < 16; i++ {
		floor = append(floor, byte(i%10)+64)
	}
	return floor
}

func main() {
	hasher := sha512.New()
	hasher.Write([]byte("string"))

	//sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	fmt.Println("世界[" + getWhereAmI() + "] = " + string(generateFloor()))
}
