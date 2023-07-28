package main

import (
	"os"

	control_pkg "github.com/kubesimplify/ksctl/api/controllers"
	"github.com/kubesimplify/ksctl/api/resources"
	"github.com/kubesimplify/ksctl/api/resources/controllers"
)

func NewCli(cmd *resources.CobraCmd) {
	cmd.Version = os.Getenv("KSCTL_VERSION")

	if len(cmd.Version) == 0 {
		cmd.Version = "dummy v11001.2"
	}
}

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	cmd := &resources.CobraCmd{}
	NewCli(cmd)

	cmd.Client.Metadata.Provider = "civo"
	cmd.Client.Metadata.K8sDistro = "kubeadm"
	cmd.Client.Metadata.StateLocation = "local"
	cmd.Client.Metadata.ClusterName = "dummy-name"
	cmd.Client.Metadata.Region = "southindia"

	var controller controllers.Controller = control_pkg.GenKsctlController()
	controller.CreateHACluster(&cmd.Client)
}
