package config

// basic IP struct
type Ip struct {
	Namespace string
	Ip        string
}

// list of IPs
type IpConfiguration struct {
	Ips []Ip
}

var ipConfig IpConfiguration

func GetIpConfiguration() IpConfiguration {
	return ipConfig
}

func AddIpAndNamespaceToConfiguration(namespace string, ip string) {
	ipConfig.Ips = append(ipConfig.Ips, Ip{namespace, ip})
}
