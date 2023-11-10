package out

import (
	"context"
	"my-project-name/app/domain"
)

type FindUserByEmail func(ctx context.Context, e domain.User) (domain.User, error)
