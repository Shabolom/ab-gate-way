package shortcut

import (
	"errors"
	"log"
)

var (
	ErrNotFound                   = errors.New("not found")
	ErrValidation                 = errors.New("validation error")
	ErrDuplicateKey               = errors.New("duplicate key")
	ErrForeignKeyViolation        = errors.New("foreign key violation")
	ErrCheckViolation             = errors.New("check violation")
	ErrExclusionViolation         = errors.New("exclusion violation")
	ErrUnspecifiedRequest         = errors.New("unspecified error")
	ErrInvalidRequest             = errors.New("invalid response")
	ErrUnspecifiedResponseGetUser = errors.New("unspecified response, nil user")
	ErrNotAllowedAge              = errors.New("not allowed")
	ErrFieldNotFilledName         = errors.New("field not filled name")
	ErrFieldNotFilledMail         = errors.New("field not filled mail")
	ErrFieldNotFilledPassword     = errors.New("field not filled password")
	ErrTokenNotFilled             = errors.New("token not filled")

	ErrABValidation                  = errors.New("ab validation error")
	ErrABInvalidRequest              = errors.New("ab invalid request")
	ErrABUnspecified                 = errors.New("ab unspecified error")
	ErrABEmptyResponse               = errors.New("ab empty response")
	ErrABExperimentNameRequired      = errors.New("ab experiment name required")
	ErrABExperimentLayersRequired    = errors.New("ab experiment layers required")
	ErrABExperimentGroupsMinCount    = errors.New("ab experiment groups min count is 2")
	ErrABExperimentStartDateAfterEnd = errors.New("ab experiment start_date after end_date")
	ErrABExperimentRolloutOutOfRange = errors.New("ab experiment rollout must be between 0 and 100")

	ErrStartDateAfterEnd = errors.New("start date after end")

	ErrABNamespaceNameRequired = errors.New("ab namespace name required")

	ErrABLayerNameRequired      = errors.New("ab layer name required")
	ErrABLayerNamespaceRequired = errors.New("ab layer namespace id required")

	ErrABCustomParamNameRequired      = errors.New("ab custom param name required")
	ErrABCustomParamNamespaceRequired = errors.New("ab custom param namespace id required")
	ErrABCustomParamTypeRequired      = errors.New("ab custom param type required")

	ErrABExperimentIDRequired = errors.New("ab experiment id required")

	ErrABNamespaceRequired = errors.New("ab namespace required")
	ErrABSplitIDRequired   = errors.New("ab split id required")

	ErrAccessTokenExpired  = errors.New("access token expired")
	ErrRefreshTokenExpired = errors.New("refresh token expired")

	ErrAccessTokenNotFound  = errors.New("access token not found")
	ErrRefreshTokenNotFound = errors.New("refresh token not found")

	ErrAccessTokenRevoked  = errors.New("access token revoked")
	ErrRefreshTokenRevoked = errors.New("refresh token revoked")

	ErrInvalidAccessToken  = errors.New("invalid access token")
	ErrInvalidRefreshToken = errors.New("invalid refresh token")

	ErrTokenPairMismatch = errors.New("token pair mismatch")
	ErrSessionNotFound   = errors.New("session not found")

	ErrFeatureToggleNameRequired        = errors.New("feature toggle name required")
	ErrFeatureToggleNamespaceIDRequired = errors.New("feature toggle namespace id required")

	ErrFeatureToggleRolloutRequired = errors.New("feature toggle rollout required")

	ErrFeatureToggleIOSRolloutOutOfRange     = errors.New("feature toggle ios rollout out of range")
	ErrFeatureToggleAndroidRolloutOutOfRange = errors.New("feature toggle android rollout out of range")
	ErrFeatureToggleWebRolloutOutOfRange     = errors.New("feature toggle web rollout out of range")
	ErrFeatureToggleRolloutOutOfRange        = errors.New("feature toggle rollout out of range")
	ErrFeatureToggleID                       = errors.New("feature toggle ID validate error")
	ErrFeatureNamespaceMissing               = errors.New("feature namespace missing")
	ErrFeatureUserIDMissing                  = errors.New("feature user id missing")
	ErrFeatureStatusMissing                  = errors.New("feature status missing")
)

func FatalIfErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
