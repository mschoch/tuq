package parser

import ()

type Expression interface {
}

// NOTE: this should be OK even if
// Expression eventually gets methods
// as an Expression List is not an Expression
// by itself (though it can become one inside
// a literal array)
type ExpressionList []Expression
