package cuda
import "errors"
var nilResultError                 =  errors.New("Result is nil")
var targetTypeHasNotBeenSpecified  =  errors.New("cuda:Target type has not been specified")
var pathHasNotBeenSpecified        =  errors.New("cuda:Path has not been specified for this type of target(file or directory)")
var pathHaveToBeEmpty              =  errors.New("cuda:Path have to be empty for line target")
var cantCreateNewTarget            =  errors.New("cuda:Can't create target with provided config set")
var cantAddLineForThisTypeOfTarget =  errors.New("cuda:Can't add line for this type of target")
var targetWasNotConfigured         =  errors.New("cuda:Target was not configured")
var lineIsNil                      =  errors.New("cuda:Line is nil")

