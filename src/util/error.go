package util

import "github.com/sirupsen/logrus"

//捕获异常 error
func Catch(err error) {
	if err != nil {
		logrus.WithError(err).Println("catch error")
		panic(err)
	}
}
