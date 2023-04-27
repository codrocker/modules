package modules

import (
	"errors"
	"strconv"
	"strings"
)

const (
    WinTag = "win"
    AndroidTag = "android"
    IphoneTag = "iphone"
)

type Cv struct {
    Num int
    Os string
}

func GetCv(cv string) (data Cv, err error){
    str := strings.Replace(cv, WinTag + "_", "", -1)
    str = strings.Replace(str, AndroidTag + "_", "", -1)
    str = strings.Replace(str, IphoneTag + "_", "", -1)
    str = strings.Replace(str, ".", "", -1)


    num, err := strconv.Atoi(str)
    if err != nil {
        return
    }

    data.Num = num

    if strings.Contains(cv, WinTag) == true{
        data.Os = WinTag
        return
    }
    if strings.Contains(cv, AndroidTag) == true{
        data.Os = AndroidTag
        return
    }

    if strings.Contains(cv, IphoneTag) == true{
        data.Os = IphoneTag
        return
    }

    err = errors.New("query params cv err")
    return
}
