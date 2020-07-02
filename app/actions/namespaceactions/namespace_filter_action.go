package namespaceactions

import (
	"k8s-management-go/app/models"
	"sort"
	"strings"
)

// ActionReadNamespaceWithFilter is a namespace loader and filter
func ActionReadNamespaceWithFilter(filter *string) (namespaces []string) {
	ipList := models.GetIPConfiguration().Ips
	for _, ip := range ipList {
		if filter != nil && *filter != "" {
			if strings.Contains(ip.Namespace, *filter) {
				namespaces = append(namespaces, ip.Namespace)
			}
		} else {
			namespaces = append(namespaces, ip.Namespace)
		}
	}
	sort.Strings(namespaces)
	return namespaces
}
