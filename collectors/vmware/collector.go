package vmware

import (
	"context"
	"github.com/models"
	"github.com/models/vmware"
	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/session"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
	"net/url"
	"time"
)

const (
	// Connect retry delay
	RetryDelay = time.Second * 5
	// Maximum number of objects in each update
	MaxObjectUpdates = 10000
)

type VCenterCollector struct {
	vcenter *vmware.Vcenter
	client  *govmomi.Client
}

func New(vcenter vmware.VCenter) *VCenterCollector struct {
	return &vmware.VCenter{
		vcenter: vcenter
	}
}

func (vcc *VCenterCollector) connect(ctx context.Context) error {
	if vcc.client != nil {
		_ = vcc.client.Logout(ctx)
		vcc.client = nil
	}
	url.
}
