package views

import (
	"github.com/blezzio/triniti/presentation/l10n"
	"golang.org/x/text/message"
)

type HeaderData struct {
	Lang, Desc, Kw string
}

func NewHeaderData(printer *message.Printer) HeaderData {
	return HeaderData{
		Lang: printer.Sprintf(l10n.HeaderLang),
		Desc: printer.Sprintf(l10n.HeaderDesc),
		Kw:   printer.Sprintf(l10n.HeaderKeywords),
	}
}
