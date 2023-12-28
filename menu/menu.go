package menu

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	// project imports
	fileio "github.com/rDybing/ksp2scirep/fileIO"
	science "github.com/rDybing/ksp2scirep/science"
)

const dir = string("Saves")
const divider = "------------------------------------------------------------------------------"
const delay = 1

type bodiesT struct {
	name  string
	moons []string
}

var bodies = []bodiesT{
	{
		name:  "Moho",
		moons: nil,
	},
	{
		name:  "Eve",
		moons: []string{"Gilly"},
	},
	{
		name:  "Kerbin",
		moons: []string{"Mun", "Minmus"},
	},
	{
		name:  "Duna",
		moons: []string{"Ike"},
	},
	{
		name:  "Dres",
		moons: nil,
	},
	{
		name:  "Jool",
		moons: []string{"Pol", "Bop", "Tylo", "Val", "Laythe"},
	},
	{
		name:  "Eelo",
		moons: nil,
	},
}

func MainMenu() error {
	var quit bool
	for !quit {
		saveList, err := fileio.ReadDir(dir)
		if err != nil {
			return err
		}
		var saveRaw []byte
		if len(saveList) > 1 {
			saveFile := selectSave(saveList)
			saveRaw, err = fileio.LoadSave(dir, saveFile)
		} else {
			saveRaw, err = fileio.LoadSave(dir, saveList[0])
		}
		if err != nil {
			return err
		}
		if len(saveRaw) == 0 {
			return fmt.Errorf("empty save file")
		}
		science := &science.SortedT{
			Body: make(map[string]science.ReportT),
		}
		if err := science.SetData(saveRaw); err != nil {
			return err
		}
		quit = selectMajorBody(*science)
	}
	return nil
}

func selectSave(list []string) string {
	fmt.Printf("Select save, 1 through %d\n", len(list))
	for i, v := range list {
		fmt.Printf("%d\t%s\n", i+1, v)
	}
	var inputStr, out string
	var done bool
	for !done {
		fmt.Scanf("%s\n", &inputStr)
		input, err := strconv.Atoi(inputStr)
		if err != nil {
			fmt.Println("Numbers only!")
		} else {
			if input > len(list) {
				fmt.Println("Selected save do not exist, try again!")
			} else {
				out = list[input-1]
				done = true
			}
		}
	}
	return out
}

func selectMajorBody(s science.SortedT) bool {
	var done, quit bool
	for !done {
		var inputStr, body string
		fmt.Print("\033[1;1H\033[0J")
		fmt.Println(divider)
		fmt.Println("Select Major Body: (B = Back to Save File Select, Q = Quit)")
		drawMajorBody()
		fmt.Scanf("%s\n", &inputStr)
		input, err := strconv.Atoi(inputStr)
		if err != nil {
			inputStr = strings.ToLower(inputStr)
			switch inputStr {
			case "b":
				done = true
				continue
			case "q":
				done = true
				quit = true
				continue
			default:
				fmt.Println("Not valid input, Try again...")
				time.Sleep(delay * time.Second)
				continue
			}
		}
		if input > len(bodies) {
			fmt.Println("Selected body do not exist, try again!")
			time.Sleep(delay * time.Second)
			continue
		}
		if bodies[input-1].moons != nil {
			body = selectMinorBody(input - 1)
			if body == "b" || body == "B" {
				continue
			}
		} else {
			body = bodies[input-1].name
		}
		fmt.Print("\033[1;1H\033[0J")
		fmt.Println(divider)
		if len(s.Body[body].Data) == 0 {
			fmt.Printf("No science data for %s\n", body)
			time.Sleep(delay * time.Second)
			continue
		}
		fmt.Printf("%-8s%-15s%-15s%-15s%-20s%5s\n", body, "State", "Biome", "Type", "Description", "Value")
		fmt.Println(divider)
		for _, v := range s.Body[body].Data {
			fmt.Printf("\t%-15s%-15s%-15s%-20s%4d\n", v.State, v.Biome, v.Type, v.Description, v.Value)
		}
		fmt.Println(divider)
		fmt.Println("Hit return to go back...")
		fmt.Scanf("%s\n", &inputStr)
	}
	return quit
}

func selectMinorBody(index int) string {
	var done bool
	var body string
	for !done {
		var inputStr string
		fmt.Print("\033[1;1H\033[0J")
		fmt.Println(divider)
		fmt.Println("Select Body: (B = Back to Major Body Select)")
		drawMinorBody(index)
		fmt.Scanf("%s\n", &inputStr)
		input, err := strconv.Atoi(inputStr)
		if err != nil {
			inputStr = strings.ToLower(inputStr)
			switch inputStr {
			case "b":
				done = true
				body = inputStr
				continue
			default:
				fmt.Println("Not valid input, Try again...")
				time.Sleep(delay * time.Second)
				continue
			}
		}
		if input > len(bodies[index].moons)+1 {
			fmt.Println("Selected body do not exist, try again!")
			time.Sleep(delay * time.Second)
			continue
		}
		if input == 1 {
			body = bodies[index].name
		} else {
			body = bodies[index].moons[input-2]
		}
		done = true
	}
	return body
}

func drawMajorBody() {
	fmt.Println(divider)
	for i, v := range bodies {
		fmt.Printf("%2d: %s\n", i+1, v.name)
	}
	fmt.Println(divider)
	fmt.Println()
}

func drawMinorBody(index int) {
	fmt.Println(divider)
	fmt.Println(" 1: " + bodies[index].name)
	for i, v := range bodies[index].moons {
		fmt.Printf("%2d: %s\n", i+2, v)
	}
	fmt.Println(divider)
	fmt.Println()
}
