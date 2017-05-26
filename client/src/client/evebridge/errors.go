package evebridge
import "errors"

var unrecognizable_task_state = errors.New("\nErr:Failed to recognize task state")
var targetTypeUndefined       = errors.New("\nErr:Target type is undefined")
var appWasNotConfigured       = errors.New("\nErr:App was not configured")

