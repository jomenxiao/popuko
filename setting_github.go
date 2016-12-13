package main

type GithubSetting struct {
	botName    string
	token      string
	hookSecret string
	repoList   []RepositorySetting
	repoMap    RepositoryMap
}

type RepositoryMap struct {
	inner map[string]RepositorySetting
}

func newRepositoryMap(list []RepositorySetting) *RepositoryMap {
	m := make(map[string]RepositorySetting)
	for _, item := range list {
		item.Init()

		k := item.Fullname()
		m[k] = item
	}
	return &RepositoryMap{
		m,
	}
}

func (m *RepositoryMap) Entries() map[string]RepositorySetting {
	return m.inner
}

func (m *RepositoryMap) Get(owner string, repo string) *RepositorySetting {
	k := owner + "/" + repo
	item, ok := m.inner[k]
	if ok {
		return &item
	}
	return nil
}
