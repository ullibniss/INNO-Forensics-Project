package mac

type Repository interface {
	FindVendor(prefix string) (string, error)
}
