package sevice

import (
	"IoTGateWay/model"
	"fmt"
)

var (
	DevStatusSer *model.DeviceStatusService
	DetResultSer *model.DetectResultService
)
func Init() {
	DevStatusSer = model.GetDeviceStatusHandler()
	DetResultSer = model.GetDetectResultHandler()
	fmt.Println("service init success")
}
func Log(v ...interface{}) {
	log.Println(v...)
}


func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
	}
}
