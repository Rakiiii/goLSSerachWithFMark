package main

import (
	"fmt"
	"log"
	"os"
	//"runtime"
	"strconv"
	"time"
//boolmatrixlib "github.com/Rakiiii/goBoolMatrix"
//graphlib "github.com/Rakiiii/goGraph"
klpartitionlib "github.com/Rakiiii/goKLPartition"
bipartitonlocalsearchlib "github.com/Rakiiii/goBipartitonLocalSearch"
)

func main(){

	disbalance, er := strconv.ParseFloat(os.Args[2], 64)
	if er != nil {
		log.Println(er)
		return
	}

	var graph bipartitonlocalsearchlib.Graph

	if err := graph.ParseGraph(os.Args[1]); err != nil {
		log.Println(err)
		return
	}

	groupSize := graph.AmountOfVertex()/2 - int(float64(graph.AmountOfVertex())*disbalance)

	fmt.Println("GroupSize:", groupSize)

	
	var ord []int

		ord = graph.HungryNumIndependent()

		log.Println("amount of independent:", graph.GetAmountOfIndependent(), "|amount of vertex:", graph.AmountOfVertex())

		timeStart := time.Now()

		depGraph := graph.GetDependentGraph()

		var err error
		//var for result of algorithm
		result := klpartitionlib.Result{Matrix: nil, Value: -1}
		result, err = klpartitionlib. KLPartitionigAlgorithm(&depGraph, result.Matrix)



		sol := bipartitonlocalsearchlib.Solution{Value:-1,Vector:make([]bool,graph.AmountOfVertex()),Gr:&graph} 

		sol.Init(&graph)

		sol.SetDependentAsBinnary(result.Matrix.GetNumber())
		sol.CountMark()
		sol.PartIndependent(groupSize)
		//sol.CountParameter()

		fmt.Println("param:",sol.CountParameter())

		res := bipartitonlocalsearchlib.LSPartiotionAlgorithm(&graph, &sol, groupSize, 0)

		timeEnd := time.Now()
		elapced := timeEnd.Sub(timeStart)
		timeFile, err := os.Create("time")
		defer timeFile.Close()
		if err != nil {
			fmt.Println(err)
		} else {
			timeFile.WriteString(strconv.FormatInt(elapced.Milliseconds(), 10) + "ms")
		}

		formatedRes := make([]int, len(ord))
		strRes := ""
		for i, num := range ord {
			if res.Vector[i] {
				formatedRes[num] = 1
			} else {
				formatedRes[num] = 0
			}
		}

		for _,v := range formatedRes{
			strRes += strconv.Itoa(v)
		}

		fmt.Println(ord)
		fmt.Println(res.Vector)

		f, err := os.Create("result_" + os.Args[1])
		if err != nil {
			fmt.Println("res:", formatedRes)
			fmt.Println("value", res.CountParameter())
			log.Panic(err)
		}
		defer f.Close()

		f.WriteString(strconv.FormatInt(res.Value, 10) + "\n")
		f.WriteString(strRes)
}
