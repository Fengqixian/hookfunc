package v1

type StrategyRequest struct {
	UserId     int64 `json:"userId"`
	StrategyId int64 `json:"strategyId" binding:"required"`
}

type CreateStrategyRequest struct {
	UserId            int64          `json:"userId"`
	Name              string         `json:"name" binding:"required"`
	SubscriptionState int32          `json:"subscriptionState"`
	IndexList         []IndexRequest `json:"indexList" binding:"required"`
}

type IndexRequest struct {
	IndexId      int64  `json:"indexId" binding:"required"`
	IndexConfig  string `json:"indexConfig" binding:"required"`
	WarningIndex int32  `json:"warningIndex" binding:"required"`
	Bar          string `json:"bar" binding:"required"`
}
