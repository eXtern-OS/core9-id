package tokens

import (
	"github.com/eXtern-OS/core9-common/db"
	"github.com/eXtern-OS/core9-common/models/token"
)

func NewToken(userId string) (*token.Token, *token.Token, error) {
	access := token.NewToken(userId, false)
	refresh := token.NewToken(userId, true)

	err := db.DefaultClient.InsertOne(access, "users", "access_tokens")
	if err != nil {
		return nil, nil, err
	}
	return &access, &refresh, db.DefaultClient.InsertOne(refresh, "users", "refresh_tokens")
}
