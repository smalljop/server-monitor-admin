package main

import (
	"encoding/json"
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	net2 "github.com/shirou/gopsutil/net"
	"log"
	"net"
	"time"
)

func init() {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

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

//获取客户端服务器信息
func GetStat() {
	stat := Stat{}
	getHost(&stat)
	getNet(&stat)
	getCpuInfo(&stat)
	getMemory(&stat)
	getDiskInfo(&stat)
	avg, _ := load.Avg()
	stat.Load = avg.Load1
	stat.TimeStamp = time.Now().Unix()
	marshal, _ := json.Marshal(stat)
	println(string(marshal))
	//hostInfo, err := host.Info()
	//println(hostInfo.Uptime)
	//println(hostInfo.Hostname)
	//host.PlatformInformation()
	//println(ip)
}

func getMemory(s *Stat) {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		log.Panicln(fmt.Printf("get mem info failed, err:%v", err))
	}
	s.RamRate = memInfo.UsedPercent
	swapMemInfo, err := mem.SwapMemory()
	s.SwapRate = swapMemInfo.UsedPercent
	// almost every return value is a struct
	//fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", memInfo.Total, memInfo.Free, memInfo.UsedPercent)
}

//获取cpu信息
func getCpuInfo(s *Stat) {
	cpuInfos, err := cpu.Info()
	if err != nil {
		log.Panicln(fmt.Printf("get cpu info failed, err:%v", err))
	}
	for _, ci := range cpuInfos {
		s.CPUName = ci.ModelName

	}
	counts, err := cpu.Counts(true)
	s.CPUCore = uint64(counts)
	// CPU使用率
	percent, _ := cpu.Percent(time.Second, false)
	s.CPURate = percent[0]
}

//获取网络相关
func getNet(stat *Stat) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Panicln("get net ip", err)
	}
	for _, addr := range addrs {
		ipAddr, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if ipAddr.IP.IsLoopback() {
			continue
		}
		if !ipAddr.IP.IsGlobalUnicast() {
			continue
		}
		stat.IpAddr = ipAddr.IP.String()
		break
	}
	info, _ := net2.IOCounters(true)
	if len(info) > 0 {
		stat.NetRead = FormatByteSize(int64(info[0].BytesRecv), 10)
		stat.NetWrite = FormatByteSize(int64(info[0].BytesSent), 10)
	}
	//for index, v := range info {
	//
	//	fmt.Printf("%v:%v send:%v recv:%v\n", index, v, v.BytesSent, v.BytesRecv)
	//}
}

//获取host相关信息
func getHost(stat *Stat) {
	hostInfo, err := host.Info()
	if err != nil {
		log.Panicln("get host", err)
		return
	}
	stat.HostName = hostInfo.Hostname
	stat.BootTime = hostInfo.BootTime
	stat.PlatformVersion = hostInfo.PlatformVersion
}

// disk info
func getDiskInfo(stat *Stat) {
	//parts, err := disk.Partitions(false)
	//if err != nil {
	//	fmt.Printf("get Partitions failed, err:%v\n", err)
	//	return
	//}
	//for _, part := range parts {
	//	/**/ fmt.Printf("part:%v\n", part.String())
	diskInfo, _ := disk.Usage("/")
	stat.DiskRate = diskInfo.UsedPercent
	//fmt.Printf("disk info:used:%v free:%v total:%v\n", diskInfo.UsedPercent, diskInfo.Free, diskInfo.Total)

	ioStat, _ := disk.IOCounters()
	diskRead := uint64(0)
	diskWrite := uint64(0)
	for _, v := range ioStat {
		diskRead += v.ReadBytes
		diskWrite += v.WriteBytes
		//fmt.Printf("%v:%v\n", k, v)
	}
	stat.DiskRead = FormatByteSize(int64(diskRead), 10)
	stat.DiskWrite = FormatByteSize(int64(diskWrite), 10)
}

func main() {
	//counters, _ := net2.IOCounters(false)
	//net2.Interfaces()
	//for index, v := range counters {
	//	fmt.Printf("%v:%v send:%v recv:%v\n", index, v, v.BytesSent, v.BytesRecv)
	//}
	//size := FormatByteSize(470923591, 10)
	//println(size)
	////println(size)

	//getDiskInfo(nil)
	//state := diskstate.DiskUsage("/")
	//fmt.Printf("All=%dM, Free=%dM, Available=%dM, Used=%dM, Usage=%d%%",
	//	state.All/diskstate.MB, state.Free/diskstate.MB, state.Available/diskstate.MB, state.Used/diskstate.MB, 100*state.Used/state.All)
	//getDiskInfo(&Stat{})
	GetStat()
}
