# Pipe Package for Go

The `pipe` package provides a simple yet powerful way to create Unix-like pipelines in Go. It allows developers to chain together functions in a sequence where the output of one function becomes the input to the next. This package is designed to make function composition intuitive and to support clean, readable code for complex data processing tasks.

## Features

- **Sequential Execution**: Functions in the pipeline are executed in the order they are added, with the output of one function passed as input to the next.
- **Error Handling**: The pipeline supports error propagation. If any function in the sequence returns an error, the pipeline execution stops, and the error is returned to the caller.
- **Dynamic Function Support**: Functions of varying signatures can be added to the pipeline, providing flexibility in the types of operations that can be performed.
- **Type Safety**: While the package uses reflection to dynamically invoke functions, it includes mechanisms to ensure that function inputs and outputs are correctly managed and errors are meaningfully reported.

## Installation

To use the `pipe` package in your Go project, you first need to install it:

```bash
go get github.com/devalexandre/pipe/v1
```

## Usage
Here's a simple example of how to use the pipe package to create a pipeline that processes strings:

```go
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

```

# API Reference

## Pipe
Pipe is the core function of the package. It accepts a variable number of functions as arguments and returns a Pipeline function. The Pipeline function can be executed with an input value, and the output of the pipeline is the result of executing the functions in sequence.

```go
func Pipe(fs ...interface{}) Pipeline {
```

## Pipeline
The Pipeline type is a function that accepts an input value and returns the result of executing the functions in the pipeline. If any function in the sequence returns an error, the pipeline execution stops, and the error is returned to the caller.

```go
type Pipeline func(args ...interface{}) (interface{}, error)
```

## ParseTo
The ParseTo function is a helper function that converts the output of a pipeline to a specific type. It accepts a type as an argument and returns a function that can be used to execute the pipeline and parse the result to the specified type.

```go
func ParseTo(result interface{}, target interface{}) error
```

# Contributing 
Contributions to the `pipe` package are welcome! Whether it's bug reports, feature requests, or pull requests, all forms of contributions are appreciated.

# License
This package is licensed under the MIT License. See the LICENSE file for details.