package api

import (
	"fmt"
	"runtime"
	"time"
	"tnvd/controllers/lib"
)

func init() {
	ApiNS.Router("/sysStatus", &DashBoard{}, "get:SysStatus")
}

type DashBoard struct {
	lib.AuthController
}

var startTime = time.Now()

var sysStatus struct {
	Overall struct {
		OS        string `json:"操作系统"`
		Arch      string `json:"cpu架构"`
		CPUNum    int    `json:"cpu数量"`
		GoVersion string `json:"go版本"`

		Uptime       string `json:"已运行时间"`
		NumGoroutine int    `json:"goroutine数量"`
		OtherSys     string // other system allocations
	} `json:"总览"`

	MemStates struct {
		// General statistics.
		MemAllocated string `json:"已使用内存"`    // bytes allocated and still in use
		MemTotal     string `json:"已分配内存"`    // bytes allocated (even if freed)
		MemSys       string `json:"从系统中获得内存"` // bytes obtained from system (sum of XxxSys below)
		Lookups      uint64 `json:"指针查询数量"`   // number of pointer lookups
		MemMallocs   uint64 `json:"malloc数量"` // number of mallocs
		MemFrees     uint64 `json:"free数量"`   // number of frees
	} `json:"内存信息"`
	HeapStates struct {
		// Main allocation heap statistics.
		HeapAlloc    string `json:"堆内存分配"` // bytes allocated and still in use
		HeapSys      string `json:"堆内存获得"` // bytes obtained from system
		HeapIdle     string `json:"堆内存空闲"` // bytes in idle spans
		HeapInuse    string `json:"堆内存使用"` // bytes in non-idle span
		HeapReleased string `json:"堆释放"`   // bytes released to the OS
		HeapObjects  uint64 `json:"堆对象数量"` // total number of allocated objects
	} `json:"堆信息"`
	StackStates struct {
		// Low-level fixed-size structure allocator statistics.
		//	Inuse is bytes used now.
		//	Sys is bytes obtained from system.
		StackInuse  string `json:"栈内存使用"` // bootstrap stacks
		StackSys    string `json:"栈内存分配"`
		MSpanInuse  string // mspan structures
		MSpanSys    string
		MCacheInuse string // mcache structures
		MCacheSys   string
		BuckHashSys string // profiling bucket hash table

	} `json:"栈信息"`

	GCStates struct {
		// Garbage collector statistics.
		GCSys        string `json:"GC元数据大小"`  // GC metadata
		NextGC       string `json:"下一次GC大小"`  // next run in HeapAlloc time (bytes)
		LastGC       string `json:"上一次GC距现在"` // last run in absolute time (ns)
		PauseTotalNs string `json:"GC总暂停时间"`
		PauseNs      string `json:"上一次GC暂停时间"` // circular buffer of recent GC pause times, most recent at [(NumGC+255)%256]
		NumGC        uint32 `json:"GC次数"`
	} `json:"GC信息"`
}

func (this *DashBoard) SysStatus() {

	sysStatus.Overall.GoVersion = runtime.Version()
	sysStatus.Overall.CPUNum = runtime.NumCPU()
	sysStatus.Overall.OS = runtime.GOOS
	sysStatus.Overall.Arch = runtime.GOARCH
	sysStatus.Overall.Uptime = lib.TimeSincePro(startTime)

	m := new(runtime.MemStats)
	runtime.ReadMemStats(m)
	sysStatus.Overall.NumGoroutine = runtime.NumGoroutine()

	sysStatus.MemStates.MemAllocated = lib.FileSize(int64(m.Alloc))
	sysStatus.MemStates.MemTotal = lib.FileSize(int64(m.TotalAlloc))
	sysStatus.MemStates.MemSys = lib.FileSize(int64(m.Sys))
	sysStatus.MemStates.Lookups = m.Lookups
	sysStatus.MemStates.MemMallocs = m.Mallocs
	sysStatus.MemStates.MemFrees = m.Frees

	sysStatus.HeapStates.HeapAlloc = lib.FileSize(int64(m.HeapAlloc))
	sysStatus.HeapStates.HeapSys = lib.FileSize(int64(m.HeapSys))
	sysStatus.HeapStates.HeapIdle = lib.FileSize(int64(m.HeapIdle))
	sysStatus.HeapStates.HeapInuse = lib.FileSize(int64(m.HeapInuse))
	sysStatus.HeapStates.HeapReleased = lib.FileSize(int64(m.HeapReleased))
	sysStatus.HeapStates.HeapObjects = m.HeapObjects

	sysStatus.StackStates.StackInuse = lib.FileSize(int64(m.StackInuse))
	sysStatus.StackStates.StackSys = lib.FileSize(int64(m.StackSys))
	sysStatus.StackStates.MSpanInuse = lib.FileSize(int64(m.MSpanInuse))
	sysStatus.StackStates.MSpanSys = lib.FileSize(int64(m.MSpanSys))
	sysStatus.StackStates.MCacheInuse = lib.FileSize(int64(m.MCacheInuse))
	sysStatus.StackStates.MCacheSys = lib.FileSize(int64(m.MCacheSys))
	sysStatus.StackStates.BuckHashSys = lib.FileSize(int64(m.BuckHashSys))
	sysStatus.GCStates.GCSys = lib.FileSize(int64(m.GCSys))
	sysStatus.Overall.OtherSys = lib.FileSize(int64(m.OtherSys))

	sysStatus.GCStates.NextGC = lib.FileSize(int64(m.NextGC))
	sysStatus.GCStates.LastGC = fmt.Sprintf("%.1fs", float64(time.Now().UnixNano()-int64(m.LastGC))/1000/1000/1000)
	sysStatus.GCStates.PauseTotalNs = fmt.Sprintf("%.1fs", float64(m.PauseTotalNs)/1000/1000/1000)
	sysStatus.GCStates.PauseNs = fmt.Sprintf("%.3fs", float64(m.PauseNs[(m.NumGC+255)%256])/1000/1000/1000)
	sysStatus.GCStates.NumGC = m.NumGC
	this.Data[`json`] = sysStatus
	this.ServeJson()
}
