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
