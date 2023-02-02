package plugin

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
)

func TestAWS(t *testing.T) {
	// Note: this test is simple, but serves as a smoke test.
	// The AWS() call below also catches duplicate columns and other issues
	// that may have been missed if mock tests are incomplete.
	p := AWS(aws.Config{})
	name := p.Name()
	if name != "aws" {
		t.Errorf("Name() = %q, want %q", name, "aws")
	}
}
