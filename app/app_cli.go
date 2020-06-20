// +build !darwin

package app

import "k8s-management-go/app/cli"

// start App with CLI
func StartApp(info string) {
	cli.Workflow(info, nil)
}
