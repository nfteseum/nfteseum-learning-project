// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package sqlc

import (
	"database/sql"
	"time"
)

type Comments struct {
	ID        int32     `json:"id"`
	PostID    int32     `json:"postID"`
	Author    string    `json:"author"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}

type Likes struct {
	ID      int32  `json:"id"`
	PostID  int32  `json:"postID"`
	LikedBy string `json:"likedBy"`
}

type Posts struct {
	ID           int32         `json:"id"`
	ContractAddr string        `json:"contractAddr"`
	TokenID      int32         `json:"tokenID"`
	LikeCount    sql.NullInt32 `json:"likeCount"`
	CommentCount sql.NullInt32 `json:"commentCount"`
	Author       string        `json:"author"`
	CreatedAt    time.Time     `json:"createdAt"`
}

type Users struct {
	Addr      string       `json:"addr"`
	Admin     sql.NullBool `json:"admin"`
	Name      string       `json:"name"`
	Pfp       interface{}  `json:"pfp"`
	RandomMsg string       `json:"randomMsg"`
}