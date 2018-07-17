////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Copyright (c) 2018 A Bit of Help, Inc. - All Rights Reserved, Worldwide.
// Use of this source code is governed by a BSD-style  license that can be found in the LICENSE file.
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

package dog

import (
	"fmt"
	"github.com/abitofhelp/animals/base"
)

type Dog struct {
	base.Animal
}

func (d *Dog) Sound() string {
	return "woof!"
}

func (d *Dog) MakeSound() {
	fmt.Printf("%v", d.Sound())
}

// Here is how we would build this component
func New() *Dog {
	impl1 := &Dog{}
	abs := &base.Animal{
		IAbstractAnimalMembers: impl1,
	}
	impl1.Animal = *abs
	return impl1
}
