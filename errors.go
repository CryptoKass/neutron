package neutron

import "log"

func Handle(err error) {
	if err != nil {
		panic(err)
	}
}

func Warn(err error) {
	log.Println("warning-err", err)
}
