package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Data []float64

type Report struct {
	Data
}

func (reportTemp *Data) addTemp(temp float64) {
	*reportTemp = append(*reportTemp, temp)

}

func (reportTemp Data) CalculateMean() float64 {
	var sum float64
	for _, mean := range reportTemp {
	}
		sum += mean
	return float64(sum / float64(len(reportTemp)))
}

func (reportTemp Data) CalculateMedian() float64 {
	median := 0.0
	if len(reportTemp)%2 != 0 {
		median = float64(len(reportTemp)/2 + 1)
	} else {
		temp := float64(len(reportTemp)/2) + 1
		median = (float64(len(reportTemp)/2) + temp) / 2
	}
	fmt.Println(median)
	return median
}

//Calculate the mean and median
//Mean, Median - Temperature data => Third column

func main() {
	var fileBytes *csv.Reader
	if ok := openFile("data.txt"); ok == nil {
		panic(ok)
	} else {
		fileBytes = csv.NewReader(ok)
	}
	fileBytes.Comma = '\t'
	fileBytes.TrailingComma = true
	data, err := fileBytes.ReadAll()
	tempReport, err := peekData(data, 1)
	rHumidityReport, _ := peekData(data, 4)
	windSpeedReport, err := peekData(data, 5)
	if err != nil {
		panic(err)
	}

	fmt.Println("Air temp: ", *tempReport, "Mean:", tempReport.CalculateMean(), "Median:", tempReport.CalculateMedian())
	fmt.Println("Humidity: ", *rHumidityReport)
	fmt.Println("Wind Speed: ", *windSpeedReport)
}

func openFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error occurred while opennig file")
		return nil
	}
	return file

}

func peekData(data [][]string, dataSetLocation int) (*Report, error) {
	report := &Report{}
// reportData := make([]float64, len(data)-1)
	for rows, data1 := range data {
		if rows != 0 {
			for j, data12 := range data1 {
				if j == dataSetLocation {
					f, err := strconv.ParseFloat(data12, 64)
					if err != nil {
						fmt.Println("could not parse median: ", err)
						return report, err
					}
					// reportData[rows-1] = f
					// report.Temp = reportData
					report.addTemp(f)
				}
			}
		}
	}
	return report, nil
}
