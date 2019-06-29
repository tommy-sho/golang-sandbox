package greeting

import "fmt"

var message string

func Hey(name string) string {
	return fmt.Sprintf("%s, %s", message, name)
}
