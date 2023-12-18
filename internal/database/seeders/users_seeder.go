package seeders

import (
	"database/sql"
	"fmt"
	"github.com/panjf2000/ants/v2"
	"strings"
	"sync/atomic"
)

var (
	poolSize   = 100
	batchSize  = 1000
	totalCount = 1000000
)

func SeedUsers(db *sql.DB) {
	var totalExec int32
	p, _ := ants.NewPoolWithFunc(poolSize, func(i interface{}) {
		var values []string
		query := "INSERT INTO users (first_name, last_name) VALUES "
		for j := 1; j <= i.(int); j++ {
			values = append(values, fmt.Sprintf("('%s', '%s')", "Dummy First Name", "Dummy Last Name"))
		}
		query += strings.Join(values, ",") + " ;"
		_, err := db.Exec(query)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		atomic.AddInt32(&totalExec, int32(batchSize))
	})
	for atomic.LoadInt32(&totalExec) < int32(totalCount) {
		p.Invoke(batchSize)
	}

	defer p.Release()
}
