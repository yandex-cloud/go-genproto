// Code generated by protoc-gen-goext. DO NOT EDIT.

package baremetal

func (m *StoragePartition) SetType(v StoragePartitionType) {
	m.Type = v
}

func (m *StoragePartition) SetSizeGib(v int64) {
	m.SizeGib = v
}

func (m *StoragePartition) SetMountPoint(v string) {
	m.MountPoint = v
}

type Storage_StorageType = isStorage_StorageType

func (m *Storage) SetStorageType(v Storage_StorageType) {
	m.StorageType = v
}

func (m *Storage) SetPartitions(v []*StoragePartition) {
	m.Partitions = v
}

func (m *Storage) SetDisk(v *Disk) {
	m.StorageType = &Storage_Disk{
		Disk: v,
	}
}

func (m *Storage) SetRaid(v *Raid) {
	m.StorageType = &Storage_Raid{
		Raid: v,
	}
}

func (m *Raid) SetType(v RaidType) {
	m.Type = v
}

func (m *Raid) SetDisks(v []*Disk) {
	m.Disks = v
}
