package main

import (
	"bufio"
	"strings"
)

// spaceGenerator generate space
func spaceGenerator(i int) string {
	s := ""
	for j := 0; j <= i; j++ {
		s = s + " "
	}
	return s
}

// comparepositon calculate longest column
func comparepositon(position *int, positonIndex int) {
	if positonIndex > *position {
		*position = positonIndex
	}
}

// addspace add space
func addspace(s string, space string, i int) string {
	s = s[:i] + space + s[i:]
	return s
}
func beautifyDef(def string) string {
	// keyword with space
	keyword := map[string]int{" TYPE ": 0, " LIKE ": 0, " = ": 0, " VALUE ": 0, `"`: 0}
	keywordS := []string{" TYPE ", " LIKE ", " =", " VALUE ", `"`}
	sProcess := def
	sResult := ""

	// loop in slice not map, because map won't loop by sequence
	for _, key := range keywordS {
		value := keyword[key]
		sResult = ""
		// find the longest column
		scanner := bufio.NewScanner(strings.NewReader(sProcess))
		for scanner.Scan() {
			// fmt.Println(scanner.Text())
			comparepositon(&value, strings.Index(scanner.Text(), key))
		}
		// add the space
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
