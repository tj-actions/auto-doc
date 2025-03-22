//Package utils is a package that contains all the utility functions
/*
Copyright © 2021 Tonye Jack <jtonye@ymail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Release struct {
	TagName string `json:"tag_name"`
}

type Tag struct {
	Object struct {
		SHA string `json:"sha"`
	} `json:"object"`
}

// GetTagCommitHash returns the commit hash of a tag
func GetTagCommitHash(repository, tag, token string) (string, error) {
	tagURL := fmt.Sprintf("https://api.github.com/repos/%s/git/refs/tags/%s", repository, tag)
	req, err := http.NewRequest("GET", tagURL, nil)
	if err != nil {
		return "", err
	}
	if token != "" {
		req.Header.Set("Authorization", "token "+token)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var tagObj Tag
	err = json.Unmarshal(body, &tagObj)
	if err != nil {
		return "", err
	}

	return tagObj.Object.SHA, nil
}

// GetLatestRepositoryTag returns the latest tag of a repository
func GetLatestRepositoryTag(repository, token string, useMajorVersion bool, useTagCommitHash bool) (string, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/releases/latest", repository)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	if token != "" {
		req.Header.Set("Authorization", "token "+token)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var release Release
	err = json.Unmarshal(body, &release)
	if err != nil {
		return "", err
	}

	version := release.TagName

	if useMajorVersion && strings.Contains(version, ".") {
		version = strings.Split(version, ".")[0]
	}

	if useTagCommitHash {
		commitHash, err := GetTagCommitHash(repository, version, token)
		if err != nil {
			return "", err
		}
		version = fmt.Sprintf("%s # %s", commitHash, version)
	}

	return version, nil
}
