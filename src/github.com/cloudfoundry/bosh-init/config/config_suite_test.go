package config_test

import (
	. "github.com/cloudfoundry/bosh-init/internal/github.com/onsi/ginkgo"
	. "github.com/cloudfoundry/bosh-init/internal/github.com/onsi/gomega"
	"testing"
)

func TestCmd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Config Suite")
}