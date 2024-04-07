package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <domain> <interval-seconds>\n", os.Args[0])
		os.Exit(1)
	}

	domain := os.Args[1]

	interval := 1.0
	if len(os.Args) == 3 {
		f, err := strconv.ParseFloat(os.Args[2], 64)
		if err != nil {
			fmt.Printf("failed to parse interval(%s): %v\n", os.Args[2], err)
			os.Exit(1)
		}
		interval = f
	}

	duration := time.Millisecond * time.Duration(interval*1000)

	for {
		resolveDomain(domain)
		time.Sleep(duration)
	}
}

func resolveDomain(domain string) {
	startTime := time.Now()
	_, err := net.LookupIP(domain)
	elapsed := time.Since(startTime)

	if err != nil {
		fmt.Printf("%v : %v\n", startTime.Format(time.StampMilli), err)
		return
	}

	fmt.Printf("%v : %.3f ms\n", startTime.Format(time.StampMilli), elapsed.Seconds()*1000)
}
