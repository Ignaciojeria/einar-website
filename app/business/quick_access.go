package business

import (
	"context"
	"encoding/base64"
	"my-project-name/app/adapter/out/firestore"
	"my-project-name/app/domain"
	"my-project-name/app/domain/ports/in"

	"github.com/google/uuid"
)

var QuickAccess in.QuickAccess = func(ctx context.Context, e domain.User) (err error) {
	//IMPLEMENT YOUR BUSINESS USECASE HERE
	user, err := firestore.FindUserByEmail(ctx, e)
	if err != nil {
		return err
	}

	if user == (domain.User{}) {
		user.ID = uuid.NewString()
		apiKey := uuid.NewString() + ":" + uuid.NewString()
		e.ApiKey = base64.RawStdEncoding.EncodeToString([]byte(apiKey))
		err = firestore.SaveUser(ctx, e)
	}

	if err != nil {
		return err
	}
	//TODO CHECK HOW TO SEND MAGIC LINK USING FIRESTORE!

	return nil
}
