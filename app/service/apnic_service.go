package service

import (
	"bufio"
	"docker_go_test/app/entity"
	"docker_go_test/app/exception"
	"docker_go_test/app/helper"
	"docker_go_test/app/repository"
	"fmt"
	"net"
	"os"
	"strings"
)

type ApnicService interface {
	InsertData()
	WhoisIp(inetNum string) interface{}
}

type ApnicServiceImpl struct {
	ApnicRepository repository.ApnicRepository
}

func (service ApnicServiceImpl) WhoisIp(inetNum string) interface{} {
	data := service.ApnicRepository.WhoisIp(inetNum)
	if (data == entity.ApnicInetnum{}) {
		panic(exception.ValidationError{
			Message: "Sorry the data you are looking for is not available.",
		})
	}
	return data.ApnicResponse()
}

func (service ApnicServiceImpl) InsertData() {
	fmt.Println("**** CRON IS RUNNING ****")
	/**
	Open file of data source
	*/
	dataSource, err := os.Open("apnic.db.inetnum")
	if err != nil {
		exception.PanicIfNeeded(err)
	}
	defer dataSource.Close()

	/**
	Read data from source
	*/
	scanner := bufio.NewScanner(dataSource)
	if err = scanner.Err(); err != nil {
		exception.PanicIfNeeded(err)
	}

	var apnicSource entity.ApnicInetnum
	groups := make(map[string]entity.ApnicInetnum)
	var descTemp []string
	var inetNum string
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")

		if len(parts) > 1 {
			property := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])

			if strings.ToLower(property) == "inetnum" {
				apnicSource = entity.ApnicInetnum{}

				/**
				Parse the start and end IP addresses
				*/
				partOfValue := strings.Split(value, "-")
				ipStart := net.ParseIP(strings.TrimSpace(partOfValue[0]))
				ipEnd := net.ParseIP(strings.TrimSpace(partOfValue[1]))
				if ipStart == nil || ipEnd == nil {
					helper.WriteLog("insert_data.log", "ERROR [PARSE IP] Inet Num = "+value+" - "+err.Error())
					apnicSource.Inetnum = value
				} else {
					/**
					Get the network IP and subnet mask length
					*/
					ip := ipStart.Mask(ipStart.DefaultMask())
					maskLen, _ := ip.DefaultMask().Size()

					/**
					Format the CIDR string
					*/
					inetNum = fmt.Sprintf("%s/%d", ip.String(), maskLen)
					apnicSource.Inetnum = inetNum
				}
			} else if strings.ToLower(property) == "netname" {
				/**
				Get net name
				*/
				apnicSource.Name = value
			} else if strings.ToLower(property) == "descr" {
				/**
				Get description
				*/
				if value != "" {
					descTemp = append(descTemp, value)
				}
			} else {
				if len(descTemp) > 0 {
					apnicSource.Descr = strings.Join(descTemp, ", ")
					groups[inetNum] = apnicSource
					descTemp = nil

					/**
					Insert or update data into database
					*/
					if err = service.ApnicRepository.InsertInetNum(apnicSource); err != nil {
						helper.WriteLog("insert_data.log", "ERROR [INSERT DB] Inet Num = "+apnicSource.Inetnum)
						helper.WriteLog("insert_data.log", "ERROR [INSERT DB] Param = "+apnicSource.Inetnum)
						helper.WriteLog("insert_data.log", "ERROR [INSERT DB] Message = "+err.Error())
						exception.PanicIfNeeded(err)
					}
					fmt.Println("INSERT DATA INTO DB IS RUNNING")
				}
			}
		}
	}
	fmt.Println("*** COMPLETED ***")
	helper.WriteLog("data_apnic.log", helper.PrettyPrint(groups))
}

func NewApnicService(repository *repository.ApnicRepository) ApnicService {
	return ApnicServiceImpl{
		ApnicRepository: *repository,
	}
}
