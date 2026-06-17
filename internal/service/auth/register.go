package auth

import (
	"context"
	"gate-way/internal/dto"
	authMicroserviceDto "gate-way/internal/dto/auth-microservice-dto"
	"gate-way/pkg/shortcut"

	"go.uber.org/zap"
)

func (s *Service) Register(ctx context.Context, register *authMicroserviceDto.Register) (*dto.CommonResponse, *dto.Tokens, error) {
	s.logger.Info(
		"register started",
		zap.String("mail", register.Mail),
		zap.String("name", register.Name),
	)

	switch {
	case register.Age <= 13:
		s.logger.Warn(
			"register validation failed: age is not allowed",
			zap.Int("age", register.Age),
			zap.String("mail", register.Mail),
		)
		return nil, nil, shortcut.ErrNotAllowedAge

	case register.Name == "":
		s.logger.Warn(
			"register validation failed: name is empty",
			zap.String("mail", register.Mail),
		)
		return nil, nil, shortcut.ErrFieldNotFilledName

	case register.Mail == "":
		s.logger.Warn("register validation failed: mail is empty")
		return nil, nil, shortcut.ErrFieldNotFilledMail

	case register.Password == "":
		s.logger.Warn(
			"register validation failed: password is empty",
			zap.String("mail", register.Mail),
		)
		return nil, nil, shortcut.ErrFieldNotFilledPassword
	}

	response, tokens, err := s.authAdapter.Register(ctx, register)
	if err != nil {
		s.logger.Warn(
			"register failed",
			zap.String("mail", register.Mail),
			zap.String("name", register.Name),
			zap.Error(err),
		)
		return nil, nil, err
	}

	s.logger.Info(
		"register successful",
		zap.String("mail", register.Mail),
		zap.String("name", register.Name),
	)

	return response, tokens, nil
}
