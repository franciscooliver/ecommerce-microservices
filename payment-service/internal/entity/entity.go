package entity

import "errors"

type OrderRequest struct {
	OrderId  string  `json:"order_id"`
	CardHash string  `json:"card_hash"`
	Total    float64 `json:"total"`
}

func NewOrderRequest(orderID, cardHash string, total float64) *OrderRequest {
	return &OrderRequest{
		OrderId:  orderID,
		CardHash: cardHash,
		Total:    total,
	}
}

func (o *OrderRequest) Validate() error {
	if o.OrderId == "" {
		return errors.New("order_id is required")
	}

	if o.CardHash == "" {
		return errors.New("card_hash is required")
	}

	if o.Total <= 0 {
		return errors.New("total must be greater than 0")
	}

	return nil
}

func (o *OrderRequest) Process() (*OrderResponse, error) {
	if err := o.Validate(); err != nil {
		return nil, err
	}

	orderResponse := NewOrderResponse(o.OrderId, "failed")
	if o.Total < 100.00 {
		orderResponse.Status = "paid"
	}
	return orderResponse, nil
}

type OrderResponse struct {
	OrderId string `json:"order_id"`
	Status  string `json:"status"`
}

func NewOrderResponse(orderID, status string) *OrderResponse {
	return &OrderResponse{
		OrderId: orderID,
		Status:  status,
	}
}
