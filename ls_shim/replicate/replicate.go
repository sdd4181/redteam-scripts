package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os/exec"
	"strings"
)

var systemOS string = getOS()

func main() {
	if systemOS == "" {

	}

	fmt.Println(systemOS)
	if strings.Index(systemOS, "debian") != -1 || strings.Index(systemOS, "ubuntu") != -1 {
		runCommand("apt-get", "install sshpass -y")
	}
}

func getOS(isFail ...bool) string {
	var ret_os string
	os_str := readFile("/etc/os-release")
	os_split := strings.Split(os_str, "\n")
	for i := 0; i < len(os_split); i++ {
		matchString := "ID_LIKE="
		if isFail[0] {
			matchString = "ID="
		}
		if strings.Index(os_split[i], matchString) != -1 {
			ret_os = strings.Replace(os_split[i], matchString, "", 1)
			ret_os = strings.Replace(ret_os, `"`, "", -1)
			break
		}
	}
	if ret_os == "" && !isFail[0] {
		return getOS(true)
	}
	return ret_os
}

func runCommand(binary, args string) {
	cmd := exec.Command(binary, args)
	cmd.Run()
}

func readFile(path string) string {
	dat, _ := ioutil.ReadFile(path)
	str := string(dat)
	return str
}

func GetOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	ip := localAddr.IP
	ipstr := ip.String()
	ipstr = strings.Replace(ipstr, ".", "", -1)
	return ipstr
}
