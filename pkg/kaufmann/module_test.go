package kaufmann

import (
	"context"
	"encoding/json"
	"math/big"
	"testing"

	"github.com/DIMO-Network/cloudevent"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	testChainID   = uint64(137)
	testSynthAddr = common.HexToAddress("0x4804e8D1661cd1a1e5dDdE1ff458A7f878c0aC6D")
	testVehAddr   = common.HexToAddress("0xbA5738a18d83D41847dfFbDC6101d37C69c9B0cF")
)

func testModule() Module {
	return Module{ChainID: testChainID, AftermarketContractAddr: testSynthAddr, VehicleContractAddr: testVehAddr}
}

func did(addr common.Address, tokenID int64) string {
	return cloudevent.ERC721DID{ChainID: testChainID, ContractAddress: addr, TokenID: big.NewInt(tokenID)}.String()
}

// TestCloudEventConvert_Kamaleon verifies a kam/ event becomes a single status
// event with producer = synthetic device DID, subject = vehicle DID, and the data
// passed through untouched.
func TestCloudEventConvert_Kamaleon(t *testing.T) {
	msg := []byte(`{"ds":"kam/v0/s","time":"2026-07-23T19:34:31Z","data":{"signals":[{"timestamp":"2026-07-23T19:34:31Z","name":"speed","value":42}]},"deviceTokenId":789,"vehicleTokenId":456}`)

	hdrs, data, err := testModule().CloudEventConvert(context.Background(), msg)
	require.NoError(t, err)
	require.Len(t, hdrs, 1)

	h := hdrs[0]
	assert.Equal(t, cloudevent.TypeStatus, h.Type)
	assert.Equal(t, "kam/v0/s", h.DataVersion)
	assert.Equal(t, did(testSynthAddr, 789), h.Producer)
	assert.Equal(t, did(testVehAddr, 456), h.Subject)
	assert.JSONEq(t, `{"signals":[{"timestamp":"2026-07-23T19:34:31Z","name":"speed","value":42}]}`, string(data))
}

// TestCloudEventConvert_KamaleonNoVehicleToken falls back to the producer as
// subject when no vehicle token id is present.
func TestCloudEventConvert_KamaleonNoVehicleToken(t *testing.T) {
	msg := []byte(`{"ds":"kam/v0/s","time":"2026-07-23T19:34:31Z","data":{"signals":[]},"deviceTokenId":789}`)

	hdrs, _, err := testModule().CloudEventConvert(context.Background(), msg)
	require.NoError(t, err)
	require.Len(t, hdrs, 1)
	assert.Equal(t, did(testSynthAddr, 789), hdrs[0].Producer)
	assert.Equal(t, hdrs[0].Producer, hdrs[0].Subject)
}

// TestCloudEventConvert_KamaleonMissingDeviceToken errors, matching ruptela.
func TestCloudEventConvert_KamaleonMissingDeviceToken(t *testing.T) {
	msg := []byte(`{"ds":"kam/v0/s","time":"2026-07-23T19:34:31Z","data":{"signals":[]}}`)
	_, _, err := testModule().CloudEventConvert(context.Background(), msg)
	require.Error(t, err)
}

// TestSignalConvert_Kamaleon decodes the pre-VSS payload via the default module.
func TestSignalConvert_Kamaleon(t *testing.T) {
	raw := cloudevent.RawEvent{
		CloudEventHeader: cloudevent.CloudEventHeader{DataVersion: "kam/v0/s", Type: cloudevent.TypeStatus},
		Data:             json.RawMessage(`{"signals":[{"timestamp":"2026-07-23T19:34:31Z","name":"speed","value":42}]}`),
	}

	sigs, err := testModule().SignalConvert(context.Background(), raw)
	require.NoError(t, err)
	require.Len(t, sigs, 1)
	assert.Equal(t, "speed", sigs[0].Data.Name)
	assert.Equal(t, float64(42), sigs[0].Data.ValueNumber)
}

// TestCloudEventConvert_RuptelaDelegates verifies the "r/" family still goes
// through the ruptela module (subject = vehicle DID for a status event),
// unchanged by this module.
func TestCloudEventConvert_RuptelaDelegates(t *testing.T) {
	msg := []byte(`{"ds":"r/v0/s","time":"2026-07-23T19:34:31Z","data":{"signals":{}},"deviceTokenId":789,"vehicleTokenId":456}`)

	hdrs, _, err := testModule().CloudEventConvert(context.Background(), msg)
	require.NoError(t, err)
	require.NotEmpty(t, hdrs)
	assert.Equal(t, "r/v0/s", hdrs[0].DataVersion)
	assert.Equal(t, did(testVehAddr, 456), hdrs[0].Subject)
	assert.Equal(t, did(testSynthAddr, 789), hdrs[0].Producer)
}

func TestIsKam(t *testing.T) {
	assert.True(t, isKam("kam/v0/s"))
	assert.False(t, isKam("r/v0/s"))
	assert.False(t, isKam(""))
}
