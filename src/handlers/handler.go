package handlers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ahlemarg/BMI/src/globals"
	"github.com/ahlemarg/BMI/src/model"
)

func Change(pointer string) int {
	res, _ := strconv.Atoi(pointer)
	return res
}

func SlopeRate() {
	var a, b string
	fmt.Scanf("Please Enter Pointer A(example: x,y): %s\n", &a)
	fmt.Scanf("Please Enter Pointer B(exapmle: x,y): %s\n", &b)

	pointer_a := strings.Split(strings.ReplaceAll(a, " ", ""), ",")
	pointer_b := strings.Split(strings.ReplaceAll(b, " ", ""), ",")

	k := Slope(Change(pointer_a[0]), Change(pointer_a[1]), Change(pointer_b[0]), Change(pointer_b[1]))
	globals.Rate = append(globals.Rate, k)
}

func SlopesEqual() {
	for i := 0; i < 2; i++ {
		SlopeRate()
	}

	if globals.Rate[0] == globals.Rate[1] {
		fmt.Println("两条直线的斜率是 相等 的.")
	} else {
		fmt.Println("两条直线的斜率 不相等 的.")
	}
}

func Group(t func(int, float64), age int, fat float64, funcName string) {
	if funcName == "Women" {
		t(age, fat)
	} else if funcName == "Man" {
		t(age, fat)
	}
}

func MaleBodyFat(age int, fat float64) {
	switch {
	case (age >= 18 && age <= 39):
		switch {
		case (fat > 0 && fat <= 10):
			fmt.Printf("偏瘦.\n")
		case (fat > 10 && fat <= 16):
			fmt.Printf("标准.\n")
		case (fat > 16 && fat <= 21):
			fmt.Printf("偏重.\n")
		case (fat > 21 && fat <= 26):
			fmt.Printf("肥胖.\n")
		case (fat > 26 && fat <= 100):
			fmt.Printf("严重肥胖.\n")
		default:
			fmt.Printf("胖成这个吊样, 等死吧你.\n")
		}
	case (age >= 40 && age <= 59):
		switch {
		case (fat > 0 && fat <= 11):
			fmt.Printf("偏瘦.\n")
		case (fat > 11 && fat <= 17):
			fmt.Printf("标准.\n")
		case (fat > 17 && fat <= 22):
			fmt.Printf("偏重.\n")
		case (fat > 22 && fat <= 27):
			fmt.Printf("肥胖.\n")
		case (fat > 27 && fat <= 100):
			fmt.Printf("严重肥胖.\n")
		default:
			fmt.Printf("胖成这个吊样, 等死吧你.\n")
		}
	case (age >= 60):
		switch {
		case (fat > 0 && fat <= 13):
			fmt.Printf("偏瘦.\n")
		case (fat > 13 && fat <= 19):
			fmt.Printf("标准.\n")
		case (fat > 19 && fat <= 24):
			fmt.Printf("偏重.\n")
		case (fat > 24 && fat <= 29):
			fmt.Printf("肥胖.\n")
		case (fat > 29 && fat <= 100):
			fmt.Printf("严重肥胖.\n")
		default:
			fmt.Printf("胖成这个吊样, 等死吧你.\n")
		}
	}
}

func Man(age int, fat float64) {
	Group(MaleBodyFat, age, fat, "Man")
}

func FemaleBodyFat(age int, fat float64) {
	switch {
	case (age >= 18 && age <= 39):
		switch {
		case (fat > 0 && fat <= 20):
			fmt.Printf("偏瘦.\n")
		case (fat > 20 && fat <= 27):
			fmt.Printf("标准.\n")
		case (fat > 27 && fat <= 34):
			fmt.Printf("偏重.\n")
		case (fat > 34 && fat <= 39):
			fmt.Printf("肥胖.\n")
		case (fat > 39 && fat <= 100):
			fmt.Printf("严重肥胖.\n")
		default:
			fmt.Printf("胖成这个吊样, 等死吧你.\n")
		}
	case (age >= 40 && age <= 59):
		switch {
		case (fat > 0 && fat <= 21):
			fmt.Printf("偏瘦.\n")
		case (fat > 21 && fat <= 28):
			fmt.Printf("标准.\n")
		case (fat > 28 && fat <= 35):
			fmt.Printf("偏重.\n")
		case (fat > 35 && fat <= 40):
			fmt.Printf("肥胖.\n")
		case (fat > 40 && fat <= 100):
			fmt.Printf("严重肥胖.\n")
		default:
			fmt.Printf("胖成这个吊样, 等死吧你.\n")
		}
	case (age >= 60):
		switch {
		case (fat > 0 && fat <= 22):
			fmt.Printf("偏瘦.\n")
		case (fat > 22 && fat <= 29):
			fmt.Printf("标准.\n")
		case (fat > 29 && fat <= 36):
			fmt.Printf("偏重.\n")
		case (fat > 36 && fat <= 41):
			fmt.Printf("肥胖.\n")
		case (fat > 41 && fat <= 100):
			fmt.Printf("严重肥胖.\n")
		default:
			fmt.Printf("胖成这个吊样, 等死吧你.\n")
		}
	}
}

func Women(age int, fat float64) {
	Group(FemaleBodyFat, age, fat, "Women")
}

func Calculator() {
	res := new([]model.Person)
	r := globals.DB.Select("k").Find(&res)

	sum := 0.0
	for _, v := range *res {
		sum += v.K
	}

	k := sum / float64(r.RowsAffected)
	fmt.Printf("当前数据组%d中的平均体脂为: %.2f\n", r.RowsAffected, k)
}

func SaveData(weight float64, height float64, age int, sex_code int, name string) {
	bmi := BMI(weight, height)
	k := Fat(bmi, age, sex_code)
	p := model.Person{
		Name:   name,
		Weight: weight,
		Height: height,
		Age:    age,
		Sex:    sex_code,
		K:      k,
	}
	globals.DB.Select("Name", "Weight", "Height", "Age", "Sex", "K").Create(&p)
}

func BodyFatCalculator() {
	defer fmt.Printf("欢迎下次使用BMI体脂计算器.\n\n")

	var (
		name, sex      string
		height, weight float64
		age            int
		ok             string
	)

	fmt.Println("**** BMI体脂计算器 ****")
Menu:
	fmt.Scanf("姓名: %s\n", &name)
	fmt.Scanf("性别(男/女): %s\n", &sex)
	fmt.Scanf("身高(米): %.2f\n", &height)
	fmt.Scanf("体重(公斤): %.2f\n", &weight)
	fmt.Scanf("年龄(岁): %d\n", &age)

	fmt.Scanf("是否继续输入用户信息(yes/no): %s\n", &ok)

	sex_code := 1
	switch sex {
	case "男":
		sex_code = 1
	case "女":
		sex_code = 0
	}

	SaveData(weight, height, age, sex_code, name)

	if ok == "yes" {
		goto Menu
	}

	Calculator()
}
