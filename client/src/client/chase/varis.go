package chase

//
import "time"
// //
var TGT_PER_GR int64                       = 50 // if FILES_PER_GR is very big - TargetsCount type should be modified 
var TIMEOUT_MS              time.Duration  = 800
var INHIBITION_TIMEOUT      time.Duration  = 1000
var LOG_CHANNEL_TIMEOUT_MS  time.Duration  = 1000
// //
var EMPTY_OPENING_MODE int = 0
var LAZY_OPENING_MODE  int = 01
var SAFE_OPENING_MODE  int = 02

