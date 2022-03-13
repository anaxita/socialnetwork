package faker

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

type F struct {
}

func New() *F {
	return &F{}
}

func (f *F) randString(data []string) string {
	i := rand.Intn(len(data))

	return data[i]
}
