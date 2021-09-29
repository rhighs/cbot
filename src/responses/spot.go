package responses

type NewOrderAck struct {
	Symbol        string `json:"symbol"`       
	OrderID       int64  `json:"orderId"`      
	OrderListID   int64  `json:"orderListId"`  
	ClientOrderID string `json:"clientOrderId"`
	TransactTime  int64  `json:"transactTime"` 
}

type NewOrderResult struct {
	Symbol              string `json:"symbol"`
	OrderID             int64  `json:"orderId"`
	OrderListID         int64  `json:"orderListId"`
	ClientOrderID       string `json:"clientOrderId"`
	TransactTime        int64  `json:"transactTime"`
	Price               string `json:"price"`
	OrigQty             string `json:"origQty"`
	ExecutedQty         string `json:"executedQty"`
	CummulativeQuoteQty string `json:"cummulativeQuoteQty"`
	Status              string `json:"status"`
	TimeInForce         string `json:"timeInForce"`
	Type                string `json:"type"`
	Side                string `json:"side"`
}

//=================================================================
type Fill struct {
	Price           string `json:"price"`
	Qty             string `json:"qty"`
	Commission      string `json:"commission"`
	CommissionAsset string `json:"commissionAsset"`
}

type NewOrderFull struct {
	Symbol              string `json:"symbol"`
	OrderID             int64  `json:"orderId"`
	OrderListID         int64  `json:"orderListId"`
	ClientOrderID       string `json:"clientOrderId"`
	TransactTime        int64  `json:"transactTime"`
	Price               string `json:"price"`
	OrigQty             string `json:"origQty"`
	ExecutedQty         string `json:"executedQty"`
	CummulativeQuoteQty string `json:"cummulativeQuoteQty"`
	Status              string `json:"status"`
	TimeInForce         string `json:"timeInForce"`
	Type                string `json:"type"`
	Side                string `json:"side"`
	Fills               []Fill `json:"fills"`
}
//=================================================================

type CancelOrder struct {
	Symbol              string `json:"symbol"`
	OrigClientOrderID   string `json:"origClientOrderId"`
	OrderID             int64  `json:"orderId"`
	OrderListID         int64  `json:"orderListId"`
	ClientOrderID       string `json:"clientOrderId"`
	Price               string `json:"price"`
	OrigQty             string `json:"origQty"`
	ExecutedQty         string `json:"executedQty"`
	CummulativeQuoteQty string `json:"cummulativeQuoteQty"`
	Status              string `json:"status"`
	TimeInForce         string `json:"timeInForce"`
	Type                string `json:"type"`
	Side                string `json:"side"`
}

type OrderReport struct {
	Symbol              string `json:"symbol"`
	OrigClientOrderID   string `json:"origClientOrderId"`
	OrderID             int64  `json:"orderId"`
	OrderListID         int64  `json:"orderListId"`
	ClientOrderID       string `json:"clientOrderId"`
	Price               string `json:"price"`
	OrigQty             string `json:"origQty"`
	ExecutedQty         string `json:"executedQty"`
	CummulativeQuoteQty string `json:"cummulativeQuoteQty"`
	Status              string `json:"status"`
	TimeInForce         string `json:"timeInForce"`
	Type                string `json:"type"`
	Side                string `json:"side"`
	StopPrice           string `json:"stopPrice"`
	IcebergQty          string `json:"icebergQty"`
}

type PriceTicker struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
	Time   int64  `json:"time"`
}

type QueryOrder struct {
	Symbol              string `json:"symbol"`             
	OrderID             int64  `json:"orderId"`            
	OrderListID         int64  `json:"orderListId"`        
	ClientOrderID       string `json:"clientOrderId"`      
	Price               string `json:"price"`              
	OrigQty             string `json:"origQty"`            
	ExecutedQty         string `json:"executedQty"`        
	CummulativeQuoteQty string `json:"cummulativeQuoteQty"`
	Status              string `json:"status"`             
	TimeInForce         string `json:"timeInForce"`        
	Type                string `json:"type"`               
	Side                string `json:"side"`               
	StopPrice           string `json:"stopPrice"`          
	IcebergQty          string `json:"icebergQty"`         
	Time                int64  `json:"time"`               
	UpdateTime          int64  `json:"updateTime"`         
	IsWorking           bool   `json:"isWorking"`          
	OrigQuoteOrderQty   string `json:"origQuoteOrderQty"`  
}

type OpenOrder = QueryOrder
