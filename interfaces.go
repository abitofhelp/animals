////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Copyright (c) 2018 A Bit of Help, Inc. - All Rights Reserved, Worldwide.
// Use of this source code is governed by a BSD-style  license that can be found in the LICENSE file.
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Package animals provides methods for seeing what animals sound like.
package main

// Interface IAnimal defines the methods that must be implemented to satisfy it.
type IAnimal interface {
	Sound() string
	MakeSound()
}
