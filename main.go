package main

import (
	"fmt"
	"net"
	// "time"
	// uerr "github.com/ucloud/ucloud-sdk-go/ucloud/error"
)

func main() {
	config.getConf()
	initUcloud()
	// for {
	process()
	// 	time.Sleep(time.Hour * 24)
	// }

	// getProjectId()

}

func process() {
	fmt.Println("Start to process...")
	/*
		fmt.Println("Delete old GlobalSSH...")
		for _, id := range config.Instances {
			fmt.Println("delete old id:", id)
			deleteGlobalSSH(id)
		}

		fmt.Println("Checking Github ip...")
		// ips, _ := fetchIPs()
		ips := [...]string{
			"192.30.252.0/22",
			"185.199.108.0/22",
			"140.82.112.0/20",
			"13.114.40.48/32",
			"13.229.188.59/32",
			"13.234.176.102/32",
			// "13.234.210.38/32",
			// "13.236.229.21/32",
			// "13.237.44.5/32",
			// "13.250.177.223/32",
			// "15.164.81.167/32",
			// "18.194.104.89/32",
			// "18.195.85.27/32",
			// "18.228.52.138/32",
			// "18.228.67.229/32",
			// "18.231.5.6/32",
			// "35.159.8.160/32",
			// "52.192.72.89/32",
			// "52.64.108.95/32",
			// "52.69.186.44/32",
			// "52.74.223.119/32",
			// "52.78.231.108/32"
		}

		fmt.Println("Creating ucloud GlobalSSH...")
		DomainsNew := make([]string, 0)
		instances := make([]string, 0)
		for _, ip := range ips {
			ip = parseIP(ip)
			fmt.Println("ipip...", ip)
			if ip != "" {
				domain, instance, err := newGlobalSSH(ip, area(ip))
				if err != nil {
					e := err.(uerr.Error)
					if e.Code() == 33981 {
						fmt.Println("Ucloud ip", ip, "already exists")
						instances = append(instances, instance)
					} else {
						fmt.Println("Ucloud error:", err.Error())
					}
					continue
				} else {
					fmt.Println("Ucloud new GlobalSSH:", domain)
					DomainsNew = append(DomainsNew, domain)
					instances = append(instances, instance)
				}
			}
		}
		config.Instances = instances
		config.save()
	*/
	DomainsNew := [...]string{
		"13.114.40.48.ipssh.net",
		"13.229.188.59.ipssh.net",
		"13.234.176.102.ipssh.net",
		"13.234.210.38.ipssh.net",
		"13.236.229.21.ipssh.net",
		"13.237.44.5.ipssh.net",
	}
	fmt.Println("\nSleeping 2min for dns ...")
	// time.Sleep(time.Minute * 2)
	fmt.Println("\nNow Lookup Domain...")
	newDnsRecordIps := make([]string, 0)
	for _, domain := range DomainsNew {
		ips, err := net.LookupIP(domain)
		if err != nil {
			fmt.Println("DNS lookup error:", err.Error())
		} else {
			for _, newip := range ips {
				newDnsRecordIps = append(newDnsRecordIps, newip.String())
			}
		}
	}
	// setDnsRecords(newDnsRecordIps)
	fmt.Println("newDnsRecordIps...", newDnsRecordIps)
	fmt.Println("End...")
}
