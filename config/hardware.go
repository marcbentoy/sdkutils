package sdkconfig

type HardwareConfig struct {
	Arch     string   `json:"arch"`
	Packages []string `json:"packages"`
}

var (
	Hardware_X86_64 = HardwareConfig{
		Arch:     "x86_64",
		Packages: []string{},
	}
)
