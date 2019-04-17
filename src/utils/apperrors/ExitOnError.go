package apperrors

import "log"

func ExitOnError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
