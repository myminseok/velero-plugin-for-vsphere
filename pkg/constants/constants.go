/*
Copyright 2020 the Velero contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package constants

import "time"

const (
	// supported volume type in plugin
	CnsBlockVolumeType = "ivd"
)

const (
	// Duration at which lease expires on CRs.
	LeaseDuration = 60 * time.Second

	// Duration after which leader renews its lease.
	RenewDeadline = 15 * time.Second

	// Duration after which non-leader retries to acquire lease.
	RetryPeriod = 5 * time.Second
)

const (
	// Duration after which Reflector resyncs CRs and calls UpdateFunc on each of the existing CRs.
	ResyncPeriod = 10 * time.Minute

	DefaultSecretResyncPeriod = 5 * time.Minute
)

// configuration constants for the volume snapshot plugin
const (
	// The key of SnapshotManager mode for data movement. Specifically, boolean string values are expected.
	// By default, it is "false". No data movement from local to remote storage if "true" is set.
	VolumeSnapshotterLocalMode = "LocalMode"
	// The key of SnapshotManager location
	VolumeSnapshotterManagerLocation = "SnapshotManagerLocation"
	// Valid values for the config with the VolumeSnapshotterManagerLocation key
	VolumeSnapshotterPlugin     = "Plugin"
	VolumeSnapshotterDataServer = "DataServer"
)

const (
	// Max retry limit for downloads.
	DOWNLOAD_MAX_RETRY = 5

	// Initial retry for both uploads and downloads.
	MIN_RETRY = 0

	// BACKOFF for downloads.
	DOWNLOAD_BACKOFF = 5

	// Max backoff limit for uploads.
	UPLOAD_MAX_BACKOFF = 60

	// Exceeds this number of retry, will give a warning message to ask user to fix network issue in cluster.
	RETRY_WARNING_COUNT = 8
)

// configuration constants for the S3 repository
const (
	DefaultS3RepoPrefix     = "plugins/vsphere-astrolabe-repo"
	DefaultS3BackupLocation = "default"
	AWS_ACCESS_KEY_ID       = "aws_access_key_id"
	AWS_SECRET_ACCESS_KEY   = "aws_secret_access_key"
)

const (
	// Minimum velero version number to meet velero plugin requirement
	VeleroMinVersion = "v1.5.1"

	// Minimum csi driver version number to meet velero plugin requirement
	CsiMinVersion = "v1.0.2"
)

const (
	// DefaultNamespace is the Kubernetes namespace that is used by default for
	// the Velero server and API objects.
	DefaultNamespace          = "velero"
	CloudCredentialSecretName = "cloud-credentials"
)

const (
	// Default port used to access vCenter.
	DefaultVCenterPort string = "443"
)

const (
	DataManagerForPlugin   string = "data-manager-for-plugin"
	BackupDriverForPlugin  string = "backup-driver"
	BackupDriverNamespace  string = "velero-vsphere-plugin-backupdriver"
	VeleroPluginForVsphere string = "velero-plugin-for-vsphere"
	VeleroDeployment       string = "velero"
	VSphereCSIController          = "vsphere-csi-controller"
	KubeSystemNamespace           = "kube-system"
)

const (
	S3RepositoryDriver string = "s3repository.astrolabe.vmware-tanzu.com"
)

const (
	VCSecretNs           = "kube-system"
	VCSecretNsSupervisor = "vmware-system-csi"
	VCSecret             = "vsphere-config-secret"
	VCSecretTKG          = "csi-vsphere-config"
)

const (
	TkgSupervisorService = "supervisor"
)

// Indicates the type of cluster where Plugin is installed
type ClusterFlavor string

const (
	Unknown    ClusterFlavor = "Unknown"
	Supervisor               = "Supervisor Cluster"
	TkgGuest                 = "TKG Guest Cluster"
	VSphere                  = "vSphere Kubernetes Cluster"
)

// feature flog constants
const (
	VSphereLocalModeFlag    = "local-mode"
	VSphereLocalModeFeature = "EnableLocalMode"
)

const (
	VSpherePluginFeatureStates = "velero-vsphere-plugin-feature-states"
)

// Para Virtual Cluster access for Guest Cluster
const (
	PvApiEndpoint = "supervisor.default.svc" // TODO: get it from "kubectl get cm -n vmware-system-csi pvcsi-config"
	PvPort        = "6443"
	PvSecretName  = "pvbackupdriver-provider-creds"
)

const (
	WithoutBackupRepository = "without-backup-repository"
	WithoutDeleteSnapshot   = "without-delete-snapshot"
)

const (
	ItemSnapshotLabel   = "velero-plugin-for-vsphere/item-snapshot-blob"
	PluginVersionLabel  = "velero-plugin-for-vsphere/plugin-version"
	SnapshotBackupLabel = "velero.io/backup-name"
)

const (
	RetryInterval = 5
	RetryMaximum  = 5
)

// Keys for supervisor cluster parameters
const (
	VCuuidKey                 = "vCenterUUID"
	SupervisorClusterIdKey    = "SupervisorClusterId"
	SupervisorResourcePoolKey = "SupervisorResourcePool"
)

// We are currently translating Group + Kind -> the names below.  This involves
// a singular to plural conversion.  When adding new resources, ensure that they
// work with the singular to plural rule of words ending in "y" the "y" becomes "ies" and other
// words get an "s" attached.
var ResourcesToBlock = map[string]bool{
	// Kubernetes with vSphere Supervisor Cluster resources
	"agentinstalls.installers.tmc.cloud.vmware.com":           true,
	"aviloadbalancerconfigs.netoperator.vmware.com":           true,
	"blockaffinities.crd.projectcalico.org":                   true,
	"certificaterequests.cert-manager.io":                     true,
	"certificates.cert-manager.io":                            true,
	"challenges.acme.cert-manager.io":                         true,
	"clusterissuers.cert-manager.io":                          true,
	"clusterresourcesetbindings.addons.cluster.x-k8s.io":      true,
	"clusterresourcesets.addons.cluster.x-k8s.io":             true,
	"clusters.cluster.x-k8s.io":                               true,
	"cnsnodevmattachments.cns.vmware.com":                     true,
	"cnsregistervolumes.cns.vmware.com":                       true,
	"cnsvolumemetadatas.cns.vmware.com":                       true,
	"compatibilities.run.tanzu.vmware.com":                    true,
	"contentlibraryproviders.vmoperator.vmware.com":           true,
	"contentsourcebindings.vmoperator.vmware.com":             true,
	"contentsources.vmoperator.vmware.com":                    true,
	"gatewayclasses.networking.x-k8s.io":                      true,
	"gateways.networking.x-k8s.io":                            true,
	"haproxyloadbalancerconfigs.netoperator.vmware.com":       true,
	"httproutes.networking.x-k8s.io":                          true,
	"imagedisks.imagecontroller.vmware.com":                   true,
	//"images.imagecontroller.vmware.com":                     true, // DO NOT ADD IT BACK
	"installoptions.appplatform.wcp.vmware.com":               true,
	"installrequirements.appplatform.wcp.vmware.com":          true,
	"ipamblocks.crd.projectcalico.org":                        true,
	"ipamconfigs.crd.projectcalico.org":                       true,
	"ipamhandles.crd.projectcalico.org":                       true,
	"ippools.crd.projectcalico.org":                           true,
	"ippools.netoperator.vmware.com":                          true,
	"issuers.cert-manager.io":                                 true,
	"kubeadmconfigs.bootstrap.cluster.x-k8s.io":               true,
	"kubeadmconfigtemplates.bootstrap.cluster.x-k8s.io":       true,
	"kubeadmcontrolplanes.controlplane.cluster.x-k8s.io":      true,
	"kuberneteslicenses.licenseoperator.vmware.com":           true,
	"loadbalancerconfigs.netoperator.vmware.com":              true,
	"loadbalancers.vmware.com":                                true,
	"machinedeployments.cluster.x-k8s.io":                     true,
	"machinehealthchecks.cluster.x-k8s.io":                    true,
	"machinepools.exp.cluster.x-k8s.io":                       true,
	"machines.cluster.x-k8s.io":                               true,
	"machinesets.cluster.x-k8s.io":                            true,
	"members.registryagent.vmware.com":                        true,
	"ncpconfigs.nsx.vmware.com":                               true,
	"network-attachment-definitions.k8s.cni.cncf.io":          true,
	"networkinterfaces.netoperator.vmware.com":                true,
	"networks.netoperator.vmware.com":                         true,
	"nsxerrors.nsx.vmware.com":                                true,
	//"nsxlbmonitors.vmware.com":                              true, // DO NOT ADD IT BACK
	//"nsxloadbalancermonitors.vmware.com":                    true, // DO NOT ADD IT BACK
	"nsxlocks.nsx.vmware.com":                                 true,
	"nsxnetworkinterfaces.nsx.vmware.com":                     true,
	"orders.acme.cert-manager.io":                             true,
	"persistenceinstanceinfoes.psp.wcp.vmware.com":            true,
	"projects.registryagent.vmware.com":                       true,
	"providerserviceaccounts.run.tanzu.vmware.com":            true,
	"registries.registryagent.vmware.com":                     true,
	"resourcecheckreports.psp.wcp.vmware.com":                 true,
	"resourcechecks.psp.wcp.vmware.com":                       true,
	"storagepolicies.appplatform.wcp.vmware.com":              true,
	"storagepolicies.psp.wcp.vmware.com":                      true,
	"storagepools.cns.vmware.com":                             true,
	"supervisorservices.appplatform.wcp.vmware.com":           true,
	"tanzukubernetesaddons.run.tanzu.vmware.com":              true,
	"tanzukubernetesclusters.run.tanzu.vmware.com":            true,
	"tanzukubernetesreleases.run.tanzu.vmware.com":            true,
	"tcproutes.networking.x-k8s.io":                           true,
	"tkgserviceconfigurations.run.tanzu.vmware.com":           true,
	"vcuiplugins.appplatform.wcp.vmware.com":                  true,
	"veleroservices.veleroappoperator.vmware.com":             true,
	"virtualmachineclassbindings.vmoperator.vmware.com":       true,
	"virtualmachineclasses.vmoperator.vmware.com":             true,
	"virtualmachineimages.vmoperator.vmware.com":              true,
	"virtualmachineservices.vmoperator.vmware.com":            true,
	"virtualmachinesetresourcepolicies.vmoperator.vmware.com": true,
	"virtualmachines.vmoperator.vmware.com":                   true,
	"virtualnetworkinterfaces.vmware.com":                     true,
	"virtualnetworks.vmware.com":                              true,
	"vmxnet3networkinterfaces.netoperator.vmware.com":         true,
	"vspheredistributednetworks.netoperator.vmware.com":       true,
	"wcpclusters.infrastructure.cluster.vmware.com":           true,
	"wcpnamespaces.appplatform.wcp.vmware.com":                true,
	"wcpmachines.infrastructure.cluster.vmware.com":           true,
	"wcpmachinetemplates.infrastructure.cluster.vmware.com":   true,

	// plugin resources
	"backuprepositories.backupdriver.cnsdp.vmware.com":     true,
	"backuprepositoryclaims.backupdriver.cnsdp.vmware.com": true,
	"clonefromsnapshots.backupdriver.cnsdp.vmware.com":     true,
	"deletesnapshots.backupdriver.cnsdp.vmware.com":        true,
	"downloads.datamover.cnsdp.vmware.com":                 true,
	"snapshots.backupdriver.cnsdp.vmware.com":              true,
	"uploads.datamover.cnsdp.vmware.com":                   true,
}

var ResourcesToBlockOnRestore = map[string]bool{
	// Kubernetes with vSphere Supervisor Cluster resources

	// We need to remove some metadata from the Pod resource on
	// Supervisor Cluster, i.e., annotation "vmware-system-vm-uuid"
	// before the restore as the existing VM UUID is associated with
	// the old VM that does not exist any more
	"pods": true,

	// The following resources are backed up everytime when a container
	// is backed up on Supervisor Cluster.
	// We should skip it at restore time.
	"images.imagecontroller.vmware.com": true,

	// "nsxlbmonitors.vmware.com" is the real name for this resource,
	// however, our existing name parsing mechanism for resource matches
	// with the parsed name. Adding both of them to the list.
	// The real name will be used to make sure the resource is
	// picked up in the AppliesTo func of item action plugin, while
	// the parsed name will be used to skip restoring the resource
	// in the Execute func of item action plugin.
	"nsxlbmonitors.vmware.com":           true, // real name
	"nsxloadbalancermonitors.vmware.com": true, // parsed name
}

var ResourcesToHandle = map[string]bool{
	"persistentvolumeclaims": true,
}

var PodAnnotationsToSkip = map[string]bool{
	// Kubernetes with vSphere Supervisor Cluster Pod annotations
	"kubernetes.io/psp":                 true,
	"mac":                               true,
	"vlan":                              true,
	"vmware-system-ephemeral-disk-uuid": true,
	"vmware-system-image-references":    true,
	"vmware-system-vm-moid":             true,
	"vmware-system-vm-uuid":             true,
}

// UUID of the VM on Supervisor Cluster
const VMwareSystemVMUUID = "vmware-system-vm-uuid"

// Label to set for Velero to ignore resources
const VeleroExcludeLabel = "velero.io/exclude-from-backup"

// Default time window to clean up CR which is in terminal state
const DefaultCRCleanUpWindow = 24

const (
	SnapshotParamBackupName       = "BackupName"
	SnapshotParamSvcSnapshotName  = "SvcSnapshotName"
	SnapshotParamBackupRepository = "BackupRepository"
)

// These label keys are used to identify configMap used for storage class mapping, format:
// velero.io/plugin-config: ""
// velero.io/change-storage-class: RestoreItemAction
const (
	// Plugin kind name
	PluginKindRestoreItemAction = "RestoreItemAction"
	// This label key is used to identify the name and kind of plugin that configMap is for
	ChangeStorageClassLabelKey = "velero.io/change-storage-class"
	// This label key is used to identify the ConfigMap as config for a plugin.
	PluginConfigLabelKey = "velero.io/plugin-config"
	// This is the reserved name used to map from a non storage class to a new storage class
	// at restore time. Storage class is required for restore. If no storage class is specified
	// in the PVC during backup, user can specify "com.vmware.cnsdp.emptystorageclass" as the
	// old storage class name to map to a new existing storage class name at restore time.
	EmptyStorageClass = "com.vmware.cnsdp.emptystorageclass"
)

const (
	DefaultRetryIntervalStart = time.Second
	DefaultRetryIntervalMax   = 5 * time.Minute
)

const (
	ImageRepositoryComponent = "Repository"
	ImageContainerComponent  = "Container"
	ImageVersionComponent    = "Version"
)

const VsphereVolumeSnapshotLocationProvider = "velero.io/vsphere"
