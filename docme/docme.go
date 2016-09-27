package docme

import (
	"fmt"
)

/*
Return a greeting message using name into it.
*/
func Hello(name string) string {
	return fmt.Sprintf("Hello, my name is %s", name)
}
