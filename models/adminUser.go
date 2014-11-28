package models

import (
	"fmt"
)

type adminCollection struct{}

var AdminCollection adminCollection

func (adminCollection) ValidateUser(username string, password string) (int, error) {
	if username == "fail" {
		return 0, fmt.Errorf("123")
	}
	return 0, nil
}
