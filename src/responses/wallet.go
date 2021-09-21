package responses

type Withdraw struct {
    Id string `json:"id"`
}

type Deposit struct {
	Amount        string `json:"amount"`       
	Coin          string `json:"coin"`         
	Network       string `json:"network"`      
	Status        int64  `json:"status"`       
	Address       string `json:"address"`      
	AddressTag    string `json:"addressTag"`   
	TxID          string `json:"txId"`         
	InsertTime    int64  `json:"insertTime"`   
	TransferType  int64  `json:"transferType"` 
	UnlockConfirm string `json:"unlockConfirm"`
	ConfirmTimes  string `json:"confirmTimes"` 
}

type WithdrawHistory struct {
	Address         string `json:"address"`
	Amount          string `json:"amount"`
	ApplyTime       string `json:"applyTime"`
	Coin            string `json:"coin"`
	ID              string `json:"id"`
	WithdrawOrderID string `json:"withdrawOrderId"`
	Network         string `json:"network"`
	TransferType    int64  `json:"transferType"`
	Status          int64  `json:"status"`
	TransactionFee  string `json:"transactionFee"`
	ConfirmNo       int64  `json:"confirmNo"`
	TxID            string `json:"txId"`
}

type DepositAddress struct {
	Address string `json:"address"`
	Coin    string `json:"coin"`
	Tag     string `json:"tag"`
	URL     string `json:"url"`
}

type AccountStatus struct {
    Data string `json:"data"`
}

type ApiTradingStatus struct {
	Data ApiTradingStatusData `json:"data"`
}

type ApiTradingStatusData struct {
	IsLocked           bool             `json:"isLocked"`
	PlannedRecoverTime int64            `json:"plannedRecoverTime"`
	TriggerCondition   TriggerCondition `json:"triggerCondition"`
	Indicators         Indicators       `json:"indicators"`
	UpdateTime         int64            `json:"updateTime"`
}

type Indicators struct {
	Btcusdt []Usdt `json:"BTCUSDT"`
	Ethusdt []Usdt `json:"ETHUSDT"`
}

type Usdt struct {
	I string  `json:"i"`
	C int64   `json:"c"`
	V float64 `json:"v"`
	T float64 `json:"t"`
}

type TriggerCondition struct {
	GCR  int64 `json:"GCR"`
	Ifer int64 `json:"IFER"`
	Ufr  int64 `json:"UFR"`
}
