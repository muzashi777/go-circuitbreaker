package config

import (
	"example/hello/entity"
	"fmt"
	"time"

	"github.com/sony/gobreaker/v2"
)

type GoBreaker struct {
	CB *gobreaker.CircuitBreaker[[]byte]
}

func NewGoBreaker(name string) *GoBreaker {
	fmt.Println("START INIT BREAKER: ", name)
	var st gobreaker.Settings = gobreaker.Settings{
		Name:        name,
		MaxRequests: 2,
		Interval:    8 * time.Second, // interval time to reset count in the close state
		/*
			On the half-open stage
			 - minumum times that need to success request pass
			 - defeaul equal to 1
			for example:
			 - if it set to 3, success 2 and fail 1 , then the status will back to open state
		*/
		BucketPeriod: 2 * time.Second, // use with Interval to seperate sliding window, Interval will automatically adjust to be multiple of bucket
		/*
			- ระยะเวลาพัก (Open state)
			- after open state will be half open stat.
			- default equal to 60 seconds
		*/
		Timeout: 2 * time.Second,
		/*
			- ฟังก์ชันที่จะถูกเรียกเมื่อเกิดความล้มเหลวในสถานะ Closed
			- if nil, default ConsecutiveFailures equal to 5
		*/
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			// fmt.Println("Requests:", counts.Requests)
			// fmt.Println("TotalSuccesses:", counts.TotalSuccesses)
			// fmt.Println("TotalFailures:", counts.TotalFailures)
			// fmt.Println("ConsecutiveSuccesses:", counts.ConsecutiveSuccesses) // ไม่แสดง
			// fmt.Println("ConsecutiveFailures:", counts.ConsecutiveFailures)
			return counts.ConsecutiveFailures >= 4
		},
		OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
			fmt.Printf("\n--- ⚡️ STATE CHANGE: %s -> %s ⚡️ ---\n", from, to)
		},

		// detect status: true=> pass(return and not count as fail), false=> not pass(cound as fail)
		// detecting status and bypass
		IsSuccessful: func(err error) bool {
			switch err {
			case entity.ERRForever:
				fmt.Println("❌ [Error Forever case, ➡️ Success ✅")
				return true
			case entity.ERRPermission:

				fmt.Println("❌ [Error Premission case, ➡️ Success ✅")
				return true

			case nil:
				fmt.Println(" ➡️ Success ✅")
				return true
			}

			return false
		},
	}

	return &GoBreaker{
		CB: gobreaker.NewCircuitBreaker[[]byte](st),
	}
}
