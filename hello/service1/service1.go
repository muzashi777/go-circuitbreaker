package service1

import (
	config "example/hello/_config"
	"example/hello/entity"
	"fmt"
	"time"
)

var status map[int]error = map[int]error{

	0:  nil,
	1:  entity.ERRUnusual,
	2:  entity.ERRUnusual,
	3:  entity.ERRUnusual,
	4:  entity.ERRUnusual,
	5:  entity.ERRUnusual,
	6:  entity.ERRUnusual,
	7:  entity.ERRUnusual,
	8:  entity.ERRUnusual,
	9:  nil,
	10: nil,
	11: nil,
	12: entity.ERRForever,
	13: entity.ERRForever,
	14: entity.ERRForever,
}

func mockCall(i int) ([]byte, error) {

	retrier := config.NewGoRetry(funcA).Retrier
	for {
		err := retrier.Do(
			func() error {
				return status[i]
			},
		)
		if err != nil {
			fmt.Println("END")
			return nil, status[i]
		}

		return []byte("Success"), nil
	}
}

func Service1() {
	cb := config.NewGoBreaker("demo-breaker").CB

	for i := 0; i < 15; i++ {
		result, err := cb.Execute(func() ([]byte, error) {
			return mockCall(i)
		})

		fmt.Printf("Attempt %d | State: %s | Result: %v | Error: %v\n", i+1, cb.State(), string(result), err)
		fmt.Println("---------------------")

		time.Sleep(1 * time.Second)
	}
}

func funcA(err error) bool {

	switch err.Error() {
	case entity.ERRPermission.Error():
		fmt.Println("❌ [Error Premission case, return immeadiatly")
		return false

	case entity.ERRForever.Error():
		fmt.Println("❌ [Error Forever case, return immeadiatly")
		return false
	}

	return true
}
