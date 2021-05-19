package cerror

import "golang-core/internal/domain"

func NewError(code int, err error) domain.Error{
	return domain.Error{
		Code: code,
		Err: err,
	}
}