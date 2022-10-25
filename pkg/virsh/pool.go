package virsh

import (
	"encoding/xml"
	"github.com/kube-stack/sdsctl/pkg/utils"
	libvirtxml "github.com/libvirt/libvirt-go-xml"
	"libvirt.org/go/libvirt"
)

func GetPoolInfo(name string) (*libvirt.StoragePool, error) {
	conn, err := GetConn()
	defer conn.Close()
	if err != nil {
		return nil, err
	}
	pool, err := conn.LookupStoragePoolByName(name)
	if err != nil {
		return nil, err
	}
	return pool, nil
}

func GetPoolState(state libvirt.StoragePoolState) string {
	if state == libvirt.STORAGE_POOL_RUNNING {
		return "active"
	}
	return "inactive"
}

func CreatePool(name, ptype, target string) (*libvirt.StoragePool, error) {
	pool, err := DefinePool(name, ptype, target)
	if err != nil {
		return nil, err
	}
	// start pool
	err = StartPool(name)
	if err != nil {
		return nil, err
	}
	return pool, nil
}

func DefinePool(name, ptype, target string) (*libvirt.StoragePool, error) {
	conn, err := GetConn()
	defer conn.Close()
	if err != nil {
		return nil, err
	}

	if !utils.Exists(target) {
		utils.CreateDir(target)
	}

	poolXML := libvirtxml.StoragePool{
		Type: ptype,
		Name: name,
		Target: &libvirtxml.StoragePoolTarget{
			Path: target,
		},
	}
	poolDoc, err := poolXML.Marshal()
	if err != nil {
		return nil, err
	}
	// use define instead of create in order to create xml in /etc/libvirt/storage
	return conn.StoragePoolDefineXML(poolDoc, 0)
}

func AutoStartPool(name string, autoStart bool) error {
	conn, err := GetConn()
	defer conn.Close()
	if err != nil {
		return err
	}
	pool, err := conn.LookupStoragePoolByName(name)
	return pool.SetAutostart(autoStart)
}

func DeletePool(name string) error {
	conn, err := GetConn()
	defer conn.Close()
	if err != nil {
		return err
	}
	pool, err := conn.LookupStoragePoolByName(name)
	if err != nil {
		return err
	}
	pool.Destroy()
	pool.Undefine()
	return nil
}

func StartPool(name string) error {
	conn, err := GetConn()
	defer conn.Close()
	if err != nil {
		return err
	}
	pool, err := conn.LookupStoragePoolByName(name)
	if err != nil {
		return err
	}
	return pool.Create(0)
}

func StopPool(name string) error {
	conn, err := GetConn()
	defer conn.Close()
	if err != nil {
		return err
	}
	pool, err := conn.LookupStoragePoolByName(name)
	if err != nil {
		return err
	}
	return pool.Destroy()
}

func IsPoolActive(name string) (bool, error) {
	conn, err := GetConn()
	defer conn.Close()
	if err != nil {
		return false, err
	}
	pool, err := conn.LookupStoragePoolByName(name)
	if err != nil {
		return false, err
	}
	return pool.IsActive()
}

func GetPoolTargetPath(name string) (string, error) {
	conn, err := GetConn()
	defer conn.Close()
	if err != nil {
		return "", err
	}
	pool, err := GetPoolInfo("pooltest2")
	if err != nil {
		return "", err
	}
	pxml, err := pool.GetXMLDesc(0)
	if err != nil {
		return "", err
	}
	poolObj := &libvirtxml.StoragePool{}
	err = xml.Unmarshal([]byte(pxml), poolObj)
	if err != nil {
		return "", err
	}
	return poolObj.Target.Path, nil
}
