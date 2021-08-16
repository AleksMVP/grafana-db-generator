package chart

import (
	"github.com/aleksmvp/grafanagenerator/models"
	"github.com/aleksmvp/grafanagenerator/utils"
)

type PieChart struct {
	data chartTemplate
}

func NewPieChart(title string, height, width int, targets []models.ChartTarget) (PieChart) {
	return PieChart{
		data: chartTemplate{
			Height:  height,
			Width:   width,
			Title:   title,
			Targets: targets,
		},
	}
}

func (instance *PieChart) Draw(x, y, ID int) (chart string, nextX, nextY, height int, err error) {
	instance.data.X  = x
	instance.data.Y  = y
	instance.data.ID = ID

	chart, err = utils.ExecuteTemplate(PIE_CHART_TEMPLATE, instance.data)
	if err != nil {
		return
	}

	nextY  = y
	nextX  = x + instance.data.Width
	height = instance.data.Height

	return 
}

const (
	PIE_CHART_TEMPLATE = `{
	"datasource": null,
	"fieldConfig": {
	  "defaults": {
		"color": {
		  "mode": "palette-classic"
		},
		"mappings": [],
		"thresholds": {
		  "mode": "absolute",
		  "steps": [
			{
			  "color": "green",
			  "value": null
			},
			{
			  "color": "red",
			  "value": 80
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
	  "displayLabels": [],
	  "legend": {
		"displayMode": "list",
		"placement": "right",
		"values": []
	  },
	  "pieType": "pie",
	  "reduceOptions": {
		"calcs": [
		  "lastNotNull"
		],
		"fields": "",
		"values": false
	  },
	  "text": {}
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
	"type": "piechart"
  }`
)