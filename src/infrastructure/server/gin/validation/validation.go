package validation

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin/binding"
	english "github.com/go-playground/locales/en"
	brazilian_portuguese "github.com/go-playground/locales/pt_BR"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	pt_translations "github.com/go-playground/validator/v10/translations/pt_BR"
)

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ErrorRoot struct {
	Error interface{}
}

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
)

func Init(local string) (err error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		pt_BR := brazilian_portuguese.New()
		eng := english.New()
		uni := ut.New(eng, eng, pt_BR)

		local = strings.ToLower(local)

		if local == "" || local != "pt_br" {
			local = "en"
		}

		var o bool
		trans, o = uni.GetTranslator(local)

		if !o {
			return fmt.Errorf("uni.GetTranslator(%s) failed", local)
		}

		switch local {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		case "pt_br":
			err = pt_translations.RegisterDefaultTranslations(v, trans)
		default:
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		}
		return nil
	}
	return nil
}

func Translate(err error) []ErrorMsg {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]ErrorMsg, len(ve))
		for i, fe := range ve {
			out[i] = ErrorMsg{fe.Field(), fe.Translate(trans)}
		}
		return out
	}

	return nil
}
