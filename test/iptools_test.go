package test

import (
	"iptools"
	"testing"
)

func TestGetInterFaceIps(t *testing.T) {

	ipsString := iptools.GetInterFaceIps("en0", 4)

	if ipsString[0] == "10.237.222.74" {
		t.Log(ipsString)
	} else {
		t.Fatal(ipsString)
	}
}

func TestParseIpsToString(t *testing.T) {

	ip := []string{"1.1.1.1"}
	ips := []string{"1.1.1.1", "2.2.2.2"}

	if iptools.ParseIpsToString(ip) == "1.1.1.1" {
		t.Log(ip)
	} else {
		t.Fatal(ip)
	}

	if iptools.ParseIpsToString(ips) == "1.1.1.1 2.2.2.2" {
		t.Log(ips)
	} else {
		t.Fatal(ips)
	}

}
