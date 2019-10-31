package src

import (
	"flag"
	"fmt"
	"iptools"
)

func main() {

	ipType := flag.Int("v",4,"4:ipv4 6:ipv6")
	interfaceName := flag.String("i","wlan0","interface name. eth0,en0,wlan0")
	flag.Parse()

	ips := iptools.GetInterFaceIps(*interfaceName,*ipType)
	ipString := iptools.ParseIpsToString(ips)
	fmt.Println(ipString)

}




