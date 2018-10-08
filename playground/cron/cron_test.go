package cron

import (
	"fmt"
	"testing"
	"time"

	"github.com/edgexfoundry/edgex-go/pkg/models"
	"github.com/robfig/cron"
)

type Job struct {
	schedule       models.Schedule
	scheduleEvents []models.ScheduleEvent
}

func (job Job) Run() {
	fmt.Println(job.schedule.Name, job.schedule.Frequency)
}

func TestCronWithAddJob(t *testing.T) {
	var job = Job{
		schedule: models.Schedule{
			Id:        "xxx",
			Name:      "5sec-schedule",
			Frequency: "PT5S",
		},
		scheduleEvents: []models.ScheduleEvent{},
	}
	var job2 = Job{
		schedule: models.Schedule{
			Id:        "xxx",
			Name:      "2sec-schedule",
			Frequency: "PT2S",
		},
		scheduleEvents: []models.ScheduleEvent{},
	}

	// init cron
	c := cron.New()

	// add cron job
	var spec = fmt.Sprintf("@every %v", ParseDuration(job.schedule.Frequency))
	c.AddJob(spec, job)

	spec = fmt.Sprintf("@every %v", ParseDuration(job2.schedule.Frequency))
	c.AddJob(spec, job2)

	// start cron
	c.Start()

	time.Sleep(10 * time.Second)
	// keep alive
	//select {}
}

func TestCronWithAddFunc(t *testing.T) {
	// init cron
	c := cron.New()

	// add cron job
	var duration = ParseDuration("PT2S")
	var spec = fmt.Sprintf("@every %v", duration)

	c.AddFunc(spec, func() {
		// @every 2s
		fmt.Println(spec)
	})

	// start cron
	c.Start()

	// keep alive
	select {}
}

func TestParseISO8601(t *testing.T) {
	var duration = ParseDuration("PT2S")

	// PT2S -> 2s
	fmt.Println(duration)
	// PT15M -> 15m0s
	fmt.Println(ParseDuration("PT15M"))
	// P12Y4MT15M -> 108000h15m0s
	fmt.Println(ParseDuration("P12Y4MT15M"))
}
