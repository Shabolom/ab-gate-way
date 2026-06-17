package auth

import (
	"context"
	"gate-way/internal/dto"
	authMicroserviceDto "gate-way/internal/dto/auth-microservice-dto"
	"gate-way/pkg/shortcut"

	"go.uber.org/zap"
)

func (s *Service) Login(ctx context.Context, login *authMicroserviceDto.Login) (*dto.CommonResponse, *dto.Tokens, error) {
	s.logger.Info(
		"login started",
		zap.String("mail", login.Mail),
	)

	switch {
	case login.Mail == "":
		s.logger.Warn("login validation failed: mail is empty")
		return nil, nil, shortcut.ErrFieldNotFilledMail

	case login.Password == "":
		s.logger.Warn(
			"login validation failed: password is empty",
			zap.String("mail", login.Mail),
		)
		return nil, nil, shortcut.ErrFieldNotFilledPassword
	}

	response, tokens, err := s.authAdapter.Login(ctx, login)
	if err != nil {
		s.logger.Warn(
			"login failed",
			zap.String("mail", login.Mail),
			zap.Error(err),
		)
		return nil, nil, err
	}

	s.logger.Info(
		"login successful",
		zap.String("mail", login.Mail),
	)

	return response, tokens, nil
}
