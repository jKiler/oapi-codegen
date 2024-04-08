// Package param provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.0.0-00010101000000-000000000000 DO NOT EDIT.
package param

import (
	"encoding/json"
	"errors"

	"github.com/oapi-codegen/runtime"
)

// Defines values for CatBreed.
const (
	MaineCoon CatBreed = "maine_coon"
	Ragdoll   CatBreed = "ragdoll"
)

// Defines values for DogBreed.
const (
	Beagle DogBreed = "beagle"
	Poodle DogBreed = "poodle"
)

// Cat defines model for cat.
type Cat struct {
	Breed CatBreed `json:"breed"`
}

// CatBreed defines model for Cat.Breed.
type CatBreed string

// Dog defines model for dog.
type Dog struct {
	Breed DogBreed `json:"breed"`
}

// DogBreed defines model for Dog.Breed.
type DogBreed string

// Pet defines model for pet.
type Pet struct {
	union json.RawMessage
}

// AsDog returns the union data inside the Pet as a Dog
func (t Pet) AsDog() (Dog, error) {
	var body Dog
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromDog overwrites any union data inside the Pet as the provided Dog
func (t *Pet) FromDog(v Dog) error {
	v.Breed = "beagle"
	v.Breed = "poodle"
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeDog performs a merge with any union data inside the Pet, using the provided Dog
func (t *Pet) MergeDog(v Dog) error {
	v.Breed = "beagle"
	v.Breed = "poodle"
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsCat returns the union data inside the Pet as a Cat
func (t Pet) AsCat() (Cat, error) {
	var body Cat
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromCat overwrites any union data inside the Pet as the provided Cat
func (t *Pet) FromCat(v Cat) error {
	v.Breed = "maine_coon"
	v.Breed = "ragdoll"
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeCat performs a merge with any union data inside the Pet, using the provided Cat
func (t *Pet) MergeCat(v Cat) error {
	v.Breed = "maine_coon"
	v.Breed = "ragdoll"
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t Pet) Discriminator() (string, error) {
	var discriminator struct {
		Discriminator string `json:"breed"`
	}
	err := json.Unmarshal(t.union, &discriminator)
	return discriminator.Discriminator, err
}

func (t Pet) ValueByDiscriminator() (interface{}, error) {
	discriminator, err := t.Discriminator()
	if err != nil {
		return nil, err
	}
	switch discriminator {
	case "beagle":
		return t.AsDog()
	case "maine_coon":
		return t.AsCat()
	case "poodle":
		return t.AsDog()
	case "ragdoll":
		return t.AsCat()
	default:
		return nil, errors.New("unknown discriminator value: " + discriminator)
	}
}

func (t Pet) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *Pet) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}