////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Copyright (c) 2018 A Bit of Help, Inc. - All Rights Reserved, Worldwide.
// Use of this source code is governed by a BSD-style  license that can be found in the LICENSE file.
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"errors"
	"github.com/abitofhelp/animals/cat"
	"github.com/abitofhelp/animals/dog"
)

// Default component implementation to use
const defaultName = "cat"

var instance *Factory

type Factory struct {
	// Map of constructors for the animals
	ctors map[string]func() IAnimal
}

func (f *Factory) New() IAnimal {
	ret, _ := f.Create(defaultName)
	return ret
}

func (f *Factory) Create(name string) (IAnimal, error) {
	ctor, ok := f.ctors[name]
	if !ok {
		return nil, errors.New("animal not found")
	}
	return ctor(), nil
}

func (f *Factory) Register(name string, constructor func() IAnimal) {
	f.ctors[name] = constructor
}

func AnimalFactory() *Factory {
	if instance == nil {
		instance = &Factory{ctors: map[string]func() IAnimal{}}
	}

	return instance
}

// Here we register the implementations in the factory
func init() {
	AnimalFactory().Register("cat", func() IAnimal { return cat.New() })
	AnimalFactory().Register("dog", func() IAnimal { return dog.New() })
}
