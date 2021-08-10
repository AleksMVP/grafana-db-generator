package variable

import "github.com/aleksmvp/grafanagenerator/utils"

type constVariableTemplate struct {
	Query string
	Name  string
}

type ConstVariable struct {
	data constVariableTemplate
}

func NewConstVariable(name, query string) ConstVariable {
	return ConstVariable{
		data: constVariableTemplate{
			Query: query,
			Name:  name,
		},
	}
}

func (instance *ConstVariable) Draw() (variable string, err error) {
	variable, err = utils.ExecuteTemplate(CONST_VAR_TEMPLATE, instance.data)

	return variable, err
}

const (
	CONST_VAR_TEMPLATE = `{
		"description": null,
		"error": null,
		"hide": 2,
		"label": null,
		"name": "{{ .Name }}",
		"query": "{{ .Query }}",
		"skipUrlSync": false,
		"type": "constant"
	}`
)