package http_test

func TestGetByID(t *testing.T) {
  var mockProduct models.Product
  err := faker.FakeData(&mockProduct)
  assert.NoError(t, err)

  mockUCase := new(mocks.ProductUsecase)
  mockID := int(mockProduct.ID)
  mockUCase.On("GetByID", int64(mockID)).Return(&mockProduct, nil)
  e := echo.New()
  req, err := http.NewRequest(echo.GET, "/product/" +
              strconv.Itoa(int(mockID)), strings.NewReader(""))
  assert.NoError(t, err)

  rec := httptest.NewRecorder()
  c := e.NewContext(req, rec)
  c.SetPath("product/:id")
  c.SetParamNames("id")
  c.SetParamValues(strconv.Itoa(mockID))

  handler:= productHttp.ProductHandler{
            AUsecase: mockUCase,
            Helper: httpHelper.HttpHelper{}
  }
  handler.GetByID(c)

  assert.Equal(t, http.StatusOK, rec.Code)
  mockUCase.AssertCalled(t, "GetByID", int64(mockID))
}
