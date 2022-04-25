package shared

import "log"

func ActOn(err error) {
	if err != nil {
		log.Fatal(err) // or panic or os.Exit
	}
}
