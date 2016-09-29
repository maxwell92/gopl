package main

import (
	"fmt"
	logrus "github.com/Sirupsen/logrus"
	logrus_syslog "github.com/Sirupsen/logrus/hooks/syslog"
	"log/syslog"
	"os"
	"time"
)

func main() {

	//printx()
	//test1()
	//stdlogger()
	textformat()
}

func printx() {
	fmt.Println("Printx:-----------------")
	logrus.Warnln("Warning")
	logrus.Debugln("Debug")
	logrus.Errorln("Error")
	logrus.Infoln("Info")
	//logrus.Panicln("Panic")
	//logrus.Fatalln("Fatal")

	fmt.Println("Printx:-----------------")

	logrus.WithFields(logrus.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stderr)
	logrus.SetLevel(logrus.WarnLevel)

	logrus.WithFields(logrus.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	fmt.Println("Printx:-----------------")
	contextLogger := logrus.WithFields(logrus.Fields{
		"common": "this is a common field",
		"other":  "I also should be logged always",
	})
	contextLogger.Info("I'll be logged with common and other field")

	/*
		fmt.Println("Printx:-----------------")
		log       := logrus.New()
		hook, err := logrus_syslog.NewSyslogHook("udp", "localhost:514", syslog.LOG_INFO, "")

		if err == nil {
			log.Hooks.Add(hook)
		}
	*/

	fmt.Println("Printx:-----------------")
	log := logrus.New()
	hook, err := logrus_syslog.NewSyslogHook("", "", syslog.LOG_INFO, "")

	if err == nil {
		log.Hooks.Add(hook)
	}

	logrus.WithFields(logrus.Fields{
		"code": 3,
		"msg":  "test",
	}).Info("This is magic group")

	fmt.Printf("%v\n", logrus.IsTerminal())
	fmt.Printf("%v\n", logrus.GetLevel())

}

func stdlogger() {
	/*
		s := new(logrus.StdLogger) // a pointer to interface, not a interface
			s.Print("StdLogger")
			s.Println("StdLogger")
			s.Printf("StdLogger")
			s.Fatal("StdLogger")
			s.Fatalln("StdLogger")
			s.Fatalf("StdLogger")
			s.Panic("StdLogger")
			s.Panicln("StdLogger")
			s.Panicf("StdLogger")
	*/

}

func textformat() {
	tf := &logrus.TextFormatter{
		//ForceColors:   true,
		DisableColors: true,
		FullTimestamp: true,
	}

	l := logrus.New()
	e := logrus.NewEntry(l)

	bytes, err := tf.Format(e)
	if err != nil {
		logrus.Debugln(err)
		return
	}
	logrus.Println(string(bytes))

	e.Logger.Println("textformat")
}

type myHook struct {
}

func (m *myHook) Levels() []logrus.Level {
	level := []logrus.Level{logrus.InfoLevel}
	return level
}

func (m *myHook) Fire(e *logrus.Entry) error {
	//l.Logger.Println(l.Message)
	e.Logger.Println(e.Message)
	return nil
}

func test1() {
	fmt.Println("Test1:-----------------")
	hook := new(myHook)

	l := new(logrus.Logger)

	e := &logrus.Entry{
		Logger:  l,
		Time:    time.Now(),
		Level:   logrus.InfoLevel,
		Message: "MyEntry",
	}

	hook.Fire(e)
}
