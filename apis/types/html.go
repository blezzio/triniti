package types

type HTMLIndexView struct {
	AcceptLanguage string
}

type HTMLSuccessView struct {
	AcceptLanguage string
	URL            string
}

type HTMLErrorView struct {
	AcceptLanguage string
	Code           int
	Error          error
}
