package delivery

import (
	"net/http/httptest"
	"testing"

	"github.com/monitoror/monitoror/internal/pkg/monitorable/params"
	"github.com/monitoror/monitoror/internal/pkg/validator"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type Params1 struct {
	params.Default

	Field string `query:"field" validate:"required"`
}

type Params2 struct {
	Field string `query:"field" validate:"required"`
}

func (p *Params2) Validate() []validator.Error {
	return nil
}

type Params3 struct {
	Field string `query:"field" validate:"required"`
}

func (p *Params3) Validate() []validator.Error {
	return []validator.Error{validator.NewDefaultError("Field", "", "boom")}
}

func TestBindAndValidateParams(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/api/v1/xxx", nil)
	res := httptest.NewRecorder()
	ctx := e.NewContext(req, res)

	p := &Params1{}
	err := BindAndValidateParams(ctx, p)
	assert.Error(t, err)
	assert.Equal(t, `Required "field" field is missing.`, err.Error())

	ctx.QueryParams().Add("field", "test")

	p2 := &Params2{}
	err = BindAndValidateParams(ctx, p2)
	assert.NoError(t, err)

	p3 := &Params3{}
	err = BindAndValidateParams(ctx, p3)
	assert.Error(t, err)
	assert.Equal(t, `Invalid "field" field. Must be boom.`, err.Error())

}
