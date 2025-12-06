package config

import (
	"example/hello/entity"
	"fmt"
	"time"

	"github.com/avast/retry-go/v5"
)

type GoRetry struct {
	Retrier *retry.Retrier
}

func NewGoRetry(retryIfFunc func(error) bool) *GoRetry {
	fmt.Println("START INIT RETRIER")
	// retryIf := retry.RetryIf(func(err error) bool {
	// 	if err.Error() == "special error" {
	// 		fmt.Println("2")
	// 		return false
	// 	}
	// 	return true
	// })

	retryIf := retry.RetryIf(retryIfFunc)
	retryAttempt := retry.Attempts(6)
	// retryDelay := retry.Delay(1 * time.Second)
	onretry := retry.OnRetry(func(n uint, err error) {
		// log.Printf("#%d: %s\n", n, err)
	})

	delayTypeFunc := retry.DelayTypeFunc(func(n uint, err error, config retry.DelayContext) time.Duration {

		switch err.Error() {
		case entity.ERRUnusual.Error():
			fmt.Printf("⚠️ [Error unusual case, retry immeadiatly\n")
			return 100 * time.Millisecond

		case entity.ERRBusy1.Error():
			fmt.Printf("⚠️[Error Aware Delay]Database Lock Detected: Long Delay (%v)\n", 3*time.Second)
			return 3 * time.Second

		case entity.ERRBusy2.Error():
			baseDelay := (500 * time.Millisecond)
			delay := baseDelay * time.Duration(1<<n)
			fmt.Printf("⚠️[Error Exponential] (%v)\n", delay)
			return delay

		}
		// if err.Error() == entity.ERRUnusual.Error() {
		// 	fmt.Printf("⚠️ [Error unusual case, retry immeadiatly ")
		// 	return 100 * time.Millisecond

		// } else if err.Error() == entity.ERRBusy1.Error() {
		// 	fmt.Printf("⚠️[Error Aware Delay]Database Lock Detected: Long Delay (%v)\n", 3*time.Second)
		// 	return 3 * time.Second

		// } else if err.Error() == entity.ERRBusy2.Error() {
		// 	fmt.Printf("⚠️[Error Exponential]")
		// 	baseDelay := 500 * time.Millisecond
		// 	return baseDelay * time.Duration(1<<n)

		// }

		// สำหรับ Error ชนิดอื่นๆ ที่ไม่ได้ระบุ ให้ใช้ Fixed Delay (เช่น 1 วินาที)
		return 1 * time.Second
	})

	delayType := retry.DelayType(delayTypeFunc)

	retrier := retry.New(retryAttempt, onretry, delayType, retryIf)

	return &GoRetry{
		Retrier: retrier,
	}
}
