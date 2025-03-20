package intrface

import (
	"context"

	"github.com/junaidmdv/clean_architure/BookStore/model"
)


type BookService interface{
	printBookTitle(ctx context.Context, book *model.Book)
}