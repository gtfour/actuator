package dusk

import "errors"

func TokenDoesNotExist()(err error) {
    return errors.New("token_id or user_id is invalid")
}
func DashboardDoesNotExist()(err error) {
    return errors.New("dashboard_id is invalid")
}
