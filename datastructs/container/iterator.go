package container

type Iterator interface {
	Valid() bool
	Next() Comparer
}
