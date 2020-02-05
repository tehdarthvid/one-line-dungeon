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
	//fmt.Println(os.Getenv("OLD_DEPLOYMENT_ENV"))
	return os.Getenv("OLD_DEPLOYMENT_ENV")
	//return "dev"
}

// GenerateEncounter main returner.
func GenerateEncounter() Encounter {
	var res Encounter

	res.WorldName = "tempDev"
	res.Floor = "<.\\" + ".@!.[...g..\u2640|"

	return res

}
