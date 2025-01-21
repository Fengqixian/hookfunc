package okx

type Config struct {
	Server            string
	Wallets           []PayWalletInfo
	SubscriptionPrice []int64
	Limit             int // 限制条数
	Retry             int // 重试次数
}

type PayWalletInfo struct {
	WalletAddress string
	Chain         string
	Coin          string
}

type TransactionRecordResponse struct {
	Data    []TransactionRecord `json:"data"`
	Success bool                `json:"success"`
	Meta    struct {
		At       int64 `json:"at"`
		PageSize int   `json:"page_size"`
	} `json:"meta"`
}

type TransactionRecord struct {
	TransactionId  string `json:"transaction_id"`
	BlockTimestamp int64  `json:"block_timestamp"`
	From           string `json:"from"`
	To             string `json:"to"`
	Type           string `json:"type"`
	Value          string `json:"value"`
}
