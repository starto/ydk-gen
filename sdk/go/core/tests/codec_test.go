package test

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/CiscoDevNet/ydk-go/ydk"
	"github.com/CiscoDevNet/ydk-go/ydk/models/ydktest/oc_pattern"
	ysanity_bgp "github.com/CiscoDevNet/ydk-go/ydk/models/ydktest/openconfig_bgp"
	ysanity_bgp_types "github.com/CiscoDevNet/ydk-go/ydk/models/ydktest/openconfig_bgp_types"
	ysanity "github.com/CiscoDevNet/ydk-go/ydk/models/ydktest/sanity"
	"github.com/CiscoDevNet/ydk-go/ydk/providers"
	"github.com/CiscoDevNet/ydk-go/ydk/services"
	"github.com/CiscoDevNet/ydk-go/ydk/types"
	"github.com/stretchr/testify/suite"
	"testing"
)

const (
	xmlBgpPayload = `<bgp xmlns="http://openconfig.net/yang/bgp">
  <global>
    <config>
      <as>65172</as>
      <router-id>1.2.3.4</router-id>
    </config>
    <afi-safis>
      <afi-safi>
        <afi-safi-name xmlns:oc-bgp-types="http://openconfig.net/yang/bgp-types">oc-bgp-types:IPV6_UNICAST</afi-safi-name>
        <config>
          <afi-safi-name xmlns:oc-bgp-types="http://openconfig.net/yang/bgp-types">oc-bgp-types:IPV6_UNICAST</afi-safi-name>
          <enabled>true</enabled>
        </config>
      </afi-safi>
      <afi-safi>
        <afi-safi-name xmlns:oc-bgp-types="http://openconfig.net/yang/bgp-types">oc-bgp-types:IPV4_UNICAST</afi-safi-name>
        <config>
          <afi-safi-name xmlns:oc-bgp-types="http://openconfig.net/yang/bgp-types">oc-bgp-types:IPV4_UNICAST</afi-safi-name>
          <enabled>true</enabled>
        </config>
      </afi-safi>
    </afi-safis>
  </global>
</bgp>
`

	jsonBgpPayload = `{
  "openconfig-bgp:bgp": {
    "global": {
      "afi-safis": {
        "afi-safi": [
          {
            "afi-safi-name": "openconfig-bgp-types:IPV6_UNICAST",
            "config": {
              "afi-safi-name": "openconfig-bgp-types:IPV6_UNICAST",
              "enabled": true
            }
          },
          {
            "afi-safi-name": "openconfig-bgp-types:IPV4_UNICAST",
            "config": {
              "afi-safi-name": "openconfig-bgp-types:IPV4_UNICAST",
              "enabled": true
            }
          }
        ]
      },
      "config": {
        "as": 65172,
        "router-id": "1.2.3.4"
      }
    }
  }
}
`

	xmlRunnerPayload = `<runner xmlns="http://cisco.com/ns/yang/ydktest-sanity">
  <two-list>
    <ldata>
      <number>21</number>
      <name>runner:twolist:ldata:21</name>
      <subl1>
        <number>211</number>
        <name>runner:twolist:ldata:21:subl1:211</name>
      </subl1>
      <subl1>
        <number>212</number>
        <name>runner:twolist:ldata:21:subl1:212</name>
      </subl1>
    </ldata>
    <ldata>
      <number>22</number>
      <name>runner:twolist:ldata:22</name>
      <subl1>
        <number>221</number>
        <name>runner:twolist:ldata:22:subl1:221</name>
      </subl1>
      <subl1>
        <number>222</number>
        <name>runner:twolist:ldata:22:subl1:222</name>
      </subl1>
    </ldata>
  </two-list>
</runner>
`

	xmlRunnerPayload2 = `<runner xmlns="http://cisco.com/ns/yang/ydktest-sanity">
  <ytypes>
    <built-in-t>
      <enum-value>local</enum-value>
    </built-in-t>
  </ytypes>
</runner>`

	jsonRunnerPayload = `{
  "ydktest-sanity:runner": {
    "two-list": {
      "ldata": [
        {
          "number": 21,
          "name": "runner:twolist:ldata:21",
          "subl1": [
            {
              "number": 211,
              "name": "runner:twolist:ldata:21:subl1:211"
            },
            {
              "number": 212,
              "name": "runner:twolist:ldata:21:subl1:212"
            }
          ]
        },
        {
          "number": 22,
          "name": "runner:twolist:ldata:22",
          "subl1": [
            {
              "number": 221,
              "name": "runner:twolist:ldata:22:subl1:221"
            },
            {
              "number": 222,
              "name": "runner:twolist:ldata:22:subl1:222"
            }
          ]
        }
      ]
    }
  }
}
`

	jsonRunnerPayload2 = `{
  "ydktest-sanity:built-in-t": {
    "enum-value": "local"
  }
}`

	xmlOCPatternPayload = `<oc-A xmlns="http://cisco.com/ns/yang/oc-pattern">
  <a>Hello</a>
</oc-A>
`
)

