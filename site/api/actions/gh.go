package actions

import (
	"context"

	"github.com/google/go-github/github"
)

func getFileFromGH(ctx context.Context, ghClient *github.Client, path string) (string, error) {
	// https://github.com/arschles/care-anywhere-tools/pull/507
	repoContent, _, _, err := ghClient.Repositories.GetContents(
		ctx,
		"arschles",
		"go-in-5-minutes",
		path,
		&github.RepositoryContentGetOptions{Ref: "master"},
	)
	// TODO
	// func (s *RepositoriesService) DownloadContents(ctx context.Context, owner, repo, filepath string, opt *RepositoryContentGetOptions) (io.ReadCloser, error)
	txt, err := repoContent.GetContent()
	if err != nil {
		return "", err
	}
	return txt, nil
}
