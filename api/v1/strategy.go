package v1

type StrategyRequest struct {
	UserId     int64 `json:"userId"`
	StrategyId int64 `json:"strategyId" binding:"required"`
}

type StrategyIndexRequest struct {
	UserId          int64 `json:"userId"`
	StrategyIndexId int64 `json:"strategyIndexId" binding:"required"`
}

type CreateStrategyRequest struct {
	UserId            int64          `json:"userId"`
	Name              string         `json:"name" binding:"required"`
	InstId            string         `json:"InstId" binding:"required"`
	SubscriptionState int32          `json:"subscriptionState"`
	IndexList         []IndexRequest `json:"indexList" binding:"required"`
}

type IndexRequest struct {
	IndexId       int64  `json:"indexId" binding:"required" example:"1"`
	InstId        string `json:"InstId" binding:"required" example:"DOGE-USDT-SWAP"`
	IndexConfig   string `json:"indexConfig" binding:"required" example:"[12, 26,9]"`
	WarningConfig string `json:"warningConfig" binding:"required" example:"{\"index\": 0}"`
	Bar           string `json:"bar" binding:"required" example:"5m"`
}
