package brand

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestQinghaiBrandTask16(t *testing.T) {
	s := NewService(NewRegistry(), time.Now)
	p := ProductListing{OriginRegion: "青海", IngredientLotIDs: []string{"l"}}
	require.NoError(t, s.CheckHighlandProduct(context.Background(), p, map[string]IngredientLot{"l": {ID: "l", OriginRegion: "青海"}}))
}
