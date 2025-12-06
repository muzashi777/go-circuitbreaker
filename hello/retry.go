package hello

import (
	config "example/hello/_config"
	"example/hello/entity"
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
	for {

		err := retrier.Do(
			func() error {
				res := entity.ERRBusy2.Error()
				return fmt.Errorf(res)
			},
		)
		if err != nil {
			// handle error
			fmt.Println("END")
			return
		}
	}

}
