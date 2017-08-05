package cross

var CHECK_EXIST             = 4000

// var CHECK_TABLE_EXIST       = 4002

var CREATE_NEW              = 4004
var CREATE_NEW_IFNOT        = 4006

// var CREATE_NEW_TABLE        = 4008

var EDIT                    = 4010
var GET                     = 4012
var GET_ALL                 = 4014
var INSERT_ITEM             = 4016

var REMOVE                  = 4018
var DELETE                  = 4018
var REMOVE_IFEXIST          = 4019
var DELETE_IFEXIST          = 4019
var REMOVE_ITEM             = 4020
var DELETE_ITEM             = 4020

// var REMOVE_TABLE            = 4022

var REPLACE                 = 4024
var UPDATE                  = 4026
var MODIFY                  = 4026

// 

var APPEND_TO_LIST          = 4102
var REMOVE_FROM_LIST        = 4106


// working over nested hashes

var ADD_PAIR           = 4108 // this key-value pair  can't be present in hash marked by this key.???    
var REMOVE_PAIR        = 4110 // 
var GET_PAIR           = 4112 //
var APPEND_TO_SLICE    = 4202 // пиздешь!  if pair value is array
var REMOVE_FROM_SLICE  = 4203 // if pair value is array
var GET_SLICE_ELEM     = 4204
var CREATE_EMPTY_ARRAY = 4205
var GET_ARRAY          = 4206

var CREATE_NEW_TABLE                   = 5000
var CREATE_NEW_TABLE_IF_DOESNT_EXIST   = 5001
var CHECK_TABLE_EXIST                  = 5002
var REMOVE_TABLE                       = 5003
var DELETE_TABLE                       = 5003

var TABLE_SIZE                         = 5004

// dumm shit

var ARRAY_TYPE_STRING                  = 6000
var ARRAY_TYPE_INT                     = 6002
var ARRAY_TYPE_BOOL                    = 6004

//

var TABLE_OPS = []int { CREATE_NEW_TABLE, CREATE_NEW_TABLE_IF_DOESNT_EXIST , CHECK_TABLE_EXIST , REMOVE_TABLE , DELETE_TABLE }

