package abAdapter

import (
	authv1 "gate-way/gen/proto"
	"gate-way/pkg/shortcut"
)

func stockReplyErr(reason authv1.StockReply_ERR_INFO_REASON) error {
	switch reason {
	case authv1.StockReply_STATUS_OK:
		return nil
	case authv1.StockReply_VALIDATION_ERROR:
		return shortcut.ErrABValidation
	case authv1.StockReply_INVALID_REQUEST:
		return shortcut.ErrABInvalidRequest
	case authv1.StockReply_UNSPECIFIED:
		return shortcut.ErrABUnspecified
	default:
		return shortcut.ErrABUnspecified
	}
}

func experimentsReplyErr(reason authv1.ExperimentsReply_ERR_INFO_REASON) error {
	switch reason {
	case authv1.ExperimentsReply_STATUS_OK:
		return nil
	case authv1.ExperimentsReply_VALIDATION_ERROR:
		return shortcut.ErrABValidation
	case authv1.ExperimentsReply_INVALID_REQUEST:
		return shortcut.ErrABInvalidRequest
	case authv1.ExperimentsReply_UNSPECIFIED:
		return shortcut.ErrABUnspecified
	default:
		return shortcut.ErrABUnspecified
	}
}

func isUserInFeatureReplyErr(reason authv1.IsUserInFeatureReply_ERR_INFO_REASON) error {
	switch reason {
	case authv1.IsUserInFeatureReply_STATUS_OK:
		return nil
	case authv1.IsUserInFeatureReply_VALIDATION_ERROR:
		return shortcut.ErrABValidation
	case authv1.IsUserInFeatureReply_INVALID_REQUEST:
		return shortcut.ErrABInvalidRequest
	case authv1.IsUserInFeatureReply_UNSPECIFIED:
		return shortcut.ErrABUnspecified
	default:
		return shortcut.ErrABUnspecified
	}
}

func getExperimentErr(reason authv1.GetExperimentByIDReply_ERR_INFO_REASON) error {
	switch reason {
	case authv1.GetExperimentByIDReply_STATUS_OK:
		return nil
	case authv1.GetExperimentByIDReply_VALIDATION_ERROR:
		return shortcut.ErrABValidation
	case authv1.GetExperimentByIDReply_INVALID_REQUEST:
		return shortcut.ErrABInvalidRequest
	case authv1.GetExperimentByIDReply_UNSPECIFIED:
		return shortcut.ErrABUnspecified
	default:
		return shortcut.ErrABUnspecified
	}
}

func getNamespaceErr(reason authv1.GetNamespaceByIDReply_ERR_INFO_REASON) error {
	switch reason {
	case authv1.GetNamespaceByIDReply_STATUS_OK:
		return nil
	case authv1.GetNamespaceByIDReply_VALIDATION_ERROR:
		return shortcut.ErrABValidation
	case authv1.GetNamespaceByIDReply_INVALID_REQUEST:
		return shortcut.ErrABInvalidRequest
	case authv1.GetNamespaceByIDReply_UNSPECIFIED:
		return shortcut.ErrABUnspecified
	default:
		return shortcut.ErrABUnspecified
	}
}

func getLayerErr(reason authv1.GetLayerByIDReply_ERR_INFO_REASON) error {
	switch reason {
	case authv1.GetLayerByIDReply_STATUS_OK:
		return nil
	case authv1.GetLayerByIDReply_VALIDATION_ERROR:
		return shortcut.ErrABValidation
	case authv1.GetLayerByIDReply_INVALID_REQUEST:
		return shortcut.ErrABInvalidRequest
	case authv1.GetLayerByIDReply_UNSPECIFIED:
		return shortcut.ErrABUnspecified
	default:
		return shortcut.ErrABUnspecified
	}
}

func getCustomParamErr(reason authv1.GetCustomParamByIDReply_ERR_INFO_REASON) error {
	switch reason {
	case authv1.GetCustomParamByIDReply_STATUS_OK:
		return nil
	case authv1.GetCustomParamByIDReply_VALIDATION_ERROR:
		return shortcut.ErrABValidation
	case authv1.GetCustomParamByIDReply_INVALID_REQUEST:
		return shortcut.ErrABInvalidRequest
	case authv1.GetCustomParamByIDReply_UNSPECIFIED:
		return shortcut.ErrABUnspecified
	default:
		return shortcut.ErrABUnspecified
	}
}

