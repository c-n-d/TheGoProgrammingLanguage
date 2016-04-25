/*
Exercise 7.13 - Add Strings method to Expr interface to pretty-print the syntax tree
*/

package eval

import (
    "fmt"
    "strings"
)

func (v Var) String() string {
    return string(v)
}

func (l literal) String() string {
    return fmt.Sprintf("%f", l)
}

func (u unary) String() string {
    return fmt.Sprintf("%s%s", string(u.op), u.x.String())
}

func (b binary) String() string {
    return fmt.Sprintf("%s%s%s", b.x.String(), string(b.op), b.y.String())
}

func (c call) String() string {
    var rem []string
    for _, ex := range c.args {
       rem = append(rem, ex.String())
    }
    return fmt.Sprintf("%s(%s)", c.fn, strings.Join(rem, ", "))
}
