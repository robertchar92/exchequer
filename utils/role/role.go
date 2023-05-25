package role

import (
	"reflect"

	"github.com/liip/sheriff"
	"gitlab.com/depatu/core/utils/errors"
)

var TEST_MODE = false

const (
	User      = "user"
	Admin     = "admin"
	Retail    = "retail"
	Assistant = "assistant"
)

func GetDataJSONByRole(data interface{}, groups ...string) (interface{}, error) {
	kind := reflect.ValueOf(data).Kind()
	if TEST_MODE {
		if data == nil ||
			((kind == reflect.Map ||
				kind == reflect.Slice ||
				kind == reflect.Ptr) &&
				reflect.ValueOf(data).IsNil()) {
			return nil, errors.ErrUnprocessableEntity
		}
	} else {
		if data == nil ||
			((kind == reflect.Ptr) &&
				reflect.ValueOf(data).IsNil()) {
			return nil, errors.ErrUnprocessableEntity
		}
	}

	o := &sheriff.Options{
		Groups: groups,
	}

	filtered, err := sheriff.Marshal(o, data)
	if err != nil {
		return nil, err
	}

	return filtered, nil
}
