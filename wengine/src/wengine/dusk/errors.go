package dusk

import "errors"

func TokenDoesNotExist()(err error) {
    return errors.New("token_id or user_id is invalid")
}
func DashboardDoesNotExist()(err error) {
    return errors.New("dashboard_id is invalid")
}
func DashboardGroupDoesNotExist()(err error) {
    return errors.New("dashboardgroup_id is invalid")
}

var incorrect_query_type          = errors.New("Incorrect query type")
var dashboard_group_doesnt_exist  = errors.New("Dashboard group does'nt exist")
var dashboard_doesnt_exist        = errors.New("Group does'nt exist")
var token_doesnt_exist            = errors.New("Token does'nt exist")

var empty_key                     = errors.New("Key is empty")
var empty_query                   = errors.New("Query is empty")




