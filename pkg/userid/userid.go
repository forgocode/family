package userid

import (
	"math/rand"
	"strconv"
	"time"
)

const SuperAdministrator = 10000000

func GetUserID() string {
	rand.NewSource(time.Now().Unix())
	id := rand.Intn(89999999)
	return strconv.Itoa(id + SuperAdministrator)
}
