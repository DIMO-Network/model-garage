package status_test

import (
	"cmp"
	"slices"
	"testing"
	"time"

	"github.com/DIMO-Network/model-garage/pkg/ruptela/status"
	"github.com/DIMO-Network/model-garage/pkg/vss"
	"github.com/stretchr/testify/require"
)

func TestValidDTCPayload(t *testing.T) {
	t.Parallel()
	actualSignals, err := status.SignalsFromDTCPayload([]byte(dtcInputJSON))
	require.NoErrorf(t, err, "error converting full input data: %v", err)

	// sort the signals so diffs are easier to read
	sortFunc := func(a, b vss.Signal) int {
		return cmp.Compare(a.Name, b.Name)
	}
	expected := expectedDTCSignals()
	slices.SortFunc(expected, sortFunc)
	slices.SortFunc(actualSignals, sortFunc)
	require.Equal(t, expected, actualSignals, "converted vehicle does not match expected vehicle")
}

func TestEmptyDTCPayload(t *testing.T) {
	t.Parallel()
	_, err := status.SignalsFromDTCPayload([]byte(emptyDtcInputJSON))
	require.Errorf(t, err, "error converting full input data: %v", err)
}

func TestNoDTCPayload(t *testing.T) {
	t.Parallel()
	_, err := status.SignalsFromDTCPayload([]byte(noDtcInputJSON))
	require.Errorf(t, err, "error converting full input data: %v", err)
}

func expectedDTCSignals() []vss.Signal {
	ts = time.Unix(1727360354, 0).UTC()
	return []vss.Signal{
		{TokenID: 33, Timestamp: ts, Name: "obdDTCList", ValueString: `["P0101","P0202"]`, Source: "ruptela/TODO"},
	}
}

var dtcInputJSON = `
	{
	   "subject":"did:nft:1:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF_33",
	   "source":"ruptela/TODO",
	   "data":{
	      "dtc_codes":[
	         {
	            "id":"7E8",
	            "code":"P0101"
	         },
	         {
	            "id":"7E8",
	            "code":"P0202"
	         }
	      ]
	   },
	   "ds":"r/v0/dtc",
	   "signature":"0xb6b130b31b4cd73182008d286fe878bc311a2259b8cfc1ba785495d9c88a028c55e489608191518b3ad26e8226c35a4dfd032f03aac930712ac038e2afeeefc81c",
	   "time":"2024-09-26T14:19:14Z"
	}`

var emptyDtcInputJSON = `
	{
	   "subject":"did:nft:1:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF_33",
	   "source":"ruptela/TODO",
	   "data":{
	      "dtc_codes":[
	        
	      ]
	   },
	   "ds":"r/v0/dtc",
	   "signature":"0xb6b130b31b4cd73182008d286fe878bc311a2259b8cfc1ba785495d9c88a028c55e489608191518b3ad26e8226c35a4dfd032f03aac930712ac038e2afeeefc81c",
	   "time":"2024-09-26T14:19:14Z"
	}`

var noDtcInputJSON = `
	{
	   "subject":"did:nft:1:0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF_33",
	   "source":"ruptela/TODO",
	   "data":{
	   },
	   "ds":"r/v0/dtc",
	   "signature":"0xb6b130b31b4cd73182008d286fe878bc311a2259b8cfc1ba785495d9c88a028c55e489608191518b3ad26e8226c35a4dfd032f03aac930712ac038e2afeeefc81c",
	   "time":"2024-09-26T14:19:14Z"
	}`
