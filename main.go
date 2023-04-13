package main

import (
	"cb/pkg/client"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

const monthsAgoNum = 3

var (
	maxValue = -1.0
	maxName  = ""
	maxDate  = ""
	minValue = 99999.0
	minName  = ""
	minDate  = ""
	sumValue float64
	count    = 0
)

func main() {
	client := client.NewClient()

	for i := 0; i < 90; i++ {
		tmp := time.Now().AddDate(0, 0, -i)
		rate, err := client.GetRate(tmp)
		if err != nil {
			log.Fatalf("%s = %d", err, i)
		}

		for _, valute := range rate.Currencies {
			// Находим максимальное и минимальное значение курса валюты
			properFormattedValue := strings.Replace(valute.Value, ",", ".", -1)
			floatValue, _ := strconv.ParseFloat(properFormattedValue, 64)
			//if err != nil {
			//	return 0, err
			//}
			valCur := floatValue / float64(valute.Nom)
			if valCur > maxValue {
				maxValue = valCur
				maxDate = rate.Date
				maxName = valute.Name
			}
			if valCur < minValue {
				minValue = valCur
				minDate = rate.Date
				minName = valute.Name
			}

			// Добавляем значение курса валюты к сумме для расчета среднего значения
			sumValue += valCur
			count++
		}

		//fmt.Println(rate)
	}
	avgValue := sumValue / float64(count)

	// Выводим результаты
	fmt.Printf("Максимальный курс: %.5f рублей %s %s\n", maxValue, maxName, maxDate)
	fmt.Printf("Минимальный курс: %.5f рублей за %s %s\n", minValue, minName, minDate)
	fmt.Printf("Средний курс за последние 90 дней: %.2f рублей", avgValue)

}
