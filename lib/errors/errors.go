package errors

import (
	"fmt"
)

type CustomError struct {
	Desc   string
	RawErr error
}

func (c CustomError) Error() string {
	if c.RawErr == nil {
		return c.Desc
	}

	return fmt.Sprintf(`%s caused by raw error "%s"`, c.Desc, c.RawErr.Error())
}
