package igcgo


import (
    "os"
    "bufio"
    "strings"
    "strconv"
    "errors"
    "time"
    "io"
)

func ParseIGC(fp *os.File)(*IGC,error){
    igc,err := ParseIGCReader(fp)
    return igc,err
}

func ParseIGCReader(r io.Reader)(*IGC,error){
    scanner := bufio.NewScanner(r)
    igc := &IGC{
        Config : make(map[string]string),
        Plots : make([]Plot,0),
    }
    var befplot Plot
    for scanner.Scan(){
        line := scanner.Text()
        line = strings.ToLower(line)
        hasp := strings.HasPrefix(line,"b")
        if hasp{
            plot := parse_record(line,befplot)
            igc.Plots = append(igc.Plots,plot)
            befplot = plot
        }else{
        }
    }
    if len(igc.Plots) == 0{
        return igc,errors.New("Faild to parse")
    }
    return igc,nil
}

func parse_record(line string,befplot Plot)Plot{
    times := line[1:7]
    lats := line[7:15]
    longs := line[15:24]
    //has_alt := line[24]
    pressure_alts := line[25:30]
    gps_alts := line[30:35]

    plot := Plot{
        Time : times_time(times),
        Latitude : lat_float(lats),
        Longitude : long_float(longs),
        PressureAltitude:alt_int(pressure_alts),
        GPSAltitude:alt_int(gps_alts),
    }
    plot.VerticalVelocity = vertical_velocity(plot,befplot)
    return plot
}

func lat_float(lat string)float64{
    degree,_ := strconv.ParseFloat(lat[0:2],64)
    minutes,_ := strconv.ParseFloat(lat[2:7],64)
    decimal := degree+minutes/60000.
    if strings.HasSuffix(lat,"n"){
        return decimal
    }else{
        return -1*decimal
    }
}

func long_float(long string)float64{
    degree,_ := strconv.ParseFloat(long[0:3],64)
    minutes,_ := strconv.ParseFloat(long[3:8],64)
    decimal := degree+minutes/60000.
    if strings.HasSuffix(long,"e"){
        return decimal
    }else{
        return -1*decimal
    }
}

func alt_int(alt string)int{
    i,_ := strconv.ParseInt(alt,10,0)
    return int(i)
}

var time_format = "150405"
func times_time(times string)time.Time{
    t,err := time.Parse(time_format,times)
    if err != nil{
    }
    return t
}

func vertical_velocity(plot1,plot2 Plot)float64{
    if &plot2 == nil{
        return 0
    }else{
        duration := float64(plot1.Time.Sub(plot2.Time))/float64(1000000000)
        altsub := float64(plot1.GPSAltitude - plot2.GPSAltitude)
        v := altsub/duration
        return v
    }
}


