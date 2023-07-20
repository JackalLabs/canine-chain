package gmp_testing

import ibctesting "github.com/cosmos/ibc-go/v4/testing"

type GMPTestSuite struct {
	GMPTestHelper

	coordinator *ibctesting.Coordinator
}
