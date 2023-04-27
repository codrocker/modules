package modules

import (
	"strconv"
)

// 根据取模方式，获取表分表后的名称。table表示没有加后缀前的表名，count表示总共分成多少份，id表示用来分表的输入数字，newTable表示增加后缀之后的表名，比如record_01
func AddTableSuffix(table string, count int, id int) (newTable string){
	suffixNum := strconv.Itoa(id % count)

	prefixLen := len(strconv.Itoa(count - 1)) - len(suffixNum)
	for prefixLen > 0{
		suffixNum = "0" + suffixNum
		prefixLen --
	}

	return table + "_" + suffixNum
}