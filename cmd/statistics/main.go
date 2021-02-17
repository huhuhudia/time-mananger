package main

import (
	"encoding/json"
	"fmt"
	"github.com/huhuhudia/time-manager/def"
	"github.com/huhuhudia/time-manager/tool/model"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

func main(){
	start , err := time.ParseInLocation(def.TimeTmpl, "2020-01-02T15:04:05", time.Local)
	if err != nil{
		panic(err)
	}
	end , err := time.ParseInLocation(def.TimeTmpl, "2021-02-24T15:04:05", time.Local)
	if err != nil{
		panic(err)
	}
	res := []*model.Journal{}
	filepath.Walk("../../data", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			fmt.Println(info.Name())
			return nil
		}
		data ,err := ioutil.ReadFile(path)
		if err != nil{
			panic(err)
		}
		fmt.Println(string(data))
		tmp := new(model.Journal)
		err = json.Unmarshal(data, tmp)
		if err != nil{
			panic(err)
		}
		if tmp.DateUnixTime >= start.Unix() && tmp.DateUnixTime <= end.Unix(){
			res = append(res, tmp)
			fmt.Println(path)
		}
		return nil
	})
	statistics(res)
}

func statistics(journals []*model.Journal){
	m := make(map[string]float64,len(journals))
	for _, journal := range journals{
		for _, task := range  journal.Tasks{
			duration := task.EndTime - task.StartTime
			for _, tag := range task.Tags{
				m[tag] += float64(duration)/(60.0*60.0)
			}
		}
	}
	fmt.Println(m)
}
