package sqltyping

import (
	"html"
	"strings"
)

func SanitizeString(str *string) {
	// Change ' character with ` character
	*str = html.EscapeString(*str)
	*str = strings.Replace(*str, "&#34;", "\"", -1)
}
