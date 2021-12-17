package book

type Service interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	Create(bookRequest BookRequest) (Book, error)
	Update(ID int, bookRequest BookRequest) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	books, err := s.repository.FindAll()
	return books, err
}

func (s *service) FindByID(id int) (Book, error) {
	book, err := s.repository.FindByID(id)
	return book, err
}

func (s *service) Create(bookRequest BookRequest) (Book, error) {

	price, _ := bookRequest.Price.Int64()
	discount, _ := bookRequest.Discount.Int64()
	rating, _ := bookRequest.Rating.Int64()
	book := Book{
		Title:       bookRequest.Title,
		Author:      bookRequest.Author,
		Description: bookRequest.Description,
		Rating:      int(rating),
		Discount:    int(discount),
		Price:       int(price),
	}

	newbook, err := s.repository.Create(book)

	return newbook, err
}

func (s *service) Update(ID int, bookRequest BookRequest) (Book, error) {

	book_, err := s.repository.FindByID(ID)

	price, _ := bookRequest.Price.Int64()
	discount, _ := bookRequest.Discount.Int64()
	rating, _ := bookRequest.Rating.Int64()

	book_.Title = bookRequest.Title
	book_.Author = bookRequest.Author
	book_.Description = bookRequest.Description
	book_.Rating = int(rating)
	book_.Discount = int(discount)
	book_.Price = int(price)

	newbook, err := s.repository.Update(book_)

	return newbook, err
}
