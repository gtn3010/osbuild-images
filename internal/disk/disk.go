// Disk package contains abstract data-types to define disk-related entities.
//
// PartitionTable, Partition and Filesystem types are currently defined.
// All of them can be 1:1 converted to osbuild.QEMUAssemblerOptions.
package disk

import "github.com/osbuild/osbuild-composer/internal/osbuild"

type PartitionTable struct {
	// Size of the disk.
	Size uint64
	UUID string
	// Partition table type, e.g. dos, gpt.
	Type       string
	Partitions []Partition
}

type Partition struct {
	Start    uint64
	Size     uint64
	Type     string
	Bootable bool
	// ID of the partition, dos doesn't use traditional UUIDs, therefore this
	// is just a string.
	UUID string
	// If nil, the partition is raw; It doesn't contain a filesystem.
	Filesystem *Filesystem
}

type Filesystem struct {
	Type string
	// ID of the filesystem, vfat doesn't use traditional UUIDs, therefore this
	// is just a string.
	UUID       string
	Label      string
	Mountpoint string
}

// Converts PartitionTable to osbuild.QEMUAssemblerOptions that encode
// the same partition table.
func (pt PartitionTable) QEMUAssemblerOptions() osbuild.QEMUAssemblerOptions {
	var partitions []osbuild.QEMUPartition
	for _, p := range pt.Partitions {
		partitions = append(partitions, p.QEMUPartition())
	}

	return osbuild.QEMUAssemblerOptions{
		Size:       pt.Size,
		PTUUID:     pt.UUID,
		PTType:     pt.Type,
		Partitions: partitions,
	}
}

// Converts Partition to osbuild.QEMUPartition that encodes the same partition.
func (p Partition) QEMUPartition() osbuild.QEMUPartition {
	var fs *osbuild.QEMUFilesystem
	if p.Filesystem != nil {
		f := p.Filesystem.QEMUFilesystem()
		fs = &f
	}
	return osbuild.QEMUPartition{
		Start:      p.Start,
		Size:       p.Size,
		Type:       p.Type,
		Bootable:   p.Bootable,
		UUID:       p.UUID,
		Filesystem: fs,
	}
}

// Converts Filesystem to osbuild.QEMUFilesystem that encodes the same fs.
func (fs Filesystem) QEMUFilesystem() osbuild.QEMUFilesystem {
	return osbuild.QEMUFilesystem{
		Type:       fs.Type,
		UUID:       fs.UUID,
		Label:      fs.Label,
		Mountpoint: fs.Mountpoint,
	}
}
