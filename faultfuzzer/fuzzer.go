package faultfuzzer

import (
	"bufio"
	"fmt"
	"github.com/google/syzkaller/pkg/log"
	"github.com/scylladb/go-set/strset"
	"io/ioutil"
	"os"
	"os/exec"
	//"strconv"
)

var queue = make([]string, 0)
var history = strset.New()
var current string
var max_cov int = 0
var flag int = 0
var maxfault = 0x3fffffff

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
func num_cov(cov string) int {
	num := 0
	for _, char := range cov {
		for i := 0; i < 8; i++ {
			if (char & (1 << i)) != 0 {
				num += 1
			}
		}
	}
	return num
}
func no_zero_cov(cov string) bool {
	for _, char := range cov {
		if char != 0 {
			return true
		}
	}
	return false
}

func reduce_ehc(ecov int) {
	num := ecov
	if num > max_cov {
		max_cov = num
	}
	f, err := os.OpenFile("/root/ehc_rec", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(fmt.Sprintf("%d\n", max_cov)); err != nil {
		panic(err)
	}
}

func Get_cover() {

	if flag == 0 {
		flag = 1
//		ecmd("~/reset_ehc_site")
	}
//	res := ecmd("~/get_ehc_fault_site")
//	ecov, _ := strconv.Atoi(res)
//	log.Logf(0, "ecov-----------%v\n", ecov)
//	reduce_ehc(ecov)

	ecmd("~/get_fault_site")

	cov, _ := ioutil.ReadFile("/dev/fault")
	log.Logf(0, "Rrooach: Covvvv")
	exists := history.Has(string(cov))
	if !exists && no_zero_cov(string(cov)) {
		queue = append(queue, string(cov))
		history.Add(string(cov))
		log.Logf(0, "--------------------------")
		log.Logf(0, "new cov")
	}
	//if num_bits(cov) > num_bits(current) {
	//	mutate()
	//}
}

func write_fault(fault string) {
	file, _ := os.Create("/dev/fault")
	writer := bufio.NewWriter(file)
	writer.WriteString(fault)
	writer.Flush()
}

func Set_fault() int {
	var fault string
	if len(queue) == 0 {
		log.Logf(0, "--------------------------")
		log.Logf(0, "queue is empty")
		fault = ""

		write_fault(fault)

		ecmd("~/set_fault")
		log.Logf(0, "--------------------------")
		log.Logf(0, "set fault")
		return 1
	}
	fault = queue[0]
	queue = queue[1:]
	current = fault
	write_fault(fault)
	ecmd("~/set_fault")
	log.Logf(0, "--------------------------")
	log.Logf(0, "set fault")
	return 0
}
