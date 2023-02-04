package cluster

import (
	"cli/cluster/executors"
	"cli/cluster/provisioner"
	"cli/cluster/provisioner/terraform"
	"cli/utils/file"
	"path/filepath"
)

const (
	DefaultConfigDir    = "config"
	DefaultShareDir     = "share"
	DefaultTerraformDir = DefaultConfigDir + "/terraform"

	DefaultNewConfigFilename     = "kubitect.yaml"
	DefaultAppliedConfigFilename = "kubitect-applied.yaml"
	DefaultInfraConfigFilename   = "infrastructure.yaml"

	DefaultTerraformStateFilename = "terraform.tfstate"
	DefaultKubeconfigFilename     = "admin.conf"
)

type ClusterContext interface {
	WorkingDir() string
	HomeDir() string
	ShareDir() string
	ClustersDir() string
	LocalClustersDir() string

	Local() bool
	ShowTerraformPlan() bool
}

type ClusterMeta struct {
	ClusterContext

	Name  string
	Path  string
	Local bool

	exec executors.Executor
	prov provisioner.Provisioner
}

// func NewClusterMeta(ctx ClusterContext, clusterPath string) (*ClusterMeta, error) {
// 	cpStat, err := os.Stat(clusterPath)
// 	if err != nil {
// 		return nil, fmt.Errorf("cluster meta: %v", err)
// 	}

// 	meta := &ClusterMeta{
// 		ClusterContext: ctx,
// 		Name:           cpStat.Name(),
// 		Path:           filepath.Join(clusterPath, cpStat.Name()),
// 		Local:          ctx.Local(),
// 	}

// 	// if !cpStat.IsDir() || !meta.ContainsArchiveFile() {
// 	// 	return nil, fmt.Errorf("cluster meta: %s is not a cluster directory", err)
// 	// }

// 	return meta, nil
// }

func (c ClusterMeta) AppliedConfigPath() string {
	return filepath.Join(c.Path, DefaultConfigDir, DefaultAppliedConfigFilename)
}

func (c ClusterMeta) InfrastructureConfigPath() string {
	return filepath.Join(c.Path, DefaultConfigDir, DefaultInfraConfigFilename)
}

func (c ClusterMeta) TfStatePath() string {
	return filepath.Join(c.Path, DefaultTerraformDir, DefaultTerraformStateFilename)
}

func (c ClusterMeta) KubeconfigPath() string {
	return filepath.Join(c.Path, DefaultConfigDir, DefaultKubeconfigFilename)
}

func (c ClusterMeta) PrivateSshKeyPath() string {
	return filepath.Join(c.Path, DefaultConfigDir, ".ssh", "id_rsa")
}

func (c ClusterMeta) ContainsAppliedConfig() bool {
	return file.Exists(c.AppliedConfigPath())
}

func (c ClusterMeta) ContainsTfStateConfig() bool {
	return file.Exists(c.TfStatePath())
}

func (c ClusterMeta) ContainsKubeconfig() bool {
	return file.Exists(c.KubeconfigPath())
}

func (c *ClusterMeta) Provisioner() provisioner.Provisioner {
	if c.prov != nil {
		return c.prov
	}

	c.prov = terraform.NewTerraformProvisioner(
		c.Path,
		c.ShareDir(),
		c.ShowTerraformPlan(),
		nil,
	)

	return c.prov
}
