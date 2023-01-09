package service_test

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type gowansServiceTestSuite struct {
	suite.Suite
}

func TestGowansService(t *testing.T) {
	suite.Run(t, new(gowansServiceTestSuite))
}

func (suite *gowansServiceTestSuite) SetupSuite() {}

func (suite *gowansServiceTestSuite) TestService__() {
	// TODO: TEST CASE n
}
