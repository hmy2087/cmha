package check

import (
	"fmt"
//	"log"
	"os"
	"strconv"
	"strings"
	"time"
	//	"strings"
)

func IsPingType(user, password string) {
	host, port, checktime_string, timeout, defaultDb, ping_type := GetConfig()
	if ping_type == "select,replication" || ping_type == "select" {
		TrySelectCheckTime(user, password, host, port, defaultDb, checktime_string, ping_type, timeout)
	} else if ping_type == "update,replication" || ping_type == "update" {
		TryUpdateCheckTime(user, password, host, port, defaultDb, checktime_string, ping_type, timeout)
	}else{
		fmt.Println("Configuration error")
		os.Exit(2)
	}

}

func TrySelectCheckTime(user, password, host, port, defaultDb, checktime_string, ping_type, timeout string) {
	checktime, _ := strconv.Atoi(checktime_string)
	for {
		if checktime == 0{
			break
		}else{
			checktime--
			MYSQL_OK := SelectCheckMysqlHealth(user, password, host, port, defaultDb, timeout,checktime)
			//log.Println("MYSQL_OK:", MYSQL_OK)
			if MYSQL_OK == 0 {
				if strings.Contains(ping_type, "replication") {
					isyes, err := ShowSlave(user, password, host, port, defaultDb, timeout)
					if err != nil {
						//log.Println("exit code 2")
						fmt.Println(err)
						os.Exit(2)
					}
					if isyes == "Yes" {
						//log.Println("exit code 0")
						fmt.Println("check ok")
						os.Exit(0)
					} else {
						//log.Println("exit code 1")
						fmt.Println("check replication FAIL:", isyes)
						os.Exit(1)
					}
				} else {
					//log.Println("exit code 0")
					fmt.Println("check ok")
					os.Exit(0)
				}
			}
			if MYSQL_OK == 1 && checktime == 0 {
				//log.Println("exit code 2")
				//fmt.Println("check FAIL")
				os.Exit(2)
			}
		}
		time.Sleep(1 * time.Second)
	}
}

func TryUpdateCheckTime(user, password, host, port, defaultDb, checktime_string, ping_type, timeout string) {
	checktime, _ := strconv.Atoi(checktime_string)
	for {
		if checktime == 0{
			break
		}else{
			checktime--
			MYSQL_OK := CheckMysqlHealth(user, password, host, port, defaultDb, timeout,checktime)
			//log.Println("MYSQL_OK:", MYSQL_OK)
			if MYSQL_OK == 0 {
				if strings.Contains(ping_type, "replication") {
					isyes, err := ShowSlave(user, password, host, port, defaultDb, timeout)
					if err != nil {
						//log.Println("exit code 2")
						fmt.Println(err)
						os.Exit(2)
					}
					if isyes == "Yes" {
						//log.Println("exit code 0")
						fmt.Println("check ok")
						os.Exit(0)
					} else {
						//log.Println("exit code 1")
						fmt.Println("check replication FAIL:", isyes)
						os.Exit(1)
					}
				} else {
					//log.Println("exit code 0")
					fmt.Println("check ok")
					os.Exit(0)
				}
			}
			if MYSQL_OK == 1 && checktime == 0 {
				//log.Println("exit code 2")
				//log.Println("mysql check time:", checktime_string)
				os.Exit(2)
			}
		}
		time.Sleep(1 * time.Second)
	}
}
