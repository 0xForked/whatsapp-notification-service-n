package whatsapp_test

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/aasumitro/gowa/configs"
	"github.com/aasumitro/gowa/pkg/whatsapp"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"regexp"
	"testing"
)

const getAllDevicesQuery = `
SELECT jid, registration_id, noise_key, identity_key,
       signed_pre_key, signed_pre_key_id, signed_pre_key_sig,
       adv_key, adv_details, adv_account_sig, adv_account_sig_key, adv_device_sig,
       platform, business_name, push_name
FROM whatsmeow_device
`

type whatsappClientTestSuite struct {
	suite.Suite
	mock sqlmock.Sqlmock
}

func (suite *whatsappClientTestSuite) SetupSuite() {
	var err error
	viper.Reset()
	viper.SetConfigFile("../../.example.env")
	viper.SetConfigType("dotenv")
	configs.LoadEnv()
	configs.Instance.DBDsnURL = "./db/local-data.db"
	configs.Instance.DBDriver = "sqlite3"
	configs.DbPool, suite.mock, err = sqlmock.New(
		sqlmock.QueryMatcherOption(
			sqlmock.QueryMatcherRegexp))
	require.NoError(suite.T(), err)
}

func (suite *whatsappClientTestSuite) TestClient_MakeConnection_ShouldErrorContainer() {
	query := regexp.QuoteMeta(getAllDevicesQuery)
	suite.mock.ExpectQuery(query).WillReturnError(errors.New(""))
	whatsappInstance := whatsapp.Client{}
	assert.Panics(suite.T(), func() { whatsappInstance.MakeConnection() })
}

func (suite *whatsappClientTestSuite) TestClient_MakeConnection_ShouldErrorGetDevice() {
	data := suite.mock.
		NewRows([]string{"jid", "registration_id", "noise_key", "identity_key", "signed_pre_key", "signed_pre_key_id", "signed_pre_key_sig", "adv_key", "adv_details", "adv_account_sig", "adv_device_sig", "platform", "business_name", "push_name", "adv_account_sig_key"}).
		AddRow("123", 123, "0x123", "02xds", "02xds", 123, "0x123", "0x123", "0x123", "0x123", "0x123", "asd", "qwe", "lorem", "0x123")
	query := regexp.QuoteMeta(getAllDevicesQuery)
	suite.mock.ExpectQuery(query).WillReturnRows(data).WillReturnError(nil)
	whatsappInstance := whatsapp.Client{}
	assert.Panics(suite.T(), func() { whatsappInstance.MakeConnection() })
}

func (suite *whatsappClientTestSuite) TestClient_MakeConnection_ShouldSuccessGetDevice() {
	// SKIPP
}

func TestWhatsappClient(t *testing.T) {
	suite.Run(t, new(whatsappClientTestSuite))
}
