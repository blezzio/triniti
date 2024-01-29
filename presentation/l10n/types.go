package l10n

import "golang.org/x/text/language"

type entry struct {
	tag language.Tag
	k   string
	// msg accept string or plural.Selectf
	msg any
}
