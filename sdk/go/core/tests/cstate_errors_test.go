package test

import (
	"fmt"
	oc_bgp "github.com/CiscoDevNet/ydk-go/ydk/models/ydktest/openconfig_bgp"
	"github.com/CiscoDevNet/ydk-go/ydk"
	"github.com/CiscoDevNet/ydk-go/ydk/providers"
	"github.com/CiscoDevNet/ydk-go/ydk/services"
	"github.com/CiscoDevNet/ydk-go/ydk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

const (
	invalidXMLBgpPayload1 = `<bgp xmlns="http://openconfig.net/yang/bgp">
  <global>
    <config>
      <as>65172</as>
      <router-id>1.2.3.4</router-id>
    </afi-safis>
  </global>
</bgp>
`
	invalidXMLBgpPayload2 = `<bgp xmlns="http://openconfig.net/yang/bgp">
  <global>
    <config>
      <as>65172</as>
      <router-id>wrong router id</router-id>
    </config>
  </global>
</bgp>`
)

type CStateErrorsTestSuite struct {
	suite.Suite
}

func (suite *CStateErrorsTestSuite) SetupSuite() {
}

func (suite *CStateErrorsTestSuite) BeforeTest(suiteName, testName string) {
	fmt.Printf("%v: %v ...\n", suiteName, testName)
}

func (suite *CStateErrorsTestSuite) TearDownSuite() {
}

func (suite *CStateErrorsTestSuite) TestRepositoryWrongPath() {
	repo := types.Repository{}
	repo.Path = "/wrong/path"
	provider := providers.NetconfServiceProvider{
		Repo:     repo,
		Address:  "127.0.0.1",
		Username: "admin",
		Password: "admin",
		Port:     12022}
	defer provider.Disconnect()
	errMsg := fmt.Sprintf("YGOInvalidArgumentError: path %s is not a valid directory", repo.Path)
	assert.PanicsWithValue(suite.T(), errMsg, func() { provider.Connect() })
}

func (suite *CStateErrorsTestSuite) TestNetconfWrongPort() {
	provider := providers.NetconfServiceProvider{
		Address:  "127.0.0.1",
		Username: "wrong",
		Password: "admin",
		Port:     12000}
	defer provider.Disconnect()
	errMsg := fmt.Sprintf("YGOClientError: Could not connect to %s", provider.Address)
	assert.PanicsWithValue(suite.T(), errMsg, func() { provider.Connect() })
}

func (suite *CStateErrorsTestSuite) TestNetconfWrongCredentials() {
	provider := providers.NetconfServiceProvider{
		Address:  "127.0.0.1",
		Username: "wrong",
		Password: "admin",
		Port:     12022}
	defer provider.Disconnect()
	errMsg := fmt.Sprintf("YGOClientError: Could not connect to %s", provider.Address)
	assert.PanicsWithValue(suite.T(), errMsg, func() { provider.Connect() })
}

func (suite *CStateErrorsTestSuite) TestNetconfWrongProtocol() {
	provider := providers.NetconfServiceProvider{
		Address:  "127.0.0.1",
		Username: "admin",
		Password: "admin",
		Port:     12022,
		Protocol: "/dev/null"}
	defer provider.Disconnect()
	errMsg := "YGOOperationNotSupportedError: Protocol is not supported!"
	assert.PanicsWithValue(suite.T(), errMsg, func() { provider.Connect() })
}

func (suite *CStateErrorsTestSuite) TestOpenDaylightNotSupportedProtocol() {
	odlProvider := providers.OpenDaylightServiceProvider{
		Path:           TestHome,
		Address:        "localhost",
		Username:       "admin",
		Password:       "admin",
		Port:           12306,
		EncodingFormat: types.XML,
		Protocol:       types.Netconf}
	defer odlProvider.Disconnect()
	errMsg := "YGOServiceProviderError: Netconf protocol currently not supported"
	assert.PanicsWithValue(suite.T(), errMsg, func() { odlProvider.Connect() })
}

func (suite *CStateErrorsTestSuite) TestOpenDaylightInvalidNodeID() {
	odlProvider := providers.OpenDaylightServiceProvider{
		Path:           TestHome,
		Address:        "localhost",
		Username:       "admin",
		Password:       "admin",
		Port:           12306,
		EncodingFormat: types.JSON,
		Protocol:       types.Restconf}
	defer odlProvider.Disconnect()
	odlProvider.Connect()

	nodeIDs := odlProvider.GetNodeIDs()
	fmt.Printf("nodeIDS: %v\n", nodeIDs)

	nodeID := "xe"
	errMsg := fmt.Sprintf("YGOServiceProviderError: Invalid node id %v", nodeID)
	assert.PanicsWithValue(suite.T(), errMsg, func() { odlProvider.GetNodeProvider(nodeID) })
}

func (suite *CStateErrorsTestSuite) TestCodecInvalidEncode() {
	codec := services.CodecService{}
	provider := providers.CodecServiceProvider{}

	bgp := oc_bgp.Bgp{}
	bgp.Global.Config.As = 65172
	bgp.Global.Config.RouterId = ""

	errMsg := `YGOModelError: Value "" does not satisfy the constraint ` +
		`"(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}` +
		`([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])" (range, length, or pattern). ` +
		`Path: /openconfig-bgp:bgp/global/config/router-id`
	assert.PanicsWithValue(suite.T(), errMsg, func() { codec.Encode(&provider, &bgp) })
}

func (suite *CStateErrorsTestSuite) TestCodecInvalidDecode1() {
	codec := services.CodecService{}
	provider := providers.CodecServiceProvider{}
	assert.Panics(suite.T(), func() { codec.Decode(&provider, invalidXMLBgpPayload1) })
}

func (suite *CStateErrorsTestSuite) TestCodecInvalidDecode2() {
	codec := services.CodecService{}
	provider := providers.CodecServiceProvider{}
	provider.Encoding = types.XML

	errMsg := `YGOModelError: Value "wrong router id" does not satisfy the constraint ` +
		`"(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}` +
		`([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])" (range, length, or pattern). ` +
		`Path: /openconfig-bgp:bgp/global/config/router-id`
	assert.PanicsWithValue(suite.T(), errMsg, func() { codec.Decode(&provider, invalidXMLBgpPayload2) })
}

func TestCStateErrorsTestSuite(t *testing.T) {
	if testing.Verbose() {
		ydk.EnableLogging(ydk.Debug)
	}
	suite.Run(t, new(CStateErrorsTestSuite))
}
