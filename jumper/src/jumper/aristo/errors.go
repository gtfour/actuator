package aristo

import "errors"

var base_word                = "aristo"

func errwrap(in string)(string) {
    return base_word+":"+in
}

var access_allowed           = errors.New(errwrap("access_allowed"))
var access_denied            = errors.New(errwrap("access_denied"))

var id_isnot_specified       = errors.New(errwrap("id is'not specified"))
var id_change_is_not_allowed = errors.New(errwrap("changing id is not allowed"))
var prop_is_empty            = errors.New(errwrap("prop is empty"))
var group_list_is_empty      = errors.New(errwrap("group list is empty"))

var group_invalid            = errors.New(errwrap("group is invalid"))
var member_invalid           = errors.New(errwrap("member is invalid"))


