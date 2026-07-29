package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"unsafe"
)

type visitInfo struct {
	specDoc string
	date    string
}

type magazineStruct struct {
	magazine map[string][]visitInfo
	scanner  *bufio.Scanner
}

type PatientNotFoundError struct {
	Message string
}

func NewMagazine(magazine map[string][]visitInfo, scanner *bufio.Scanner) magazineStruct {
	if magazine != nil {
		return magazineStruct{magazine: magazine, scanner: scanner}
	}
	return magazineStruct{}
}

func (m *magazineStruct) Save() {
	var namePerson, specDoc, date string
	for i := 0; i < 3; i++ { // для возможности вернуться на предыдущий ввод параметра
		switch i {
		case 0:
			fmt.Println("Insert name Person")
			m.scanner.Scan()
			namePerson = strings.ToLower(m.scanner.Text())
		case 1:
			fmt.Println("Insert specialisation.")
			m.scanner.Scan()
			specDoc = m.scanner.Text()
		case 2:
			fmt.Println("Insert Date. Format is YYYY-MM-DD")
			m.scanner.Scan()
			date = m.scanner.Text()
		}
	}
	m.magazine[namePerson] = append(m.magazine[namePerson], visitInfo{specDoc: specDoc, date: date})
}

func (m *magazineStruct) GetHistory() error {
	fmt.Println("Insert name.")
	m.scanner.Scan()
	name := strings.ToLower(m.scanner.Text())
	if _, ok := m.magazine[name]; !ok {
		return PatientNotFoundError{Message: "patient not found"}
	}
	history := m.magazine[name]
	for _, v := range history {
		fmt.Printf("%s %s\n", v.specDoc, v.date)
	}

	return nil
}

func (m *magazineStruct) GetLastVisit() error {
	fmt.Println("Insert name")
	m.scanner.Scan()
	name := strings.ToLower(m.scanner.Text())
	fmt.Println("Insert doc specialisation")
	m.scanner.Scan()
	specDoc := m.scanner.Text()

	if _, ok := m.magazine[name]; !ok {
		return PatientNotFoundError{Message: "patient not found"}
	}
	history := m.magazine[name]

	var lastVisit visitInfo
	for i := 0; i < len(history); i++ {
		if history[i].specDoc != specDoc {
			continue
		}
		if history[i].specDoc == specDoc && unsafe.Sizeof(lastVisit) == 0 {
			lastVisit = history[i]
			continue
		}
		lastDate, _ := time.Parse("2006-01-01", lastVisit.date)
		t1, _ := time.Parse("2006-01-01", history[i].date)
		if t1.After(lastDate) {
			lastVisit = history[i]
		}
	}

	if unsafe.Sizeof(lastVisit) == 0 {
		return nil
	}

	fmt.Println(lastVisit.date)

	return nil
}

func (e PatientNotFoundError) Error() string {
	return e.Message
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	magazine := NewMagazine(make(map[string][]visitInfo), scanner)

	var err error

	for {
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		switch input {
		case "Save":
			magazine.Save()
		case "GetHistory":
			if err = magazine.GetHistory(); err != nil {
				fmt.Println(err)
			}
		case "GetLastVisit":
			if err = magazine.GetLastVisit(); err != nil {
				fmt.Println(err)
			}
		case "exit", "\\q", "q":
			fmt.Println("Good luck!")
			return
		default:
			fmt.Println("Invalid Function")
		}
	}

}
