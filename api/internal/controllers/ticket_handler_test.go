package controllers

import (
    "api/internal/domain"
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
)

type MockTicketGetAllUseCase struct {
    mock.Mock
}

func (m *MockTicketGetAllUseCase) Execute() ([]domain.Ticket, error) {
    args := m.Called()
    return args.Get(0).([]domain.Ticket), args.Error(1)
}

type MockTicketCreateUseCase struct {
    mock.Mock
}

func (m *MockTicketCreateUseCase) Execute(req domain.CreateTicketRequest) (domain.Ticket, error) {
    args := m.Called(req)
    return args.Get(0).(domain.Ticket), args.Error(1)
}

type MockTicketGetUseCase struct {
    mock.Mock
}

func (m *MockTicketGetUseCase) Execute(ticketID int) (domain.Ticket, error) {
    args := m.Called(ticketID)
    return args.Get(0).(domain.Ticket), args.Error(1)
}

type MockTicketPurchaseUseCase struct {
    mock.Mock
}

func (m *MockTicketPurchaseUseCase) Execute(ticketID int, req domain.PurchaseRequest) error {
    args := m.Called(ticketID, req)
    return args.Error(0)
}

func setupRouter(getAllUseCase *MockTicketGetAllUseCase,
    createUseCase *MockTicketCreateUseCase,
    getUseCase *MockTicketGetUseCase,
    purchaseUseCase *MockTicketPurchaseUseCase) *gin.Engine {

    r := gin.Default()

    if getAllUseCase == nil {
        getAllUseCase = new(MockTicketGetAllUseCase)
    }
    if createUseCase == nil {
        createUseCase = new(MockTicketCreateUseCase)
    }
    if getUseCase == nil {
        getUseCase = new(MockTicketGetUseCase)
    }
    if purchaseUseCase == nil {
        purchaseUseCase = new(MockTicketPurchaseUseCase)
    }

    handler := NewTicketHandler(getAllUseCase, createUseCase, getUseCase, purchaseUseCase)

    r.GET("/api/v1/tickets", handler.GetAllTickets)
    r.GET("/api/v1/tickets/:id", handler.GetTicket)
    r.POST("/api/v1/tickets", handler.CreateTicket)
    r.POST("/api/v1/tickets/:id/purchases", handler.PurchaseTicket)

    return r
}

func TestGetAllTickets(t *testing.T) {
    mockGetAll := new(MockTicketGetAllUseCase)

    tickets := []domain.Ticket{
        {ID: 1, Name: "Concert A", Desc: "Desc A", Allocation: 100},
        {ID: 2, Name: "Concert B", Desc: "Desc B", Allocation: 100},
    }

    mockGetAll.On("Execute").Return(tickets, nil)

    router := setupRouter(mockGetAll, nil, nil, nil)

    req, _ := http.NewRequest("GET", "/api/v1/tickets", nil)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
    assert.Contains(t, w.Body.String(), "Concert A")
    assert.Contains(t, w.Body.String(), "Concert B")
    mockGetAll.AssertExpectations(t)
}

func TestGetTicket(t *testing.T) {
    mockGet := new(MockTicketGetUseCase)

    ticket := domain.Ticket{
        ID: 1, Name: "Concert A", Desc: "Desc A", Allocation: 100,
    }

    mockGet.On("Execute", 1).Return(ticket, nil)

    router := setupRouter(nil, nil, mockGet, nil)

    req, _ := http.NewRequest("GET", "/api/v1/tickets/1", nil)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
    assert.Contains(t, w.Body.String(), "Concert A")
    mockGet.AssertExpectations(t)
}

func TestCreateTicket(t *testing.T) {
    mockCreate := new(MockTicketCreateUseCase)

    createReq := domain.CreateTicketRequest{Name: "Concert A", Desc: "Desc A", Allocation: 100}
    ticket := domain.Ticket{ID: 1, Name: "Concert A", Desc: "Desc A", Allocation: 100}

    mockCreate.On("Execute", createReq).Return(ticket, nil)

    router := setupRouter(nil, mockCreate, nil, nil)

    body, _ := json.Marshal(createReq)
    req, _ := http.NewRequest("POST", "/api/v1/tickets", bytes.NewBuffer(body))
    req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusCreated, w.Code)
    assert.Contains(t, w.Body.String(), "Concert A")
    mockCreate.AssertExpectations(t)
}

func TestPurchaseTicket(t *testing.T) {
    mockPurchase := new(MockTicketPurchaseUseCase)

    purchaseReq := domain.PurchaseRequest{Quantity: 1}

    mockPurchase.On("Execute", 1, purchaseReq).Return(nil)

    router := setupRouter(nil, nil, nil, mockPurchase)

    body, _ := json.Marshal(purchaseReq)
    req, _ := http.NewRequest("POST", "/api/v1/tickets/1/purchases", bytes.NewBuffer(body))
    req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusNoContent, w.Code)
    mockPurchase.AssertExpectations(t)
}
