package grafanagenerator

type IDashboard interface {
	AddUIElement(element IUIElement) error
	AddVariable(variable IVariable)  error
	Draw() (string, error)
	EndLine()
}