package Log

import (
	"os"
	"time"

	Colour "github.com/alan59a/a08/Log/Color"
)

type Log struct {
	Error     string
	Err       error
	Warning   string
	Success   string
	Operation string
	Time      int64
}

type LogBook struct {
	ID   int64
	Logs []Log
	Type string
}

func NewLog(err, war, suc string, err2 error) *Log {
	log := Log{
		Error:     err,
		Err:       err2,
		Warning:   war,
		Success:   suc,
		Operation: "",
		Time:      0,
	}
	if err != "" && err2 != nil {
		log.Error = "[_red_][black] ERROR [reset] " + err + ": \n" + err2.Error()
		log.Err = err2
	}
	if err != "" && err2 == nil {
		log.Error = "[_red_][black]  ERROR  [reset] " + err
		log.Err = nil
	}
	if war != "" {
		log.Warning = "[_yellow_][black] WARNING [reset] " + war + "."
	}
	if suc != "" {
		log.Success = "[_green_][black] SUCCESS [reset] " + suc + "."
	}
	log.Time = time.Now().UnixNano()
	return &log
}

func NewLogBook(functionType string) *LogBook {
	return &LogBook{
		ID:   time.Now().UnixMilli(),
		Logs: []Log{},
		Type: functionType,
	}
}

func (b *LogBook) LogEvent(log *Log) {
	b.Logs = append(b.Logs, *log)
}

func (b *LogBook) LogError(err error, errorText string) {
	l := NewLog(errorText, "", "", err)
	b.LogEvent(l)
}

func (b *LogBook) LogWarning(warning string) {
	l := NewLog("", warning, "", nil)
	b.LogEvent(l)
}

func (b *LogBook) LogSuccess(success string) {
	l := NewLog("", "", success, nil)
	b.LogEvent(l)
}

func (b *LogBook) Logger(err error, errorText, successText string) {
	if err != nil {
		b.LogError(err, errorText)
		return
	}
	b.LogSuccess(successText)
}

func (b *LogBook) ReadLogBook() {
	for _, a := range b.Logs {
		Colour.Print("[cyan] " + time.Unix(0, a.Time).Format("2006-01-02 15:04:05 -0700") + " " + "[reset] - ")
		if a.Error != "" {
			Colour.Println(a.Error)
		}
		if a.Warning != "" {
			Colour.Println(a.Warning)
		}
		if a.Success != "" {
			Colour.Println(a.Success)
		}
	}
}

func (l *Log) Report() {
	Colour.Print("[cyan] " + time.Unix(0, l.Time).Format("2006-01-02 15:04:05 -0700") + " " + "[reset] - ")
	if l.Error != "" {
		Colour.Println(l.Error)
	}
	if l.Warning != "" {
		Colour.Println(l.Warning)
	}
	if l.Success != "" {
		Colour.Println(l.Success)
	}
}

func (l *Log) FatalReport() {
	Colour.Print("[cyan] " + time.Unix(0, l.Time).Format("2006-01-02 15:04:05 -0700") + " " + "[reset] - ")
	if l.Error != "" {
		Colour.Println(l.Error)
		os.Exit(1)
	}
	if l.Warning != "" {
		Colour.Println(l.Warning)
	}
	if l.Success != "" {
		Colour.Println(l.Success)
	}
}
