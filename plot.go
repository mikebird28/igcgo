package igcgo

import(
    "fmt"
    "time"
)

type Plot struct{
    Time time.Time `json:"t"`
    Latitude float64 `json:"la"`
    Longitude float64 `json:"lo"`
    PressureAltitude int `json:"p"`
    GPSAltitude int `json:"g"`
    VerticalVelocity float64 `json:"v"`
}

func (p *Plot)String() string{
    str := fmt.Sprintf("{Lat:%f,Long:%f,Alt:%d,Valio:%f},",p.Latitude,p.Longitude,p.GPSAltitude,p.VerticalVelocity)
    return str
}
