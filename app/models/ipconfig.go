package models

// IP is the basic IP structure
type IP struct {
	Namespace string
	IP        string
}

// IPConfiguration contains a list of IPs
type IPConfiguration struct {
	Ips []IP
}

var ipConfig IPConfiguration

// GetIPConfiguration returns the current IP configuration
func GetIPConfiguration() IPConfiguration {
	return ipConfig
}

// AddIPAndNamespaceToConfiguration adds IP and namespace to the IP configuration
func AddIPAndNamespaceToConfiguration(namespace string, ip string) {
	ipConfig.Ips = append(ipConfig.Ips, IP{namespace, ip})
}
