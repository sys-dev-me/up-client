// retriving processes, lookup by name
package main


import "os/exec"
import "time"
import "fmt"
import "runtime"
import "os"
import "strings"

type PS struct {
	List		[]Process
	Update	time.Time
} 

type Process struct {
	
	Name			string
	Owner			string
	PID				int64
	pcpu			int
	pmem			int
	vsz				int64
	rss				int64
	tt				int64
	stat			string	
	started		string
	uptime		string

}


func (this *PS) Gathering() {
	cmd := exec.Command("ps", "aux")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if runtime.GOOS == "windows" {
		cmd = exec.Command("tasklist")
	}

	err := cmd.Run()
	if err != nil {
		fmt.Printf ( "Unable to retrive process list" )
		return
	}

	s := strings.Split( fmt.Sprintf ( "%v", cmd.Stdout ), "\r")

	for idx, v := range s {
		fmt.Printf ( "[%d]\t\t%20s\n", idx, v)	
	}
} 
