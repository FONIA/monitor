package tool

import (
	"fmt"
	"os/exec"
	"runtime"
	"time"
	"strings"
	"net"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/disk"
)


//根据进程名判断进程是否运行
func CheckProRunning(serverName string) (bool, error) {
	a := `ps ux | awk '/` + serverName + `/ && !/awk/ {print $2}'`
	pid, err := RunCommand(a)
	if err != nil {
		return false, err
	}
	return pid != "", nil
}
//根据进程名称获取进程ID
func GetPid(serverName string) (string, error) {
	a := `ps aux | awk '/` + serverName + `/ && !/awk/ {print $2}'`
	pid, err := RunCommand(a)
	return pid , err
}

//杀掉进程
func  Kill(serverName string)(string,error)  {
	a:="pkill  "+serverName
	data,err:=RunCommand(a)
	return data,err
}




func RunCommand(cmd string) (string, error) {
	if runtime.GOOS == "windows" {
		return runInWindows(cmd)
	} else {
		return runInLinux(cmd)
	}
}
func runInWindows(cmd string) (string, error) {
	result, err := exec.Command("cmd", "/c", cmd).Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(result)), err
}

func runInLinux(cmd string) (string, error) {
	fmt.Println("Running Linux cmd:" + cmd)
	result, err := exec.Command("/bin/sh", "-c", cmd).Output()
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	return strings.TrimSpace(string(result)), err
}

func getLocalIPv4Address() (ipv4Address string, err error){
	//获取所有网卡
	addrs, err := net.InterfaceAddrs()

	//遍历
	for _, addr := range addrs {
		//取网络地址的网卡的信息
		ipNet, isIpNet := addr.(*net.IPNet)
		//是网卡并且不是本地环回网卡
		if isIpNet && !ipNet.IP.IsLoopback() {
			ipv4 := ipNet.IP.To4()
			//能正常转成ipv4
			if ipv4 != nil {
				return ipv4.String(), nil
			}
		}
	}

	return
}

func GetCpuPercent() float64 {
	percent, _:= cpu.Percent(time.Second, false)
	return percent[0]
}

func GetMemPercent()float64 {
	memInfo, _ := mem.VirtualMemory()
	return memInfo.UsedPercent
}

func GetDiskPercent() float64 {
	parts, _ := disk.Partitions(true)
	diskInfo, _ := disk.Usage(parts[0].Mountpoint)
	return diskInfo.UsedPercent
}