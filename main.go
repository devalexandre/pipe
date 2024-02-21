package main

import (
	"fmt"
	"github.com/devalexandre/pipe/v1"
	"strings"
)

func ToUpper(s string) (string, error) {
	return strings.ToUpper(s), nil
}

func Trim(s string) (string, error) {
	return strings.TrimSpace(s), nil
}

func AddPrefix(prefix string, s string) (string, error) {
	return fmt.Sprintf("%s%s", prefix, s), nil
}

func main() {
	// Captura o prefixo como uma variável externa.
	prefix := "Prefix: "
	processText := v1.Pipe(
		ToUpper,
		Trim,
		// Usa uma função anônima para adaptar AddPrefix ao pipeline.
		func(s string) (string, error) {
			return AddPrefix(prefix, s)
		},
	)

	// Ajuste na chamada para desempacotar o resultado e o erro.
	resultInterface, err := processText("   go is awesome   ")
	if err != nil {
		fmt.Println("Erro ao processar texto:", err)
		return
	}

	// Converte o resultado de volta para string.
	var result string
	err = v1.ParseTo(resultInterface, &result)
	if err != nil {
		fmt.Println("Erro ao converter resultado:", err)
		return
	}
	fmt.Println("Resultado:", result)
}
