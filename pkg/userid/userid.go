package userid

import (
	"math/rand"
	"time"
)

const SuperAdministrator = 10000000

func GetUserID() int32 {
	rand.NewSource(time.Now().Unix())
	id := rand.Intn(89999999)
	return int32(id) + SuperAdministrator
}
