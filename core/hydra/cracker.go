package hydra

import (
	"Qscan/core/hydra/ftp"
	"Qscan/core/hydra/mongodb"
	"Qscan/core/hydra/mssql"
	"Qscan/core/hydra/mysql"
	"Qscan/core/hydra/oracle"
	"Qscan/core/hydra/postgresql"
	"Qscan/core/hydra/rdp"
	"Qscan/core/hydra/redis"
	"Qscan/core/hydra/smb"
	"Qscan/core/hydra/ssh"
	"Qscan/core/hydra/telnet"
	"Qscan/core/hydra/web"
	"Qscan/lib/gotelnet"
	"Qscan/lib/grdp"
	"fmt"
)

func rdpCracker(IPAddr string, port int) func(interface{}) error {
	target := fmt.Sprintf("%s:%d", IPAddr, port)
	protocol := grdp.VerifyProtocol(target)
	//slog.Println(slog.DEBUG, "rdp protocol is :", protocol)
	return func(i interface{}) error {
		info := i.(AuthInfo)
		domain := ""
		return rdp.Check(info.IPAddr, domain, info.Auth.Username, info.Auth.Password, info.Port, protocol)
	}
}

func smbCracker(i interface{}) error {
	info := i.(AuthInfo)
	domain := ""
	return smb.Check(info.IPAddr, domain, info.Auth.Username, info.Auth.Password, info.Port)
}

func webCracker(i interface{}) error {
	info := i.(AuthInfo)
	domain := ""
	return web.Check(info.IPAddr, domain, info.Auth.Username, info.Auth.Password, info.Port)
}

func sshCracker(i interface{}) error {
	info := i.(AuthInfo)
	return ssh.Check(info.IPAddr, info.Auth.Username, info.Auth.Password, info.Port)
}

func telnetCracker(serverType int) func(interface{}) error {
	return func(i interface{}) error {
		info := i.(AuthInfo)
		return telnet.Check(info.IPAddr, info.Auth.Username, info.Auth.Password, info.Port, serverType)
	}
}

func getTelnetServerType(ip string, port int) int {
	client := gotelnet.New(ip, port)
	err := client.Connect()
	if err != nil {
		return gotelnet.Closed
	}
	defer client.Close()
	return client.MakeServerType()
}

func mysqlCracker(i interface{}) error {
	info := i.(AuthInfo)
	return mysql.Check(info.IPAddr, info.Auth.Username, info.Auth.Password, info.Port)
}

func mssqlCracker(i interface{}) error {
	info := i.(AuthInfo)
	return mssql.Check(info.IPAddr, info.Auth.Username, info.Auth.Password, info.Port)
}

func redisCracker(i interface{}) error {
	info := i.(AuthInfo)
	return redis.Check(info.IPAddr, info.Auth.Password, info.Port)
}

func ftpCracker(i interface{}) error {
	info := i.(AuthInfo)
	return ftp.Check(info.IPAddr, info.Auth.Username, info.Auth.Password, info.Port)
}

func postgresqlCracker(i interface{}) error {
	info := i.(AuthInfo)
	return postgresql.Check(info.IPAddr, info.Auth.Username, info.Auth.Password, info.Port)
}

func oracleCracker(sid string) func(interface{}) error {
	return func(i interface{}) error {
		info := i.(AuthInfo)
		info.Auth.Other["SID"] = sid
		return oracle.Check(info.IPAddr, info.Auth.Username, info.Auth.Password, info.Port, sid)
	}
}

func mongodbCracker(i interface{}) error {
	info := i.(AuthInfo)
	return mongodb.Check(info.IPAddr, info.Auth.Username, info.Auth.Password, info.Port)
}
