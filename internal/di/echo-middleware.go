package di

import (
	"context"
	"net/http"

	"gate-way/internal/render"
	"gate-way/pkg/shortcut"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

const refreshTokenHeader = "Refresh-Token"

func (d *DI) AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get(echo.HeaderAuthorization)
		refreshToken := c.Request().Header.Get(refreshTokenHeader)

		ctx := metadata.AppendToOutgoingContext(
			c.Request().Context(),
			"client-ip", c.RealIP(),
			"user-agent", c.Request().UserAgent(),
		)

		d.Logger().Info(
			"auth middleware route check",
			zap.String("method", c.Request().Method),
			zap.String("path", c.Path()),
			zap.String("url_path", c.Request().URL.Path),
		)

		if isPublicRoute(c) {
			c.SetRequest(c.Request().WithContext(ctx))
			return next(c)
		}

		if token == "" {
			d.Logger().Warn(
				"auth middleware failed: token is empty",
				zap.String("method", c.Request().Method),
				zap.String("path", c.Path()),
			)

			return render.BadRequest(c, shortcut.ErrAccessTokenNotFound)
		}

		if refreshToken == "" {
			d.Logger().Warn(
				"auth middleware failed: refresh-token is empty",
				zap.String("method", c.Request().Method),
				zap.String("path", c.Path()),
			)

			return render.BadRequest(c, shortcut.ErrRefreshTokenNotFound)
		}

		ctx = withTokensMetadata(ctx, token, refreshToken)
		c.SetRequest(c.Request().WithContext(ctx))

		ctx, err := d.GetAuthService().Check(ctx)
		if err != nil {
			d.Logger().Warn(
				"auth middleware failed: token check failed",
				zap.String("method", c.Request().Method),
				zap.String("path", c.Path()),
				zap.Error(err),
			)

			return render.FromError(c, shortcut.ErrInvalidAccessToken)
		}

		c.SetRequest(c.Request().WithContext(ctx))
		setTokensToHTTPResponse(c)

		d.Logger().Info(
			"auth middleware passed",
			zap.String("method", c.Request().Method),
			zap.String("path", c.Path()),
		)

		err = next(c)

		return err
	}
}

func isPublicRoute(c echo.Context) bool {
	path := c.Path()
	method := c.Request().Method

	switch {
	case method == http.MethodPost && path == "/v1/auth/login":
		return true

	case method == http.MethodPost && path == "/v1/auth/register":
		return true

	default:
		return false
	}
}

func withTokensMetadata(ctx context.Context, accessToken, refreshToken string) context.Context {
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		md = metadata.New(nil)
	} else {
		md = md.Copy()
	}

	md.Set("authorization", accessToken)
	md.Set("refresh-token", refreshToken)

	return metadata.NewOutgoingContext(ctx, md)
}

func setTokensToHTTPResponse(c echo.Context) {
	md, ok := metadata.FromOutgoingContext(c.Request().Context())
	if !ok {
		return
	}

	accessTokens := md.Get("authorization")
	if len(accessTokens) > 0 && accessTokens[0] != "" {
		c.Response().Header().Set(echo.HeaderAuthorization, accessTokens[0])
	}

	refreshTokens := md.Get("refresh-token")
	if len(refreshTokens) > 0 && refreshTokens[0] != "" {
		c.Response().Header().Set(refreshTokenHeader, refreshTokens[0])
	}
}
