package constants

var LinkPageList = map[string]string{
	// Login URL
	"WebLoginURL": WebLoginURL,

	// SIS URL
	"WebSISHome":    WebSISHomeURL,
	"WebSISUser":    WebSISUserURL,
	"WebSISMember":  WebSISMemberURL,
	"WebSISProduct": WebSISProductURL,

	// Web POS Constanta
	"WebPOSHome": WebPOSHomeURL,
}

type WebData struct {
	Title        string
	BaseURL      string
	CurrentURL   string
	StaticUrl    string
	LinkPageList map[string]string
}
