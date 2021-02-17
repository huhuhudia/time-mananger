package model

var Today = "2021-02-17T"

func init() {

	for _, task := range Template.Tasks {
		task.StartTimeStr = Today + task.StartTimeStr
		task.EndTimeStr = Today + task.EndTimeStr
	}

}

var Template = Journal{
	Weather: "☀️",
	Tasks: []*Task{
		&Task{
			Desc: "protobuf",
			Tags: []string{
				Net,
			},
			StartTimeStr: "11:00:00",
			EndTimeStr:   "12:00:00",
		},
		&Task{
			Desc: "跑步，思考记忆方法，创业",
			Tags: []string{
				Think,
				Run,
				Exercise,
			},
			StartTimeStr: "14:00:00",
			EndTimeStr:   "18:00:00",
		},
		&Task{
			Desc: "golang 垃圾回收 栈内存管理",
			Tags: []string{
				Golang,
			},
			StartTimeStr: "19:00:00",
			EndTimeStr:   "20:30:00",
		},
		&Task{
			Desc: "运动休息",
			Tags: []string{
				SitUps,
				Exercise,
			},
			StartTimeStr: "20:30:00",
			EndTimeStr:   "20:45:00",
		},
		&Task{
			Desc: "二分查找 堆 avl 红黑树 b树",
			Tags: []string{
				Algorithm,
			},
			StartTimeStr: "21:45:00",
			EndTimeStr:   "20:45:00",
		},

		&Task{
			Desc: "mysql面试题",
			Tags: []string{
				DB,
			},
			StartTimeStr: "21:45:00",
			EndTimeStr:   "20:45:00",
		},
	},
}
