package data


import (
	"fmt"
	"io/ioutil"
	"os"
	"encoding/json"
	
	_ "github.com/lib/pq"
)

type Info struct {
	Name		string `json:"name_ja"`
	Birthday	string `json:"birthday"`
	Height		string `json:"height"`
	BloodType	string `json:"blood_type"`
	Generation	string `json:"generation"`
	BlogUrl		string `json:"blog_url"`
	ImgUrl		string `json:"img_url"`
}

type Position struct {
	NameEn 		string `json:"name_en"`
	NameJa		string `json:"name_ja"`
	Position 	string `json:"position"`
}

type Formation struct {
	Single		string `json:"single"`
	Title		string `json:"title"`
	Position	[]Position `json:"positions"`
}


func LoadMemberInfoFile(group string) []*Info {
	// 実行ファイルからの相対パスなのでこの方式にしてる
	raw, err := ioutil.ReadFile("./db/data/" + group + ".json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	
	data := make([]*Info, 0)

	err = json.Unmarshal(raw, &data)
    if err != nil {
        fmt.Println(err)
    }

	return data
}


func LoadFormationFile(group string) []*Formation {
	// 実行ファイルからの相対パスなのでこの方式にしてる
	raw, err := ioutil.ReadFile("./db/data/positions_" + group + ".json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	
	formations := make([]*Formation, 0)

	err = json.Unmarshal(raw, &formations)
    if err != nil {
        fmt.Println(err)
    }

	return formations
}
