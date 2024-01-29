package views

const (
	indexTemplName   = "INDEX"
	successTemplName = "SUCCESS"
	failureTemplName = "FAILURE"
)

var (
	indexTemplFN   = []string{"templates/header.gohtml", "templates/index.gohtml"}
	successTemplFN = []string{"templates/header.gohtml", "templates/success.gohtml"}
	failureTemplFN = []string{"templates/header.gohtml", "templates/failure.gohtml"}
)
