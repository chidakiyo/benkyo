package lib_a

import "time"

func TimeString() string  {
	return time.Now().Format("2006-01-02")
}
