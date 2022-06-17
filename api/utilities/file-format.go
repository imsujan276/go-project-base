package utilities

import (
	"path"
	"strings"

	"github.com/twinj/uuid"
)

func UniqueFileName(fn string) string {
	//path.Ext() get the extension of the file
	fileName := strings.TrimSuffix(fn, path.Ext(fn))
	extension := path.Ext(fn)
	u := uuid.NewV4()
	newFileName := fileName + "-" + u.String() + extension

	return newFileName

}

func CustomFileName(fn, newFn string) string {
	//path.Ext() get the extension of the file
	extension := path.Ext(fn)
	newFileName := newFn + extension

	return newFileName

}
