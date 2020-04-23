package faultfuzzer

import (
	"github.com/google/syzkaller/pkg/log"
	"os/exec"
	"strconv"
)

var queue = make([]int64, 0)
var history = make(map[int64]bool)
var current int64

func ecmd(cmd string) string {
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		panic("some error found")
	}
	return string(out)
}

func mutate() {

}

func num_bits(num int64) int {
	var ret int
	for i := 0; i < 64; i++ {
		if num&(1<<i) != 0 {
			ret++
		}
	}
	return ret
}

func Get_cover() {
	res := ecmd("~/get_fault_site")
	cov, err := strconv.ParseInt(res, 10, 64)
	_ = err
	//exists := history[cov]
	//if !exists {
    queue = append(queue, cov)
    history[cov] = true
    log.Logf(0, "--------------------------")
    log.Logf(0, "old cov: %v",current)
    log.Logf(0, "new cov: %v",cov)
	//}
	if num_bits(cov) > num_bits(current) {
		mutate()
	}
}

func Set_fault() int {
	var fault int64
	if len(queue) == 0 {
        log.Logf(0, "--------------------------")
        log.Logf(0, "queue is empty")
		queue = append(queue, 0)
		return 1
	}
	fault = queue[0]
	queue = queue[1:]
	current = fault
	ecmd("~/set_fault " + strconv.FormatInt(int64(fault), 10))
	log.Logf(0, "--------------------------")
	log.Logf(0, "set fault to %v", fault)
	return 0
}
