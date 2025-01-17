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
	IndexId       int64  `json:"indexId" binding:"required"`
	InstId        string `json:"InstId" binding:"required"`
	IndexConfig   string `json:"indexConfig" binding:"required"`
	WarningConfig string `json:"warningConfig" binding:"required"`
	Bar           string `json:"bar" binding:"required"`
}
