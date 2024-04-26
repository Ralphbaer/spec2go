package model

type TemplateData struct {
	Apis       []API
	ServerPort string
}

type API struct {
	ClassName string
	HasMore   bool
}
