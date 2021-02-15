package model

var Template  = Journal{
	Weather: "☀️",
	Tasks: []*Task{
		&Task{
			Desc: "golang内存分配， 虚拟内存",
			Tags: []string{
				Os,
				Golang,
			},
			StartTimeStr: "2021-02-15T15:00:00",
			EndTimeStr:"2021-02-15T16:00:00",
		},
		&Task{
			Desc: "计算机网络 epoll select 原理 ",
			Tags: []string{
				Net,
				Linux,
			},
			StartTimeStr: "2021-02-15T17:00:00",
			EndTimeStr:"2021-02-15T18:00:00",
		},
	},
}
