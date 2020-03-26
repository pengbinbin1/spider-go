package parser

import (
	//	"log"
	"spider-go/engine"
	"spider-go/model"
	"spider-go/regex"
	"strconv"
)

const agematch = `<td width="180"><span class="grayL">年龄：</span>([\d]+)</td>`
const addresMatch = `<td><span class="grayL">居住地：</span>([^<]+)</td>`
const marrageMatch = `<td width="180"><span class="grayL">婚况：</span>([^<]+)</td>`
const heightMatch = `<td width="180"><span class="grayL">身&nbsp;&nbsp;&nbsp;高：</span>([^<]+)</td>`
const incomeMatch = `<td><span class="grayL">月 薪：</span>([^<]+)</td>`
const genderMatch = `<td width="180"><span class="grayL">性别：</span>([^<]+)</td> `
const nameMatch = `<th><a href="http://album.zhenai.com/u/1333514341" target="_blank">([^<]+)</a></th>`

func ParseProfile(content []byte) engine.ParseRes {

	var person model.Detail

	ageStr := reg.DoProfile(content, agematch)
	person.Age, _ = strconv.Atoi(ageStr)
	person.Address = reg.DoProfile(content, addresMatch)
	person.Height = reg.DoProfile(content, heightMatch)
	person.Marrage = reg.DoProfile(content, marrageMatch)
	person.Income = reg.DoProfile(content, incomeMatch)
	person.Gender = reg.DoProfile(content, genderMatch)
	person.Name = reg.DoProfile(content, nameMatch)

	parRes := engine.ParseRes{
		Items: []interface{}{person},
	}

	return parRes
}
