package bossy

import "fmt"

// Err - the type used within my application for error handling
type Error struct {
	Code     int
	path     string
	valType  string
	value    interface{}
	innerErr error
}

// Upgrades an external error making it a error
func UpgradeError(e error) Error {
	return Error{Code: ErrUpgradedError, deepErr: e}
}

// the function `Error` to make my custom errors work
func (e NewErr) Error() string {
	switch {
	case e.Code == ErrUpgradedError:
		return e.innerError.Error()
	case e.path != "" && e.valType != "" && e.value != nil && e.innerErr != nil:
		return fmt.Sprintf(errMsg[e.Code], e.path, e.valType, e.value, e.innerErr)
	case e.path != "" && e.valType != "" && e.value != nil:
		return fmt.Sprintf(errMsg[e.Code], e.path, e.valType, e.value)
	case e.path != "" && e.valType != "" && e.innerErr != nil:
		return fmt.Sprintf(errMsg[e.Code], e.path, e.valType, e.innerErr)
	case e.path != "" && e.value != nil && e.innerErr != nil:
		return fmt.Sprintf(errMsg[e.Code], e.path, e.value, e.innerErr)
	case e.valType != "" && e.value != nil && e.innerErr != nil:
		return fmt.Sprintf(errMsg[e.Code], e.valType, e.value, e.innerErr)
	case e.path != "" && e.valType != "":
		return fmt.Sprintf(errMsg[e.Code], e.path, e.valType)
	case e.path != "" && e.value != nil:
		return fmt.Sprintf(errMsg[e.Code], e.path, e.value)
	case e.valType != "" && e.value != nil:
		return fmt.Sprintf(errMsg[e.Code], e.valType, e.value)
	case e.innerErr != nil:
		return fmt.Sprintf(errMsg[e.Code], e.innerErr)
	case e.value != nil:
		return fmt.Sprintf(errMsg[e.Code], e.value)
	case e.valType != "":
		return fmt.Sprintf(errMsg[e.Code], e.valType)
	case e.path != "":
		return fmt.Sprintf(errMsg[e.Code], e.path)
	default:
		return errMsg[e.Code]
	}
}

// assign a unique id to each error
const (
	ErrUpgradedError = 1 << iota
	ErrBadType
)

// specify the error message for each error
var errMsg = map[int]string{
	ErrUpgradedError: "nothing to see here",
	ErrBadType:       "value at address [%s] is of the wrong type [%s]",
}
