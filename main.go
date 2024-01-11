package main

import (
	"github.com/buding00/bScan/utils"
	"golang.org/x/net/icmp"
)

func main() {
	//start := time.Now()
	//time.Sleep(1 * time.Second)
	utils.Banner()
	//fmt.Printf("[*] 扫描结束,耗时: %s\n", time.Since(start))
	conn, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
}
