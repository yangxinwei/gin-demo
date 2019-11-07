package utils

import (
	"bytes"
	"fmt"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
)

const (
	FieldKeyMsg        = "msg"
	FieldKeyLevel      = "level"
	FieldKeyHttpStatus = "status"
	FieldKeyIP         = "ip"
	FieldKeyPath       = "path"
)

type CustomFormatter struct {
	Context         *gin.Context
	TimestampFormat string
}

func NewLogger() *logrus.Logger {
	//f, err := os.OpenFile("/var/log/logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	//if err != nil {
	//	panic(err)
	//}
	custom := new(CustomFormatter)
	custom.TimestampFormat = "2006-01-02 15:04:05"
	logger := logrus.New()
	logger.Out = os.Stdout
	logger.Formatter = custom
	return logger
}

func (f *CustomFormatter) appendKeyValue(b *bytes.Buffer, key string, value interface{}) {
	if b.Len() > 0 {
		b.WriteByte(' ')
	}
	b.WriteString(key)
	b.WriteByte('=')
	f.appendValue(b, value)
}
func (f *CustomFormatter) appendValue(b *bytes.Buffer, value interface{}) {
	stringVal, ok := value.(string)
	if !ok {
		stringVal = fmt.Sprint(value)
	}

	b.WriteString(stringVal)
}
func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	keys := make([]string, 0, len(entry.Data))
	for k := range entry.Data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	b.WriteString(entry.Time.Format(f.TimestampFormat))
	f.appendKeyValue(b, FieldKeyLevel, entry.Level.String())
	/*
		f.appendKeyValue(b, FieldKeyPath, f.Context.Request.URL.Path)
		f.appendKeyValue(b, FieldKeyHttpStatus, f.Context.Writer.Status())
		f.appendKeyValue(b, FieldKeyIP, f.Context.ClientIP())
	*/
	if entry.Message != "" {
		f.appendKeyValue(b, FieldKeyMsg, entry.Message)
	}
	for _, key := range keys {
		f.appendKeyValue(b, key, entry.Data[key])
	}
	b.WriteByte('\n')
	return b.Bytes(), nil
	/*
		b.WriteByte(' ')
		b.WriteByte('[')
		b.WriteString(strings.ToUpper(entry.Level.String()))
		b.WriteString("]:")

		if len(entry.Message) != 0 {
			b.WriteString(" | ")
			b.WriteString(entry.Message)
		}

		if len(entry.Data) > 0 {
			b.WriteString(" || ")
		}
		for key, value := range entry.Data {
			b.WriteString(key)
			b.WriteByte('=')
			b.WriteByte('{')
			fmt.Fprint(b, value)
			b.WriteString("}, ")
		}

		b.WriteByte('\n')
		return b.Bytes(), nil
	*/
}
