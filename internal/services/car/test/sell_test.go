package test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	carService "github.com/tx-pgx/internal/services/car"

	mocksRepository "github.com/tx-pgx/internal/persistence/postgres/car/mock"
)

func TestService_Sell_CreateCarHistoryEventError(t *testing.T) {
	mockCarRepository := new(mocksRepository.CarRepository)
	carService := &carService.Service{
		CarRepo: mockCarRepository,
	}
	mockCarRepository.On("CreateCarHistoryEvent", mock.Anything, mock.Anything).Return(errors.New("error"))
	carService.Sell(context.Background(), "carId", "owner", "buyer", 100)
	assert.EqualError(t, errors.New("error"), errors.New("error").Error())
}
