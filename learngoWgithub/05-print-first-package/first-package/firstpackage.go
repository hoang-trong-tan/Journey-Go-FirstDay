package firstpackage

import "runtime"

func Version() string {
	return runtime.Version()
}
