package entity

type Product struct {
	Id    int
	Name  string
	Price int
	Stock int
}

// gunakan pointer untuk mereferensi struct pada memory
// hal ini  lebih hemat memory
// daripada mendefenisikan untuk setiap objek secara berulang
func (p *Product) StockStatus() string {
	var status string
	if p.Stock < 3 {
		status = "Stock Hampir Habis"
	} else if p.Stock < 10 {
		status = "Stock Terbatas"
	}
	return status
}
