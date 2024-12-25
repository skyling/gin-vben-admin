package lock

import (
	"github.com/sirupsen/logrus"
	"sync"
	"testing"
	"time"
)

func TestLock(t *testing.T) {
	Init()

	wg := sync.WaitGroup{}
	go func() {
		wg.Add(1)
		time.Sleep(1 * time.Second)
		locker, ctx, err := LockOpt("test", 30*time.Second, 1*time.Second, 2, false)
		logrus.Info("=========release lock", locker, ctx, err)
		wg.Done()
	}()
	locker, ctx, err := Lock("test", 10*time.Second, false)
	logrus.Info("=========test lock", locker, ctx, err)
	time.Sleep(2 * time.Second)
	locker.Release(ctx)
	logrus.Info("=========test r lock", locker, ctx, err)

	wg.Wait()
}
