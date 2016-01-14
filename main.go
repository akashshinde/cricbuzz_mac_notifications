package main

import (
	"encoding/xml"
	"fmt"
	"github.com/everdev/mack"
	"io/ioutil"
	"net/http"
	"time"
)

type Result struct {
	TeamA       string
	TeamB       string
	BattingTeam string
	Score       string
	Overs       string
	Wicket      string
}

type MatchData struct {
	Root Main `xml:"mchdata"`
}

type Main struct {
	Matches []Match `xml:"match"`
}

type Match struct {
	Mscr     MSCR   `xml:"mscr"`
	TeamName string `xml:"sName,attr"`
	Status   State  `xml:"state"`
}

type State struct {
	Status string `xml:"status,attr"`
}

type BattingTeam struct {
	Inngs    Inning `xml:"Inngs"`
	TeamName string `xml:"sName,attr"`
}

type MSCR struct {
	Team     BattingTeam `xml:"btTm"`
	TeamName string      `xml:"sName,attr"`
}

type Inning struct {
	R       string `xml:"r,attr"`
	Overs   string `xml:"ovrs,attr"`
	Wickets string `xml:"wkts,attr"`
}

type Root struct {
	XMLName xml.Name `xml:"mchdata"`
	Matches []Match  `xml:"match"`
}

func main() {
	// do stuff
	//mack.Notify("Complete")
	for {
		resp, err := http.Get("http://synd.cricbuzz.com/j2me/1.0/livematches.xml")
		if err != nil {
			panic("Failed")
		}
		str, _ := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		//decoder := xml.NewEncoder(resp.Body)
		var v Root
		err = xml.Unmarshal(str, &v)
		if err != nil {
			fmt.Println("Error is : ", err)
		}
		fmt.Println("%+v", v)
		for key, value := range v.Matches {
			if key == 0 {
				title := fmt.Sprintf("AUS : %s/%s  ", value.Mscr.Team.Inngs.R, value.Mscr.Team.Inngs.Wickets)
				subTitle := fmt.Sprint("Overs : ", value.Mscr.Team.Inngs.Overs)
				fmt.Println("title is : ", title)
				mack.Notify(subTitle, title)
			}
		}
		time.Sleep(1 * time.Minute)
	}
}
