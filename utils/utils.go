package utils

import (
	"bytes"
	"strings"
	"text/template"
)

func ExecuteTemplateFromFile(templatePath string, data interface{}) (result string, err error) {
	name := strings.Split(templatePath, "/")
	t := template.New(name[len(name) - 1])

	t, err = t.ParseFiles(templatePath)
	if err != nil {
		return result, err
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, data); err != nil {
		return result, err
	}

	return tpl.String(), nil
}

func ExecuteTemplate(templ string, data interface{}) (result string, err error) {
  	t, err := template.New("templ").Parse(templ)
	if err != nil {
		return result, err
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, data); err != nil {
		return result, err
	}

	return tpl.String(), nil
}

func SetToSlice(set map[string]bool) (result []string) {
	for key := range set {
		result = append(result, key)
	}

	return result
}

func CutLastElement(str string) string {
	runes := []rune(str)
	if len(runes) == 0 {
		return ""
	}

	return string(runes[:len(runes) - 1])
}