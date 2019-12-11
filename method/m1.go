package method
import (
    "log"
    "../util"
)
func init () {
    util.M("m1", func (body util.Body) []byte {
        log.Printf("m1 called，data: %v", string(body.Data))
        return body.Data
    })
}
