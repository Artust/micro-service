package err

import (
	"fmt"

	"github.com/pkg/errors"
)

const ERR_NO_RECORD = "Result contains no more records"

func RecoverPanic(function func() error) func() error {
	return func() error {
		defer func() {
			if err, ok := recover().(error); ok {
				if err != nil {
					if ok {
						fmt.Println(errors.Wrapf(err, "unexpected panic"))
					} else {
						fmt.Println(errors.Errorf("unexpected panic: %+v", recover()))
					}
				}
			}
		}()
		return function()
	}
}
