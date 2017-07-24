package VPN

import (
	"github.com/xiaokangwang/AndroidLibV2ray/CoreI"
	"github.com/xiaokangwang/AndroidLibV2ray/Process/Escort"
	"github.com/xiaokangwang/AndroidLibV2ray/configure"

	"golang.org/x/sys/unix"

	"v2ray.com/core/transport/internet"
)

/*VpnSupportReady VpnSupportReady*/
func (v *VPNSupport) VpnSupportReady() {
	if !v.status.VpnSupportnodup {
		/*
			v.VpnSupportnodup = true
			//Surpress Network Interruption Notifiction
			go func() {
				time.Sleep(5 * time.Second)
				v.VpnSupportnodup = false
			}()*/
		v.VpnSupportSet.Setup(v.Conf.Service.VPNSetupArg)
		v.setV2RayDialer()
		v.startVPNRequire()
	}
}
func (v *VPNSupport) startVPNRequire() {
	e := Escort.NewEscort()
	go e.EscortRun(v.Conf.Service.Target, v.Conf.Service.Args, false, v.VpnSupportSet.GetVPNFd())
}

func (v *VPNSupport) askSupportSetInit() {
	v.VpnSupportSet.Prepare()
}

func (v *VPNSupport) vpnSetup() {
	if v.Conf.Service.VPNSetupArg != "" {
		v.prepareDomainName()

		v.askSupportSetInit()
	}
}
func (v *VPNSupport) vpnShutdown() {

	if v.Conf.Service.VPNSetupArg != "" {
		/*
			BUG DISCOVERED!

			v.VpnSupportnodup can have unexpected value cause VPN failed to revoke.
			more testing needed.

		*/

		//if v.VpnSupportnodup {
		err := unix.Close(v.VpnSupportSet.GetVPNFd())
		println(err)
		//}
		v.VpnSupportSet.Shutdown()
	}
	v.status.VpnSupportnodup = false
}

func (v *VPNSupport) setV2RayDialer() {
	protectedDialer := &vpnProtectedDialer{vp: v}
	internet.UseAlternativeSystemDialer(internet.WithAdapter(protectedDialer))
}

type VPNSupport struct {
	prepareddomain preparedDomain
	VpnSupportSet  V2RayVPNServiceSupportsSet
	status         *CoreI.Status
	Conf           configure.VPNConfig
}

type V2RayVPNServiceSupportsSet interface {
	GetVPNFd() int
	Setup(Conf string) int
	Prepare() int
	Shutdown() int
	Protect(int) int
}
