package entity

type Friendship struct {
	FirstUserId      string `json:"first_user_id"`
	SecondUserId     string `json:"second_user_id"`
	FriendshipStatus string `json:"friendship_status"`
}
