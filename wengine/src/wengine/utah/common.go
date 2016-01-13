package utah

var ACTION_PERM_CREATE    int = 1000
var ACTION_PERM_EDIT      int = 1001
var ACTION_PERM_DELETE    int = 1002
var ACTION_PERM_RUN       int = 1003

var FEATURE_PERM_CREATE   int = 1010
var FEATURE_PERM_EDIT     int = 1011
var FEATURE_PERM_DELETE   int = 1012

var TRIGGER_PERM_CREATE   int = 1020
var TRIGGER_PERM_EDIT     int = 1021
var TRIGGER_PERM_DELETE   int = 1022

var DASHBOARD_PERM_CREATE int = 1030
var DASHBOARD_PERM_EDIT   int = 1031
var DASHBOARD_PERM_DELETE int = 1032




type User struct {

    Id         string
    Name       string
    SecondName string
    Password   string
    Admin      bool
    Perms      []int

}

type Group struct {

    Id   string
    Name string

}

type Token struct {

    Id             string
    UserId         string
    ExpirationTime string
    CreatedAt      string
    ExpiredAt      string

}


func ( u *User ) CheckPerm (int)(bool) {

    return true

}


