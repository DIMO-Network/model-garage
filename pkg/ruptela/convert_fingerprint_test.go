package ruptela_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/DIMO-Network/cloudevent"
	"github.com/DIMO-Network/model-garage/pkg/ruptela"
	"github.com/stretchr/testify/require"
)

func TestFullFPFromDataConversion(t *testing.T) {
	t.Parallel()
	expectedVIN := "UALLAAAF3AA444482"
	var event cloudevent.RawEvent
	err := json.Unmarshal([]byte(fullFPInputJSON), &event)
	require.NoError(t, err)
	fp, err := ruptela.DecodeFingerprint(event)
	require.NoError(t, err, "error decoding fingerprint")
	require.Equal(t, expectedVIN, fp.VIN, "decoded VIN does not match expected VIN")
}

var fullFPInputJSON = `
{
	"source": "ruptela/TODO",
	"data": {
		"pos": {
			"alt": 1048,
			"dir": 19730,
			"hdop": 6,
			"lat": 522721466,
			"lon": -9014316,
			"sat": 20,
			"spd": 0
		},
		"prt": 0,
		"signals": {
			"102": "0",
			"103": "0",
			"104": "55414C4C41414146",
			"105": "3341413434343438",
			"106": "3200000000000000",
			"107": "0",
			"108": "0",
			"114": "0",
			"135": "0",
			"136": "0",
			"137": "14",
			"173": "1",
			"205": "0",
			"207": "0",
			"29": "37FF",
			"30": "1080",
			"409": "1",
			"49": "FE",
			"50": "FA",
			"5005": "31",
			"5060": "6597",
			"51": "ED",
			"525": "A502A",
			"525_1": "A502A",
			"642": "FFFF",
			"645": "FFFFFFFF",
			"722": "FF",
			"723": "FFFF",
			"754": "FB8F",
			"92": "0",
			"93": "0",
			"94": "0",
			"95": "0",
			"950": "0",
			"96": "FF",
			"97": "FF",
			"98": "0",
			"985": "0",
			"99": "1",
			"999": "0"
		},
		"trigger": 7
	},
	"ds": "r/v0/s",
	"signature": "0x6fb5849e21e66f3e0619f148bc032153aa4c90be4cd175e83c1f959e1bc551d940d516fe74f50aed380e432406675c583e75155bf1c77b9ec0761b1dbe1ab87e1c",
	"subject": "did:erc721:1:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF:33",
	"time": "2024-09-27T08:33:26Z",
	"topic": "devices/0xf47f6579029a1c53388e4d8776413aa3f993cb94/status"
}`

var japanVinSignals = `{
	"source": "ruptela/TODO",
	"data": {
		"pos": {
			"lat": 522721466,
			"lon": -9014316
		},
		"prt": 0,
		"signals": {
			"104": "XXX104",
			"105": "XXX105",
			"106": "XXX106"
		},
		"trigger": 7
	},
	"ds": "r/v0/s"
}`

func TestDecodeFingerprint_JapanVIN(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		vinFrame1   string
		vinFrame2   string
		vinFrame3   string
		expectedVIN string
	}{
		{
			name:        "Japan VIN ZWR90-8000186",
			expectedVIN: "ZWR90-8000186",
			vinFrame1:   "5A575239302D3830",
			vinFrame2:   "3030313836000000",
			vinFrame3:   "0000000000000000",
		},
		{
			name:        "Japan VIN GRX120-3043102",
			expectedVIN: "GRX120-3043102",
			vinFrame1:   "4752583132302D33",
			vinFrame2:   "3034333130320000",
			vinFrame3:   "0000000000000000",
		},
		{
			name:        "Japan VIN GFC27-183060",
			expectedVIN: "GFC27-183060",
			vinFrame1:   "47464332372D3138",
			vinFrame2:   "3330363000000000",
			vinFrame3:   "0000000000000000",
		},
		{
			name:        "Japan VIN DJLAS203662",
			expectedVIN: "DJLAS203662",
			vinFrame1:   "444A4C4153323033",
			vinFrame2:   "3636320000000000",
			vinFrame3:   "0000000000000000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var event cloudevent.RawEvent
			signals := strings.Replace(japanVinSignals, "XXX104", tt.vinFrame1, 1)
			signals = strings.Replace(signals, "XXX105", tt.vinFrame2, 1)
			signals = strings.Replace(signals, "XXX106", tt.vinFrame3, 1)

			err := json.Unmarshal([]byte(signals), &event)
			require.NoError(t, err)

			got, err := ruptela.DecodeFingerprint(event)
			require.NoError(t, err, "error decoding fingerprint")
			require.Equal(t, tt.expectedVIN, got.VIN, "decoded VIN does not match expected VIN")
		})
	}
}

