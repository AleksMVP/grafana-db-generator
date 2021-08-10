package variable

import (
	"encoding/json"
	"fmt"

	"github.com/aleksmvp/grafanagenerator/utils"
)

type option struct {
	Selected bool   `json:"selected"`
	Text	 string `json:"text"`
	Value	 string `json:"value"`
}

type current struct {
	Selected bool     `json:"selected"`
	Tags     []string `json:"tags"`
	Text     []string `json:"text"`
	Value    []string `json:"value"`
}

type customVariableTemplate struct {
	Current string
	Options string 
	Query   string
}

type CustomVariable struct {
	data customVariableTemplate
}

func NewCustomVariable(book map[string]string) CustomVariable {
	opts := []option{}
	cur  := current{
		Selected: true,
		Tags:     []string{},
	}
	query   := ""

	for key, value := range book {
		opts = append(opts, option{
			Selected: true,
			Text:     key,
			Value:    value,
		})

		cur.Text  = append(cur.Text, key)
		cur.Value = append(cur.Value, value)

		query += fmt.Sprintf("%v : %v,", key, value)
	}

	optsJson, _ := json.Marshal(opts)
	curJson, _  := json.Marshal(cur)

	return CustomVariable{
		data: customVariableTemplate{
			Current: string(curJson),
			Options: string(optsJson),
			Query:   query,
		},
	}
}

func (instance *CustomVariable) Draw() (variable string, err error) {
	variable, err = utils.ExecuteTemplate(CUSTOM_VAR_TEMPLATE, instance.data)

	return variable, err
}

const (
	CUSTOM_VAR_TEMPLATE = `{
		"allValue": null,
		"current": {{ .Current }},
		"description": null,
		"error": null,
		"hide": 0,
		"includeAll": false,
		"label": null,
		"multi": true,
		"name": "exchange",
		"options": {{ .Options }},
		"query": "{{ .Query }}",
		"queryValue": "",
		"skipUrlSync": false,
		"type": "custom"
	}`
)