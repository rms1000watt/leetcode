package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Dino struct {
	Diet         string
	LegLength    float64
	Name         string
	Stance       string
	StrideLength float64
	Speed        float64
}

type DinosArr []*Dino

func (d DinosArr) Len() int           { return len(d) }
func (d DinosArr) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }
func (d DinosArr) Less(i, j int) bool { return d[i].Speed < d[j].Speed }

func main() {
	dinos := map[string]*Dino{}

	// if err := loadDinosFromFiles(dinos); err != nil {
	if err := loadDinosFromStrings(dinos); err != nil {
		fmt.Println("ERROR: failed loading dinos:", err)
		return
	}

	bipedalDinos := filterStance(dinos, "bipedal")
	calculateSpeedAndClean(dinos)

	printDinosSortedBySpeed(bipedalDinos)
}

func printDinosSortedBySpeed(dinos map[string]*Dino) {
	dinosArr := DinosArr{}
	for _, dino := range dinos {
		dinosArr = append(dinosArr, dino)
	}

	sort.Sort(dinosArr)

	for i := len(dinosArr) - 1; i >= 0; i-- {
		fmt.Println(dinosArr[i])
	}
}

func printDinos(dinos map[string]*Dino) {
	for _, dino := range dinos {
		fmt.Println(dino)
	}
}

func calculateSpeedAndClean(dinos map[string]*Dino) {
	// Given the following formula, speed = ((STRIDE_LENGTH / LEG_LENGTH) - 1) * SQRT(LEG_LENGTH * g)
	// Where g = 9.8 m/s^2 (gravitational constant)

	g := float64(9.8)
	for _, dino := range dinos {
		if dino.StrideLength == 0 || dino.LegLength == 0 {
			continue
		}
		dino.Speed = ((dino.StrideLength / dino.LegLength) - 1) * math.Sqrt(dino.LegLength*g)
		dinos[dino.Name] = dino
	}
}

func filterStance(dinos map[string]*Dino, stance string) (out map[string]*Dino) {
	out = map[string]*Dino{}

	for _, dino := range dinos {
		if dino.Stance == stance {
			out[dino.Name] = dino
		}
	}

	return
}

func loadDinosFromStrings(dinos map[string]*Dino) (err error) {
	dinos1String := `NAME,LEG_LENGTH,DIET
  Hadrosaurus,1.4,herbivore
  Struthiomimus,0.72,omnivore
  Velociraptor,1.8,carnivore
  Triceratops,0.47,herbivore
  Euoplocephalus,2.6,herbivore
  Stegosaurus,1.50,herbivore
  Tyrannosaurus Rex,6.5,carnivore`

	reader1 := csv.NewReader(strings.NewReader(dinos1String))
	reader1.Read() // skip columns row

	for {
		dinoRow, err := reader1.Read()
		if err != nil {
			// end of file
			if err == io.EOF {
				break
			}
			fmt.Println("ERROR: failed reading file", err)
			return err
		}

		dino := &Dino{}
		if dinos[dinoRow[0]] != nil {
			dino = dinos[dinoRow[0]]
		}

		legLength, err := strconv.ParseFloat(dinoRow[1], 64)
		if err != nil {
			fmt.Println("ERROR: failed parsing float for leg length:", err)
			return err
		}
		dino.Name = dinoRow[0]
		dino.LegLength = float64(legLength)
		dino.Diet = dinoRow[2]

		dinos[dinoRow[0]] = dino
	}

	dinos2String := `NAME,STRIDE_LENGTH,STANCE
  Euoplocephalus,1.97,quadrupedal
  Stegosaurus,1.70,quadrupedal
  Tyrannosaurus Rex,4.76,bipedal
  Hadrosaurus,1.3,bipedal
  Deinonychus,1.11,bipedal
  Struthiomimus,1.24,bipedal
  Velociraptor,2.62,bipedal`

	reader2 := csv.NewReader(strings.NewReader(dinos2String))
	reader2.Read() // skip columns row

	for {
		dinoRow, err := reader2.Read()
		if err != nil {
			// end of file
			if err == io.EOF {
				break
			}
			fmt.Println("ERROR: failed reading file", err)
			return err
		}

		dino := &Dino{}
		if dinos[dinoRow[0]] != nil {
			dino = dinos[dinoRow[0]]
		}

		strideLength, err := strconv.ParseFloat(dinoRow[1], 64)
		if err != nil {
			fmt.Println("ERROR: failed parsing float for leg length:", err)
			return err
		}
		dino.Name = dinoRow[0]
		dino.StrideLength = float64(strideLength)
		dino.Stance = dinoRow[2]
		dinos[dinoRow[0]] = dino
	}

	return
}

func loadDinosFromFiles(dinos map[string]*Dino) (err error) {

	// file 1
	file1, err := os.Open("05_dino1.csv")
	if err != nil {
		fmt.Println("ERROR: failed opening file1", err)
		return err
	}
	defer file1.Close()

	reader1 := csv.NewReader(file1)
	reader1.Read() // skip columns row

	for {
		dinoRow, err := reader1.Read()
		if err != nil {
			// end of file
			if err == io.EOF {
				break
			}
			fmt.Println("ERROR: failed reading file", err)
			return err
		}

		dino := &Dino{}
		if dinos[dinoRow[0]] != nil {
			dino = dinos[dinoRow[0]]
		}

		dino.Name = dinoRow[0]
		legLength, err := strconv.ParseFloat(dinoRow[1], 64)
		if err != nil {
			fmt.Println("ERROR: failed parsing float for leg length:", err)
			return err
		}
		dino.LegLength = float64(legLength)
		dino.Diet = dinoRow[2]
		dinos[dinoRow[0]] = dino
	}

	// file 2
	file2, err := os.Open("05_dino2.csv")
	if err != nil {
		fmt.Println("ERROR: failed opening file1", err)
		return err
	}
	defer file2.Close()

	reader2 := csv.NewReader(file2)
	reader2.Read() // skip columns row

	for {
		dinoRow, err := reader2.Read()
		if err != nil {
			// end of file
			if err == io.EOF {
				break
			}
			fmt.Println("ERROR: failed reading file", err)
			return err
		}

		dino := &Dino{}
		if dinos[dinoRow[0]] != nil {
			dino = dinos[dinoRow[0]]
		}

		dino.Name = dinoRow[0]
		strideLength, err := strconv.ParseFloat(dinoRow[1], 64)
		if err != nil {
			fmt.Println("ERROR: failed parsing float for leg length:", err)
			return err
		}
		dino.StrideLength = float64(strideLength)
		dino.Stance = dinoRow[2]
		dinos[dinoRow[0]] = dino
	}

	return
}
