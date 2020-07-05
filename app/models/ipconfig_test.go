package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddIPAndNamespaceToConfiguration(t *testing.T) {
	var ipOne = "1.2.3.4"
	var nsOne = "projectA"
	var ipTwo = "1.2.3.5"
	var nsTwo = "projectB"
	AddIPAndNamespaceToConfiguration(nsOne, ipOne)
	AddIPAndNamespaceToConfiguration(nsTwo, ipTwo)

	assert.Len(t, GetIPConfiguration().IPs, 2)
	assert.Equal(t, GetIPConfiguration().IPs[0].IP, ipOne)
	assert.Equal(t, GetIPConfiguration().IPs[0].Namespace, nsOne)
	assert.Equal(t, GetIPConfiguration().IPs[1].IP, ipTwo)
	assert.Equal(t, GetIPConfiguration().IPs[1].Namespace, nsTwo)
	ResetIPAndNamespaces()
}

func TestResetIPAndNamespaces(t *testing.T) {
	AddIPAndNamespaceToConfiguration("prjA", "1.2.3.4")
	AddIPAndNamespaceToConfiguration("prjB", "1.2.3.5")

	assert.Len(t, GetIPConfiguration().IPs, 2)
	ResetIPAndNamespaces()
	assert.Nil(t, GetIPConfiguration().IPs)
}
