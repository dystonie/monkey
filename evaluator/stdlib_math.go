package evaluator

import (
	"math/rand"
	"monkey/object"
	"time"
)

func mathAbs(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1",
			len(args))
	}
	switch arg := args[0].(type) {
	case *object.Integer:
		v := arg.Value
		if v < 0 {
			v = v * -1
		}
		return &object.Integer{Value: v}
	case *object.Float:
		v := arg.Value
		if v < 0 {
			v = v * -1
		}
		return &object.Float{Value: v}
	default:
		return newError("argument to `type` not supported, got=%s",
			args[0].Type())
	}

	return NULL
}

// val = math.random()
func mathRandom(args ...object.Object) object.Object {
	return &object.Float{Value: rand.Float64()}
}

func init() {
	//
	// Setup our random seed.
	//
	rand.Seed(time.Now().UnixNano())
	registerBuiltin("math.abs",
		func(args ...object.Object) object.Object {
			return (mathAbs(args...))
		})
	registerBuiltin("math.random",
		func(args ...object.Object) object.Object {
			return (mathRandom(args...))
		})
}