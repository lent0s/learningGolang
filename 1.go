package main

import (
	logrus "github.com/sirupsen/logrus"
)

func main() {

	var a int

	for i := 0; i <= 10; i++ {
		a += i * i
		//logrus.Infof("i = %v a = %v", i, a)
		logrus.WithFields(logrus.Fields{
			"i": i,
			"a": a,
		}).Info("значения переменных")
	}

} //{"level":"info","msg":"i = 0 a =  0","time":"2022-06-20T03:16:28+04:00"}

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
}
