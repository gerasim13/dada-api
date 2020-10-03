## Example

## POST `/aggrInfo`

- from: `USDT`
- to: `DAI`
- amount: `1200000`

```json
{
  "from_name": "ETH",
  "to_name": "DAI",
  "from_addr": "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE",
  "to_addr": "0x6b175474e89094c44da98b954eedeac495271d0f",
  "exchange_pairs": [
    {
      "contract_name": "1inch",
      "ratio": "743934090251009241497",
      "tx_fee": "614915073000000000"
    },
    {
      "contract_name": "Sushiswap",
      "ratio": "743356707489010273018",
      "tx_fee": "84454919861122738"
    },
    {
      "contract_name": "Kyber",
      "ratio": "743356707489010273018",
      "tx_fee": ""
    },
    {
      "contract_name": "Mooniswap",
      "ratio": "742954730453393981837",
      "tx_fee": "153637461231001546"
    },
    {
      "contract_name": "UniswapV2",
      "ratio": "741948272651386670794",
      "tx_fee": "73077349968450170"
    },
    {
      "contract_name": "Bancor",
      "ratio": "284990699382090366116",
      "tx_fee": "34590062323383505"
    },
    {
      "contract_name": "Paraswap",
      "ratio": "0",
      "tx_fee": "88697086144942528"
    },
    {
      "contract_name": "Dforce",
      "ratio": "0",
      "tx_fee": "20592359068442936"
    },
    {
      "contract_name": "ZeroX",
      "ratio": "-9223372036854775808",
      "tx_fee": ""
    }
  ]
}
```

### Get: `tokenlist`

```
[
  {
    "name": "ETH",
    "addr": "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"
  },
  {
    "name": "USDC",
    "addr": "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"
  },
  {
    "name": "YFI",
    "addr": "0x0bc529c00c6401aef6d220be8c6ea1667f6ad93e"
  }
  ...
]
```

### POST: `swapInfo`

- `contract`: `UniswapV2`
- `from`: `ETH`
- `to`: `DAI`
- `amount`: `10000000`
- `user`: `0xcd69c8CbBFe5b1219C0f8911204aA961294E74e3` (用户地址)
- `slippage`: `5` (滑点为 0.05%)

```
{
  "data": "0x7ff36ab500000000000000000000000000000000000000000000000000000000009882f80000000000000000000000000000000000000000000000000000000000000080000000000000000000000000cd69c8cbbfe5b1219c0f8911204aa961294e74e3000000000000000000000000000000000000000000000000000000005f6304ff0000000000000000000000000000000000000000000000000000000000000002000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc20000000000000000000000006b175474e89094c44da98b954eedeac495271d0f",
  "tx_fee": "76880097964862601",
  "contract_addr": "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D",
  "from_token_amount": "10000000",
  "to_token_amount": "3750340403",
  "from_token_addr": "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
  "allowance": "0",
  "allowance_satisfied": true,
  "allowance_data": "0x095ea7b300000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000989680"
}
```

- 当 `from` 为 ETH 时，不用去执行 approve

- 当`from`为其他 ERC20 token 时，如果 allowance_satisfied 为 false，需要执行 approve，执行的合约地址为`from_token_addr`, call_data 为：`allowance_data`
- `contract` 可以选择为：`UniswapV2` ，`Bancor`，`Dforce`，`Kyber`，`Mooniswap`，`Sushiswap`

### GET: `/invest/list`

```json
{
  "address": "0x0d4a11d5eeaac28ec3f61d100daf4d40471f1852",
  "platform": "UniswapV2",
  "liquidity": "396157271.3394341760047949599262313",
  "reserves": ["588180.540502564651528039", "198623914.846887"],
  "tokenprices": [
    "0.002961277552886694139943719263862072",
    "337.6920880061331057763689995205936"
  ],
  "volumes": ["6201321.981781429014459259", "2355360952.854604"],
  "reserveUSD": "396157271.3394341760047949599262313",
  "reserveETH": "1176361.081005129303056078",
  "totalSupply": "8.739720560464409088",
  "volumeUSD": "2339564927.901699586387007576320898",
  "tokens": [
    {
      "address": "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
      "name": "Wrapped Ether",
      "symbol": "WETH",
      "logo": "https://1inch.exchange/assets/tokens/0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2.png"
    },
    {
      "address": "0xdac17f958d2ee523a2206206994597c13d831ec7",
      "name": "Tether USD",
      "symbol": "USDT",
      "logo": "https://1inch.exchange/assets/tokens/0xdac17f958d2ee523a2206206994597c13d831ec7.png"
    }
  ]
}
```

