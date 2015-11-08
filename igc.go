package igcgo

import(
    "time"
    "io"
    "encoding/json"
    "bytes"
)

type IGC struct{
    ID int `json:id'`
    Pilot Pilot `json:pi`
    Date time.Time `json:d`
    Note string `json:n`
    Config map[string]string `json:c`
    Glider Glider `json:"g"`
    Plots []Plot `json:pl`
}

func (igc *IGC)Json()string{
    bytes,_ := json.Marshal(igc)
    json := string(bytes)
    return json
}

func (igc *IGC)JsonReader()io.Reader{
    b,_ := json.Marshal(igc)
    reader := bytes.NewReader(b)
    return reader
}


