package cuda
import "errors"
var nilResultError                = errors.New("Result is nil")
var targetTypeHasNotBeenSpecified = errors.New("cuda:Target type has not been specified")
var pathHasNotBeenSpecified       = errors.New("cuda:Path has not been specified for this type of target(file or directory)")
var pathHaveToBeEmpty             = errors.New("cuda:Path have to be empty for line target")
var cantCreateNewTarget           = errors.New("cuda:Can't create target with provided config set")

