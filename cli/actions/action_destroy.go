package actions

import (
	"cli/tools/terraform"
	"cli/ui"
	"fmt"
	"os"
)

func (c *Cluster) Destroy() error {
	if !c.ContainsTfStateConfig() {
		return fmt.Errorf("cluster '%s' is already destroyed (or not yet initialized).", c.Name)
	}

	fmt.Printf("Cluster '%s' will be destroyed.\n", c.Name)

	if err := ui.Ask(); err != nil {
		return err
	}

	fmt.Printf("Destroying cluster '%s'...\n", c.Name)

	t, err := terraform.NewTerraform(c.Ctx, c.Path)

	if err != nil {
		return err
	}

	if err := t.Destroy(); err != nil {
		return err
	}

	if err := os.Remove(c.TfStatePath()); err != nil {
		return err
	}

	return os.Remove(c.KubeconfigPath())
}
