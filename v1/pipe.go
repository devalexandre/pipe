package v1

import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
)

type Pipeline func(args ...interface{}) (interface{}, error)

func Pipe(fs ...interface{}) Pipeline {
	return func(args ...interface{}) (interface{}, error) {
		var result interface{}
		var err error

		inputs := make([]reflect.Value, len(args))
		for i, arg := range args {
			inputs[i] = reflect.ValueOf(arg)
		}

		for i, f := range fs {
			funcValue := reflect.ValueOf(f)
			funcType := funcValue.Type()

			if len(inputs) != funcType.NumIn() {
				funcName := runtime.FuncForPC(funcValue.Pointer()).Name()
				return nil, fmt.Errorf("número incorreto de argumentos para a função: %s", funcName)
			}

			outputs := funcValue.Call(inputs)

			if len(outputs) > 0 {
				lastOutput := outputs[len(outputs)-1]
				if lastOutput.Type().Implements(reflect.TypeOf((*error)(nil)).Elem()) && !lastOutput.IsNil() {
					return nil, lastOutput.Interface().(error)
				}
				if i == len(fs)-1 && len(outputs) > 0 {
					result = outputs[0].Interface()
				}
			}

			if len(outputs) > 1 {
				inputs = outputs[:len(outputs)-1]
			} else {
				inputs = []reflect.Value{}
			}
		}

		return result, err
	}
}

func ParseTo(result interface{}, target interface{}) error {

	if reflect.TypeOf(target).Kind() != reflect.Ptr {
		return errors.New("target deve ser um ponteiro")
	}

	resultValue := reflect.ValueOf(result)
	targetValue := reflect.ValueOf(target).Elem()

	if !resultValue.Type().ConvertibleTo(targetValue.Type()) {
		return fmt.Errorf("não é possível converter o resultado do tipo %s para o tipo %s",
			resultValue.Type(), targetValue.Type())
	}

	convertedValue := resultValue.Convert(targetValue.Type())
	targetValue.Set(convertedValue)

	return nil
}
