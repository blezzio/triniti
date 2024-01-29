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

type Failure struct {
	printers   map[language.Tag]*message.Printer
	defPrinter *message.Printer
	templ      *template.Template
}

func NewFailure(fs embed.FS) *Failure {
	templ := template.Must(template.New(failureTemplName).ParseFS(fs, failureTemplFN...))
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

	return &Failure{
		printers:   prts,
		defPrinter: defprt,
		templ:      templ,
	}
}

func (t *Failure) Exec(wr http.ResponseWriter, data any) error {
	validData, ok := data.(*types.HTMLErrorView)
	if !ok {
		return utils.Trace(fmt.Errorf("expected type %T got %T", &types.HTMLErrorView{}, data), "failed to parse data type")
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
		Oopsie, AnErrorOccurred, Trace, Error string
	}{
		HeaderData:      NewHeaderData(printer),
		Oopsie:          printer.Sprintf(l10n.ErrOopsie, validData.Code),
		AnErrorOccurred: printer.Sprintf(l10n.ErrAnErrorOccurred),
		Trace:           printer.Sprintf(l10n.ErrTrace),
		Error:           validData.Error.Error(),
	}

	t.addHeaders(wr, validData.Code)
	if err := t.templ.Execute(wr, param); err != nil {
		return utils.Trace(err, "failed to execute template")
	}
	return nil
}

func (t *Failure) addHeaders(wr http.ResponseWriter, code int) {
	wr.Header().Add("Content-Type", "text/html")
	wr.WriteHeader(code)
}
