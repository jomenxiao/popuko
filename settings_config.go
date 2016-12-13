package main

type RepositorySetting struct {
	owner        string
	name         string
	reviewerList []string
	reviewers    ReviewerSet

	shouldMergeAutomatically bool
	shouldDeleteMerged       bool
	// Use `OWNERS` file as reviewer list in the repository's root.
	useOwnersFile bool
}

func (s *RepositorySetting) Init() {
	set := newReviewerSet(s.reviewerList, false)
	s.reviewerList = nil
	s.reviewers = *set
}

func (s *RepositorySetting) Owner() string {
	return s.owner
}

func (s *RepositorySetting) Name() string {
	return s.name
}

func (s *RepositorySetting) Fullname() string {
	return s.owner + "/" + s.name
}

func (s *RepositorySetting) Reviewers() (ok bool, set *ReviewerSet) {
	return true, &s.reviewers
}

func (r *RepositorySetting) ShouldMergeAutomatically() bool {
	return r.shouldMergeAutomatically
}

func (r *RepositorySetting) ShouldDeleteMerged() bool {
	return r.shouldDeleteMerged
}

func (r *RepositorySetting) UseOwnersFile() bool {
	return r.useOwnersFile
}

type ReviewerSet struct {
	regardAllAsReviewer bool
	set                 map[string]*interface{}
}

func (s *ReviewerSet) Has(person string) bool {
	if s.regardAllAsReviewer {
		return true
	}

	_, ok := s.set[person]
	return ok
}

func (s *ReviewerSet) Entries() []string {
	list := make([]string, 0)
	for k := range s.set {
		list = append(list, k)
	}
	return list
}

func newReviewerSet(list []string, regardAllAsReviewer bool) *ReviewerSet {
	if regardAllAsReviewer {
		return &ReviewerSet{
			true,
			nil,
		}
	}

	s := make(map[string]*interface{})
	for _, name := range list {
		s[name] = nil
	}

	return &ReviewerSet{
		false,
		s,
	}
}
