package postgres

import (
	"time"

	"github.com/google/uuid"
)

type item struct {
	// use any postgres lib you wish
	//bun.BaseModel `bun:"item"`

	UUID        uuid.UUID //`sortable:"true" bun:"type:uuid,pk,notnull,default:gen_random_uuid()"`
	Name        string    //`sortable:"true" bun:"type:text,notnull"`
	Description string
	Price       float64

	//Medias []media `bun:"rel:has-many,join:uuid=owner_uuid"`

	Updated time.Time //`sortable:"true" bun:"type:timestamptz,nullzero,notnull,default:current_timestamp"`
	Created time.Time //`sortable:"true" bun:"type:timestamptz,nullzero,notnull,default:current_timestamp"`

}
