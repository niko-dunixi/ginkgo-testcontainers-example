package nginx

import (
	"net/http"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func WithMyTestContainer(text string, specs ...func(nginxGetter Supplier[*nginxContainer])) {
	GinkgoHelper()
	Context(text, func() {

		var nginxContainerInstance *nginxContainer

		BeforeEach(func(ctx SpecContext) {
			container, err := startContainer(ctx)
			Expect(err).ShouldNot(HaveOccurred(), "could not setup nginx testcontainer")
			Expect(container).ToNot(BeNil(), "nginx container should not be nil if there is no err")
			nginxContainerInstance = container
			DeferCleanup(func(ctx SpecContext) {
				err := container.Terminate(ctx)
				Expect(err).ShouldNot(HaveOccurred(), "could not cleanup nginx testcontainer")
			}, NodeTimeout(time.Second*30))
		}, NodeTimeout(time.Minute*5))

		for _, spec := range specs {
			spec(func() *nginxContainer {
				// This must be a higher order function, otherwise we cannot pass the initialized
				// value from the `BeforeEach` to the `It` with during the Test Execution phase
				return nginxContainerInstance
			})
		}
	})
}

var _ = Describe("some scenario where I do things", func() {
	WithMyTestContainer("and the other service is available", func(nginxGetter Supplier[*nginxContainer]) {
		// Using a BeforeEach for variable assignment.
		// This helps tests conform to other ginkgo examples
		// in the wild and help with debugability in a given test.
		var nginx *nginxContainer
		BeforeEach(func() {
			nginx = nginxGetter()
		})

		It("should succeed", func(ctx SpecContext) {
			response, err := http.Get(nginx.URI)
			Expect(err).ShouldNot(HaveOccurred(), "nginx isn't working")
			Expect(response.StatusCode).To(Equal(http.StatusOK), "nginx isn 't working.")
		})
	})
})
