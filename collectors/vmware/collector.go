package vmware

import (
	"context"
	"net/url"
	"time"

	modelsVmware "github.com/konveyor/forklift-inventory/models/vmware"
	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/session"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/soap"
)

const (
	// RetryDelay delay between connection retries
	RetryDelay = time.Second * 5
	// MaxObjectUpdates maximum number of objects per update
	MaxObjectUpdates = 10000
)

// VCenterCollector contains a reference to the VCenter definition and to the govmomi client
type VCenterCollector struct {
	vcenter *modelsVmware.VCenter
	client  *govmomi.Client
}

// New returns a VCenterCollect struct
func New(vc modelsVmware.VCenter) *VCenterCollector {
	return &VCenterCollector{
		vcenter: &vc,
	}
}

func (vcc *VCenterCollector) connect(ctx context.Context) error {
	if vcc.client != nil {
		_ = vcc.client.Logout(ctx)
		vcc.client = nil
	}
	url := url.URL{
		Scheme: "https",
		Host:   vcc.vcenter.HostnameOrIp,
		Path:   "/sdk",
		User: url.UserPassword(
			vcc.vcenter.Username,
			vcc.vcenter.Password,
		),
	}
	soapClient := soap.NewClient(&url, false)
	soapClient.SetThumbprint(url.Host, vcc.vcenter.SslThumbprint)
	vimClient, err := vim25.NewClient(ctx, soapClient)
	if err != nil {
		return err
	}
	vcc.client = &govmomi.Client{
		SessionManager: session.NewManager(vimClient),
		Client:         vimClient,
	}
	err = vcc.client.Login(ctx, url.User)
	if err != nil {
		return err
	}

	return nil
}
