package goctp

// 公共-连接
type OnFrontConnectedType func()

// 公共-断开
type OnFrontDisConnectedType func(reson int)

// 公共-登录
type OnRspUserLoginType func(loginField *RspUserLoginField, info *RspInfoField)

// 行情
type OnTickType func(tick *TickField)

// 交易-委托响应
type OnRtnOrderType func(field *OrderField)

// 交易-错误委托
type OnRtnErrOrderType func(field *OrderField, info *RspInfoField)

// 交易-错误撤单
type OnRtnErrActionType func(orderID string, info *RspInfoField)

// 交易-成交响应
type OnRtnTradeType func(field *TradeField)

// 交易-合约状态响应
type OnRtnInstrumentStatusType func(field *InstrumentStatus)