func getFeatureToggleErr(reason authv1.GetFeatureToggleByIDReply_ERR_INFO_REASON) error {
	switch reason {
	case authv1.GetFeatureToggleByIDReply_STATUS_OK:
		return nil
	case authv1.GetFeatureToggleByIDReply_VALIDATION_ERROR:
		return shortcut.ErrABValidation
	case authv1.GetFeatureToggleByIDReply_INVALID_REQUEST:
		return shortcut.ErrABInvalidRequest
	case authv1.GetFeatureToggleByIDReply_UNSPECIFIED:
		return shortcut.ErrABUnspecified
	default:
		return shortcut.ErrABUnspecified
	}
}

func getExperimentsErr(reason authv1.GetExperimentsReply_ERR_INFO_REASON) error {
	switch reason {
	case authv1.GetExperimentsReply_STATUS_OK:
		return nil
	case authv1.GetExperimentsReply_VALIDATION_ERROR:
		return shortcut.ErrABValidation
	case authv1.GetExperimentsReply_INVALID_REQUEST:
		return shortcut.ErrABInvalidRequest
	case authv1.GetExperimentsReply_UNSPECIFIED:
		return shortcut.ErrABUnspecified
	default:
		return shortcut.ErrABUnspecified
	}
}

func getNamespacesErr(reason authv1.GetNamespaceBysReply_ERR_INFO_REASON) error {
	switch reason {
	case authv1.GetNamespaceBysReply_STATUS_OK:
		return nil
	case authv1.GetNamespaceBysReply_VALIDATION_ERROR:
		return shortcut.ErrABValidation
	case authv1.GetNamespaceBysReply_INVALID_REQUEST:
		return shortcut.ErrABInvalidRequest
	case authv1.GetNamespaceBysReply_UNSPECIFIED:
		return shortcut.ErrABUnspecified
	default:
		return shortcut.ErrABUnspecified
	}
}

func getLayersErr(reason authv1.GetLayerBysReply_ERR_INFO_REASON) error {
	switch reason {
	case authv1.GetLayerBysReply_STATUS_OK:
		return nil
	case authv1.GetLayerBysReply_VALIDATION_ERROR:
		return shortcut.ErrABValidation
	case authv1.GetLayerBysReply_INVALID_REQUEST:
		return shortcut.ErrABInvalidRequest
	case authv1.GetLayerBysReply_UNSPECIFIED:
		return shortcut.ErrABUnspecified
	default:
		return shortcut.ErrABUnspecified
	}
}

func getCustomParamsErr(reason authv1.GetCustomParamBysReply_ERR_INFO_REASON) error {
	switch reason {
	case authv1.GetCustomParamBysReply_STATUS_OK:
		return nil
	case authv1.GetCustomParamBysReply_VALIDATION_ERROR:
		return shortcut.ErrABValidation
	case authv1.GetCustomParamBysReply_INVALID_REQUEST:
		return shortcut.ErrABInvalidRequest
	case authv1.GetCustomParamBysReply_UNSPECIFIED:
		return shortcut.ErrABUnspecified
	default:
		return shortcut.ErrABUnspecified
	}
}

func getFeatureTogglesErr(reason authv1.GetFeatureToggleBysReply_ERR_INFO_REASON) error {
	switch reason {
	case authv1.GetFeatureToggleBysReply_STATUS_OK:
		return nil
	case authv1.GetFeatureToggleBysReply_VALIDATION_ERROR:
		return shortcut.ErrABValidation
	case authv1.GetFeatureToggleBysReply_INVALID_REQUEST:
		return shortcut.ErrABInvalidRequest
	case authv1.GetFeatureToggleBysReply_UNSPECIFIED:
		return shortcut.ErrABUnspecified
	default:
		return shortcut.ErrABUnspecified
	}
}
