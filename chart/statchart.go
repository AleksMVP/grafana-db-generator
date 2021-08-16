package chart

import (
	"github.com/aleksmvp/grafanagenerator/models"
	"github.com/aleksmvp/grafanagenerator/utils"
)

type StatChart struct {
	data chartTemplate
}

func NewStatChart(title string, height, width int, targets []models.ChartTarget) (Chart) {
	return Chart{
		data: chartTemplate{
			Height:  height,
			Width:   width,
			Title:   title,
			Targets: targets,
		},
	}
}

func (instance *StatChart) Draw(x, y, ID int) (chart string, nextX, nextY, height int, err error) {
	instance.data.X  = x
	instance.data.Y  = y
	instance.data.ID = ID

	chart, err = utils.ExecuteTemplate(STAT_CHART_TEMPLATE, instance.data)
	if err != nil {
		return
	}

	nextY  = y
	nextX  = x + instance.data.Width
	height = instance.data.Height

	return 
}

const (
	STAT_CHART_TEMPLATE = `{
		"datasource": null,
		"description": "",
		"fieldConfig": {
			"defaults": {
				"color": {
					"mode": "thresholds"
				},
				"mappings": [],
				"thresholds": {
					"mode": "absolute",
					"steps": [
						{
							"color": "red",
							"value": null
						},
						{
							"color": "green",
							"value": 0
						}
					]
				}
			},
			"overrides": []
		},
		"gridPos": {
			"h": {{ .Height }},
			"w": {{ .Width }},
			"x": {{ .X }},
			"y": {{ .Y }}
		},
		"id": {{ .ID }},
		"options": {
			"colorMode": "value",
			"graphMode": "area",
			"justifyMode": "center",
			"orientation": "auto",
			"reduceOptions": {
				"calcs": [
					"lastNotNull"
				],
				"fields": "",
				"values": false
			},
			"text": {},
			"textMode": "auto"
		},
		"pluginVersion": "7.5.2",
		"targets": [
			{{ range $index, $item := .Targets }}
				{{ if ne $index 0 }}
				  ,
				{{ end }}
				{
				  "exemplar": true,
				  "expr": "{{ .Expr }}",
				  "interval": "",
				  "legendFormat": "{{ .Legend }}",
				  "refId": "{{ .RefId }}"
				}
			{{ end }}
		],
		"timeFrom": null,
		"timeShift": null,
		"title": "{{ .Title }}",
		"type": "stat"
	}`
)