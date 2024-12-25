package crontab

import (
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"time"
)

var c *cron.Cron

func Init() {
	c = cron.New()
	id, err := c.AddFunc("*/10 * * * *", CrontabTest)
	logrus.WithFields(map[string]interface{}{
		"id":  id,
		"err": err,
	}).Info("crontab start")
	c.Start()
}

func Close() {
	if c != nil {
		c.Stop()
	}
}
func CrontabTest() {
	logrus.Info("start CrontabTest")
	time.Sleep(1 * time.Second)
	logrus.Info("end CrontabTest")
}
