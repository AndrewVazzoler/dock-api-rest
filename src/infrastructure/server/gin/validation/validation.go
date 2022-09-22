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
	validate_doc "github.com/paemuri/brdoc/v2"
)

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ErrorRoot struct {
	Error interface{} `json:"errors"`
}

var (
	uni   *ut.UniversalTranslator
	trans ut.Translator
)

func Init(local string) (err error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {

		v.RegisterValidation(`isCpf`, ValidateCPF)

		pt_BR := brazilian_portuguese.New()
		eng := english.New()
		uni = ut.New(eng, eng, pt_BR)

		local = strings.ToLower(local)

		if local == "" || local != "pt_br" {
			local = "en"
		}

		var o bool
		trans, o = uni.GetTranslator(local)

		if !o {
			return fmt.Errorf("uni.GetTranslator(%s) failed", local)
		}

		v.RegisterTranslation("isCpf", trans, func(ut ut.Translator) error {
			if ut.Locale() == "pt_BR" {
				return ut.Add("isCpf", "{0} deve ser um cpf valido!", true)
			}
			return ut.Add("isCpf", "{0} must be a valid cpf!", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("isCpf", fe.Field())
			return t
		})

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

func ValidateCPF(fl validator.FieldLevel) bool {
	return validate_doc.IsCPF(fl.Field().String())
}
