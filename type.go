package bossy

import "strconv"

type config struct {
	// LoadJSON() Error
	// Export()
}

func (c *config) Export() {

}

func atob(in string) bool {
	checkedValues := []string{
		"true",
		"checked",
		"on",
		"yes",
		"y",
		"1",
	}

	for _, each := range checkedValues {
		if each == in {
			return true
		}
	}
	return false
}

func (c *config) changeItem(address []string, loc, val interface{}) Error {
	if len(address) < 1 {
		return NewError{Code: BadAddressLocation}
	}
	switch loc := loc.(type) {
	case map[string]interface{}:
		// no conversion needed
		if len(address) > 1 {
			childLoc, ok := loc[address[0]]
			if !ok {
				return NewError{Code: BadAddressLocation}
			}
			return c.setItem(address[1:], childLoc, val)
		}
		loc[address[0]] = val
	case []interface{}:
		i, err := strconv.Atoi(address[0])
		if err != nil {
			return NewError{Code: ErrBadAddressStructure}
		}
		if i < 0 || i > len(loc)-1 {
			return NewError{Code: ErrBadAddressIndex}
		}
		if len(address) > 1 {
			return c.setItem(address[1:], loc[i], val)
		}
		loc[i] = val
	case map[int]interface{}:
		i, err := strconv.Atoi(address[0])
		if err != nil {
			return NewError{Code: ErrBadAddressStructure}
		}
		childLoc, ok := loc[i]
		if !ok {
			return NewError{Code: ErrBadAddressIndex}
		}
		if len(address) > 1 {
			return c.setItem(address[1:], childLoc, val)
		}
		loc[i] = val
	case map[bool]interface{}:
		v := atob(address[0])
		childLoc, ok := loc[v]
		if !ok {
			return NewError{Code: ErrBadAddressIndex}
		}
		if len(address) > 1 {
			return c.setItem(address[1:], childLoc, val)
		}
		loc[i] = val
	case map[float64]interface{}:
		i, err := strconv.ParseFloat(address[0], 64)
		if err != nil {
			return NewError{Code: ErrBadAddressStructure}
		}
		childLoc, ok := loc[i]
		if !ok {
			return NewError{Code: ErrBadAddressIndex}
		}
		if len(address) > 1 {
			return c.setItem(address[1:], childLoc, val)
		}
		loc[i] = val
	}
	return nil
}

func (c *config) setItem(address []string, loc, val interface{}) Error {
	if len(address) < 1 {
		return NewError{Code: BadAddressLocation}
	}
	switch loc := loc.(type) {
	case map[string]interface{}:
		// no conversion needed
		if len(address) > 1 {
			return c.setItem(address[1:], loc[address[0]], val)
		}
		loc[address[0]] = val
	case []interface{}:
		switch address[0] {
		case "":
			if len(address) > 1 {
				return NewError{Code: ErrUnimplemented}
			}
			loc = append(loc, val)
		case "-1":
			if len(address) > 1 {
				return NewError{Code: ErrUnimplemented}
			}
			loc = append([]interface {
				val
			}, loc...)
		case "*":
			for i := range loc {
				if len(address) > 1 {
					if err := c.SetItem(address[1:], loc[i], val); err != nil {
						return err
					}
				} else {
					loc[i] = val
				}
			}
		default:
			i, err := strconv.Atoi(address[0])
			if err != nil {
				return NewError{Code: ErrBadAddressStructure}
			}
			if i < 0 || i > len(loc) {
				return NewError{Code: ErrBadAddressIndex}
			}

			if i < len(loc)-1 {
				if len(address) > 1 {
					return c.setItem(address[1:], loc[i], val)
				}
				loc[i] = val
			} else {
				if len(address) > 1 {
					// TODO determine setItem behavior for slices
					return NewError{Code: ErrUnimplemented}
				}
				loc = append(loc, val)
			}
		}
	case map[int]interface{}:
		i, err := strconv.Atoi(address[0])
		if err != nil {
			return NewError{Code: ErrBadAddressStructure}
		}
		if len(address) > 1 {
			return c.setItem(address[1:], loc[i], val)
		}
		loc[i] = val
	case map[bool]interface{}:
		v := atob(address[0])
		if len(address) > 1 {
			return c.setItem(address[1:], loc[v], val)
		}
		loc[i] = val
	case map[float64]interface{}:
		n, err := strconv.ParseFloat(address[0], 64)
		if err != nil {
			return NewError{Code: ErrBadAddressStructure}
		}
		if len(address) > 1 {
			return c.setItem(address[1:], loc[n], val)
		}
		loc[n] = val
	default:
		// i need a default case maybe?
		return NewError{Code: ErrUnsupportedType}
	}
	return nil
}

func (c *config) getItem(address []string, loc interface{}) (interface{}, Error) {
	if len(address) < 1 {
		return NewError{Code: BadAddressLocation}
	}
	switch loc := loc.(type) {
	case map[string]interface{}:
		// no conversion needed
		if len(address) > 1 {
			return c.getItem(address[1:], loc[address[0]], val)
		}
		return loc[address[0]], nil
	case []interface{}:
		i, err := strconv.Atoi(address[0])
		if err != nil {
			return NewError{Code: ErrBadAddressStructure}
		}
		if i < 0 || i > len(loc) {
			return NewError{Code: ErrBadAddressIndex}
		}
		if len(address) > 1 {
			return c.getItem(address[1:], loc[i], val)
		}
		return loc[i]
	case map[int]interface{}:
		i, err := strconv.Atoi(address[0])
		if err != nil {
			return nil, NewError{Code: ErrBadAddressStructure}
		}
		if _, exists := loc[i]; !exists {
			return nil, NewErr{Code: ErrBadAddressIndex}
		}
		if len(address) > 1 {
			return c.getItem(address[1:], loc[i], val), nil
		}
		return loc[i], nil
	case map[bool]interface{}:
		v := atob(address[0])
		if _, exists := loc[i]; !exists {
			return nil, NewErr{Code: ErrBadAddressIndex}
		}
		if len(address) > 1 {
			return c.getItem(address[1:], loc[v], val)
		}
		return loc[i], nil
	case map[float64]interface{}:
		n, err := strconv.ParseFloat(address[0], 64)
		if err != nil {
			return NewError{Code: ErrBadAddressStructure}
		}
		if _, exists := loc[i]; !exists {
			return nil, NewErr{Code: ErrBadAddressIndex}
		}
		if len(address) > 1 {
			return c.setItem(address[1:], loc[n], val)
		}
		return loc[n], nil
		loc[n] = val
	}
	return nil, NewError{Code: ErrUnsupportedType}
}
