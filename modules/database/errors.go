package database

import (
	"errors"
	"strings"

	"github.com/aldiset/purchase-order/business"

	"gorm.io/gorm"
)

func WrapError(err error) error {
	if strings.Contains(err.Error(), "Duplicate entry") {
		return &business.ErrDuplicateEntry{}

	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return &business.ErrNotFound{}

	} else {
		return err
	}
}
