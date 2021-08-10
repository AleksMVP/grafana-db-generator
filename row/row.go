package row

import (
	"github.com/aleksmvp/grafanagenerator/utils"
)

type rowTemplate struct {
	Y     int
	ID    int
	Title string
}

type Row struct {
	data rowTemplate
}

func NewRow(title string) (Row) {
	return Row {
		data: rowTemplate{
			Title: title,
		},
	}
}

func (instance *Row) Draw(x, y, ID int) (row string, nextX, nextY, height int, err error) {
	instance.data.Y  = y
	instance.data.ID = ID

	row, err = utils.ExecuteTemplate(ROW_TEMPLATE, instance.data)
	nextX  = 0
	nextY  = y + 1
	height = 1

	return 
}

const (
	ROW_TEMPLATE = `{
		"collapsed": false,
		"datasource": null,
		"gridPos": {
			"h": 1,
			"w": 24,
			"x": 0,
			"y": {{ .Y }}
		},
		"id": {{ .ID }},
		"panels": [],
		"title": "{{ .Title }}",
		"type": "row"
	}`
)