package constant

import (
	"net/http"
)

// grpc status & msg
const SocketPath = "/var/lib/libvirt/ks.sock"

const (
	STATUS_OK         = http.StatusOK
	STATUS_BADREQUEST = http.StatusBadRequest
	STATUS_ERR        = http.StatusInternalServerError
)

const (
	MESSAGE_OK = "ok"
)

// log
const (
	DefualtLogPath = "/var/log/sdsctl.log"
)

// pool type
const (
	PoolCephType  = "ceph"
	PoolNFSType   = "nfs"
	PoolNetfsType = "netfs"
	PoolDirType   = "dir"
)

// rook
const (
	DefaultMdsNamespace = "myfs"
	DefaultName         = "admin"
)

// k8s GVK
const (
	// crd group & version
	DefaultGroup     = "doslab.io"
	DefaultVersion   = "v1"
	DefaultNamespace = "default"

	// rook related
	RookNamespace = "rook-ceph"

	// crd kind
	VMD_Kind     = "VirtualMachineDisk"
	VMDS_Kind    = "virtualmachinedisks"
	VMP_Kind     = "VirtualMachinePool"
	VMPS_Kind    = "virtualmachinepools"
	VMDSN_Kind   = "VirtualMachineDiskSnapshot"
	VMDSNS_Kinds = "virtualmachinedisksnapshots"
	VMDI_KIND    = "VirtualMachineDiskImage"
	VMDIS_KINDS  = "virtualmachinediskimages"

	// spec key
	CRD_Pool_Key     = "pool"
	CRD_Volume_Key   = "volume"
	CRD_NodeName_Keu = "nodeName"

	// vm pool status
	CRD_Pool_Active   = "active"
	CRD_Pool_Inactive = "inactive"

	CRD_Ready_Msg    = "The resource is ready."
	CRD_Ready_Reason = "Ready"
)
