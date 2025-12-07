package entity

import "fmt"

var ERRPermission error = fmt.Errorf("err permision")
var ERRUnusual error = fmt.Errorf("err unusual")
var ERRBusy1 error = fmt.Errorf("err busy1")
var ERRBusy2 error = fmt.Errorf("err busy2")
var ERRForever error = fmt.Errorf("err forever")
