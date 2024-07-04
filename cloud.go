package cloud

import (
	"os"
	"sync/atomic"
)

const envCloudMode = "CLOUD_MODE"

const (
	AWS         = "aws"
	GCP         = "gcp"
	UNSPECIFIED = "unspecified"
)

const (
	awsCode = iota
	gcpCode
	unspecifiedCode
)

var (
	appMode  int32 = unspecifiedCode
	modeName atomic.Value
)

func init() {
	mode := os.Getenv(envCloudMode)
	SetMode(mode)
}

// SetMode sets mode according to input string.
func SetMode(value string) {
	if value == "" {
		value = UNSPECIFIED
	}

	switch value {
	case AWS:
		atomic.StoreInt32(&appMode, awsCode)
	case GCP:
		atomic.StoreInt32(&appMode, gcpCode)
	case UNSPECIFIED:
		atomic.StoreInt32(&appMode, unspecifiedCode)
	default:
		panic("mode unknown: " + value)
	}

	modeName.Store(value)
}

// Mode returns current gin mode.
func Mode() string {
	return modeName.Load().(string)
}
