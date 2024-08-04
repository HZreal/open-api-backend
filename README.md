# 统一接口代理服务平台

## 简介
类似于短信、图片、天气等接口服务平台，但更抽象统一化的接口代理服务平台（产品）。提供用户接口服务，吸引接口服务商入驻上架接口，收取服务费

## 核心业务

- 用户浏览、试用、调用接口；
- 管理员接口管理、上下架、统计；
- 服务商入驻、接口注册；
- 接口付费

## 系统设计

- 接口统一管理服务（平台）
  - https://github.com/HZreal/open-api-backend
- 接口提供服务（实际接口提供源）
  - https://github.com/HZreal/open-api-provider
- 通用业务服务（统一模型、服务类）
  - https://github.com/HZreal/open-api-common
- 网关服务（统一鉴权、路由、访问控制、统计计费等）
  - https://github.com/HZreal/open-api-gateway
- 接口客户端 SDK（用户客户端多语言接口调用封装包）
  - https://github.com/code-elastic/open-api-client-sdk-go
  - https://github.com/code-elastic/open-api-client-sdk-python
  - https://github.com/code-elastic/open-api-client-sdk-node

## 技术点

- 实现接口签名认证算法生成 ak、sk，用于用户接口认证；
- 采用 timestamp + nonce 机制防止接口重放；
- 不同语言客户端接口 SDK 开发；
- 采用 Spring Cloud Gateway 配置化方式进行网关鉴权、限流、染色、业务统计、日志记录；
- 采用 gRPC 远程服务调用进行接口模拟调用