func TestCANBasedVINOnly(t *testing.T) {
	t.Parallel()
	expectedVIN := "W1NGM2BB7RA057440"

	// Build an event that only has CAN-based VIN parts in signals 123/124/125
	var event cloudevent.RawEvent
	canVINEvent := `{
        "source": "ruptela/TODO",
        "data": {
            "signals": {
                "123": "57314e474d324242",
                "124": "3752413035373434",
                "125": "3000000000000000"
            }
        }
    }`
	err := json.Unmarshal([]byte(canVINEvent), &event)
	require.NoError(t, err)
	fp, err := ruptela.DecodeFingerprint(event)
	require.NoError(t, err, "error decoding fingerprint")
	require.Equal(t, expectedVIN, fp.VIN, "decoded VIN does not match expected VIN")
}

func TestPreferStandardOverCAN(t *testing.T) {
	t.Parallel()
	// When both sets exist, we should prefer the standard 104/105/106
	expectedVIN := "UALLAAAF3AA444482" // same as in TestFullFPFromDataConversion
	var event cloudevent.RawEvent
	eventJSON := `{
        "source": "ruptela/TODO",
        "data": {
            "signals": {
                "104": "55414C4C41414146",
                "105": "3341413434343438",
                "106": "3200000000000000",
                "123": "57314e474d324242",
                "124": "3752413035373434",
                "125": "3000000000000000"
            }
        }
    }`
	err := json.Unmarshal([]byte(eventJSON), &event)
	require.NoError(t, err)
	fp, err := ruptela.DecodeFingerprint(event)
	require.NoError(t, err, "error decoding fingerprint")
	require.Equal(t, expectedVIN, fp.VIN, "decoded VIN should prefer standard fields over CAN-based ones")
}

// zeroStandardVINWithCANVINPayload is a real payload where 104/105/106 are all zeros
// and 123/124/125 contain a CAN VIN. Decoder treats zero 104/105/106 as empty and uses CAN.
var zeroStandardVINWithCANVINPayload = `{
  "id": "38wq8Hhcm2OzMxaZJlN96DufRlQ",
  "source": "0x8D8cDB2B26423c8fDbb27321aF20b4659Ce919fD",
  "producer": "did:erc721:137:0x4804e8D1661cd1a1e5dDdE1ff458A7f878c0aC6D:168286",
  "specversion": "1.0",
  "subject": "did:erc721:137:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF:188899",
  "time": "2026-01-29T21:08:18.001Z",
  "type": "dimo.fingerprint",
  "datacontenttype": "application/json",
  "dataversion": "r/v0/s",
  "data": {
    "trigger": 5,
    "prt": 0,
    "pos": {
      "lat": -333597333,
      "lon": -705170183,
      "alt": 8175,
      "dir": 24200,
      "spd": 0,
      "sat": 0,
      "hdop": 0
    },
    "signals": {
      "104": "0000000000000000",
      "105": "0000",
      "106": "0000000000000000",
      "123": "57314E474D324242",
      "124": "3752413035373434",
      "125": "3000000000000000"
    }
  }
}`

func TestZeroStandardVINWithCANVIN_ProducesVINFromCAN(t *testing.T) {
	t.Parallel()
	// 104/105/106 are all zeros; 123/124/125 have CAN VIN → decoder must use CAN and produce VIN.
	expectedVIN := "W1NGM2BB7RA057440"
	var event cloudevent.RawEvent
	err := json.Unmarshal([]byte(zeroStandardVINWithCANVINPayload), &event)
	require.NoError(t, err)
	fp, err := ruptela.DecodeFingerprint(event)
	require.NoError(t, err)
	require.Equal(t, expectedVIN, fp.VIN)
}

