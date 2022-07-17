package constants

import "errors"

const ErrNoDataFound = "data tidak ditemukan"

var ErrWrongPassword = errors.New("password yang anda masukkan salah")
var ErrUserNotFound = errors.New("user yang anda masukkan tidak ada di dalam database")
