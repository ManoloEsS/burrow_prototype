package httprequest

// struct used to capture user inputs in tui
type UserRequest struct {
	Method string
	Auth
	Body
	Url     string
	Params  map[string]string
	Headers map[string]string
}

// struct that stores auth type and auth fields
type Auth struct {
	Type        string
	BearerToken string
	User        string
	Password    string
}

// struct that stores body type and content
type Body struct {
	Type    string
	Content string
}

// error message struct that is returned to tui when an error occurs
type RequestErrorMsg struct {
	Err error
}

