package utils

import (
	"net/http"

	"backend/business"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/labstack/echo/v4"
)

// use a single instance , it caches struct info
var (
	uni *ut.UniversalTranslator
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if uni == nil {
		en := en.New()
		uni = ut.New(en, en)
	}
	trans, _ := uni.GetTranslator("en")
	en_translations.RegisterDefaultTranslations(cv.Validator, trans)

	err := cv.Validator.Struct(i)
	if err != nil {
		message := translateIndividual(trans, err)
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"status_code": business.BadRequest,
			"message":     message,
			"data":        map[string]interface{}{},
		})
	}

	return nil
}

func translateIndividual(trans ut.Translator, err error) string {

	errs := err.(validator.ValidationErrors)

	for _, e := range errs {
		return e.Translate(trans)
	}

	return "validation error"
}
