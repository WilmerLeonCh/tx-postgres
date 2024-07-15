package car

import (
	"github.com/gin-gonic/gin"

	"github.com/tx-pgx/internal/domain"
)

func (h *Handler) Register(ginCtx *gin.Context) {
	var req domain.Car
	if err := ginCtx.ShouldBindJSON(&req); err != nil {
		ginCtx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := h.validateRegisterCar(req); err != nil {
		ginCtx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := h.CarService.SaveCar(ginCtx, &req); err != nil {
		ginCtx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ginCtx.JSON(200, gin.H{"message": "Car registered successfully"})
}

func (h *Handler) validateRegisterCar(req domain.Car) error {
	if req.Name == "" {
		return domain.ErrMissingName
	}
	if req.Version == "" {
		return domain.ErrMissingVersion
	}
	if req.Owner == "" {
		return domain.ErrMissingOwner
	}
	if req.ManufactureDate == 0 {
		return domain.ErrMissingManufactureDate
	}
	return nil
}
