package swapfactory

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"

	"github.com/y2labs-0sh/aggregator_info/box"
	"github.com/y2labs-0sh/aggregator_info/data"
	estimatetxfee "github.com/y2labs-0sh/aggregator_info/estimate_tx_fee"
	estimatetxrate "github.com/y2labs-0sh/aggregator_info/estimate_tx_rate"
	"github.com/y2labs-0sh/aggregator_info/types"
)

// DforceSwap 返回swap交易所需参数
// amount 应该是乘以精度的量比如1ETH，则amount为1000000000000000000
// slippage 比如滑点0.05%,则应该传5
func DforceSwap(fromToken, toToken, userAddr string, slippage int64, amount *big.Int) (types.SwapTx, error) {

	amountIn := big.NewInt(0)
	var valueInput []byte
	var ok bool
	aSwapTx := types.SwapTx{}

	parsedABI, err := abi.JSON(bytes.NewReader(box.Get("abi/uniswapv2.abi")))
	if err != nil {
		log.Error(err)
		return aSwapTx, err
	}

	// swap(address source, address dest, uint256 sourceAmount)
	valueInput, err = parsedABI.Pack(
		"swap",
		common.HexToAddress(data.TokenInfos[fromToken].Address),
		common.HexToAddress(data.TokenInfos[toToken].Address),
		amountIn,
	)
	if err != nil {
		log.Error(err)
		return aSwapTx, err
	}

	toTokenAmount, err := estimatetxrate.DforceHandler(fromToken, toToken, amount)
	if err != nil {
		log.Error(err)
		return aSwapTx, err
	}

	amountConvertRatio := big.NewInt(0)
	amountConvertRatio, ok = amountConvertRatio.SetString(toTokenAmount.Ratio, 10)
	if !ok {
		log.Error("convert amount Ratio to big.Int error")
		return aSwapTx, err
	}

	aCheckAllowanceResult, err := CheckAllowance(fromToken, data.Dforce, userAddr, amount)
	if err != nil {
		log.Error(err)
		return aSwapTx, err
	}

	aSwapTx = types.SwapTx{
		Data:               fmt.Sprintf("0x%x", valueInput),
		TxFee:              estimatetxfee.TxFeeOfContract["Dforce"],
		ContractAddr:       data.Dforce,
		FromTokenAmount:    amount.String(),
		ToTokenAmount:      amountConvertRatio.String(),
		FromTokenAddr:      data.TokenInfos[fromToken].Address,
		Allowance:          aCheckAllowanceResult.AllowanceAmount.String(),
		AllowanceSatisfied: aCheckAllowanceResult.IsSatisfied,
		AllowanceData:      aCheckAllowanceResult.AllowanceData,
	}

	return aSwapTx, nil
}
