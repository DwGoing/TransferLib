# transfer_lib

### 使用说明
``` shell
# 运行服务
git clone https://github.com/DwGoing/transfer_lib/
cd transfer_lib
go mod tidy
go build -v -o ./build/transfer_lib transfer_lib.go
./build/transfer_lib

# 使用依赖
go get -u github.com/DwGoing/transfer_lib/pkg
```

### 配置说明
``` yaml
Mode: dev
Log:
  Encoding: plain
Name: transfer_lib.rpc
ListenOn: 0.0.0.0:8080
Timeout: 0
Eth:
  Nodes:
    https://ethereum-holesky-rpc.publicnode.com: 100 # 节点Url及权重
  Currencies:
    ETH: # 主链币种
      Contract: # 必须为空
      Decimals: 18
    USDT: # 合约币种
      Contract: 4555Ed1F6D9cb6CC1D52BB88C7525b17a06da0Dd # 合约地址
      Decimals: 18
Tron:
  Nodes:
    grpc.nile.trongrid.io:50051: 100
  ApiKeys: [d9b77ec9-39e0-4765-98d8-2c59188344a0] # 节点ApiKey列表
  Currencies:
    TRX:
      Contract:
      Decimals: 6
    USDT:
      Contract: TXLAQ63Xg1NAzckPwKHvzw7CSEmLMEqcdj
      Decimals: 6
Bsc:
  Nodes:
    https://data-seed-prebsc-2-s1.binance.org:8545: 100
  Currencies:
    BNB:
      Contract:
      Decimals: 18
    USDT:
      Contract: 337610d27c682E347C9cD60BD4b3b107C9d34dDd
      Decimals: 18
```

### 功能列表
|     **功能**     |   **接口名**   | **说明** |    **链支持**    |
| :--------------: | :------------: | :------: | :--------------: |
| **获取账户信息** |   GetAccount   |          | BTC/ETH/TRON/BSC |
|   **获取余额**   |   GetBalance   |          |   ETH/TRON/BSC   |
|     **转账**     |    Transfer    |          |   ETH/TRON/BSC   |
| **查询交易信息** | GetTransaction |          |   ETH/TRON/BSC   |
