package delivery

import (
	"github.com/labstack/echo/v4"

	uiConfigModels "github.com/monitoror/monitoror/api/config/models"
	coreModels "github.com/monitoror/monitoror/models"
)

func BindAndValidateRequestParams(ctx echo.Context, v uiConfigModels.ParamsValidator) error {
	if err := ctx.Bind(v); err != nil {
		return coreModels.ParamsError
	}

	if errors := v.Validate(); len(errors) > 0 {
		return &coreModels.MonitororError{Message: errors[0].Error()}
	}

	return nil
}
