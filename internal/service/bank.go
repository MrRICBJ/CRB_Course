package service

import (
	"cb/internal/service/model"
	"cb/pkg/client"
	"log"
	"strconv"
	"strings"
	"time"
)

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

type service struct {
	client client.Client
}

type Service interface {
	GetMinMaxAvg() (model.Response, error)
}

func NewService(cli client.Client) Service {
	return &service{
		client: cli,
	}
}

func (s *service) GetMinMaxAvg() (model.Response, error) {
	for i := 0; i < 90; i++ {
		tmp := time.Now().AddDate(0, 0, -i)
		rate, err := s.client.GetRate(tmp)
		if err != nil {
			log.Fatalf("%s = %d", err, i)
		}

		for _, valute := range rate.Currencies {
			// Находим максимальное и минимальное значение курса валюты
			properFormattedValue := strings.Replace(valute.Value, ",", ".", -1)
			floatValue, err := strconv.ParseFloat(properFormattedValue, 64)
			if err != nil {
				return model.Response{}, err
			}
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

	}
	avgValue := sumValue / float64(count)

	return model.Response{
		maxValue,
		sumValue,
		minValue,
		avgValue,
		maxName,
		maxDate,
		minName,
		minDate,
		count,
	}, nil
}
