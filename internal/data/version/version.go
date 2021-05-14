package version

import (
	"fmt"
	"github.com/cstria0106/tamagotchi/internal/util"
	"gopkg.in/yaml.v2"
)

type Version struct {
	Major        uint32
	Minor        uint32
	Distribution string
}

func (v *Version) String() string {
	return fmt.Sprintf("%d.%d", v.Major, v.Minor)
}

func (v *Version) Buffer() ([]byte, uint32) {
	distributionLength := len(v.Distribution)
	size := 8 + distributionLength

	buffer := make([]byte, size)
	copy(buffer[0:4], util.EncodeU32(v.Major))
	copy(buffer[4:8], util.EncodeU32(v.Minor))
	copy(buffer[8:], v.Distribution)

	return buffer, uint32(size)
}

func FromYML(buffer []byte) (*Version, error) {
	v := &Version{}
	err := yaml.Unmarshal(buffer, v)

	if err != nil {
		return nil, err
	}

	return v, nil
}

func FromBuffer(buffer []byte) *Version {
	major := util.DecodeU32(buffer[0:4])
	minor := util.DecodeU32(buffer[4:8])
	distribution := string(buffer[8:])

	return &Version{
		Major:        major,
		Minor:        minor,
		Distribution: distribution,
	}
}

func (v *Version) Equals(other *Version) bool {
	return v.Major == other.Major && v.Minor == other.Minor
}
