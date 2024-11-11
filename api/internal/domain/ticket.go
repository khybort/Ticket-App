package domain

type Ticket struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	Allocation int    `json:"allocation"`
}

type CreateTicketRequest struct {
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	Allocation int    `json:"allocation"`
}

type PurchaseRequest struct {
	Quantity int    `json:"quantity"`
	UserID   string `json:"user_id"`
}

type JSONResponse struct {
    Message string `json:"message"`
}