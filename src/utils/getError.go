package utils

import "log"

func GetErro(err error) error {
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
