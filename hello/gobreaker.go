package hello

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
	5:  nil,
	6:  nil,
	7:  nil,
	8:  entity.ERRUnusual,
	9:  nil,
	10: nil,
	11: nil,
	12: entity.ERRUnusual,
	13: entity.ERRUnusual,
	14: entity.ERRUnusual,

	// 0:  nil,
	// 1:  entity.ERRForever,
	// 2:  entity.ERRForever,
	// 3:  entity.ERRForever,
	// 4:  entity.ERRForever,
	// 5:  entity.ERRForever,
	// 6:  entity.ERRForever,
	// 7:  entity.ERRForever,
	// 8:  entity.ERRForever,
	// 9:  nil,
	// 10: nil,
	// 11: nil,
	// 12: entity.ERRForever,
	// 13: entity.ERRForever,
	// 14: entity.ERRForever,
}

func mockCall(i int) ([]byte, error) {
	err := status[i]
	if err != nil {
		return nil, err
	}

	return []byte("Success"), nil
}

func GoBreakerWithGo() {
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
