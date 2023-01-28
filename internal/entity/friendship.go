package entity

import uuid "github.com/satori/go.uuid"

type Friendship struct {
	FirstUserId      uuid.UUID `json:"first_user_id"`
	SecondUserId     uuid.UUID `json:"second_user_id"`
	FriendshipStatus string    `json:"friendship_status"`
}
