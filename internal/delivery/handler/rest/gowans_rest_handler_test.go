package rest

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type gowansRESTHandlerTestSuite struct {
	suite.Suite
}

func (suite *gowansRESTHandlerTestSuite) SetupSuite() {
	// TODO ...
}

func (suite *gowansRESTHandlerTestSuite) TestHandler__() {
	// TODO ...
}

func TestGowansRESTHandler(t *testing.T) {
	suite.Run(t, new(gowansRESTHandlerTestSuite))
}
