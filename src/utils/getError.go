package utils

import "log"

func GetErro(err error) error {
	log.Fatal(err.Error())
	return err
}
