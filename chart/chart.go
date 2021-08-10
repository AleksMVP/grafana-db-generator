package chart

import (
	"github.com/aleksmvp/grafanagenerator/models"
	"github.com/aleksmvp/grafanagenerator/utils"
)

type Chart struct {
	data chartTemplate
}

type chartTemplate struct {
	Height  int
	Width   int
	X       int
	Y       int
	ID      int
	Title   string
	Targets []models.ChartTarget
}

func NewChart(title string, height, width int, targets []models.ChartTarget) (Chart) {
	return Chart{
		data: chartTemplate{
			Height:  height,
			Width:   width,
			Title:   title,
			Targets: targets,
		},
	}
}

func (instance *Chart)Draw(x, y, ID int) (chart string, nextX, nextY, height int, err error) {
	instance.data.X  = x
	instance.data.Y  = y
	instance.data.ID = ID

	chart, err = utils.ExecuteTemplate(CHART_TEMPLATE, instance.data)
	if err != nil {
		return
	}

	nextY  = y
	nextX  = x + instance.data.Width
	height = instance.data.Height

	return 
}


const (
	CHART_TEMPLATE = `{
		"aliasColors": {},
		"bars": false,
		"dashLength": 10,
		"dashes": false,
		"datasource": null,
		"fieldConfig": {
		  "defaults": {},
		  "overrides": []
		},
		"fill": 1,
		"fillGradient": 0,
		"gridPos": {
		  "h": {{ .Height }},
		  "w": {{ .Width  }},
		  "x": {{ .X }},
		  "y": {{ .Y }}
		},
		"hiddenSeries": false,
		"id": {{ .ID }},
		"legend": {
		  "avg": false,
		  "current": false,
		  "max": false,
		  "min": false,
		  "show": true,
		  "total": false,
		  "values": false
		},
		"lines": true,
		"linewidth": 1,
		"nullPointMode": "null",
		"options": {
		  "alertThreshold": true
		},
		"percentage": false,
		"pluginVersion": "7.5.2",
		"pointradius": 2,
		"points": false,
		"renderer": "flot",
		"seriesOverrides": [],
		"spaceLength": 10,
		"stack": false,
		"steppedLine": false,
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
		"thresholds": [],
		"timeFrom": null,
		"timeRegions": [],
		"timeShift": null,
		"title": "{{ .Title }}",
		"tooltip": {
		  "shared": true,
		  "sort": 0,
		  "value_type": "individual"
		},
		"type": "graph",
		"xaxis": {
		  "buckets": null,
		  "mode": "time",
		  "name": null,
		  "show": true,
		  "values": []
		},
		"yaxes": [
		  {
			"format": "short",
			"label": null,
			"logBase": 1,
			"max": null,
			"min": null,
			"show": true
		  },
		  {
			"format": "short",
			"label": null,
			"logBase": 1,
			"max": null,
			"min": null,
			"show": true
		  }
		],
		"yaxis": {
		  "align": false,
		  "alignLevel": null
		}
	  }`
)