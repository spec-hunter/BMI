package globals

import (
	"gorm.io/gorm"

	"github.com/ahlemarg/BMI/src/config"
)

var (
	DB       *gorm.DB
	DBconfig config.DBConfig

	Exit string

	Rate []float64
)

func init() {
	Rate = make([]float64, 0, 4)
}
