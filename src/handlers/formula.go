package handlers

import (
	"math"
)

// BMI=体重（公斤）÷（身高×身高）（米）
func BMI(weight, height float64) float64 {
	BMI := weight / math.Pow(height, 2.0)
	return BMI
}

//体脂率：1.2×BMI+0.23×年龄-5.4-10.8×性别（男为1，女为0）
func Fat(bmi float64, age, sex int) float64 {
	fat := 1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(sex)
	return fat
}

// 直线的斜率 = (y2-y1) / (x2-x1)
func Slope(x1, y1, x2, y2 int) float64 {
	k := float64((y2 - y1) / (x2 - x1))
	return k
}
