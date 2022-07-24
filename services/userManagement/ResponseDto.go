package userManagement

type statusMessage string

const (
	success statusMessage = "SUCCESS"
	error                 = "ERROR"
)

type Response struct {
	Id int
	//TODO refactor to http status code
	Status  statusMessage
	Message string
	Token   string
}
