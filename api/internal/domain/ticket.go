package domain


type Ticket struct {
	ID         uint
	Name       string
	Desc       string
	Allocation int
}

type TicketRequest struct {
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	Allocation int    `json:"allocation"`
}

type TicketResponse struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	Allocation int    `json:"allocation"`
}

func (t *Ticket) ToTicketResponse() TicketResponse {
    return TicketResponse{
        ID:         t.ID,
        Name:       t.Name,
        Desc:       t.Desc,
        Allocation: t.Allocation,
    }
}

type PurchaseRequest struct {
	Quantity int    `json:"quantity"`
	UserID   string `json:"user_id"`
}

type JSONResponse struct {
    Message string `json:"message"`
}