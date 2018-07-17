////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Copyright (c) 2018 A Bit of Help, Inc. - All Rights Reserved, Worldwide.
// Use of this source code is governed by a BSD-style  license that can be found in the LICENSE file.
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

package cat

import "github.com/abitofhelp/animals/base"

type Cat struct {
	base.Animal
}

func (c *Cat) Sound() string {
	return "meeeoooooowwwwww!"
}

// Here is how we would build a Cat...
func New() *Cat {
	cat := &Cat{}
	abs := &base.Animal{
		IAbstractAnimalMembers: cat,
	}
	cat.Animal = *abs
	return cat
}
