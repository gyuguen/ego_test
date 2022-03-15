package main

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	panaceatypes "github.com/medibloc/panacea-core/v2/x/market/types"
)

func main() {
	fmt.Println(sdk.AccAddress("panacea1dw3ev60v44q68qrrsz8lmwrx68f70hyrykw5ry").String())
	fmt.Println(panaceatypes.Deal{})
}
