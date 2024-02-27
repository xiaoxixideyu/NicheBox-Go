package model

type LeafInterface interface {
	CreateLeafTX(leaf *Leaf) error
	NextTX(bizTag string, step int) (int64, int64, error)
}