func equalPayload(s1, s2 string, um func([]uint8, interface{}) error, m func(interface{}) ([]byte, error)) bool {
	var o1 interface{}
	var o2 interface{}

	var err error
	err = um([]byte(s1), &o1)
	if err != nil {
		panic(fmt.Sprintf("Error unmashalling string 1: %s", err.Error()))
	}
	err = um([]byte(s2), &o2)
	if err != nil {
		panic(fmt.Sprintf("Error unmashalling string 1: %s", err.Error()))
	}

	// Unmarshal and Marshal the payload in case list items in payload shows in random order
	p1, err := m(o1)
	if err != nil {
		panic(fmt.Sprintf("Error mashalling object 1: %s", err.Error()))
	}
	p2, err := m(o1)
	if err != nil {
		panic(fmt.Sprintf("Error mashalling object 1: %s", err.Error()))
	}

	return string(p1) == string(p2)
}

type CodecTestSuite struct {
	suite.Suite
	Provider providers.CodecServiceProvider
	Codec    services.CodecService
}

func (suite *CodecTestSuite) SetupSuite() {
	suite.Codec = services.CodecService{}
	suite.Provider = providers.CodecServiceProvider{}
}

func (suite *CodecTestSuite) BeforeTest(suiteName, testName string) {
	fmt.Printf("%v: %v ...\n", suiteName, testName)
}

func (suite *CodecTestSuite) TearDownSuite() {
}

func config(bgp *ysanity_bgp.Bgp) {
	bgp.Global.Config.As = 65172
	bgp.Global.Config.RouterId = "1.2.3.4"

	ipv6Afisafi := ysanity_bgp.Bgp_Global_AfiSafis_AfiSafi{}
	ipv6Afisafi.AfiSafiName = &ysanity_bgp_types.Ipv6_Unicast{}
	ipv6Afisafi.Config.AfiSafiName = &ysanity_bgp_types.Ipv6_Unicast{}
	ipv6Afisafi.Config.Enabled = true
	bgp.Global.AfiSafis.AfiSafi = append(bgp.Global.AfiSafis.AfiSafi, ipv6Afisafi)

	ipv4Afisafi := ysanity_bgp.Bgp_Global_AfiSafis_AfiSafi{}
	ipv4Afisafi.AfiSafiName = &ysanity_bgp_types.Ipv4_Unicast{}
	ipv4Afisafi.Config.AfiSafiName = &ysanity_bgp_types.Ipv4_Unicast{}
	ipv4Afisafi.Config.Enabled = true
	bgp.Global.AfiSafis.AfiSafi = append(bgp.Global.AfiSafis.AfiSafi, ipv4Afisafi)
}

func configRunner(runner *ysanity.Runner) {
	elem1 := ysanity.Runner_TwoList_Ldata{}
	elem2 := ysanity.Runner_TwoList_Ldata{}

	elem1.Number = 21
	elem1.Name = "runner:twolist:ldata:21"

	elem2.Number = 22
	elem2.Name = "runner:twolist:ldata:22"

	elem11 := ysanity.Runner_TwoList_Ldata_Subl1{}
	elem12 := ysanity.Runner_TwoList_Ldata_Subl1{}

	elem11.Number = 211
	elem11.Name = "runner:twolist:ldata:21:subl1:211"
	elem12.Number = 212
	elem12.Name = "runner:twolist:ldata:21:subl1:212"

	elem1.Subl1 = append(elem1.Subl1, elem11)
	elem1.Subl1 = append(elem1.Subl1, elem12)

	elem21 := ysanity.Runner_TwoList_Ldata_Subl1{}
	elem22 := ysanity.Runner_TwoList_Ldata_Subl1{}

	elem21.Number = 221
	elem21.Name = "runner:twolist:ldata:22:subl1:221"
	elem22.Number = 222
	elem22.Name = "runner:twolist:ldata:22:subl1:222"

	elem2.Subl1 = append(elem2.Subl1, elem21)
	elem2.Subl1 = append(elem2.Subl1, elem22)

	runner.TwoList.Ldata = append(runner.TwoList.Ldata, elem1)
	runner.TwoList.Ldata = append(runner.TwoList.Ldata, elem2)
}

func (suite *CodecTestSuite) TestXMLEncode() {
	bgp := ysanity_bgp.Bgp{}
	config(&bgp)

	suite.Provider.Encoding = types.XML
	payload := suite.Codec.Encode(&suite.Provider, &bgp)

	result := equalPayload(payload, xmlBgpPayload, xml.Unmarshal, xml.Marshal)
	suite.Equal(result, true)
}

func (suite *CodecTestSuite) TestXMLDecode() {

	for i := 0; i < 100; i++ {
		bgp := ysanity_bgp.Bgp{}
		config(&bgp)

		suite.Provider.Encoding = types.XML

		entity := suite.Codec.Decode(&suite.Provider, xmlBgpPayload)
		bgpDecodec := entity.(*ysanity_bgp.Bgp)

		suite.Equal(types.EntityEqual(&bgp, bgpDecodec), true)
	}
}

