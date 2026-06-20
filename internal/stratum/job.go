package stratum

import (
	"github.com/techobg/prl-forge/internal/pool"
)

var (
	jobManager   = NewJobManager()
	shareManager = NewShareManager()
)

// Job е alias към pool.Job.
// TODO(M3): След миграцията всички файлове в stratum ще използват
// директно pool.Job и този alias ще бъде премахнат.
type Job = pool.Job