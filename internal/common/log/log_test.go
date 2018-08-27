package log

import (
	"errors"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

var tmpLogger *logrus.Logger

func init() {
	InitMock()
	//init logger without output
	InitNoOutput()

	//init tmpLogger
	tmpLogger = logrus.New()
	formatter := &logrus.TextFormatter{
		FullTimestamp: true,
	}

	//Set Info File
	infoFile, err := os.OpenFile("/dev/null", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println("Could not open log file - switching to normal output")
	}

	if infoFile != nil {
		tmpLogger.Out = infoFile
	}
	tmpLogger.Formatter = formatter

	log.SetOutput(infoFile)
}

func TestLog(t *testing.T) {
	Printf("test format: %v", "success")
	Println("basic println")

	Error(fmt.Errorf("this is err: %v", errors.New("intentional error")))
	Errorf(fmt.Errorf("this is err: %v", errors.New("intentional error")), "error with formatting %v", "lol")

	Debugf("similar to println but this is %v", "debug")
	Debugln("similar to println but this is debug")

	_ = SetVariable(map[string]interface{}{"test": "test"})

}

func BenchmarkLogInterface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Println("basic println")
	}
}

func BenchmarkLogrus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmpLogger.Println("basic println")
	}
}

func BenchmarkLogGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		log.Println("basic println")
	}
}
