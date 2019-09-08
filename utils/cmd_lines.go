package utils
import (
	"flag"
)
func ReadConf() string {
		var conf_file string
		flag.StringVar(&conf_file, "conf_file", '/etc/log-archiver.json', "Conf file to be parsed")
		flag.StringVar(&conf_file, "c", '/etc/log-archiver.json', "Conf file to be parsed")
		flag.Parse()
	return *conf_file
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
