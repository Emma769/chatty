package validator

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/emma769/chatty/pkg/funclib"
)

type Validationfn func(reflect.Value, []string, string) error

func requireString(v reflect.Value, _ []string, name string) error {
	val := v.String()

	if !funclib.ValidString(val) {
		return fmt.Errorf("%s cannot be blank", name)
	}

	return nil
}

func minlength(v reflect.Value, args []string, name string) error {
	val := v.String()

	if len(args) != 1 {
		panic(ErrNoArg)
	}

	arg, err := strconv.Atoi(args[0])
	if err != nil {
		panic(err)
	}

	if !funclib.Gte(len(val), arg) {
		return fmt.Errorf("%s must have at least %d characters", name, arg)
	}

	return nil
}

func validEmail(v reflect.Value, _ []string, _ string) error {
	val := v.String()

	if !funclib.ValidEmail(val) {
		return fmt.Errorf("provide a valid email")
	}

	return nil
}
