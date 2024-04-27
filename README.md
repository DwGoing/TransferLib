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

### 功能列表
|     **功能**     |   **接口名**   | **说明** |    **链支持**    |
| :--------------: | :------------: | :------: | :--------------: |
| **获取账户信息** |   GetAccount   |          | BTC/ETH/TRON/BSC |
|   **获取余额**   |   GetBalance   |          |   ETH/TRON/BSC   |
|     **转账**     |    Transfer    |          |   ETH/TRON/BSC   |
| **查询交易信息** | GetTransaction |          |   ETH/TRON/BSC   |
