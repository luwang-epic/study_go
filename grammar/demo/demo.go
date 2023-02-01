package main

import (
	"fmt"
	"syscall"
)

func main() {

	fs := syscall.Statfs_t{}
	syscall.Statfs("/", &fs)

	spaceTotal := uint64(fs.Bsize) * (fs.Blocks - fs.Bfree + fs.Bavail)

	fmt.Println(spaceTotal)


	spaceTags := make(map[string]string, 4)
	spaceTags["is_machine"] = "false"
	spaceTags["fs_name"] = "ext4"
	spaceTags["mount_dir"] = "Root"
	spaceTags["fs_type"] = "linux"
	dimensionValueMap := make(map[string][]*DimensionObject, 100)

	AddToMutiDimensionMap(0, dimensionValueMap, "spaceUsedPercent", 0.19, "float", spaceTags)


	for _, d := range dimensionValueMap["spaceUsedPercent"] {
		fmt.Println(d)
	}


}

type DimensionObject struct {
	name string
	value interface{}
	VType string
	Tags map[string]string
}

func AddToMutiDimensionMap(index int64, dimensionValueMap map[string][]*DimensionObject, key string, v interface{}, t string, tags map[string]string) {
	tmp := NewDimensionObject(key, v, t, tags)
	if index == 0 {
		dimensionValueMap[key] = []*DimensionObject{tmp}
		//glog.Infof("length %v", len(d.dimensionValueMap))
	} else {
		dimensionValueMap[key] = append(dimensionValueMap[key], tmp)
		//glog.Infof("length %v", len(d.dimensionValueMap))
	}
}

func NewDimensionObject(n string, v interface{}, t string, tags map[string]string) *DimensionObject {
	return &DimensionObject {
		name: 	n,
		value:	v,
		VType:	t,
		Tags:	tags,
	}
}