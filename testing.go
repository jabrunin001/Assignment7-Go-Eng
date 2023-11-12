package main

import (
	"fmt"
	"log"

	"github.com/e-XpertSolutions/go-iforest/iforest"
	"github.com/petar/GoMNIST"
)

func main() {
	// Load the MNIST data
	inputData, err := loadMNIST("/Users/jamesbruning/Desktop/Data Eng with Go/Module 7/Week_7_Assignment/Go_Testing/data/train-images-idx3-ubyte.gz")
	if err != nil {
		log.Fatal("Failed to load MNIST data:", err)
	}

	// Setup the Isolation Forest model
	treesNumber := 100
	subsampleSize := 256
	outliersRatio := 0.01
	routinesNumber := 10

	forest := iforest.NewForest(treesNumber, subsampleSize, outliersRatio)

	// Train the model
	forest.Train(inputData)

	// Test the model
	forest.TestParallel(inputData, routinesNumber)

	// Print the anomaly threshold
	fmt.Println("Anomaly Threshold:", forest.AnomalyBound)

	// Predict anomalies
	labels, scores, err := forest.Predict(inputData)
	if err != nil {
		log.Fatal("Failed to predict anomalies:", err)
	}

	// Print anomaly scores for the first 10 instances
	fmt.Println("Anomaly Scores for first 10 instances:")
	for i := 0; i < 10; i++ {
		fmt.Printf("Instance %d: Score = %f, Label = %d\n", i, scores[i], labels[i])
	}
}

func loadMNIST(filename string) ([][]float64, error) {
	set, err := GoMNIST.ReadSet(filename, "")
	if err != nil {
		return nil, err
	}

	data := make([][]float64, set.Count())
	for i := range data {
		img := set.Images[i]
		normalized := make([]float64, len(img))
		for j, pixel := range img {
			normalized[j] = float64(pixel) / 255.0
		}
		data[i] = normalized
	}

	return data, nil
}
