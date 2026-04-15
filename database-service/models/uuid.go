package models

import (
	"database/sql/driver"
	"fmt"
)

// Custom type for UUID to handle database/sql driver
type UUID string

func (u UUID) Value() (driver.Value, error) {
	if u == "" {
		return nil, nil
	}
	return string(u), nil
}

func (u *UUID) Scan(value interface{}) error {
	if value == nil {
		*u = ""
		return nil
	}

	switch v := value.(type) {
	case string:
		*u = UUID(v)
	case []byte:
		*u = UUID(v)
	default:
		return fmt.Errorf("cannot scan %T into UUID", value)
	}
	return nil
}
