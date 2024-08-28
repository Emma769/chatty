package validator

import (
	"errors"
	"reflect"
	"strings"
	"sync"
)

var (
	ErrUnknownValidationFn = errors.New("unknown validation function")
	ErrNoArg               = errors.New("no arg")
)

type Validator struct {
	mu    *sync.Mutex
	rules map[string]Validationfn
	errs  map[string]string
}

func New() *Validator {
	rules := map[string]Validationfn{
		"required": requireString,
		"min":      minlength,
		"email":    validEmail,
	}

	return &Validator{
		&sync.Mutex{},
		rules,
		map[string]string{},
	}
}

func (v Validator) valid() bool {
	return len(v.errs) == 0
}

func (v *Validator) add(key, msg string) {
	v.mu.Lock()
	defer v.mu.Unlock()

	if _, ok := v.errs[key]; !ok {
		v.errs[key] = msg
	}
}

func (validator *Validator) ValidateStruct(v any) map[string]string {
	val := reflect.ValueOf(v)

	for i := 0; i < val.NumField(); i++ {
		value := val.Field(i)
		validationTags := val.Type().Field(i).Tag.Get("validate")
		jsonTag := val.Type().Field(i).Tag.Get("json")

		if validationTags == "" {
			continue
		}

		validationRules := strings.Split(validationTags, ",")

		for _, validationRule := range validationRules {
			parts := strings.Split(validationRule, "=")
			fnName, args := parts[0], parts[1:]
			validator.apply(value, fnName, args, jsonTag)
		}
	}

	if !validator.valid() {
		return validator.errs
	}

	return nil
}

func (validator *Validator) apply(v reflect.Value, fnName string, args []string, name string) {
	fn, ok := validator.rules[fnName]

	if !ok {
		panic(ErrUnknownValidationFn)
	}

	if err := fn(v, args, name); err != nil {
		validator.add(name, err.Error())
	}
}
