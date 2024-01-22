package helpers

import "fmt"

func FormMessageHTMXResponse(kind string, message string) string {
	if kind == "ok" {
		return fmt.Sprintf("<div id='form-message' class='text-green-500 text-sm mb-4'>%s</div>", message)
	} else {
		return fmt.Sprintf("<div id='form-message' class='text-red-500 text-sm mb-4'>%s</div>", message)
	}
}
