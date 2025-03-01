// Copyright 2019 Yunion
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package options

import (
	"os"

	common_options "yunion.io/x/onecloud/pkg/cloudcommon/options"
	"yunion.io/x/onecloud/pkg/util/fileutils2"
)

type SHostOptions struct {
	common_options.HostCommonOptions
	common_options.EtcdOptions

	HostType        string   `help:"Host server type, either hypervisor or kubelet" default:"hypervisor"`
	ListenInterface string   `help:"Master address of host server"`
	BridgeDriver    string   `help:"Bridge driver, bridge or openvswitch" default:"openvswitch"`
	Networks        []string `help:"Network interface information"`
	Rack            string   `help:"Rack of host (optional)"`
	Slots           string   `help:"Slots of host (optional)"`
	Hostname        string   `help:"Customized host name"`

	ServersPath         string `help:"Path for virtual server configuration files" default:"/opt/cloud/workspace/servers"`
	ImageCachePath      string `help:"Path for storing image caches" default:"/opt/cloud/workspace/disks/image_cache"`
	MemorySnapshotsPath string `help:"Path for memory snapshot stat files" default:"/opt/cloud/workspace/memory_snapshots"`
	// ImageCacheLimit int    `help:"Maximal storage space for image caching, in GB" default:"20"`
	AgentTempPath  string `help:"Path for ESXi agent"`
	AgentTempLimit int    `help:"Maximal storage space for ESXi agent, in GB" default:"10"`

	RecycleDiskfile         bool `help:"Recycle instead of remove deleted disk file" default:"true"`
	RecycleDiskfileKeepDays int  `help:"How long recycled files kept, default 28 days" default:"28"`

	ZeroCleanDiskData bool `help:"Clean disk data by writing zeros" default:"false"`

	EnableTemplateBacking    bool `help:"Use template as backing file"`
	AutoMergeBackingTemplate bool `help:"Automatically stream merging backing file"`
	AutoMergeDelaySeconds    int  `help:"Seconds to delay mergeing backing file after VM start, default 15 minutes" default:"900"`
	EnableFallocateDisk      bool `help:"Automatically allocate all spaces using fallocate"`

	EnableMonitor  bool `help:"Enable monitor"`
	ReportInterval int  `help:"Report interval in seconds" default:"60"`

	BwDownloadBandwidth int `help:"Default ingress bandwidth in mbit (0 disabled)" default:"10"`

	DnsServer       string `help:"Address of host DNS server"`
	DnsServerLegacy string `help:"Deprecated Address of host DNS server"`

	ChntpwPath           string `help:"path to chntpw tool" default:"/usr/local/bin/chntpw.static"`
	OvmfPath             string `help:"Path to OVMF.fd" default:"/opt/cloud/contrib/OVMF.fd"`
	LinuxDefaultRootUser bool   `help:"Default account for linux system is root"`

	BlockIoScheduler string `help:"Block IO scheduler, deadline or cfq" default:"deadline"`
	EnableKsm        bool   `help:"Enable Kernel Same Page Merging"`
	HugepagesOption  string `help:"Hugepages option: disable|native|transparent" default:"transparent"`
	EnableQmpMonitor bool   `help:"Enable qmp monitor" default:"true"`

	PrivatePrefixes []string `help:"IPv4 private prefixes"`
	LocalImagePath  []string `help:"Local image storage paths"`
	SharedStorages  []string `help:"Path of shared storages"`

	DefaultQemuVersion string `help:"Default qemu version" default:"4.2.0"`

	DhcpRelay       []string `help:"DHCP relay upstream"`
	DhcpLeaseTime   int      `default:"100663296" help:"DHCP lease time in seconds"`
	DhcpRenewalTime int      `default:"67108864" help:"DHCP renewal time in seconds"`

	TunnelPaddingBytes int64 `help:"Specify tunnel padding bytes" default:"0"`

	CheckSystemServices bool `help:"Check system services (ntpd, telegraf) on startup" default:"true"`

	DhcpServerPort     int    `help:"Host dhcp server bind port" default:"67"`
	DiskIsSsd          bool   `default:"false"`
	FetcherfsPath      string `default:"/opt/yunion/fetchclient/bin/fetcherfs" help:"Fuse fetcherfs path"`
	FetcherfsBlockSize int    `default:"16" help:"Fuse fetcherfs fetch chunk_size MB"`

	DefaultImageSaveFormat string `default:"qcow2" help:"Default image save format, default is qcow2, canbe vmdk"`

	DefaultReadBpsPerCpu   int  `default:"163840000" help:"Default read bps per cpu for hard IO limit"`
	DefaultReadIopsPerCpu  int  `default:"1250" help:"Default read iops per cpu for hard IO limit"`
	DefaultWriteBpsPerCpu  int  `default:"54525952" help:"Default write bps per cpu for hard IO limit"`
	DefaultWriteIopsPerCpu int  `default:"416" help:"Default write iops per cpu for hard IO limit"`
	SetVncPassword         bool `default:"true" help:"Auto set vnc password after monitor connected"`
	UseBootVga             bool `default:"false" help:"Use boot VGA GPU for guest"`

	EnableCpuBinding         bool `default:"true" help:"Enable cpu binding and rebalance"`
	EnableOpenflowController bool `default:"false"`

	PingRegionInterval     int      `default:"60" help:"interval to ping region, deefault is 1 minute"`
	ManageNtpConfiguration bool     `default:"true"`
	LogSystemdUnits        []string `help:"Systemd units log collected by fluent-bit"`
	// 更改默认带宽限速为400GBps, qiujian
	BandwidthLimit int `default:"400000" help:"Bandwidth upper bound when migrating disk image in MB/sec, default 400GBps"`
	// 热迁移带宽，预期不低于8MBps, 1G Memory takes 128 seconds
	MigrateExpectRate        int `default:"32" help:"Expected memory migration rate in MB/sec, default 32MBps"`
	MinMigrateTimeoutSeconds int `default:"30" help:"minimal timeout for a migration process, default 30 seconds"`

	SnapshotDirSuffix  string `help:"Snapshot dir name equal diskId concat snapshot dir suffix" default:"_snap"`
	SnapshotRecycleDay int    `default:"1" help:"Snapshot Recycle delete Duration day"`

	EnableTelegraf          bool `default:"true" help:"enable send monitoring data to telegraf"`
	WindowsDefaultAdminUser bool `default:"true" help:"Default account for Windows system is Administrator"`

	HostCpuPassthrough bool `default:"true" help:"if it is true, set qemu cpu type as -cpu host, otherwise, qemu64. default is true"`
	DisableSetCgroup   bool `default:"false" help:"disable cgroup for guests"`

	MaxReservedMemory int `default:"10240" help:"host reserved memory"`

	DefaultRequestWorkerCount int `default:"8" help:"default request worker count"`

	CommonConfigFile string `help:"common config file for container"`

	AllowSwitchVMs bool `help:"allow machines run as switch (spoof mac)" default:"true"`
	AllowRouterVMs bool `help:"allow machines run as router (spoof ip)" default:"true"`

	SdnPidFile        string `help:"pid file for sdnagent" default:"$SDN_PID_FILE|/var/run/onecloud/sdnagent.pid"`
	SdnSocketPath     string `help:"sdnagent listen socket path" default:"/var/run/onecloud/sdnagent.sock"`
	SdnEnableGuestMan bool   `help:"enable guest network manager in sdnagent" default:"$SDN_ENABLE_GUEST_MAN|true"`
	SdnEnableEipMan   bool   `help:"enable eip network manager in sdnagent" default:"$SDN_ENABLE_EIP_MAN|false"`
	SdnEnableTcMan    bool   `help:"enable TC manager in sdnagent" default:"$SDN_ENABLE_TC_MAN|true"`

	SdnAllowConntrackInvalid bool `help:"allow packets marked by conntrack as INVALID to pass" default:"$SDN_ALLOW_CONNTRACK_INVALID|false"`

	OvnSouthDatabase          string `help:"address for accessing ovn south database" default:"$HOST_OVN_SOUTH_DATABASE|unix:/var/run/openvswitch/ovnsb_db.sock"`
	OvnEncapIpDetectionMethod string `help:"detection method for ovn_encap_ip" default:"$HOST_OVN_ENCAP_IP_DETECTION_METHOD"`
	OvnEncapIp                string `help:"encap ip for ovn datapath.  Default to src address of default route" default:"$HOST_OVN_ENCAP_IP"`
	OvnIntegrationBridge      string `help:"name of integration bridge for logical ports" default:"$HOST_OVN_INTEGRATION_BRIDGE|brvpc"`
	OvnMappedBridge           string `help:"name of bridge for mapped traffic management" default:"$HOST_OVN_MAPPED_BRIDGE|brmapped"`
	OvnEipBridge              string `help:"name of bridge for eip traffic management" default:"$HOST_OVN_EIP_BRIDGE|breip"`
	OvnUnderlayMtu            int    `help:"mtu of ovn underlay network" default:"1500"`

	EnableRemoteExecutor bool   `help:"Enable remote executor" default:"false"`
	EnableHealthChecker  bool   `help:"enable host health checker" default:"false"`
	HealthDriver         string `help:"Component save host health state" default:"etcd"`
	HostHealthTimeout    int    `help:"host health timeout" default:"30"`
	HostLeaseTimeout     int    `help:"lease timeout" default:"10"`

	SyncStorageInfoDurationSecond int  `help:"sync storage size duration, unit is second" default:"60"`
	StartHostIgnoreSysError       bool `help:"start host agent ignore sys error" default:"false"`

	DisableProbeKubelet bool   `help:"Disable probe kubelet config" default:"false"`
	KubeletRunDirectory string `help:"Kubelet config file path" default:"/var/lib/kubelet"`

	DisableKVM bool `help:"force disable KVM" default:"false" json:"disable_kvm"`

	DisableGPU bool `help:"force disable GPU detect" default:"false" json:"disable_gpu"`
	DisableUSB bool `help:"force disable USB detect" default:"true" json:"disable_usb"`

	EthtoolEnableGso bool `help:"use ethtool to turn on or off GSO(generic segment offloading)" default:"false" json:"ethtool_enable_gso"`

	EnableVmUuid bool `help:"enable vm UUID" default:"true" json:"enable_vm_uuid"`

	EnableVirtioRngDevice bool `help:"enable qemu virtio-rng device" default:"true"`

	RestrictQemuImgConvertWorker bool `help:"restrict qemu-img convert worker" default:"false"`

	DefaultLiveMigrateDowntime float32 `help:"allow downtime in seconds for live migrate" default:"5.0"`
}

var (
	HostOptions SHostOptions
)

func Parse() (hostOpts SHostOptions) {
	common_options.ParseOptions(&hostOpts, os.Args, "host.conf", "host")
	if len(hostOpts.CommonConfigFile) > 0 && fileutils2.Exists(hostOpts.CommonConfigFile) {
		commonCfg := &common_options.HostCommonOptions{}
		commonCfg.Config = hostOpts.CommonConfigFile
		common_options.ParseOptions(commonCfg, []string{os.Args[0]}, "common.conf", "host")
		baseOpt := hostOpts.BaseOptions.BaseOptions
		hostOpts.HostCommonOptions = *commonCfg
		// keep base options
		hostOpts.BaseOptions.BaseOptions = baseOpt
	}
	return hostOpts
}

func Init() {
	HostOptions = Parse()
}
