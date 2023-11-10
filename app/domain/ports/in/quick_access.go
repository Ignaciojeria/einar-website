package in

import (
	"context"
	"my-project-name/app/domain"
)

type QuickAccess func(ctx context.Context, e domain.User) error
