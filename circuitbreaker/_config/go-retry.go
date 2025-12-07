package config

import (
	"example/circuitbreaker/entity"
	"fmt"
	"log"
	"time"

	"github.com/avast/retry-go/v5"
)

type GoRetry struct {
	Retrier *retry.Retrier
}

func NewGoRetry(retryIfFunc func(error) bool) *GoRetry {
	fmt.Println("---- START INIT RETRIER ----")
	delay := retry.Delay(1 * time.Second)
	maxJitter := retry.MaxJitter(100 * time.Millisecond)
	maxDelay := retry.MaxDelay(1 * time.Second)
	delayType := retry.DelayType(retry.DelayTypeFunc(delayFunction)) // retry.BackOffDelay() retry.FixedDelay()

	retryIf := retry.RetryIf(retryIfFunc)
	retryAttempt := retry.Attempts(6)
	onretry := retry.OnRetry(func(n uint, err error) {
		log.Printf("Do someting  #%d: %s\n", n+1, err)
	})
	attemptsForError := retry.AttemptsForError(3, entity.ERRUnusual)

	retrier := retry.New(
		delay,
		maxJitter,
		maxDelay,
		retryAttempt,
		attemptsForError,
		delayType,
		retryIf,
		onretry,
	)

	return &GoRetry{
		Retrier: retrier,
	}
}

func delayFunction(n uint, err error, config retry.DelayContext) time.Duration {

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
	return 1 * time.Second
}
