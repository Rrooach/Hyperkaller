package faultfuzzer

import (
	//"github.com/google/syzkaller/pkg/log"
	"os/exec"
    "strconv"
)


func ecmd(cmd string) string {
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		panic("some error found")
	}
	return string(out)
}
func Get_cover() {
	res := ecmd("~/get_fault_site")
    strconv.Atoi(res)
}

func Set_fault() int {
    var fault int64
    fault=0
    ecmd("~/set_fault "+strconv.FormatInt(int64(fault),10))
    return 0
}
