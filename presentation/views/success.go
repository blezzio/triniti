package views

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"

	"github.com/blezzio/triniti/apis/types"
	"github.com/blezzio/triniti/presentation/l10n"
	"github.com/blezzio/triniti/utils"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type Success struct {
	printers   map[language.Tag]*message.Printer
	defPrinter *message.Printer
	templ      *template.Template
}

func NewSucccess(fs embed.FS) *Success {
	templ := template.Must(template.New(successTemplName).ParseFS(fs, successTemplFN...))
	prts := make(map[language.Tag]*message.Printer)
	var defprt *message.Printer
	first := true
	for tag := range l10n.SupportedLocale {
		prts[tag] = message.NewPrinter(tag)
		if first {
			defprt = prts[tag]
			first = false
		}
	}

	return &Success{
		printers:   prts,
		defPrinter: defprt,
		templ:      templ,
	}
}

func (t *Success) Exec(wr http.ResponseWriter, data any) error {
	validData, ok := data.(*types.HTMLSuccessView)
	if !ok {
		return utils.Trace(fmt.Errorf("expected type %T got %T", &types.HTMLSuccessView{}, data), "failed to parse data type")
	}

	lang := language.English
	if accept, err := utils.GetLanguage(validData.AcceptLanguage, l10n.SupportedLocale); err == nil {
		lang = accept
	}

	printer := t.defPrinter
	if p, ok := t.printers[lang]; ok {
		printer = p
	}

	param := struct {
		HeaderData
		URL, Shortened, ClickButtonToCopy, Copied, Copy string
	}{
		HeaderData:        NewHeaderData(printer),
		URL:               validData.URL,
		Shortened:         printer.Sprintf(l10n.SuccessShortened),
		ClickButtonToCopy: printer.Sprintf(l10n.SuccessClickButtonToCopy),
		Copied:            printer.Sprintf(l10n.SuccessCopied),
		Copy:              printer.Sprintf(l10n.SuccessCopy),
	}

	err := t.templ.Execute(wr, param)
	return utils.Trace(err, "failed to execute template")
}

func (t *Success) AddHeaders(wr http.ResponseWriter) {
	wr.Header().Add("Content-Type", "text/html")
}
