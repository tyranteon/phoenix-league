package log

import (
	"fmt"
	"os"
)

func Error(err error) {
	fmt.Fprintf(os.Stderr, "ERROR: "+err.Error())
}
