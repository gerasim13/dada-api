package swap_factory

import (
	"bytes"
	"fmt"
	"math"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	"github.com/y2labs-0sh/dada-api/box"
	"github.com/y2labs-0sh/dada-api/daemons"
	"github.com/y2labs-0sh/dada-api/data"
	estimatetxfee "github.com/y2labs-0sh/dada-api/estimate_tx_fee"
	estimatetxrate "github.com/y2labs-0sh/dada-api/estimate_tx_rate"
	log "github.com/y2labs-0sh/dada-api/logger"
	"github.com/y2labs-0sh/dada-api/types"
)

const uniswapSwapExpireTime = 3600

// UniswapSwap 返回swap交易所需参数
// amount 应该是乘以精度的量比如1ETH，则amount为1000000000000000000
// slippage 比如滑点0.05%,则应该传5
func UniswapSwap(fromToken, toToken, userAddr string, slippage int64, amount *big.Int) (types.SwapTx, error) {
	tld, _ := daemons.Get(daemons.DaemonNameTokenList)
	tokenInfos := tld.GetData().(daemons.TokenInfos)
	var swapFunc string
	var valueInput []byte

	amountOutMin := big.NewInt(0)
	aSwapTx := types.SwapTx{}

	if fromToken == "ETH" {
		fromToken = "WETH"
		swapFunc = "swapExactETHForTokens"
	} else if toToken == "ETH" {
		toToken = "WETH"
		swapFunc = "swapExactTokensForETH"
	} else {
		swapFunc = "swapExactTokensForTokens"
	}

	toTokenAmount, err := estimatetxrate.UniswapV2Handler(fromToken, toToken, amount)
	if err != nil {
		log.Error(err)()
		return aSwapTx, err
	}

	amountOutMin = amountOutMin.Mul(toTokenAmount.AmountOut, big.NewInt(10000-slippage))
	amountOutMin = amountOutMin.Div(amountOutMin, big.NewInt(10000))

	amountOutMin = amountOutMin.Div(amountOutMin, big.NewInt(int64(math.Pow10((18 - tokenInfos[toToken].Decimals)))))

	parsedABI, err := abi.JSON(bytes.NewReader(box.Get("abi/uniswapv2.abi")))
	if err != nil {
		log.Error(err)()
		return aSwapTx, err
	}

	if swapFunc == "swapExactETHForTokens" {
		valueInput, err = parsedABI.Pack(
			swapFunc,
			amountOutMin, // receive_token_amount 乘以滑点
			[]common.Address{common.HexToAddress(tokenInfos[fromToken].Address), common.HexToAddress(tokenInfos[toToken].Address)},
			common.HexToAddress(userAddr),
			big.NewInt(time.Now().Unix()+uniswapSwapExpireTime),
		)
		if err != nil {
			log.Error(err)()
			return aSwapTx, err
		}

	} else {
		// function swapExactTokensForTokens( uint amountIn, uint amountOutMin, address[] calldata path, address to, uint deadline )
		// function swapExactTokensForETH(uint amountIn, uint amountOutMin, address[] calldata path, address to, uint deadline)
		// swapFunc == "swapExactTokensForETH" or "swapExactTokensForTokens"
		valueInput, err = parsedABI.Pack(
			swapFunc,
			amount,
			amountOutMin, // receive_token_amount 乘以滑点
			[]common.Address{common.HexToAddress(tokenInfos[fromToken].Address), common.HexToAddress(tokenInfos[toToken].Address)},
			common.HexToAddress(userAddr),
			big.NewInt(time.Now().Unix()+uniswapSwapExpireTime),
		)
		if err != nil {
			log.Error(err)()
			return aSwapTx, err
		}
	}

	aCheckAllowanceResult, err := CheckAllowance(fromToken, data.UniswapV2, userAddr, amount)
	if err != nil {
		log.Error(err)()
		return aSwapTx, err
	}

	aSwapTx = types.SwapTx{
		Data:               fmt.Sprintf("0x%x", valueInput),
		TxFee:              estimatetxfee.TxFeeOfContract["UniswapV2"].String(),
		ContractAddr:       data.UniswapV2,
		FromTokenAmount:    amount.String(),
		ToTokenAmount:      toTokenAmount.AmountOut.String(),
		FromTokenAddr:      tokenInfos[fromToken].Address,
		Allowance:          aCheckAllowanceResult.AllowanceAmount.String(),
		AllowanceSatisfied: aCheckAllowanceResult.IsSatisfied,
		AllowanceData:      aCheckAllowanceResult.AllowanceData,
	}

	return aSwapTx, nil
}
