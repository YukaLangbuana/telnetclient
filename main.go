package main

import (
	"net"
	"os"
	"strconv"
	"time"
)

type Options struct {
	Url      string
	Port     string
	Duration time.Duration
}

func (opt *Options) parseOptions(args []string) {
	arguments := map[string]bool{
		"-p":         true,
		"--port":     true,
		"-d":         true,
		"--duration": true,
	}

	for i, v := range os.Args {
		if arguments[v] && (v == "-p" || v == "--port") {
			opt.Port = os.Args[i+1]
		} else if arguments[v] && (v == "-d" || v == "--duration") {
			dur, _ := strconv.Atoi(os.Args[i+1])
			opt.Duration = time.Duration(dur)
		}
	}
}

func NewOptions() *Options {

	opt := Options{
		Url:      os.Args[1],
		Port:     "23",
		Duration: time.Duration(30 * time.Second),
	}

	return &opt

}

func main() {
	if len(os.Args) < 2 {
		intro()
		return
	}

	opt := NewOptions()
	opt.parseOptions(os.Args)

	println("Dialing: ", opt.Url+" at port "+opt.Port)
	conn, err := net.Dial("tcp", opt.Url+":"+opt.Port)
	if err != nil {
		println("Error!", err.Error())
		return
	}
	println("Connected!")
	println("Sending packet for", opt.Duration.Seconds(), "s")
	start := time.Now()
	if time.Since(start) < opt.Duration {
		conn.Write([]byte{65})
		println("Heartbeat sent. Time since start", time.Since(start).Seconds())
		time.Sleep(1 * time.Second)
	}
}

func intro() {
	println("usage:")
	println("	telnetclient <URL>")
	println("")
	println("Flags:")
	println("	-p, --port	int		destination port (default: 23)")
	println("	-d, --duration	int		test duration in second(s) (default: 30s)")
}
