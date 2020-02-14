package encounter

import (
	"crypto/sha512"
	"math/rand"
	"os"
	"time"
)

// Encounter that is randomly generated.
type Encounter struct {
	WorldName string
	Seed      int64
	Floor     string // for now, put everything on the floor
	Checksum  []byte // not really necessary, but I want to put in some math calculaitons
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func getTime() int64 {
	return time.Now().UnixNano()
}

func getWhereAmI() string {
	return os.Getenv("OLD_TARGET_ENV")
}

// Create and return an Encounter.
func Create() Encounter {

	//rand.Seed(time.Now().UnixNano())
	res := CreateFromSeed(rand.Int63())

	return res

}

// CreateFromSeed generate deterministic Encounter from seed
func CreateFromSeed(seed int64) Encounter {
	var res Encounter

	res.Seed = seed
	res.WorldName = getWhereAmI()
	res.Floor = "<.\\" + ".@!.[..." + string(rune((seed%26)+97)) + "..\u2640|"

	//fmt.Print(fmt.Sprintf("%v", res))
	hasher := sha512.New()
	//hasher.Write([]byte(fmt.Sprintf("%v", res)))
	res.Checksum = hasher.Sum(nil)

	return res

}
