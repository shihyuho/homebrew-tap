package gh

import "fmt"

type Repo struct {
	Owner, Name string
}

func (r *Repo) String() string {
	return fmt.Sprintf("%s/%s", r.Owner, r.Name)
}
