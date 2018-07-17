////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Copyright (c) 2018 A Bit of Help, Inc. - All Rights Reserved, Worldwide.
// Use of this source code is governed by a BSD-style  license that can be found in the LICENSE file.
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

package base

import "fmt"

// Type BaseAnimal is an abstract base class that provides a default
// implementation of MakeSound().
type Animal struct {
	IAbstractAnimalMembers
}

// Function MakeSound() is the default implementation for subclasses.
func (ba *Animal) MakeSound() {
	fmt.Printf("%v", ba.Sound())
}
