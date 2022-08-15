package collections_test

import (
	"testing"

	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
)

func TestUnitBundler(t *testing.T) {
	suite := spec.New("collections", spec.Report(report.Terminal{}))
	suite("Collections", testCollections)
	suite("Set", testSet)
	suite.Run(t)
}
