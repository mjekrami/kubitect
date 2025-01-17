// Package env provides constants for all other packages to consume,
// without creating import cycles.
//
// This package should not import any other packages.
package env

// Project related constants
const (
	ConstProjectUrl        = "https://github.com/MusicDin/kubitect"
	ConstProjectVersion    = "v3.3.1"
	ConstKubesprayUrl      = "https://github.com/kubernetes-sigs/kubespray"
	ConstKubesprayVersion  = "v2.24.1"
	ConstKubernetesVersion = "v1.28.6"
	ConstTerraformVersion  = "1.5.2"
)

// ProjectRequiredApps define applications that Kubitect depends on.
var ProjectRequiredApps = []string{
	"virtualenv",
	"python3",
	"git",
}

// ProjectRequiredFiles define required files/directories that are copied
// from embedded resources, when cluster is created.
var ProjectRequiredFiles = []string{
	"ansible/",
	"terraform/",
}

// ProjectApplyActions define options for "apply --action" command.
var ProjectApplyActions = [...]string{
	"create",
	"upgrade",
	"scale",
}

// ProjectK8sVersions define supported Kubernetes versions.
var ProjectK8sVersions = []string{
	"v1.28.0 - v1.28.6",
	"v1.27.0 - v1.27.10",
	"v1.26.0 - v1.26.13",
}

// ProjectOsPresets is a list of available OS distros.
var ProjectOsPresets = map[string]struct {
	Source           string
	NetworkInterface string
}{
	"ubuntu20": {
		Source:           "https://cloud-images.ubuntu.com/releases/focal/release/ubuntu-20.04-server-cloudimg-amd64.img",
		NetworkInterface: "ens3",
	},
	"ubuntu22": {
		Source:           "https://cloud-images.ubuntu.com/releases/jammy/release/ubuntu-22.04-server-cloudimg-amd64.img",
		NetworkInterface: "ens3",
	},
	"debian11": {
		Source:           "https://cloud.debian.org/images/cloud/bullseye/latest/debian-11-genericcloud-amd64.qcow2",
		NetworkInterface: "ens3",
	},
	"debian12": {
		Source:           "https://cloud.debian.org/images/cloud/bookworm/latest/debian-12-genericcloud-amd64.qcow2",
		NetworkInterface: "ens3",
	},
	"centos9": {
		Source:           "https://cloud.centos.org/centos/9-stream/x86_64/images/CentOS-Stream-GenericCloud-9-latest.x86_64.qcow2",
		NetworkInterface: "eth0",
	},
	"rocky9": {
		Source:           "https://dl.rockylinux.org/pub/rocky/9/images/x86_64/Rocky-9-GenericCloud-Base.latest.x86_64.qcow2",
		NetworkInterface: "eth0",
	},
}
