package grafanagenerator

type IUIElement interface {
	Draw(x, y, ID int) (chart string, nextX, nextY, height int, err error)
}