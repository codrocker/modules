package modules

import (
    "github.com/mojocn/base64Captcha"
)


func GetCaptchaCode() (captchaId string, base64Png string) {
    var configD = base64Captcha.ConfigDigit{
        Height:     80,
        Width:      240,
        MaxSkew:    0.7,
        DotCount:   80,
        CaptchaLen: 5,
    }

    captchaId, digitCap := base64Captcha.GenerateCaptcha("", configD)
    base64Png = base64Captcha.CaptchaWriteToBase64Encoding(digitCap)

    return
}

func IsCaptchaCodeValid (captchaId string, captcha_code string) (isValid bool){
    return base64Captcha.VerifyCaptcha(captchaId, captcha_code)
}
