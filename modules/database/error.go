package database

import (
	"errors"
	"strings"

	"github.com/alitdarmaputra/belanja-project/bussiness"
	"gorm.io/gorm"
)

func WrapError(err error) error {
	if err == nil {
		return nil
	}

	if strings.Contains(err.Error(), "Duplicate entry") {
		return bussiness.NewDuplicateEntryError("Duplicate entry")
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return bussiness.NewNotFoundError("Data not found")
	}

	return err
}
