package internal

import (
	"cb/internal/config"
	"cb/internal/service"
	"cb/internal/service/model"
	"cb/pkg/client"
	"fmt"
	tele "gopkg.in/telebot.v3"
	"log"
	"time"
)

type app struct {
	cfg  *config.Config
	serv service.Service
	bot  *tele.Bot
}

type App interface {
	Run()
}

func NewApp(cfg *config.Config) (App, error) {
	return &app{
		cfg:  cfg,
		serv: service.NewService(client.NewClient()),
	}, nil
}

func (a *app) Run() {
	bot, err := a.createBot()
	if err != nil {
		return
	}
	a.bot = bot
	a.bot.Start()
}

func (a *app) createBot() (*tele.Bot, error) {
	pref := tele.Settings{
		Token:  a.cfg.Token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	b.Handle("/course", func(c tele.Context) error {
		result, err := a.serv.GetMinMaxAvg()
		if err != nil {
			log.Println(err)
			return c.Send("Повторите запрос чуть позже")
		}
		return c.Send(a.toString(result))
	})

	b.Handle("/start", func(c tele.Context) error {
		return c.Send("Hi. This is a bot that shows:\n" +
			"-the value of the maximum currency rate, the name of this currency and the date of this maximum value\n" +
			"-the value of the minimum currency exchange rate, the name of this currency and the date of this minimum value\n" +
			"-the average value of the ruble exchange rate for the entire period for all currencies\n\n" +
			"If you write a command - /course")
	})

	return b, nil
}

func (a *app) toString(res model.Response) string {
	v1 := fmt.Sprintf("Максимальный курс: %.5f рублей %s %s\n", res.MaxValue, res.MaxName, res.MaxDate)
	v2 := fmt.Sprintf("Минимальный курс: %.5f рублей за %s %s\n", res.MinValue, res.MinName, res.MinDate)
	v3 := fmt.Sprintf("Средний курс за последние 90 дней: %.2f рублей", res.AvgValue)
	return v1 + v2 + v3
}
