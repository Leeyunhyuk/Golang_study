package main

import (
    "fmt"
    "github.com/robfig/cron"
    "time"
)

//crontab ==> 매 분 30초 마다 실행
const (
    CronSpec = "30 * * * * *"
)

var (
    Data int
)

func PrintData() {
    t := time.Now().Format(time.ANSIC)
    fmt.Println(t, " : ", Data)
    Data++
}

func main() {
    c := cron.New()
    c.AddFunc(CronSpec, PrintData)
    c.Start()

	//10분간 출력
    time.Sleep(time.Minute * 10)
    c.Stop()
}