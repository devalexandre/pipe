package v1

import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
)

// Pipeline agora retorna um resultado do tipo interface{} e um error.
type Pipeline func(args ...interface{}) (interface{}, error)

// Pipe foi ajustada para suportar o retorno de um resultado e um erro.
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

			// Prepara inputs para a próxima função, excluindo erros.
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
	// Verifica se target é um ponteiro.
	if reflect.TypeOf(target).Kind() != reflect.Ptr {
		return errors.New("target deve ser um ponteiro")
	}

	// Obtém o valor do resultado e o valor do target.
	resultValue := reflect.ValueOf(result)
	targetValue := reflect.ValueOf(target).Elem()

	// Verifica se o tipo do resultado é conversível para o tipo do target.
	if !resultValue.Type().ConvertibleTo(targetValue.Type()) {
		return fmt.Errorf("não é possível converter o resultado do tipo %s para o tipo %s",
			resultValue.Type(), targetValue.Type())
	}

	// Converte o resultado para o tipo do target e define o valor do target.
	convertedValue := resultValue.Convert(targetValue.Type())
	targetValue.Set(convertedValue)

	return nil
}