func TestDecodeFingerprint_NoVINData(t *testing.T) {
	t.Parallel()
	// Missing or incomplete VIN fields → error.
	for i, payload := range []string{
		`{"data":{"signals":{}}}`,
		`{"data":{"signals":{"104":"","105":"","106":"","123":"","124":"","125":""}}}`,
		`{"data":{"signals":{"104":"0000000000000000","105":"0000","106":"0000000000000000"}}}`,
	} {
		var event cloudevent.RawEvent
		err := json.Unmarshal([]byte(payload), &event)
		require.NoError(t, err)
		_, err = ruptela.DecodeFingerprint(event)
		require.Error(t, err, "payload %d should return error when no VIN data", i)
		require.Contains(t, err.Error(), "missing fingerprint data")
	}
	// All six fields present but all zero hex → success with empty VIN.
	var event cloudevent.RawEvent
	err := json.Unmarshal([]byte(`{"data":{"signals":{"104":"0000000000000000","105":"0000","106":"0000000000000000","123":"0000000000000000","124":"0000","125":"0000000000000000"}}}`), &event)
	require.NoError(t, err)
	fp, err := ruptela.DecodeFingerprint(event)
	require.NoError(t, err)
	require.Empty(t, fp.VIN, "all fields empty must decode to empty VIN")
}

var standardVINPayload = `{
  "id": "39V7aTeOxpybSxW5bRjH0PzntyZ",
  "source": "0xF26421509Efe92861a587482100c6d728aBf1CD0",
  "producer": "did:erc721:137:0x9c94C395cBcBDe662235E0A9d3bB87Ad708561BA:31694",
  "specversion": "1.0",
  "subject": "did:erc721:137:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF:186612",
  "time": "2026-02-11T00:24:03Z",
  "type": "dimo.fingerprint",
  "datacontenttype": "application/json",
  "dataversion": "r/v0/s",
  "data": {
    "trigger": 409,
    "prt": 1,
    "signals": {
      "27": "E",
      "29": "3A10",
      "30": "106C",
      "49": "F6",
      "93": "0",
      "94": "0",
      "95": "0",
      "96": "75",
      "97": "2A",
      "98": "E7",
      "99": "1",
      "100": "0",
      "101": "0",
      "102": "0",
      "103": "0",
      "104": "3146544657374C44",
      "105": "3353464131323637",
      "106": "3500000000000000",
      "107": "6AD",
      "108": "0",
      "135": "0",
      "136": "0",
      "137": "0",
      "143": "0",
      "169": "0",
      "173": "1",
      "402": "0",
      "403": "0",
      "404": "0",
      "407": "0",
      "408": "0",
      "409": "0",
      "418": "1",
      "525": "7F42F",
      "642": "FFFF",
      "644": "4",
      "645": "D3F",
      "722": "27",
      "723": "FFFF",
      "754": "136170",
      "755": "1",
      "762": "8",
      "763": "5",
      "905": "B",
      "950": "0",
      "960": "2DE",
      "961": "2DE",
      "962": "2D9",
      "963": "2DE",
      "964": "FF",
      "965": "0",
      "966": "0",
      "967": "0",
      "968": "0",
      "985": "0",
      "999": "C",
      "1148": "FF",
      "1149": "0",
      "1150": "FF",
      "1155": "0",
      "5005": "0",
      "5060": "0",
      "49_1": "F6",
      "49_2": "F6",
      "49_3": "F6",
      "49_4": "F6",
      "102_1": "0"
    }
  }
}`

func TestDecodeFingerprint_StandardVINPayload(t *testing.T) {
	t.Parallel()
	// Real payload with standard VIN in 104/105/106 → decodes to 17-char VIN.
	expectedVIN := "1FTFW7LD3SFA12675"
	var event cloudevent.RawEvent
	err := json.Unmarshal([]byte(standardVINPayload), &event)
	require.NoError(t, err)
	fp, err := ruptela.DecodeFingerprint(event)
	require.NoError(t, err)
	require.Equal(t, expectedVIN, fp.VIN)
}
