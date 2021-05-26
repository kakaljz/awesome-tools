package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	metricPrefix := "cortex_ingester"
	fd, _ := os.Open("metrics.txt")
	defer fd.Close()
	scanner := bufio.NewScanner(fd)
	helpStr := ""
	typeStr := ""
	metrisList := make(map[string]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		l := strings.Fields(line)
		if l[0] == "#" && l[1] == "HELP" {
			helpStr = strings.Join(l[3:len(l)], " ")
		}
		if l[0] == "#" && l[1] == "TYPE" {
			typeStr = strings.Join(l[3:len(l)], " ")
		}
		if strings.HasPrefix(line, metricPrefix) {
			metricNameOne := strings.Split(l[0], " ")[0]
			metricNameTwo := strings.Split(metricNameOne, "{")
			if _, ok := metrisList[metricNameTwo[0]]; ok {
				continue
			}
			metrisList[metricNameTwo[0]] = 0
			if typeStr == "histogram" && (strings.HasSuffix(metricNameTwo[0], "_sum") || strings.HasSuffix(metricNameTwo[0], "_count")) {
				fmt.Printf("%s|%s|%s\n", metricNameTwo[0], "", "")
			} else {
				fmt.Printf("%s|%s|%s\n", metricNameTwo[0], helpStr, typeStr)
			}
		}
	}
}
