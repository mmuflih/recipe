package main

import (
	"fmt"

	"github.com/mmuflih/recipe/api/container"
	"go.uber.org/dig"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
**/

var _ = dig.Name

func main() {
	c := container.BuildContainer()

	if err := c.Invoke(container.InvokeRoute); err != nil {
		panic(err)
	}

	if err := c.Provide(container.NewRoute); err != nil {
		panic(err)
	}

	if err := c.Invoke(func(s *container.ServerRoute) {
		s.Run()
	}); err != nil {
		fmt.Println(err)
	}
}
