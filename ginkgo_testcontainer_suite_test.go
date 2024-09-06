package nginx_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGinkgoTestcontainer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GinkgoTestcontainerExample Suite")
}
