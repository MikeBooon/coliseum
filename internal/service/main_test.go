package service_test

import (
	"testing"

	"github.com/MikeBooon/coliseum/test"
)

var testDeps *test.IntegrationTestDeps

func TestMain(m *testing.M) {
	test.RunTestWithIntegrationDeps(&testDeps, m)
}
