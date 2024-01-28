package views

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"

	"github.com/blezzio/triniti/handlers/types"
	"github.com/blezzio/triniti/presentation/l10n"
	"github.com/blezzio/triniti/utils"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type Index struct {
	printers   map[language.Tag]*message.Printer
	defPrinter *message.Printer
	templ      *template.Template
}

func NewIndex(fs embed.FS) *Index {
	templ := template.Must(template.New(indexTemplName).ParseFS(fs, indexTemplFN...))
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

	return &Index{
		printers:   prts,
		defPrinter: defprt,
		templ:      templ,
	}
}

func (t *Index) Exec(wr http.ResponseWriter, data any) error {
	validData, ok := data.(*types.HTMLIndexView)
	if !ok {
		return utils.Trace(fmt.Errorf("expected type %T got %T", &types.HTMLIndexView{}, data), "failed to parse data type")
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
		Text string
	}{
		HeaderData: NewHeaderData(printer),
		Text:       printer.Sprintf(l10n.IndexGreeting),
	}

	err := t.templ.Execute(wr, param)
	return utils.Trace(err, "failed to execute template")
}

func (t *Index) AddHeaders(wr http.ResponseWriter) {
	wr.Header().Add("Content-Type", "text/html")
}
