//go:build mage
// +build mage

package main

import (
	"context"
	"fmt"

	"github.com/magefile/mage/sh"
	"golang.org/x/sync/errgroup"
)

func BuildAll(ctx context.Context) error {
	skips := map[int]bool{
		// 23: false,
	}
	const lastEpisode = 32
	g, ctx := errgroup.WithContext(ctx)
	for i := 0; i <= lastEpisode; i++ {
		if _, ok := skips[i]; ok {
			continue
		}
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

func TestOne(ctx context.Context, episodeNum int) error {
	return sh.RunV(
		"go",
		"test",
		fmt.Sprintf("./episode%d", episodeNum),
	)
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
