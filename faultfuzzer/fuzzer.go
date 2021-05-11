package faultfuzzer
// package main

import (
	// "bufio"
	// "math/rand" 
	// "bytes"
	// "encoding/binary"
	// "os"
	// "os/exec"
	// "github.com/google/syzkaller/pkg/log"
	// "github.com/scylladb/go-set/strset"
	"io/ioutil"
	"fmt" 
	"time"
    "strconv"

)

var ori_fault_seq string
var fault_set []string
var fault_len = 0
var fault_seq string 
var cnt = 0
var neg_cnt = 0

func fault_seq_cmp(cur_fault_seq string) bool {
	var cur_neg_cnt = 0
	for i := 0; i < fault_len; i++ {
		if cur_fault_seq[i] == '2' {
			cur_neg_cnt++
		}
	}
	if cur_neg_cnt < neg_cnt {
		ori_fault_seq = cur_fault_seq
		neg_cnt = cur_neg_cnt
		return true
	}
	return false
}


func BytesToInt(b []byte) int {
	var x = 0
	fmt.Println(b) 
	for _,i := range b {
		if i == 10 {
			break;
		}
		fmt.Println(int(i)-48)
		
		x = x*10
		x += (int(i)-48)
	}
	return int(x)
}
 

// read fault_len from file
func read_len() int {
	bytes, err := ioutil.ReadFile("/root/cover_uid")
	if err != nil {
		panic(err)
	} 
	cover := BytesToInt(bytes) 
	fmt.Println(cover)
	return cover
}

// generate fault_seq from file 
func gen_fault_seq() string{
	var fault_seq = ""
	// rand.Seed(time.Now().UnixNano())

	for i := 0; i < fault_len; i++ {
		// tmp_int :=rand.Intn(2)   
		tmp_int := 2 
		tmp_str :=  strconv.Itoa(tmp_int)
		fault_seq += tmp_str 
	} 

	fault_set = append(fault_set, fault_seq)
	return fault_seq
}

func generate_permutation(pre_rune []rune, left, right int) {
    if left == right {
		tmp_str := string(pre_rune)
		fault_set = append(fault_set, tmp_str)
        fmt.Println(string(pre_rune))
    } else {
        for i := left; i <= right; i++ {
            pre_rune[left], pre_rune[i] = pre_rune[i], pre_rune[left]
            generate_permutation(pre_rune, left+1, right)
            pre_rune[left], pre_rune[i] = pre_rune[i], pre_rune[left]
        }
    }
}

func Mutate_seq() { 
}


//mutate fault_seq when has new and write it into file
func mutate_seq() { 
    pre_rune := []rune(fault_seq)
    generate_permutation(pre_rune, 0, len(pre_rune)-1)
}

func write_to_file() {  
	fmt.Printf("tmp = %d ", cnt)
	err := ioutil.WriteFile("/dev/fault_seq", []byte(fault_set[cnt]), 0777)
	cnt++
    if err != nil {
        fmt.Printf("ioutil.WriteFile failure, err=[%v]\n", err)
    }
	// bytes, _ := ioutil.ReadFile("/dev/fault_seq")
	// fmt.Printf("faultseq = ")
	// fmt.Println(len(bytes))

}

//caculate time diff 
func Time_diff(pre_time int64) int {
	cur_time := time.Now().Unix()
	if pre_time - cur_time >= 300 {
		return 1
	} else {
		return 0
	}
}


func Init() {
	fault_len = read_len()
	neg_cnt = fault_len
	fault_seq = gen_fault_seq()
	write_to_file()
}
	