package validator

import "k8s-management-go/app/models"

// check selected namespace against namespace list
func ValidateNamespace(namespaceToValidate string) bool {
	for _, ip := range models.GetIpConfiguration().Ips {
		if ip.Namespace == namespaceToValidate {
			return true
		}
	}
	return false
}
