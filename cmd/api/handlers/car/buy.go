package car

import (
	"context"

	"github.com/gin-gonic/gin"

	"github.com/tx-pgx/internal/domain"
)

func (h *Handler) BuyCar(ginCtx *gin.Context) {
	var req domain.BuyCarRequest
	if err := ginCtx.ShouldBindJSON(&req); err != nil {
		ginCtx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.validateBuyer(ginCtx, req.Buyer, req.Price); err != nil {
		ginCtx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// read owner domain
	car, err := h.CarService.ReadCar(ginCtx, req.CarID)
	if err != nil {
		ginCtx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if errTx := h.CarService.WithinTransaction(ginCtx, func(txCtx context.Context) error {
		return h.CarService.Sell(txCtx, car.ID, car.Owner, req.Buyer, req.Price)
	}); errTx != nil {
		ginCtx.JSON(400, gin.H{"error": errTx.Error()})
		return
	}
	ginCtx.JSON(200, gin.H{"message": "Car bought successfully"})
}

func (h *Handler) validateBuyer(ctx context.Context, id string, price int) error {
	return nil
}
