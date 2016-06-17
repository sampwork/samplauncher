package main

import (
	"bufio"
	"flag"
	"os"
	"regexp"
	"strings"
)

const configFile string = "server.cfg"

func ChangeServerPort(oldPort, newPort string) {
	regex, err := regexp.Compile(oldPort)
	if err != nil {
		return
	}

	fh, err := os.OpenFile(configFile, 0, 0666)
	f := bufio.NewReader(fh)

	if err != nil {
		return
	}
	defer fh.Close()

	buf := make([]byte, 1024)

	var cfgstr string

	for {
		buf, _, err = f.ReadLine()

		if err != nil {
			break
		}

		s := string(buf)

		result := regex.ReplaceAllString(s, newPort)

		cfgstr += result + "\n"
	}

	fo, _ := os.Create(configFile)

	writer := bufio.NewWriter(fo)

	defer fo.Close()
	defer writer.Flush()

	writer.WriteString(cfgstr)
}

func GetServerPort(filename string) string {
	fh, err := os.Open(filename)
	f := bufio.NewReader(fh)

	if err != nil {
		return "Error"
	}
	defer fh.Close()

	buf := make([]byte, 1024)
	for {
		buf, _, err = f.ReadLine()
		if err != nil {
			return "ReadLine error"
		}

		s := string(buf)

		if strings.Contains(s, "port") {
			oldPort := regexp.MustCompile("\\d+")
			return oldPort.FindString(s)
		}
	}
}

func main() {
	flag.Parse()
	if flag.NArg() != 0 {
		for x := 0; flag.NArg() != x; x++ {
			if x == 0 {
				ChangeServerPort(GetServerPort(configFile), flag.Arg(x))
			} else {
				ChangeServerPort(flag.Arg(x-1), flag.Arg(x))
			}

			LaunchServer()
		}
	} else {
		PrintUsageMessage()
	}
}
