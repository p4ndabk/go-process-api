package mylog

import (
    "log"
    "os"
)

var LogFile *os.File

func InitLog() *os.File{
    logfile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Fatal(err)
    }

    log.SetOutput(logfile)

    LogFile = logfile
    return logfile
}