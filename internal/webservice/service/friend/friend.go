package friend

import (
	"github.com/forgocode/family/internal/webservice/database/mysql"
	"github.com/forgocode/family/internal/webservice/model"
)

func createFriendShip(ship model.FriendShip) error {
	c := mysql.GetClient()
	return c.C.Model(&model.FriendShip{}).Create(&ship).Error
}

func removeFriendShip(uid, friendUID string) error {
	c := mysql.GetClient()
	return c.C.Where("uid = ? and friendUID = ?", uid, friendUID).Delete(&model.FriendShip{}).Error
}

func agreeFriendShip(uid, friendUID string) error {
	c := mysql.GetClient()
	return c.C.Model(&model.FriendShip{}).Where("uid = ? and friendUID = ?", uid, friendUID).Update("status", model.FriendShipAgree).Error
}

func blackFriendShip(uid, friendUID string) error {
	c := mysql.GetClient()
	return c.C.Model(&model.FriendShip{}).Where("uid = ? and friendUID = ?", uid, friendUID).Update("status", model.FriendShipBlack).Error
}
