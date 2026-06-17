package auth

import (
	"context"
	"gate-way/internal/dto"
	authMicroserviceDto "gate-way/internal/dto/auth-microservice-dto"
	"gate-way/pkg/shortcut"

	"go.uber.org/zap"
)

func (s *Service) UpdateUser(ctx context.Context, updateUser *authMicroserviceDto.UpdateUser) (*dto.CommonResponse, error) {
	if updateUser == nil {
		s.logger.Warn("update user validation failed: request is nil")
		return nil, shortcut.ErrInvalidRequest
	}

	s.logger.Info(
		"update user started",
		zap.String("mail", updateUser.Mail),
		zap.String("name", updateUser.Name),
		zap.Int("age", updateUser.Age),
	)

	switch {
	case updateUser.Mail == "":
		s.logger.Warn("update user validation failed: mail is empty")
		return nil, shortcut.ErrFieldNotFilledMail

	case updateUser.Name == "":
		s.logger.Warn(
			"update user validation failed: name is empty",
			zap.String("mail", updateUser.Mail),
		)
		return nil, shortcut.ErrFieldNotFilledName

	case updateUser.Age <= 13:
		s.logger.Warn(
			"update user validation failed: age is not allowed",
			zap.String("mail", updateUser.Mail),
			zap.Int("age", updateUser.Age),
		)
		return nil, shortcut.ErrNotAllowedAge
	}

	response, err := s.authAdapter.UpdateUsers(ctx, updateUser)
	if err != nil {
		s.logger.Warn(
			"update user failed",
			zap.String("mail", updateUser.Mail),
			zap.Error(err),
		)
		return nil, err
	}

	s.logger.Info(
		"update user successful",
		zap.String("mail", updateUser.Mail),
	)

	return response, nil
}
