package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"
)

type httpHandler struct {
	DB    *sql.DB
	Token string
}

type userReporting struct {
	Name  string
	Hours int
}

func (m *httpHandler) help(w http.ResponseWriter, r *http.Request) {
	text := `
	Welcome to team presence tracking bot!

	Date format: 2018-03-24
	Commands:
	For average presence - /average <daily/weekly/monthly> <startDate> <endDate> 
	For total presence - /total <startDate> <endDate>
	For this help - /?
	`

	fmt.Fprintf(w, text)
}

func (m *httpHandler) total(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form.", http.StatusBadRequest)
		return
	}

	request := r.Form.Get("text")
	dates := strings.Fields(request)
	var text string
	if len(dates) != 2 {
		text = "Wrong request"
		fmt.Fprintf(w, text)
		return
	}

	t1, err := time.Parse("2006-01-02", dates[0])
	if err != nil {
		text = "Wrong time format"
		fmt.Fprintf(w, text)
		return
	}

	t2, err := time.Parse("2006-01-02", dates[1])
	if err != nil {
		text = "Wrong time format"
		fmt.Fprintf(w, text)
		return
	}

	users, err := readUsersDataFromDB(m.DB, t1.Unix(), t2.Unix())
	if err != nil {
		text = "Something went wrong inside me..."
		log.Println(err)
		fmt.Fprintf(w, text)
		return
	}

	usersActivityCount := make(map[string]int)
	for _, v := range users {
		usersActivityCount[v.ID]++
	}

	usersRaw, err := checkUsersPresence(m.Token)
	if err != nil {
		text = "Something went wrong inside me..."
		log.Println(err)
		fmt.Fprintf(w, text)
		return
	}

	var usersReport []userReporting

	for _, user := range usersRaw {
		for id, hours := range usersActivityCount {
			if id == user.ID {
				usersReport = append(usersReport, userReporting{
					Name:  user.Name,
					Hours: hours,
				})
			}
		}
	}

	sort.Slice(usersReport, func(i, j int) bool {
		return usersReport[i].Hours > usersReport[j].Hours
	})

	for _, user := range usersReport {
		text = fmt.Sprintf("%s%s: %d h\n", text, user.Name, user.Hours/6)
	}

	fmt.Fprintf(w, text)
}
