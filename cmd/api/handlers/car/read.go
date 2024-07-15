package car

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ReadCar(ginCtx *gin.Context) {
	id := ginCtx.Param("id")
	car, errReadCar := h.CarService.ReadCar(ginCtx, id)
	if errReadCar != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": errReadCar.Error()})
		return
	}

	ginCtx.JSON(http.StatusOK, gin.H{"car": car})
}
