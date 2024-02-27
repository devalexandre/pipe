package v1

import (
	"errors"
	"reflect"
)

// Pipeline type.
type Pipeline func(...interface{}) (interface{}, error)

// Pipe creates a pipeline from a series of functions.
func Pipe(fs ...interface{}) Pipeline {
	return func(initialArgs ...interface{}) (interface{}, error) {
		var currentArgs []interface{} = initialArgs
		for _, f := range fs {
			fnVal := reflect.ValueOf(f)
			fnType := fnVal.Type()
			numIn := fnType.NumIn()

			// Prepare the input arguments for the current function.
			in := make([]reflect.Value, numIn)
			for i := 0; i < numIn; i++ {
				if i < len(currentArgs) {
					in[i] = reflect.ValueOf(currentArgs[i])
				} else if i < len(initialArgs) { // Allow passing manual arguments if not enough currentArgs.
					in[i] = reflect.ValueOf(initialArgs[i])
				} else {
					// If there are not enough arguments to pass to the function, return an error.
					return nil, errors.New("not enough arguments to pass to function")
				}
			}

			// Call the current function in the pipeline.
			results := fnVal.Call(in)

			// Assume the last function call results will be used as next input.
			currentArgs = []interface{}{} // Reset currentArgs for next function.
			for _, result := range results {
				if result.Type().Implements(reflect.TypeOf((*error)(nil)).Elem()) {
					if !result.IsNil() { // If the result is an error, return it.
						return nil, result.Interface().(error)
					}
					// If it's a nil error, ignore it for the output.
				} else {
					currentArgs = append(currentArgs, result.Interface())
				}
			}
		}

		// Return the final result which should match the last function's output type.
		if len(currentArgs) == 1 {
			return currentArgs[0], nil // Return single value if only one result.
		}
		return currentArgs, nil // Return as slice if multiple values.
	}
}
