package threat

import "errors"

var (
	// ErrNotFound возвращается, когда угроза с указанным MD5 не найдена.
	ErrNotFound = errors.New("threat not found")
	// ErrConflict возвращается, когда угроза с таким MD5 уже существует.
	ErrConflict = errors.New("threat already exists")
	// ErrInvalidMD5 возвращается при неверном формате MD5-хеша.
	ErrInvalidMD5 = errors.New("invalid md5 format")
	// ErrInvalidFields возвращается, когда поля запроса не прошли валидацию.
	ErrInvalidFields = errors.New("invalid fileds")
)
