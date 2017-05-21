package template

type IContent interface{
	Input()
	Substitute() string
	Validate() bool
}

type TEngine struct {
	Content IContent
}

func NewTemplateEngine(content IContent) *TEngine {
	return &TEngine{
		Content: content,
	}
}

func (te *TEngine) validate() bool {
	return te.Content.Validate()
}

func (te *TEngine) Render() string {
	if te.validate() {
		return te.substitute()
	}
	return ""
}

func (te *TEngine) substitute() string {
	return te.Content.Substitute()
}


