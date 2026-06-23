// TN — benign template fill; validated typed fields, no AI surface.
package aisec

import (
	"errors"
	"fmt"
)

func Render(user string, amount int) (string, error) {
	if amount < 0 {
		return "", errors.New("amount must be non-negative")
	}
	return fmt.Sprintf("Hi %s, your balance changed by %d credits.", user, amount), nil
}
