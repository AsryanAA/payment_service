package models

type Connect struct {
	HostDB   string `json:"host_db"`
	PortDB   string `json:"port_db"`
	User     string `json:"user"`
	Password string `json:"password"`
	DataBase string `json:"data_base"`
}
