package ofuscatestruct_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestOfuscateStruct(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "OfuscateStruct Suite")
}
