package model

type Socket struct {
	Id   int    `orm:"auto"`
	Hash string `orm:"size(16)"`
	Port uint
}

type ProjectPorts struct {
	Id        int `orm:"auto"`
	PortStart uint
	PortEnd   uint
}
