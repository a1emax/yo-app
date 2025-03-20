package account

const lastVersion = 1

type Data struct {
	Version int64 `yaml:"version"`
	Debug   bool  `yaml:"debug"`
}

func (d *Data) Check() {
	switch d.Version {
	case 0:
		d.Debug = true

		fallthrough
	default:
	}

	d.Version = lastVersion
}
