////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Copyright (c) 2018 A Bit of Help, Inc. - All Rights Reserved, Worldwide.
// Use of this source code is governed by a BSD-style  license that can be found in the LICENSE file.
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"github.com/abitofhelp/animals/cat"
	"github.com/abitofhelp/animals/dog"
)

func main() {

	c := cat.New()
	fmt.Printf("\nMy kitty says ")
	c.MakeSound()

	d := dog.New()
	fmt.Printf("\nMy doggy says ")
	d.MakeSound()

	factory := AnimalFactory()
	kitty, err := factory.Create("cat")
	if err == nil {
		fmt.Printf("\nMy noisy kitty says ")
		kitty.MakeSound()
	} else {
		fmt.Printf("\nMy kitty is too tired to say anything! (%v)", err)
	}

}
