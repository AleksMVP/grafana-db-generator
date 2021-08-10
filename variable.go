package grafanagenerator

type IVariable interface {
	Draw() (variable string, err error)
}