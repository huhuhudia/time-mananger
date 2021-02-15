package main

import (
	"encoding/json"
	"fmt"
	"github.com/huhuhudia/time-manager/def"
	"github.com/huhuhudia/time-manager/tool/model"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"time"
)

func main() {
	t := model.Template
	ParseTime(&t)
	sort.SliceStable(t.Tasks, func(i, j int) bool {
		return t.Tasks[i].StartTime < t.Tasks[j].StartTime
	})
	Persist(&t)
}

func ParseTime(journal *model.Journal) {
	journal.DateStr = fmt.Sprintf("%04d-%02d-%02dT00:00:00", time.Now().Year(), time.Now().Month(), time.Now().Day())
	dateTime, err := time.ParseInLocation(def.TimeTmpl, journal.DateStr, time.Local)
	if err != nil {
		panic(err)
	}
	journal.DateUnixTime = dateTime.Unix()
	for _, task := range journal.Tasks {
		startTime, err := time.ParseInLocation(def.TimeTmpl, task.StartTimeStr, time.Local)
		if err != nil {
			panic(err)
		}
		endTime, err := time.ParseInLocation(def.TimeTmpl, task.EndTimeStr, time.Local)
		if err != nil {
			panic(err)
		}
		task.StartTime = startTime.Unix()
		task.EndTime = endTime.Unix()
	}
}

func Persist(journal *model.Journal) {
	js, _ := json.MarshalIndent(journal, " ", "\t")

	dateTime, err := time.Parse(def.TimeTmpl, journal.DateStr)
	if err != nil {
		panic(err)
	}

	year, month, day := dateTime.Year(), dateTime.Month(), dateTime.Day()
	baseDir := "../data"
	fname := filepath.Join(baseDir, fmt.Sprintf("%04d/%02d/%02d.json", year, month, day))

	dirname := filepath.Dir(fname)
	// check dir exist
	if _ , err := os.Stat(dirname);err != nil{
		err := os.MkdirAll(dirname, 0777)
		if err != nil{
			panic(err)
		}
	}
	err = ioutil.WriteFile(fname, js, 0777)
	if err != nil {
		panic(err)
	}
}