func (suite *CodecTestSuite) TestXMLEncodeDecode() {
	bgp := ysanity_bgp.Bgp{}
	config(&bgp)

	suite.Provider.Encoding = types.XML

	payload := suite.Codec.Encode(&suite.Provider, &bgp)
	entity := suite.Codec.Decode(&suite.Provider, payload)
	bgpDecodec := entity.(*ysanity_bgp.Bgp)

	suite.Equal(types.EntityEqual(&bgp, bgpDecodec), true)
}

func (suite *CodecTestSuite) TestXMLDecodeEncode() {
	suite.Provider.Encoding = types.XML

	entity := suite.Codec.Decode(&suite.Provider, xmlBgpPayload)
	bgpDecodec := entity.(*ysanity_bgp.Bgp)
	payload := suite.Codec.Encode(&suite.Provider, bgpDecodec)

	result := equalPayload(payload, xmlBgpPayload, xml.Unmarshal, xml.Marshal)
	suite.Equal(result, true)
}

func (suite *CodecTestSuite) TestJSONEncode() {
	bgp := ysanity_bgp.Bgp{}
	config(&bgp)

	suite.Provider.Encoding = types.JSON
	payload := suite.Codec.Encode(&suite.Provider, &bgp)

	result := equalPayload(payload, jsonBgpPayload, json.Unmarshal, json.Marshal)
	suite.Equal(result, true)
}

func (suite *CodecTestSuite) TestJSONDecode() {
	bgp := ysanity_bgp.Bgp{}
	config(&bgp)

	suite.Provider.Encoding = types.JSON

	entity := suite.Codec.Decode(&suite.Provider, jsonBgpPayload)
	bgpDecodec := entity.(*ysanity_bgp.Bgp)

	suite.Equal(types.EntityEqual(&bgp, bgpDecodec), true)
}

func (suite *CodecTestSuite) TestJSONDecodeEncode() {
	bgp := ysanity_bgp.Bgp{}
	config(&bgp)

	suite.Provider.Encoding = types.JSON

	entity := suite.Codec.Decode(&suite.Provider, jsonBgpPayload)
	bgpDecodec := entity.(*ysanity_bgp.Bgp)
	payload := suite.Codec.Encode(&suite.Provider, bgpDecodec)

	result := equalPayload(payload, jsonBgpPayload, json.Unmarshal, json.Marshal)
	suite.Equal(result, true)
}

func (suite *CodecTestSuite) TestJSONEncodeDecode() {
	bgp := ysanity_bgp.Bgp{}
	config(&bgp)

	suite.Provider.Encoding = types.JSON

	payload := suite.Codec.Encode(&suite.Provider, &bgp)
	entity := suite.Codec.Decode(&suite.Provider, payload)
	bgpDecodec := entity.(*ysanity_bgp.Bgp)

	suite.Equal(types.EntityEqual(&bgp, bgpDecodec), true)
}

func (suite *CodecTestSuite) TestXMLEncode1() {
	suite.Provider.Encoding = types.XML
	runner := ysanity.Runner{}
	configRunner(&runner)

	payload := suite.Codec.Encode(&suite.Provider, &runner)

	result := equalPayload(payload, xmlRunnerPayload, xml.Unmarshal, xml.Marshal)
	suite.Equal(result, true)
}

func (suite *CodecTestSuite) TestXMLEncode2() {
	suite.Provider.Encoding = types.XML
	r1 := ysanity.Runner{}
	r1.Ytypes.BuiltInT.EnumValue = ysanity.YdkEnumTest_local

	payload := suite.Codec.Encode(&suite.Provider, &r1)
	result := equalPayload(payload, xmlRunnerPayload2, xml.Unmarshal, xml.Marshal)
	suite.Equal(result, true)
	fmt.Println(payload)
}

func (suite *CodecTestSuite) TestJSONEncode1() {
	suite.Provider.Encoding = types.JSON
	runner := ysanity.Runner{}
	configRunner(&runner)

	payload := suite.Codec.Encode(&suite.Provider, &runner)

	result := equalPayload(payload, jsonRunnerPayload, json.Unmarshal, json.Marshal)
	suite.Equal(result, true)
}

func (suite *CodecTestSuite) TestJSONEncode2() {
	suite.Provider.Encoding = types.JSON
	r1 := ysanity.Runner{}
	r1.Ytypes.BuiltInT.EnumValue = ysanity.YdkEnumTest_local

	payload := suite.Codec.Encode(&suite.Provider, &r1)
	result := equalPayload(payload, jsonRunnerPayload2, json.Unmarshal, json.Marshal)
	suite.Equal(result, true)
}

func (suite *CodecTestSuite) TestXMLDecodeOCPattern() {
	suite.Provider.Encoding = types.XML
	objA := oc_pattern.OcA{}

	objA.A = "Hello"

	entity := suite.Codec.Decode(&suite.Provider, xmlOCPatternPayload)
	ocA := entity.(*oc_pattern.OcA)

	suite.Equal(types.EntityEqual(entity, ocA), true)
}

func TestCodecTestSuite(t *testing.T) {
	if testing.Verbose() {
		ydk.EnableLogging(ydk.Debug)
	}
	suite.Run(t, new(CodecTestSuite))
}
