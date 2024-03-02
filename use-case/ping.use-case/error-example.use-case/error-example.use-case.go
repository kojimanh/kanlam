package error_example_use_case

import "errors"

func ReturnError() (any, error) {
	return nil, errors.New("Example error")
}
