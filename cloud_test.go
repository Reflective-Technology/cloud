package cloud_test

import (
	"testing"

	"gitlab.tp.zuso.arpa/zuso-rd-team/go-pkg/cloud.git"
)

const (
	envCloudMode = "CLOUD_MODE"
	AWS          = "aws"
	GCP          = "gcp"
	errorResult  = "Expected Mode() to return %s, but got %s"
)

func TestAWS(t *testing.T) {
	cloud.SetMode(cloud.AWS)
	if cloud.Mode() != cloud.AWS {
		t.Errorf(errorResult, AWS, cloud.Mode())
	}
}

func TestGCP(t *testing.T) {
	cloud.SetMode(cloud.GCP)
	if cloud.Mode() != cloud.GCP {
		t.Errorf(errorResult, GCP, cloud.Mode())
	}
}

func TestEmpty(t *testing.T) {
	cloud.SetMode("")
	if cloud.Mode() != cloud.UNSPECIFIED {
		t.Errorf(errorResult, cloud.UNSPECIFIED, cloud.Mode())
	}
}

func TestUnknown(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	cloud.SetMode("unknown")
}
