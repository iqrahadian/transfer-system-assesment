package ctx

import (
	"context"
	"time"
)

type Carrier struct {
	Ctx context.Context
	// Tx           *repo.Tx
	UserId       *string
	CommitCalled bool
	origin       string
	creationTime time.Time
	Parent       *Carrier
}

// func (c *Carrier) Commit() common.Error {
// 	c.CommitCalled = true
// 	return c.Tx.Commit()
// }

// func (c *Carrier) Close() {

// 	// close if tx still active
// 	if !c.CommitCalled && c.Tx != nil {
// 		err := c.Tx.Commit()
// 		if err.Error != nil {
// 			logger.Error("FAILED TO Close Trx", err)
// 		}
// 	}

// 	// avoid unnecessary computation to be executed
// 	if config.ENV.Log.Level == "DEBUG" {
// 		endTime := time.Now()
// 		timeDiff := util.GetMillisDiff(c.creationTime, endTime)

// 		functionName := strings.Split(c.origin, "/")

// 		logger.Debug(fmt.Sprintf("[%s] finished in %d millisecond", functionName[len(functionName)-1], timeDiff), nil)
// 	}
// }
