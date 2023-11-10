package firestore

import (
	"context"
	"my-project-name/app/domain"
	"my-project-name/app/domain/ports/out"
	einar "my-project-name/app/shared/archetype/firestore"

	"cloud.google.com/go/firestore"
)

var FindUserByEmail out.FindUserByEmail = func(ctx context.Context, e domain.User) (domain.User, error) {
	var _ *firestore.CollectionRef = einar.Collection("INSERT_YOUR_COLLECTION_CONSTANT_HERE")
	return domain.User{}, nil
}
