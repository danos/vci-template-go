// SPDX-License-Identifier: MIT-0

package main

import (
	"log"

	vci_template_go "github.com/danos/vci-template-go"
	"github.com/danos/vci"
)

func main() {
	comp := vci.NewComponent("net.vyatta.eng.vci.template-go")
	comp.Model("net.vyatta.eng.vci.template-go.v1").
		Config(vci_template_go.NewConfig()).
		State(vci_template_go.NewState()).
		RPC("vci-template-go-v1", vci_template_go.NewRPC())
	err := comp.Run()
	if err != nil {
		log.Fatal(err)
	}
	comp.Wait()
}
