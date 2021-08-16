package grafanagenerator

import (
	"fmt"
	"testing"

	"github.com/aleksmvp/grafanagenerator/chart"
	"github.com/aleksmvp/grafanagenerator/models"
	"github.com/aleksmvp/grafanagenerator/row"
	"github.com/aleksmvp/grafanagenerator/variable"
)

func TestCreateConfig(t *testing.T) {
	row := row.NewRow("kk")
	char := chart.NewChart("kdkd", 1, 3, []models.ChartTarget{})
	cff   := chart.NewStatChart("kkfdj", 1, 3, []models.ChartTarget{})
	pie  := chart.NewPieChart("kkfdj", 1, 3, []models.ChartTarget{})
	variable.NewConstVariable("k3k", "k3k")
	variable.NewCustomVariable("kek", map[string]string{})

	v := variable.NewIntervalVariable("kek", []string{"1m","10m","30m","1h","6h","12h","1d","7d","14d","30d"})
	fmt.Println(v.Draw())
	// db := dashboard.NewDashboard("kek")	

	// fmt.Println(db)
	fmt.Println(pie)
	fmt.Println(row)
	fmt.Println(char)
	fmt.Println(cff)
}	