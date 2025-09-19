package field_test

import (
	"fmt"

	"github.com/hollfish/gen/field"
)

func ExampleFunc() {
	expr := field.Func.UnixTimestamp()

	sql, vars := field.BuildToString(expr)
	fmt.Println(sql, vars)

	sql, vars = field.BuildToString(expr.Mul(100))
	fmt.Println(sql, vars)

	// Output:
	// UNIX_TIMESTAMP() []
	// (UNIX_TIMESTAMP())*? [100]
}
