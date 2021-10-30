module github.com/silicomdk/sts-operator

go 1.16

replace github.com/silicomdk/sts-grpc => /home/rmr/src/sts/operator/pkg/grpc/

require (
	github.com/go-logr/logr v0.4.0
	github.com/onsi/ginkgo v1.16.4
	github.com/onsi/gomega v1.15.0
	github.com/silicomdk/sts-grpc v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.38.0
	k8s.io/api v0.22.2
	k8s.io/apimachinery v0.22.2
	k8s.io/client-go v0.22.2
	sigs.k8s.io/controller-runtime v0.10.2
)
