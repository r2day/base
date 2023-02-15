package util

import (
	"fmt"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func Cap(s string) string {
	return fmt.Sprint(cases.Title(language.Und).String(s))
}