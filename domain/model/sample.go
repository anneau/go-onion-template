package model

import "fmt"

type SampleID = ID[Sample]

type Sample struct {
	SampleID  SampleID
	Name string
}

func NewSample(name string) (*Sample, error) {
	sampleID, err := NewID[Sample]()

	if err != nil {
		return nil, err
	}

	if name == "" {
		return nil, fmt.Errorf("名前は必須です")
	}

	return &Sample{
		SampleID: sampleID,
	}, nil
}

type SampleRepository interface {
	Find(id SampleID) (*Sample, error)
}
