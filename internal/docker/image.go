package docker

import (
	"context"

	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/daemon"
	"github.com/google/go-containerregistry/pkg/v1/remote"
)

func GetLocalImageDigest(imageName string) (string, error) {
	ref, err := name.ParseReference(imageName)
	if err != nil {
		return "", err
	}

	localImage, err := daemon.Image(ref)
	if err != nil {
		return "", err
	}

	localDigest, err := localImage.ConfigName()
	if err != nil {
		return "", err
	}

	return localDigest.String(), nil
}

func GetRemoteImageDigest(ctx context.Context, imageName string) (string, error) {
	ref, err := name.ParseReference(imageName)
	if err != nil {
		return "", err
	}

	remoteImage, err := remote.Head(ref, remote.WithContext(ctx))
	if err != nil {
		return "", err
	}

	return remoteImage.Digest.String(), nil
}

func CompareDigest(ctx context.Context, imageName string) (bool, error) {
	li, err := GetLocalImageDigest(imageName)
	if err != nil {
		return false, err
	}

	if err := ctx.Err(); err != nil {
		return false, err
	}

	ri, err := GetRemoteImageDigest(ctx, imageName)
	if err != nil {
		return false, err
	}

	return li == ri, nil
}
