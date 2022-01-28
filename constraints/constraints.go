// Package genc contains custom constraints.
package genc

import "constraints"

// Number is a custom constraint including constraints.Integer and constraints.Float.
type Number interface {
	constraints.Integer | constraints.Float
}
