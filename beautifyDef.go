package main

import (
	"bufio"
	"strings"
)

func spaceGenerator(i int) string {
	s := ""
	for j := 0; j <= i; j++ {
		s = s + " "
	}
	return s
}

func comparepositon(position *int, positonIndex int) {
	if positonIndex > *position {
		*position = positonIndex
	}
}

func addspace(s string, space string, i int) string {
	s = s[:i] + space + s[i:]
	return s
}
func beautifyDef(def string) string {
	keyword := map[string]int{" TYPE ": 0, " LIKE ": 0, " = ": 0, " VALUE ": 0, `"`: 0}
	keywordS := []string{" TYPE ", " LIKE ", " =", " VALUE ", `"`}
	// keyword["TYPE"] = 0
	// keyword["="] = 0
	// keyword["VALUE"] = 0
	// keyword[`"`] = 0
	sProcess := def
	sResult := ""
	// TODO
	// 1. 如果同时执行4种关键字，会导致越调越歪
	// 2. 如果定义科目里存在关键在也会调不准
	for _, key := range keywordS {
		value := keyword[key]
		sResult = ""
		scanner := bufio.NewScanner(strings.NewReader(sProcess))
		for scanner.Scan() {
			// fmt.Println(scanner.Text())
			comparepositon(&value, strings.Index(scanner.Text(), key))
		}
		scanner2 := bufio.NewScanner(strings.NewReader(sProcess))
		for scanner2.Scan() {
			// fmt.Println(scanner2.Text())
			sTemp := scanner2.Text()
			sTemp = strings.Replace(sTemp, "\t", "", -1)
			if value != 0 {
				sTemp = addspace(
					sTemp,
					spaceGenerator(value-strings.Index(sTemp, key)),
					strings.Index(sTemp, key))
			}
			sResult = sResult + sTemp + "\n"
		}
		sProcess = sResult
	}
	return sProcess
}
