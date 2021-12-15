package utils_test

import (
	"bytes"
	"encoding/json"
	"github.com/aasumitro/gowa/internal/domain/contracts"
	"github.com/aasumitro/gowa/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type testForm struct {
	Email    string `json:"email" binding:"required,email" msg:"error_invalid_email"`
	Username string `json:"username" binding:"required,alphanum,gte=6,lte=32" msg:"error_invalid_username"`
	Password string `json:"password" binding:"required,gte=6,lte=32" msg:"error_invalid_password"`
}

var requestErrorMessage = map[string]string{
	"error_invalid_email":    "please enter a valid email address",
	"error_invalid_username": "username must be alphanumeric and between 6 and 32 characters",
	"error_invalid_password": "password must be between 6 and 32 characters",
}

type validationErrorTestSuite struct {
	suite.Suite
	context         *gin.Context
	validationError contracts.ValidationErrors
}

// Mock the response writer
func MockJsonPost(context *gin.Context, content interface{}) {
	context.Request.Method = "POST"
	context.Request.Header.Set("Content-Type", "application/json")

	jsonBytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}

	context.Request.Body = io.NopCloser(bytes.NewBuffer(jsonBytes))
}

// before each test
func (suite *validationErrorTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)

	suite.validationError = utils.NewValidationErrors(requestErrorMessage)
	suite.context, _ = gin.CreateTestContext(httptest.NewRecorder())
	suite.context.Request = &http.Request{Header: make(http.Header)}
}

// TestValidationErrorTestSuite test suite for invalid username
func (suite *validationErrorTestSuite) TestRequestInvalidUsername() {
	MockJsonPost(suite.context, map[string]interface{}{
		"username": "test123!#",
		"email":    "test@email.com",
		"password": "123456",
	})

	req := testForm{}
	if err := suite.context.BindJSON(&req); err != nil {
		errors := suite.validationError.All(req, err)
		assert.True(suite.T(), len(errors) == 1)

		value, ok := errors["username"]
		assert.True(suite.T(), ok)
		assert.True(suite.T(), value == requestErrorMessage["error_invalid_username"])
	} else {
		suite.T().Errorf("expected error but got nil")
	}
}

// TestValidationErrorTestSuite test suite for invalid email
func (suite *validationErrorTestSuite) TestRequestInvalidEmail() {
	MockJsonPost(suite.context, map[string]interface{}{
		"username": "test123",
		"email":    "testemail.com",
		"password": "123456",
	})

	req := testForm{}
	if err := suite.context.BindJSON(&req); err != nil {
		errors := suite.validationError.All(req, err)
		assert.True(suite.T(), len(errors) == 1)

		value, ok := errors["email"]
		assert.True(suite.T(), ok)
		assert.True(suite.T(), value == requestErrorMessage["error_invalid_email"])
	} else {
		suite.T().Errorf("expected error but got nil")
	}
}

// TestValidationErrorTestSuite test suite for invalid password
func (suite *validationErrorTestSuite) TestRequestInvalidPassword() {
	MockJsonPost(suite.context, map[string]interface{}{
		"username": "test123",
		"email":    "test@email.com",
		"password": "12345",
	})

	req := testForm{}
	if err := suite.context.BindJSON(&req); err != nil {
		errors := suite.validationError.All(req, err)
		assert.True(suite.T(), len(errors) == 1)

		value, ok := errors["password"]
		assert.True(suite.T(), ok)
		assert.True(suite.T(), value == requestErrorMessage["error_invalid_password"])
	} else {
		suite.T().Errorf("expected error but got nil")
	}
}

// TestValidationErrorTestSuite test suite for invalid all fields
func (suite *validationErrorTestSuite) TestRequestInvalidAllForm() {
	MockJsonPost(suite.context, map[string]interface{}{
		"username": "test123!@#%",
		"email":    "testemail.com",
		"password": "12345",
	})

	req := testForm{}
	if err := suite.context.BindJSON(&req); err != nil {
		errors := suite.validationError.All(req, err)
		assert.True(suite.T(), len(errors) == 3)

		valueUsername, ok := errors["username"]
		assert.True(suite.T(), ok)
		assert.True(suite.T(), valueUsername == requestErrorMessage["error_invalid_username"])

		valueEmail, ok := errors["email"]
		assert.True(suite.T(), ok)
		assert.True(suite.T(), valueEmail == requestErrorMessage["error_invalid_email"])

		valuePassword, ok := errors["password"]
		assert.True(suite.T(), ok)
		assert.True(suite.T(), valuePassword == requestErrorMessage["error_invalid_password"])
	} else {
		suite.T().Errorf("expected error but got nil")
	}
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestValidationError(t *testing.T) {
	suite.Run(t, new(validationErrorTestSuite))
}
