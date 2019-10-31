package iptools
import (
	"net"
	"strings"
)
//get interface ipaddress
func GetInterFaceIps(interfaceName string,ipType int) []string{
	var ipv4Ips []string
	var ipv6Ips []string

	interfaceInfo, _:= net.InterfaceByName(interfaceName)

	interfaceIps,_ := interfaceInfo.Addrs()

	for _ , ip := range interfaceIps{

		address,_,_ := net.ParseCIDR(ip.String())

		if address.To4() != nil {
			ipv4Ips = append(ipv4Ips, address.String())
		}else{
			ipv6Ips = append(ipv6Ips, address.String())
		}
	}

	switch ipType {
	case 4:
		return ipv4Ips
	case 6:
		return ipv6Ips
	}
	return ipv4Ips
}

//put []ipaddress to ipaddress string
func ParseIpsToString(ips []string) string{
	if len(ips) == 1 {
		return ips[0]
	}else{
		ip := strings.Join(ips," ")
		return ip
	}
}
