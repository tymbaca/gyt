package main

type ObjectKind string

const (
	Blob   ObjectKind = "blob"
	Commit ObjectKind = "commit"
	Tree   ObjectKind = "tree"
	Tag    ObjectKind = "tag"
)

type Object struct {
	Kind ObjectKind
	Size int
	Body []byte
}
