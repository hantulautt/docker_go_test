package entity

import (
	"docker_go_test/app/exception"
	"docker_go_test/app/model"
	"fmt"
	"net"
)

type ApnicInetnum struct {
	Inetnum string `gorm:"size:50;primaryKey;autoIncrement:false"`
	Name    string `gorm:"size:100"`
	Descr   string `gorm:"type:text;"`
}

func (ApnicInetnum) TableName() string {
	return "apnic_inetnums"
}

func (data ApnicInetnum) ApnicResponse() model.ApnicSourceResponse {
	// Parse the CIDR string
	_, ipNet, err := net.ParseCIDR(data.Inetnum)
	if err != nil {
		panic(exception.ValidationError{
			Message: "Sorry something went wrong. " + err.Error(),
		})
	}

	// Get the network and broadcast IP addresses
	networkIP := ipNet.IP
	broadcastIP := make(net.IP, len(networkIP))
	copy(broadcastIP, networkIP)
	for i := range broadcastIP {
		broadcastIP[i] |= ^ipNet.Mask[i]
	}

	// Format the range string
	ipRange := fmt.Sprintf("%s - %s", networkIP.String(), broadcastIP.String())

	return model.ApnicSourceResponse{
		IpRange:     ipRange,
		Name:        data.Name,
		Description: data.Descr,
	}
}
