package provision

import (
	"github.com/docker/machine/libmachine/drivers"
)

func init() {
	Register("Fedora", &RegisteredProvisioner{
		New: NewFedoraProvisioner,
	})
}

func NewFedoraProvisioner(d drivers.Driver) Provisioner {
	return &FedoraProvisioner{
		NewRedHatProvisioner("fedora", d),
	}
}

type FedoraProvisioner struct {
	*RedHatProvisioner
}

func (provisioner *FedoraProvisioner) String() string {
	return "fedora"
}

func (provisioner *FedoraProvisioner) CompatibleWithHost() bool {
	isFedora := provisioner.OsReleaseInfo.ID == "fedora"
	isCoreOS := provisioner.OsReleaseInfo.VariantID == "coreos"
	return isFedora && !isCoreOS
}
