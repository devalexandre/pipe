package v1

import (
	"fmt"
	"strings"
	"testing"
)

func ToUpper(s string) (string, error) {
	return strings.ToUpper(s), nil
}

// Wrapper para strings.TrimSpace que se encaixa na assinatura esperada por Pipe.
func Trim(s string) (string, error) {
	return strings.TrimSpace(s), nil
}

func TestPipe(t *testing.T) {

	pipeline := Pipe(
		ToUpper,
		Trim,
	)

	input := "   go is awesome   "
	expected := "GO IS AWESOME"
	result, err := pipeline(input)
	if err != nil {
		t.Errorf("Pipeline returned an error: %v", err)
	}

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestParseTo(t *testing.T) {

	result := "test result"
	var target string

	err := ParseTo(result, &target)
	if err != nil {
		t.Errorf("ParseTo returned an error: %v", err)
	}

	if target != result {
		t.Errorf("Expected target to be %s, got %s", result, target)
	}

	var intTarget int
	err = ParseTo(result, &intTarget)
	if err == nil {
		t.Errorf("Expected ParseTo to fail when types are incompatible")
	}
}

// Sum recebe dois números e retorna a sua soma.
func Sum(a, b int) (int, error) {
	return a + b, nil
}

// IntToString recebe um número e retorna sua representação em string.
func IntToString(n int) (string, error) {
	return fmt.Sprintf("%d", n), nil
}

func TestSumPipeline(t *testing.T) {
	// Cria o pipeline.
	pipeline := Pipe(
		Sum,
		IntToString,
	)

	// Define a entrada e a saída esperada.
	input1 := 5
	input2 := 7
	expected := "12"

	// Executa o pipeline com os números de entrada.
	resultInterface, err := pipeline(input1, input2)
	if err != nil {
		t.Fatalf("Pipeline returned an error: %v", err)
	}

	// Converte o resultado para string.
	result, ok := resultInterface.(string)
	if !ok {
		t.Fatalf("Failed to convert result to string")
	}

	// Verifica se o resultado é o esperado.
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
