package utils

import (
	"fmt"
	"sort"
	"strconv"

	"errors"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	"github.com/cosmos/cosmos-sdk/types/bech32"

	"github.com/Stride-Labs/stride/x/claim/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	recordstypes "github.com/Stride-Labs/stride/x/records/types"
)

var ADMINS = map[string]bool{
	"stride1k8c2m5cn322akk5wy8lpt87dd2f4yh9azg7jlh": true, // F5
	"stride10d07y265gmmuvt4z0w9aw880jnsr700jefnezl": true, // gov module
}

func FilterDepositRecords(arr []recordstypes.DepositRecord, condition func(recordstypes.DepositRecord) bool) (ret []recordstypes.DepositRecord) {
	for _, elem := range arr {
		if condition(elem) {
			ret = append(ret, elem)
		}
	}
	return ret
}

func Int64ToCoinString(amount int64, denom string) string {
	return strconv.FormatInt(amount, 10) + denom
}

func ValidateAdminAddress(address string) error {
	if !ADMINS[address] {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid creator address (%s)", address))
	}
	return nil
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func HostZoneUnbondingKeys(m map[string]*recordstypes.HostZoneUnbonding) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func StringToIntMapKeys(m map[string]int64) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func StringToStringMapKeys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func StringToStringSliceMapKeys(m map[string][]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

//==============================  ADDRESS VERIFICATION UTILS  ================================
// ref: https://github.com/cosmos/cosmos-sdk/blob/b75c2ebcfab1a6b535723f1ac2889a2fc2509520/types/address.go#L177

var errBech32EmptyAddress = errors.New("decoding Bech32 address failed: must provide a non empty address")

// GetFromBech32 decodes a bytestring from a Bech32 encoded string.
func GetFromBech32(bech32str, prefix string) ([]byte, error) {
	if len(bech32str) == 0 {
		return nil, errBech32EmptyAddress
	}

	hrp, bz, err := bech32.DecodeAndConvert(bech32str)
	if err != nil {
		return nil, err
	}

	if hrp != prefix {
		return nil, fmt.Errorf("invalid Bech32 prefix; expected %s, got %s", prefix, hrp)
	}

	return bz, nil
}

// VerifyAddressFormat verifies that the provided bytes form a valid address
// according to the default address rules or a custom address verifier set by
// GetConfig().SetAddressVerifier().
// TODO make an issue to get rid of global Config
// ref: https://github.com/cosmos/cosmos-sdk/issues/9690
func VerifyAddressFormat(bz []byte) error {
	verifier := func(bz []byte) error {
		n := len(bz)
		if n == 20 {
			return nil
		}
		return fmt.Errorf("incorrect address length %d", n)
	}
	if verifier(bz) != nil {
		return verifier(bz)
	}

	if len(bz) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownAddress, "addresses cannot be empty")
	}

	if len(bz) > address.MaxAddrLen {
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownAddress, "address max length is %d, got %d", address.MaxAddrLen, len(bz))
	}

	return nil
}

// AccAddress a wrapper around bytes meant to represent an account address.
// When marshaled to a string or JSON, it uses Bech32.
type AccAddress []byte

// AccAddressFromBech32 creates an AccAddress from a Bech32 string.
func AccAddressFromBech32(address string, bech32prefix string) (addr AccAddress, err error) {
	if len(strings.TrimSpace(address)) == 0 {
		return AccAddress{}, errors.New("empty address string is not allowed")
	}

	bz, err := GetFromBech32(address, bech32prefix)
	if err != nil {
		return nil, err
	}

	err = VerifyAddressFormat(bz)
	if err != nil {
		return nil, err
	}

	return AccAddress(bz), nil
}

// check string array inclusion
func ContainsString(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

//==============================  AIRDROP UTILS  ================================
// max64 returns the maximum of its inputs.
func Max64(i, j int64) int64 {
	if i > j {
		return i
	}
	return j
}

// Min64 returns the minimum of its inputs.
func Min64(i, j int64) int64 {
	if i < j {
		return i
	}
	return j
}

// Insert a value in a slice at a given index
// 0 <= index <= len(a)
func Insert(a []interface{}, index int, value interface{}) []interface{} {
	if len(a) == index { // nil or empty slice or after last element
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value
	return a
}

// Compute coin amount for specific period using linear vesting calculation algorithm.
func GetVestedCoinsAt(vAt int64, vStart int64, vLength int64, vCoins sdk.Coins) sdk.Coins {
	var vestedCoins sdk.Coins

	vEnd := vStart + vLength
	if vAt <= vStart {
		return sdk.Coins{}
	} else if vAt >= vEnd {
		return vCoins
	}

	// calculate the vesting scalar
	portion := sdk.NewDec(vAt - vStart).Quo(sdk.NewDec(vLength))

	for _, ovc := range vCoins {
		vestedAmt := ovc.Amount.ToDec().Mul(portion).RoundInt()
		vestedCoins = append(vestedCoins, sdk.NewCoin(ovc.Denom, vestedAmt))
	}

	return vestedCoins
}

// Get airdrop duration for action
func GetAirdropDurationForAction(action types.Action) int64 {
	if action == types.ActionDelegateStake {
		return int64(types.DefaultVestingDurationForDelegateStake.Seconds())
	} else if action == types.ActionLiquidStake {
		return int64(types.DefaultVestingDurationForLiquidStake.Seconds())
	}
	return int64(0)
}
