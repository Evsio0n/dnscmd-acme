#!/bin/bash

dns_win_add() {
  fulldomain=$1
  txtvalue=$2
  url="http://${serverip}:8880/set"
  #Post json to http://${serverip}:8880/set
  #	Zone       string `json:"zone"`
  #	Hostname   string `json:"hostname"`
  #	Value      string `json:"value"`
  #	RecordType string `json:"record_type"`
  #body="{\"zone\":\"$fulldomain\",\"hostname\":\"$fulldomain\",\"value\":\"$txtvalue\",\"record_type\":\"TXT\"}"
  #seperate the domain with dot and remove content before the first dot
  zone=$(echo $fulldomain | cut -d. -f2-)
  #hostname should be only the first part of the domain
  hostname=$(echo $fulldomain | cut -d. -f1)
  body="{\"zone\":\"$zone\",\"hostname\":\"$hostname\",\"value\":\"$txtvalue\",\"record_type\":\"TXT\"}"
  #check http status code = 200
  result="$(_post "$body" "$url" "" "POST")"

}

# Usage: fulldomain txtvalue
# Used to remove the txt record after validation
dns_win_rm() {
  fulldomain=$1
  txtvalue=$2
  url="http://${serverip}:8880/del"
  zone=$(echo $fulldomain | cut -d. -f2-)
  hostname=$(echo $fulldomain | cut -d. -f1)
  body="{\"zone\":\"$zone\",\"hostname\":\"$hostname\",\"record_type\":\"TXT\"}"
  result="$(_post "$body" "$url" "" "POST")"

}
