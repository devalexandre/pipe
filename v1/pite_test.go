package v1

import (
	"fmt"
	"github.com/devalexandre/gofn/pipe"
	"reflect"
	"strings"
	"testing"
)

// Exemplos de funções adaptadas para usar com generics
func ToUpper(s string) (string, error) {
	return strings.ToUpper(s), nil
}

func Trim(s string) (string, error) {
	return strings.TrimSpace(s), nil
}

func ValidateCPF(cpf string) (string, error) {
	if len(cpf) != 11 {
		return "", fmt.Errorf("CPF must have 11 digits")
	}
	return cpf, nil
}

func FormatCPF(cpf string) (string, error) {
	return fmt.Sprintf("%s.%s.%s-%s", cpf[0:3], cpf[3:6], cpf[6:9], cpf[9:11]), nil
}

func Sum(a, b int) (int, error) {
	return a + b, nil
}

func Multiply(a int) (int, error) {
	return a * 2, nil
}

func ShowResult(result int) (int, error) {
	fmt.Println(result)
	return result, nil
}

// Estrutura Person e funções de validação ajustadas
type Person struct {
	Name  string
	Email string
	Phone string
}

func ValidatePersonName(p Person) (Person, error) {
	if p.Name == "" {
		return p, fmt.Errorf("Name is required")
	}
	return p, nil
}

func ValidatePersonEmail(p Person) (Person, error) {
	if p.Email == "" || !strings.Contains(p.Email, "@") {
		return p, fmt.Errorf("Valid email is required")
	}
	return p, nil
}

func ValidatePersonPhone(p Person) (Person, error) {
	if p.Phone == "" {
		return p, fmt.Errorf("Phone is required")
	}
	return p, nil
}

func SumMore(a, b, c int) (int, error) {
	return a + b + c, nil
}

func GenerateTreeNumbers(n int) (int, int, int) {
	return n, n * 2, n * 3
}

// Testes ajustados para usar a função Pipe com generics

func TestStringPipeline(t *testing.T) {
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

func TestCPFPipeline(t *testing.T) {
	pipeline := Pipe(
		ValidateCPF,
		FormatCPF,
	)

	input := "12345678901"
	expected := "123.456.789-01"
	result, err := pipeline(input)
	if err != nil {
		t.Errorf("Pipeline returned an error: %v", err)
	}

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestMathPipeline(t *testing.T) {
	pipeline := Pipe(
		Sum,
		Multiply,
		ShowResult,
	)

	expected := 18
	result, err := pipeline(5, 4)
	if err != nil {
		t.Errorf("Pipeline returned an error: %v", err)
	}

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestPersonPipeline(t *testing.T) {
	pipeline := Pipe(
		ValidatePersonName,
		ValidatePersonEmail,
		ValidatePersonPhone,
	)

	input := Person{
		Name:  "John Doe",
		Email: "john.doe@example.com",
		Phone: "1234567890",
	}
	expected := input // assumindo que a entrada é válida e esperada como saída

	result, err := pipeline(input)
	if err != nil {
		t.Errorf("Pipeline returned an error: %v", err)
	}

	if result != expected {
		t.Errorf("Expected %+v, got %+v", expected, result)
	}
}

func TestMoreArgs(t *testing.T) {
	pipeline := Pipe(
		GenerateTreeNumbers,
		SumMore,
	)

	expected := 6
	result, err := pipeline(1)
	if err != nil {
		t.Errorf("Pipeline returned an error: %v", err)
	}

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestPipeGoFn(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{4, 8, 12, 16, 20}
	// Definindo o pipeline
	process := Pipe(
		pipe.Filter(func(i int) bool { return i%2 == 0 }),
		pipe.Map(func(i int) int { return i * 2 }),
	)

	// Aplicando o pipeline aos dados
	resultInterface, err := process(data)
	if err != nil {
		fmt.Println("Erro ao processar:", err)
		return
	}

	result, ok := resultInterface.([]int)
	if !ok {
		t.Errorf("Expected result type []int, got %T", resultInterface)
	}

	if len(result) != len(expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

}
