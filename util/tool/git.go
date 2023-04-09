package tool

import (
	"blog/config"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
)

// GitClone
// @Description: 克隆仓库
// @param url
// @param path
// @return *git.Repository
// @return error
func GitCloneToMemory(url string) (*git.Repository, error) {

	return git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: "https://github.com/src-d/go-git",
	})
}
func GitCloneToDisk(url, path string) (*git.Repository, error) {
	return git.PlainClone(path, false, &git.CloneOptions{
		URL: url,
	})
}

func GitPull() {
	// 打开仓库
	repo, err := git.PlainOpen(config.BLOG_ARTICLE)

	if err != nil {
		return
	}
	if err != nil {
		fmt.Println("Error opening repository:", err)
		return
	}

	// 获取最新提交
	pullBefore, err := repo.Head()
	if err != nil {
		fmt.Println("Error getting HEAD reference:", err)
		return
	}

	fmt.Println("pullBefore is", pullBefore.Hash())
}
