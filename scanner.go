package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	host := flag.String("h", "", "Specify hostname")
	standardScan := flag.Bool("s", false, "Standard Scan (first 1500 ports)")
	allPorts := flag.Bool("a", false, "Scan all ports")
	portsList := flag.String("p", "", "Specify individual ports comma-separated (e.g., 22,80,135)")
	threads := flag.Int("t", 20, "Number of concurrent threads for scanning (maximum 50)")
	help := flag.Bool("help", false, "Show usage information")

	flag.Parse()

	if *help || len(os.Args) == 1 {
		fmt.Println("Usage:")
		fmt.Println("  -h    Specify hostname")
		fmt.Println("  -s    Standard Scan (first 1500 ports)")
		fmt.Println("  -a    Scan all ports")
		fmt.Println("  -p    Specify individual ports comma-separated (e.g., 22,80,135)")
		fmt.Println("  -t    Number of concurrent threads for scanning (maximum 50, default 20)")
		return
	}

	if *host == "" {
		fmt.Println("Error: Hostname must be specified (with -h).")
		return
	}

	var ports []int
	if *standardScan {
		for i := 1; i <= 1500; i++ {
			ports = append(ports, i)
		}
	} else if *allPorts {
		for i := 1; i <= 65535; i++ {
			ports = append(ports, i)
		}
	} else if *portsList != "" {
		portStrs := strings.Split(*portsList, ",")
		for _, portStr := range portStrs {
			portStr = strings.TrimSpace(portStr)
			port, err := strconv.Atoi(portStr)
			if err != nil {
				fmt.Printf("Invalid port: %s\n", portStr)
				return
			}
			ports = append(ports, port)
		}
	} else {
		fmt.Println("Error: Either -s, -a, or -p must be specified.")
		return
	}

	if *threads < 1 {
		fmt.Println("Error: Number of threads must be at least 1.")
		return
	} else if *threads > 50 {
		fmt.Println("Warning: Number of threads exceeds maximum (50). Setting to maximum.")
		*threads = 50
	}

	wg := &sync.WaitGroup{}
	sem := make(chan struct{}, *threads)

	for _, port := range ports {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()

			hostname := fmt.Sprintf("%s:%d", *host, port)
			conn, err := net.DialTimeout("tcp", hostname, 3*time.Second)
			if err != nil {
				return
			}
			defer conn.Close()

			fmt.Printf("Port open: %d\n", port)

			conn.SetReadDeadline(time.Now().Add(5 * time.Second))

			buffer := make([]byte, 1024)
			n, err := conn.Read(buffer)
			if err != nil {
				return
			}

			banner := string(buffer[:n])
			fmt.Printf("Port open: %d, Banner: %s\n", port, banner)
		}(port)
	}

	wg.Wait()
}
