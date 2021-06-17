package models

//TemplateData holds data from handler to the template
type TemplateData struct {
	StringMap  map[string]string
	IntMap     map[string]int
	FloatMap   map[string]float32
	Data       map[string]interface{}
	CRSFSToken string
	Flash      string
	Warning    string
	Error      string
}
