// +build ignore

package app

import "k8s-management-go/app/cli"

// start App with CLI
func StartApp(info string) {
	cli.Workflow(info, nil)
}

// start App with CLI
func StartCli(info string) {
	cli.Workflow(info, nil)
}
