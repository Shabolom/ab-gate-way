package utils

import (
	"gate-way/internal/dto"
	"gate-way/pkg/shortcut"
)

func ValidateTokens(tokens *dto.Tokens) error {
	if tokens == nil {
		return shortcut.ErrTokenNotFilled
	}

	switch {
	case tokens.AccessToken == "":
		return shortcut.ErrTokenNotFilled

	case tokens.RefreshToken == "":
		return shortcut.ErrTokenNotFilled
	}

	return nil
}
