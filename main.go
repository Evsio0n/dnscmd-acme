package main

import (
	"github.com/evsio0n/log"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
)

type DnsAddReqParams struct {
	Zone       string `json:"zone"`
	Hostname   string `json:"hostname"`
	Value      string `json:"value"`
	RecordType string `json:"record_type"`
}

type DnsDelReqParams struct {
	Zone       string `json:"zone"`
	Hostname   string `json:"hostname"`
	RecordType string `json:"record_type"`
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := runConfig()
	//disable gin output
	//run on internal port 8080
	r.Run("172.17.6.16:8880")

}

func runConfig() *gin.Engine {
	r := gin.New()
	r.POST("/set", handleSet)
	r.POST("/del", handleDel)
	return r
}

func handleSet(c *gin.Context) {
	//get json data
	var dnsResponse DnsAddReqParams
	err := c.ShouldBindJSON(&dnsResponse)
	if err == nil {
		//set dns record
		err = SetDnsRecord(dnsResponse.Zone, dnsResponse.Hostname, dnsResponse.Value, dnsResponse.RecordType)
		if err == nil {
			c.JSON(200, gin.H{
				"message": "success",
			})
			return
		}
	}
	log.Info(err.Error())
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": "fail",
		"err":     err.Error(),
	})
}

func handleDel(c *gin.Context) {
	//get json data
	var dnsResponse DnsDelReqParams
	err := c.ShouldBindJSON(&dnsResponse)
	if err == nil {
		//set dns record
		err = DelDnsRecord(dnsResponse.Zone, dnsResponse.Hostname, dnsResponse.RecordType)
		if err == nil {
			c.JSON(200, gin.H{
				"message": "success",
			})
			return
		}
	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": "fail",
		"err":     err.Error(),
	})

}

func SetDnsRecord(zone string, hostname string, value string, recordType string) error {
	// run cmd on Windows
	//eg: dnscmd . /RecordAdd {zone} {hostname} {record type} {IP address}
	cmd := exec.Command("cmd", "/C", "dnscmd . /RecordAdd "+zone+" "+hostname+" "+recordType+" "+value)
	//get output and print
	//seeing if exit code is 0
	output, err := cmd.CombinedOutput()
	log.Info(string(output))
	return err
}

func DelDnsRecord(zone string, hostname string, recordType string) error {
	// run cmd on Windows
	//eg: dnscmd . /RecordDelete {zone} {hostname} {record type}
	cmd := exec.Command("cmd", "/C", "dnscmd . /RecordDelete "+zone+" "+hostname+" "+recordType+" /f")
	// run cmd and enter y to confirm
	output, err := cmd.CombinedOutput()
	log.Info(string(output))
	return err
}
