package encounter

import (
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

func getTime() int64 {
	return time.Now().UnixNano()
}

func getWhereAmI() string {
	return os.Getenv("OLD_TARGET_ENV")
}

// Create and return an Encounter.
func Create() Encounter {
	var res Encounter

	res.WorldName = getWhereAmI()
	res.Floor = "<.\\" + ".@!.[...g..\u2640|"

	return res

}

// CreateFromSeed generate deterministic Encounter from seed
func CreateFromSeed(seed int64) Encounter {
	var res Encounter

	res.WorldName = getWhereAmI()
	res.Floor = "<.\\" + ".@!.[..." + string(rune((seed%26)+97)) + "..\u2640|"

	return res

}
