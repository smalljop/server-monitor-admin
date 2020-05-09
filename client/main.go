package main

import "time"

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
	t := time.NewTicker(5 * time.Second)
	for range t.C {
		go sendStatToServer()
	}
}
