//go:build systemdefaultregistry

package airgap

import (
	"fmt"

	"github.com/rancher/distros-test-framework/pkg/assert"
	"github.com/rancher/distros-test-framework/pkg/testcase"
	"github.com/rancher/distros-test-framework/pkg/testcase/support"

	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("Test Airgap Cluster with System Default Registry:", Ordered, func() {
	BeforeAll(func() {
		support.BuildAirgapCluster(cluster)
	})

	It("Installs and validates product on private nodes:", func() {
		testcase.TestSystemDefaultRegistry(cluster, flags)
	})

	It("Validates Nodes", func() {
		testcase.TestNodeStatusUsingBastion(
			cluster,
			assert.NodeAssertReadyStatus(),
			nil,
		)
	})

	It("Validates Pods", func() {
		testcase.TestPodStatusUsingBastion(
			cluster,
			assert.PodAssertRestart(),
			assert.PodAssertReady())
	})

	AfterAll(func() {
		support.LogClusterInfoUsingBastion(cluster)
	})

	// TODO: Validate deployment, eg: cluster-ip
})

var _ = AfterEach(func() {
	if CurrentSpecReport().Failed() {
		fmt.Printf("\nFAILED! %s\n", CurrentSpecReport().FullText())
	} else {
		fmt.Printf("\nPASSED! %s\n", CurrentSpecReport().FullText())
	}
})
