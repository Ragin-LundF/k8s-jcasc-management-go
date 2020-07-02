// +build ignore

package app

import "k8s-management-go/app/cli"

// StartApp will start App with CLI
func StartApp(info string) {
	cli.Workflow(info, nil)
}

// StartCli will start App with CLI
func StartCli(info string) {
	cli.Workflow(info, nil)
}
