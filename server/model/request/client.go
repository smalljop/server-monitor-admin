package request

//客户端服务器信息
type Stat struct {
	HostName        string  `json:"hostName"`        //Host名称
	IpAddr          string  `json:"ipAddr"`          //ip地址
	TimeStamp       int64   `json:"timeStamp"`       //同步时间
	PlatformVersion string  `json:"platformVersion"` //平台版本
	CPUName         string  `json:"cpuName"`         //CPu名称
	CPUCore         uint64  `json:"cpuCore"`         //cpu核心
	BootTime        uint64  `json:"uptime"`          // 运行时间
	Load            float64 `json:"load"`            //系统负载  任务队列长度
	CPURate         float64 `json:"cpuRate"`         //Cpu使用率
	RamRate         float64 `json:"ramRate"`         //内存使用率
	SwapRate        float64 `json:"swapRate"`        // 虚拟内存使用率
	DiskRate        float64 `json:"diskRate"`        //硬盘使用
	DiskRead        string  `json:"diskRead"`        //硬盘度
	DiskWrite       string  `json:"diskWrite"`       //硬盘写率
	NetRead         string  `json:"netRead"`         //上行网络 带单位
	NetWrite        string  `json:"netWrite"`        //下行网络
}