### POST: `/invest/estimate`

- dex: "UniswapV2"
- pool: "0x39444e8Ee494c6212054CFaDF67abDBE97e70207"
- amount: "1.6"
- token: "ETH"
- slippage: "300"(滑点为 3%)

```json
{
  "LP": "60881492234176327580",
  "tokens": {
    "MOO": ["1937805605546836037", "1.937805606"],
    "WETH": ["800000000000000000", "0.8"]
  },
  "invest_amount": "1600000000000000000"
}
```

### POST: `/invest/prepare`

- dex: "UniswapV2"
- pool: "0x39444e8Ee494c6212054CFaDF67abDBE97e70207"
- amount: "1.6"
- token: "ETH"
- slippage: "300"(滑点为 3%)
- user: "0x92E73408801e713f8371f8A8c31a40130ae61a40"

```json
{
  "data": "0x1d57232000000000000000000000000092e73408801e713f8371f8a8c31a40130ae61a400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000b93152b59e65a6de8d3464061bcc1d68f6749f98000000000000000000000000c778417e063141139fce010982780140aa0cd5ab00000000000000000000000000000000000000000000000016345785d8a000000000000000000000000000000000000000000000000000000000000000000000",
  "tx_fee": "",
  "contract_addr": "0xf092A8ff521B39af211DC426e8Eb61c0726147A2",
  "from_token_addr": "",
  "from_token_amount": "1600000000000000000",
  "allowance": "",
  "allowance_satisfied": true,
  "allowance_data": ""
}
```

返回类似 swapinfo

### POST: `/invest/estimate_prepare`

- dex: "UniswapV2"
- pool: "0x39444e8Ee494c6212054CFaDF67abDBE97e70207"
- amount: "1.6"
- token: "ETH"
- slippage: "300"(滑点为 3%)
- user: "0x92E73408801e713f8371f8A8c31a40130ae61a40"

```json
{
  "prepare": {
    "data": "0x1d57232000000000000000000000000092e73408801e713f8371f8a8c31a40130ae61a400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000b93152b59e65a6de8d3464061bcc1d68f6749f98000000000000000000000000c778417e063141139fce010982780140aa0cd5ab00000000000000000000000000000000000000000000000016345785d8a000000000000000000000000000000000000000000000000000000000000000000000",
    "tx_fee": "",
    "contract_addr": "0xf092A8ff521B39af211DC426e8Eb61c0726147A2",
    "from_token_addr": "",
    "from_token_amount": "1600000000000000000",
    "allowance": "",
    "allowance_satisfied": true,
    "allowance_data": ""
  },
  "estimate": {
    "LP": "60881492234176327580",
    "tokens": {
      "MOO": ["1937805605546836037", "1.937805606"],
      "WETH": ["800000000000000000", "0.8"]
    },
    "invest_amount": "1600000000000000000"
  }
}
```

返回类似 swapinfo

### POST: `/staking/prepare`

- dex: "UniswapV2"
- pool: "0xa1484C3aa22a66C62b77E0AE78E15258bd0cB711"
- amount: "1.6"
- user: "0x92E73408801e713f8371f8A8c31a40130ae61a40"

```json
{
  "data": "0xa694fc3a0000000000000000000000000000000000000000000000056bc75e2d63100000",
  "tx_fee": "",
  "contract_addr": "0xa1484C3aa22a66C62b77E0AE78E15258bd0cB711",
  "from_token_addr": "0xA478c2975Ab1Ea89e8196811F51A7B7Ade33eB11",
  "from_token_amount": "100000000000000000000",
  "allowance": "100000000000000000000",
  "allowance_satisfied": false,
  "allowance_data": "0x095ea7b3000000000000000000000000a1484c3aa22a66c62b77e0ae78e15258bd0cb7110000000000000000000000000000000000000000000000056bc75e2d63100000"
}
```
