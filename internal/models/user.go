package models

import (
	"fmt"
	"time"
)

type User struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	Email     string   `json:"email"`
	CreatedAt JSONTime `json:"created_at"`
	UpdatedAt JSONTime `json:"updated_at"`
}

type JSONTime time.Time

func (t JSONTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

func (t *JSONTime) Scan(value interface{}) error {
	if value == nil {
		*t = JSONTime(time.Time{})
		return nil
	}

	switch v := value.(type) {
	case time.Time:
		*t = JSONTime(v)
		return nil
	case []byte: // если строка в байтах
		tt, err := time.Parse("2006-01-02 15:04:05", string(v))
		if err != nil {
			return err
		}
		*t = JSONTime(tt)
		return nil
	case string: // если просто строка
		tt, err := time.Parse("2006-01-02 15:04:05", v)
		if err != nil {
			return err
		}
		*t = JSONTime(tt)
		return nil
	default:
		return fmt.Errorf("unsupported type %T for JSONTime", value)
	}
}
