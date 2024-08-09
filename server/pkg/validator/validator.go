package validator

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/emma769/chatty/pkg/funclib"
)

func ValidateStruct(v any) error {
	val := reflect.ValueOf(v)

	for i := 0; i < val.NumField(); i++ {
		fieldvalue := val.Field(i)

		validateTags := val.Type().Field(i).Tag.Get("validate")
		jsonTag := val.Type().Field(i).Tag.Get("json")

		if validateTags == "" {
			continue
		}

		rules := strings.Split(validateTags, ",")

		for _, rule := range rules {
			parts := strings.Split(rule, "=")
			validationfn, args := parts[0], parts[1:]

			switch validationfn {
			case "required":
				if !funclib.ValidString(fieldvalue.String()) {
					return fmt.Errorf("%s cannot be blank", jsonTag)
				}
			case "email":
				if !funclib.ValidEmail(fieldvalue.String()) {
					return errors.New("provide a valid email")
				}
			case "min":
				if len(args) == 0 {
					panic("invalid validation function usage: <min=[value]>")
				}

				arg, _ := strconv.Atoi(args[0])
				if !funclib.Gte(len(fieldvalue.String()), arg) {
					return fmt.Errorf("%s cannot be less than %d", jsonTag, arg)
				}
			}
		}
	}

	return nil
}
