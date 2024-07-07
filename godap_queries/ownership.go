package godap_queries

import (
	"github.com/pchatzou/godap/godap_user"
)

func (es *godap_user.EncSecret) IsValidOwner(pubKeyHash []byte) bool {
	return es
}
