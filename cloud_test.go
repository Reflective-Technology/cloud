package cloud_test

import (
	"testing"

	"github.com/Reflective-Technology/cloud"
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

func TestAzure(t *testing.T) {
	cloud.SetMode(cloud.AZURE)
	if cloud.Mode() != cloud.AZURE {
		t.Errorf(errorResult, cloud.AZURE, cloud.Mode())
	}
}

func TestOnPremise(t *testing.T) {
	cloud.SetMode(cloud.ON_PREMISE)
	if cloud.Mode() != cloud.ON_PREMISE {
		t.Errorf(errorResult, cloud.ON_PREMISE, cloud.Mode())
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
