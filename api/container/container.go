package container

import (
	"github.com/mmuflih/recipe/api/container/provider"
	"go.uber.org/dig"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
**/

func BuildHandlerProvider(c *dig.Container) *dig.Container {
	for _, h := range provider.Handlers() {
		if err := c.Provide(h); err != nil {
			panic(err)
		}
	}
	return c
}

func BuildRepositoryProvider(c *dig.Container) *dig.Container {
	for _, h := range provider.Repositories() {
		if err := c.Provide(h); err != nil {
			panic(err)
		}
	}
	return c
}

func BuildUsecaseProvider(c *dig.Container) *dig.Container {
	for _, h := range provider.Usecases() {
		if err := c.Provide(h); err != nil {
			panic(err)
		}
	}
	return c
}

func BuildContainer() *dig.Container {
	c := dig.New()

	c = provider.BuildConfigProvider(c)
	c = BuildRepositoryProvider(c)
	c = BuildUsecaseProvider(c)
	c = BuildHandlerProvider(c)

	return c
}
