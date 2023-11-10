package firestore

import (
	"context"
	"my-project-name/app/domain"
	"my-project-name/app/domain/ports/out"
	einar "my-project-name/app/shared/archetype/firestore"

	"cloud.google.com/go/firestore"
)

var SaveUser out.SaveUser = func(ctx context.Context, e domain.User) error {
	var _ *firestore.CollectionRef = einar.Collection("INSERT_YOUR_COLLECTION_CONSTANT_HERE")
	return nil
}
