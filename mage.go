//+build mage

package main

import (
	// mage:import
	"github.com/LUSHDigital/core-mage/targets"
)

func init() {
	targets.ProjectName = "timestables"
	targets.ProjectType = "service"
	targets.DockerComposeDevDependencies = []string{"redis"}
	targets.DockerComposeTestDependencies = []string{"redis"}
	targets.ProtoServices = []string{"timestables"}
}