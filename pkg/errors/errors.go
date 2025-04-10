package errors

import (
	"errors"
	"syscall"
)

const kernelVersionErrCode = 22

var (
	ErrFailToLoadConfig        = System(nil, "fail to load config", "MES:001")
	ErrFailEnsureConfig        = System(nil, "fail to ensure config", "MES:002")
	ConfigFileWasCreated       = Business("a new config file was created (%s)", "MES:003")
	ErrFailToMarshalConfig     = System(nil, "fail to marshal config", "MES:004")
	ErrMissingDatasourceConfig = Business("datasource must be defined", "MES:005")
	ErrDeliveryNotInitialized  = errors.New("delivery not initialized")
	ErrKernelVersion           = syscall.Errno(kernelVersionErrCode)
)
