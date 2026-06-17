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
