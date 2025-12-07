package hello

import (
	config "example/circuitbreaker/_config"
	"example/circuitbreaker/entity"
	"fmt"
)

func GoRetryWithGo() {
	funcA := func(err error) bool {

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

	retrier := config.NewGoRetry(funcA).Retrier
	data := ""
	for {

		err := retrier.Do(
			func() error {
				return entity.ERRUnusual
			},
		)
		if err != nil {
			// handle error
			fmt.Println("END")
			return
		}
		fmt.Println("END with success:", data)

		return
	}

}
