package entity

type FriendShip struct {
	FirstUserId      int    `json:"firstUserId"`
	SecondUserId     int    `json:"secondUserId"`
	FriendShipStatus string `json:"friendShipStatus"`
}
