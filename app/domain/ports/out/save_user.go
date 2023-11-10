package out

import (
	"context"
	"my-project-name/app/domain"
)

type SaveUser func(ctx context.Context, e domain.User) error
