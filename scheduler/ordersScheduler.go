package scheduler

import (
	"context"
	"github.com/procyon-projects/chrono"
)

type OrdersScheduler struct {
	cronExpression string
	task           chrono.ScheduledTask
}

func NewOrdersScheduler(cronExpression string) OrdersScheduler {
	return OrdersScheduler{cronExpression: cronExpression}
}

func (o *OrdersScheduler) Schedule() error {
	scheduler := chrono.NewDefaultTaskScheduler()
	task, err := scheduler.ScheduleWithCron(func(ctx context.Context) {

	}, o.cronExpression)
	if err != nil {
		o.task = task
	}
	return err
}

func (o *OrdersScheduler) Shutdown() {
	if o.task != nil {
		o.task.Cancel()
		o.task = nil
	}
}
