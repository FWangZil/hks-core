package user

import "github.com/FWangZil/hks-core/pkg"

// byDate implements sort.Interface.
type byDate []pkg.User

func (d byDate) Len() int { return len(d) }

//func (d byDate) Less(i, j int) bool { return d[i].Date.After(d[j].Date) }
func (d byDate) Swap(i, j int) { d[i], d[j] = d[j], d[i] }
