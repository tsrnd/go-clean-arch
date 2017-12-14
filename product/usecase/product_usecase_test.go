package usecase_test

import (
	"strconv"
	"testing"

	"github/monstar-lab/fr-circle-api/models"
	"github/monstar-lab/fr-circle-api/usecase"

	"github.com/bxcodec/faker"
	"github.com/tsrnd/go-clean-arch/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFetch(t *testing.T) {
	mockProductRepo := new(mocks.ProductRepository)
	var mockProduct models.Product
	err := faker.FakeData(&mockProduct)
	assert.NoError(t, err)

	mockListProduct := make([]*models.Product, 0)
	mockListProduct = append(mockListProduct, &mockProduct)
	mockProductRepo.On("Fetch", mock.AnythingOfType("string"), mock.AnythingOfType("int64")).Return(mockListProduct, nil)
	uc := usecase.NewProductUsecase(mockProductRepo)
	limit := int64(1)
	offset := "12"
	list, err := uc.Fetch(offset, limit)
	offsetExpected := strconv.Itoa(int(mockProduct.ID))
	assert.NoError(t, err)
	assert.Len(t, list, len(mockListProduct))

	mockProductRepo.AssertCalled(t, "Fetch", mock.AnythingOfType("string"), mock.AnythingOfType("int64"))
}
