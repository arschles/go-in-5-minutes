//go:build mage
// +build mage

package main

import (
	"context"
	"fmt"

	"github.com/magefile/mage/sh"
	"golang.org/x/sync/errgroup"
)

func Build(ctx context.Context) error {
	const lastEpisode = 32
	g, ctx := errgroup.WithContext(ctx)
	for i := 0; i <= lastEpisode; i++ {
		idx := i
		g.Go(func() error {
			return sh.RunV(
				"go",
				"build",
				"-o",
				fmt.Sprintf("bin/episode%d", idx),
				fmt.Sprintf("./episode%d", idx),
			)
		})
	}
	return g.Wait()
}

func BuildOne(ctx context.Context, episodeNum int) error {
	return sh.RunV(
		"go",
		"build",
		"-o",
		fmt.Sprintf("bin/episode%d", episodeNum),
		fmt.Sprintf("./episode%d", episodeNum),
	)
}
