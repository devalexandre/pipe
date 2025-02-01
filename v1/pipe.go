package v1

import (
	"fmt"
	"reflect"
)

// Pipeline define um pipeline que aceita argumentos iniciais e retorna um resultado ou erro.
type Pipeline func(args ...interface{}) (interface{}, error)

// Pipe cria um pipeline a partir de uma série de funções.
// Cada função deve ter uma assinatura compatível com o encadeamento:
//   - Seus parâmetros serão preenchidos com os resultados da função anterior.
//   - Se uma função retornar um error não-nulo, o pipeline é interrompido.
func Pipe(fs ...interface{}) Pipeline {
	// Valida se todos os elementos são funções.
	for i, f := range fs {
		if reflect.TypeOf(f).Kind() != reflect.Func {
			panic(fmt.Sprintf("elemento na posição %d não é uma função", i))
		}
	}

	// Obter o tipo que representa o interface error (para evitar repetir esse cálculo).
	errorType := reflect.TypeOf((*error)(nil)).Elem()

	return func(initialArgs ...interface{}) (interface{}, error) {
		currentArgs := initialArgs

		// Processa cada função do pipeline.
		for idx, f := range fs {
			fnVal := reflect.ValueOf(f)
			fnType := fnVal.Type()
			numIn := fnType.NumIn()

			// Prepara os argumentos para a função atual.
			in := make([]reflect.Value, numIn)
			for i := 0; i < numIn; i++ {
				if i < len(currentArgs) {
					in[i] = reflect.ValueOf(currentArgs[i])
				} else if i < len(initialArgs) { // fallback para os argumentos iniciais, se necessário
					in[i] = reflect.ValueOf(initialArgs[i])
				} else {
					return nil, fmt.Errorf("argumentos insuficientes para a função na posição %d", idx)
				}
			}

			// Chama a função e obtém os resultados.
			results := fnVal.Call(in)

			// Prepara os argumentos para a próxima função, resetando currentArgs.
			currentArgs = make([]interface{}, 0, len(results))
			for _, res := range results {
				// Se o resultado implementa error e não é nil, interrompe o pipeline.
				if res.Type().Implements(errorType) {
					if !res.IsNil() {
						return nil, res.Interface().(error)
					}
					// Caso o error seja nil, não o adiciona aos resultados.
				} else {
					currentArgs = append(currentArgs, res.Interface())
				}
			}
		}

		// Se houver apenas um resultado, retorna-o diretamente; caso contrário, retorna um slice.
		if len(currentArgs) == 1 {
			return currentArgs[0], nil
		}
		return currentArgs, nil
	}
}
