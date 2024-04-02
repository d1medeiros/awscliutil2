package service

import (
	"api-customer/internal/model"
	"math/rand"
	"time"
)

var mapper = map[string]model.Customer{
	"1":  {Id: "1", Name: "diego", Document: "1234"},
	"2":  {Id: "2", Name: "natalia", Document: "4321"},
	"3":  {Id: "3", Name: "john", Document: "5678"},
	"4":  {Id: "4", Name: "emma", Document: "8765"},
	"5":  {Id: "5", Name: "alex", Document: "2345"},
	"6":  {Id: "6", Name: "sara", Document: "6789"},
	"7":  {Id: "7", Name: "michael", Document: "5432"},
	"8":  {Id: "8", Name: "lucy", Document: "9876"},
	"9":  {Id: "9", Name: "james", Document: "3456"},
	"10": {Id: "10", Name: "sophia", Document: "7890"},
	"11": {Id: "11", Name: "oliver", Document: "9870"},
	"12": {Id: "12", Name: "chloe", Document: "5670"},
	"13": {Id: "13", Name: "william", Document: "4320"},
	"14": {Id: "14", Name: "mia", Document: "7894"},
	"15": {Id: "15", Name: "benjamin", Document: "6547"},
	"16": {Id: "16", Name: "ava", Document: "2319"},
	"17": {Id: "17", Name: "logan", Document: "8902"},
	"18": {Id: "18", Name: "harper", Document: "4567"},
	"19": {Id: "19", Name: "jackson", Document: "1236"},
	"20": {Id: "20", Name: "amelia", Document: "9081"},
	"21": {Id: "21", Name: "liam", Document: "1235"},
	"22": {Id: "22", Name: "olivia", Document: "5670"},
	"23": {Id: "23", Name: "ethan", Document: "7890"},
	"24": {Id: "24", Name: "ava", Document: "0987"},
	"25": {Id: "25", Name: "noah", Document: "6754"},
	"26": {Id: "26", Name: "sophia", Document: "4567"},
	"27": {Id: "27", Name: "mason", Document: "0987"},
	"28": {Id: "28", Name: "isabella", Document: "4321"},
	"29": {Id: "29", Name: "william", Document: "1234"},
	"30": {Id: "30", Name: "emma", Document: "5678"},
	"31": {Id: "31", Name: "james", Document: "3456"},
	"32": {Id: "32", Name: "ava", Document: "9087"},
	"33": {Id: "33", Name: "charlotte", Document: "7654"},
	"34": {Id: "34", Name: "michael", Document: "1234"},
	"35": {Id: "35", Name: "sophia", Document: "9876"},
	"36": {Id: "36", Name: "william", Document: "5432"},
	"37": {Id: "37", Name: "mia", Document: "5678"},
	"38": {Id: "38", Name: "alexander", Document: "9087"},
	"39": {Id: "39", Name: "amelia", Document: "2345"},
	"40": {Id: "40", Name: "jacob", Document: "6789"},
	"41": {Id: "41", Name: "emily", Document: "5432"},
	"42": {Id: "42", Name: "michael", Document: "9876"},
	"43": {Id: "43", Name: "olivia", Document: "4567"},
	"44": {Id: "44", Name: "james", Document: "3456"},
	"45": {Id: "45", Name: "elizabeth", Document: "1235"},
	"46": {Id: "46", Name: "william", Document: "5678"},
	"47": {Id: "47", Name: "charlotte", Document: "9087"},
	"48": {Id: "48", Name: "mason", Document: "2345"},
	"49": {Id: "49", Name: "ava", Document: "6789"},
	"50": {Id: "50", Name: "oliver", Document: "4321"},
}

func GetCustomersAll() []model.Customer {
	var listC []model.Customer
	for _, c := range mapper {
		listC = append(listC, c)
	}
	//time.Sleep(time.Duration(randRange(400, 700)) * time.Millisecond)
	time.Sleep(400 * time.Millisecond)
	return listC
}
func randRange(min, max int) int {
	return rand.Intn(max-min) + min
}
