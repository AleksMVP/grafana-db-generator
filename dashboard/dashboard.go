package dashboard

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/aleksmvp/grafanagenerator"
	"github.com/aleksmvp/grafanagenerator/utils"
)

const (
	MAX_PAGE_WIDTH int = 24
)

type dashboardTemplate struct {
	Panels    string
	Variables string
	UID       string
	Title     string
}

type Dashboard struct {
	currentX     int
	currentY     int
	currentID    int
	prevHeight   int
	elements     []string
	vars         []string

	data         dashboardTemplate
}

func NewDashboard(title string) Dashboard {
	return Dashboard{
		data: dashboardTemplate{
			Title: title,
		},
	}
}

func (instance *Dashboard) getID() int {
	instance.currentID++

	return instance.currentID
}

func (instance *Dashboard) AddUIElement(element grafanagenerator.IUIElement) (err error) {
	elem   := ""
	elem, instance.currentX, instance.currentY, instance.prevHeight, err = element.Draw(instance.currentX, instance.currentY, instance.getID())

	if err != nil {
		return err
	}

	if instance.currentX >= MAX_PAGE_WIDTH {
		instance.currentX = 0
		instance.currentY += instance.prevHeight
	}

	instance.elements = append(instance.elements, elem)

	return nil
}

func (instance *Dashboard) AddVariable(variable grafanagenerator.IVariable)  (err error) {
	variable_, err := variable.Draw()
	if err != nil {
		return err
	}

	instance.vars = append(instance.vars, variable_)

	return nil
}

func (instance *Dashboard) EndLine() {
	if instance.currentX == 0 {
		return
	}
	instance.currentX = 0
	instance.currentY += instance.prevHeight
}

func (instance *Dashboard) Draw() (string, error) {
	for index, element := range instance.elements {
		instance.data.Panels += element
		if index < len(instance.elements) - 1 {
			instance.data.Panels += ","
		}
	}

	for index, variable := range instance.vars {
		instance.data.Variables += variable
		if index < len(instance.vars) - 1 {
			instance.data.Variables += ","
		}
	}

	hash := sha256.Sum256([]byte(instance.data.Title + instance.data.Panels + instance.data.Variables))
	instance.data.UID = hex.EncodeToString(hash[:16])

	dashboard, err := utils.ExecuteTemplate(BASE_TEPLATE, instance.data)

	return dashboard, err
}

const (
	BASE_TEPLATE = `{
		"annotations": {
			"list": [
				{
					"builtIn": 1,
					"datasource": "-- Grafana --",
					"enable": true,
					"hide": true,
					"iconColor": "rgba(0, 211, 255, 1)",
					"name": "Annotations & Alerts",
					"type": "dashboard"
				}
			]
		},
		"editable": true,
		"gnetId": null,
		"graphTooltip": 0,
		"id": null,
		"links": [],
		"panels": [{{ .Panels }}],
		"schemaVersion": 27,
		"refresh": "5s",
		"style": "dark",
		"tags": [],
		"templating": {
			"list": [{{ .Variables }}]
		},
		"time": {
			"from": "now-1h",
			"to": "now"
		},
		"timepicker": {},
		"timezone": "",
		"title": "{{ .Title }}",
		"uid": "{{ .UID }}",
		"version": 2
	}
	`
)