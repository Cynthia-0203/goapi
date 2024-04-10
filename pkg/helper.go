package pkg

import (
	"crypto/rand"
	"fmt"
	"io"
	"time"

	"github.com/spf13/viper"
)

func MicrosecondsStr(elapsed time.Duration)string{
	return fmt.Sprintf("%.3fms",float64(elapsed.Nanoseconds())/1e6)
}
func RanddomNumber(length int)string{
	table:=[...]byte{'1','2','3','4','5','6','7','8','9','0'}
	b:=make([]byte,length)
	n,err:=io.ReadAtLeast(rand.Reader,b,length)
	if n!=length{
		panic(err)
	}
	for i:=0;i<len(b);i++{
		b[i]=table[int(b[i])%len(table)]
	}
	return string(b)
}

func TimenowInTimezone()time.Time{
	chinaTimezone,_:=time.LoadLocation(viper.GetString("app.timezone"))
	return time.Now().In(chinaTimezone)
}

// FirstElement 安全地获取 args[0]，避免 panic: runtime error: index out of range
func FirstElement(args []string) string {
    if len(args) > 0 {
        return args[0]
    }
    return ""
}