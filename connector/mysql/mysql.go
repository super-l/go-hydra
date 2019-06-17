package mysql

type MysqlProtocol struct {
	dst      string
}

func Create(address, port string) *MysqlProtocol {
	var dst string
	if port == "0" {
		dst = address + ":3306"
	} else {
		dst = address + ":" + port
	}
	return &MysqlProtocol{dst}
}