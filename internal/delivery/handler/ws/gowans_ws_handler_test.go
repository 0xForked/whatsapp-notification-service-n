package ws

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type gowansWSHandlerTestSuite struct {
	suite.Suite
}

func (suite *gowansWSHandlerTestSuite) SetupSuite() {
	// TODO ...
}

func (suite *gowansWSHandlerTestSuite) TestHandler__() {
	// TODO ...
}

func TestGowansWSHandler(t *testing.T) {
	suite.Run(t, new(gowansWSHandlerTestSuite))
}
