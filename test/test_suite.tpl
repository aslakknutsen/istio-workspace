package {{.Package}}

import (
	"go.uber.org/goleak"
	"testing"

	. "github.com/maistra/istio-workspace/test"

	{{.GinkgoImport}}
	{{.GomegaImport}}
)

func Test{{.FormattedName}}(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecWithJUnitReporter(t, "{{.FormattedName}} Suite")
}

var _ = SynchronizedAfterSuite(func() {}, func() {
	goleak.VerifyNone(GinkgoT(),
		goleak.IgnoreTopFunction("k8s.io/klog/v2.(*loggingT).flushDaemon"),
		goleak.IgnoreTopFunction("github.com/onsi/ginkgo/internal/specrunner.(*SpecRunner).registerForInterrupts"),
	)
})
