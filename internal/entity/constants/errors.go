package constants

import "errors"

const ErrNoDataFound = "data tidak ditemukan"
const ErrInvalidORExpiredJWT = "Invalid or expired JWT"
const ErrMissingOrMalformedJWT = "Missing or malformed JWT"

var ErrWrongPassword = errors.New("password yang anda masukkan salah")
var ErrUserNotFound = errors.New("user yang anda masukkan tidak ada di dalam database")
