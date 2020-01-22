package main


import "github.com/shirou/gopsutil/mem"

type RAM struct {

    Total		uint64
    Used		uint64
    Free		uint64
}

func ReadMemory() *RAM {

	v, _ := mem.VirtualMemory()
	a := new(RAM)
	a.Total = v.Total
	a.Used = v.Used
	a.Free = v.Free
	return a
}
