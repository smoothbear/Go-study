package model

import "database/sql/driver"

type uuid string

func UUID(s string) uuid                         { return uuid(s) }
func (u uuid) Value() (driver.Value, error)      { return string(u), nil }
func (u *uuid) Scan(src interface{}) (err error) { *u = uuid(src.([]uint8)); return }
func (u uuid) KeyName() string                   { return "uuid" }