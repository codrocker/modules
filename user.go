package modules

type UserInfoSql struct {
    Uid      int64  `gorm:"uid"`
    ClassId  int    `gorm:"class_id"`
    Nick     string `gorm:"nick"`
    AppNick     string `gorm:"app_nick"`
    Portrait string `gorm:"portrait"`
    Gender int `gorm:"gender"`
    Birth int `gorm:"birth"`
    Phone string `gorm:"phone"`
    VcLevel  int    `gorm:"vc_level"`
    PkLevel  int    `gorm:"pk_level"`
    PkCoins  int    `gorm:"pk_coins"`
    PkWin    int    `gorm:"pk_win"`
    PkLose   int    `gorm:"pk_lose"`
    PkStars  int    `gorm:"pk_stars"`
}

type UserInfo struct{
    Uid int64 `json:"uid"`
    ClassId  int    `json:"class_id"`
    Nick string `json:"nick"`
    Portrait string `json:"portrait"`
    Gender int `json:"gender"`
    Birth int `json:"birth"`
    Phone string `json:"phone"`
    VcLevel int `json:"vc_level"`
    PkLevel int `json:"pk_level"`
    PkCoins int `json:"pk_coins"`
    PkWin int `json:"pk_win"`
    PkLose int `json:"pk_lose"`
    PkStars  int    `json:"pk_stars"`
    VcLevelTitle string `json:"vc_level_title"`
    PkLevelTitle string `json:"pk_level_title"`
}
