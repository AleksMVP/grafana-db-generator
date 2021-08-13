package variable

import (
	"encoding/json"
	"fmt"

	"github.com/aleksmvp/grafanagenerator/utils"
)

type intervalVariableTemplate struct {
	Name    string
	Current string
	Options string 
	Query   string
}

type IntervalVariable struct {
	data intervalVariableTemplate
}

func NewIntervalVariable(name string, intervals []string) IntervalVariable {
	opts := []option{}
	cur  := option{
		Selected: false,
		Text:     intervals[0],
		Value:    intervals[0],
	}

	query   := ""

	for _, value := range intervals {
		opts = append(opts, option{
			Selected: false,
			Text:     value,
			Value:    value,
		})

		query += fmt.Sprintf("%v,", value)
	}

	opts[0].Selected = true

	query = utils.CutLastElement(query)

	optsJson, _ := json.Marshal(opts)
	curJson, _  := json.Marshal(cur)

	return IntervalVariable{
		data: intervalVariableTemplate{
			Name:    name,
			Current: string(curJson),
			Options: string(optsJson),
			Query:   query,
		},
	}
}

func (instance *IntervalVariable) Draw()(variable string, err error) {
	variable, err = utils.ExecuteTemplate(INTERVAL_VAR_TEMPLATE, instance.data)

	return variable, err
}

const (
	INTERVAL_VAR_TEMPLATE = `{
		"auto": false,
		"auto_count": 30,
		"auto_min": "10s",
		"current": {{ .Current }},
		"description": null,
		"error": null,
		"hide": 0,
		"label": null,
		"name": "{{ .Name }}",
		"options": {{ .Options }},
		"query": "{{ .Query }}",
		"queryValue": "",
		"refresh": 2,
		"skipUrlSync": false,
		"type": "interval"
	}`
)