package auth

import (
	"context"
)

func (s *Service) Check(ctx context.Context) (context.Context, error) {
	return s.authAdapter.Checker(ctx)
}
