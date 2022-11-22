package main

type MRSatan struct{}

func NewMRSatan() *MRSatan {
	return &MRSatan{}
}

func (m MRSatan) Kick() {
	println("MRSatan kicks")
}

func (m MRSatan) Punch() {
	println("MRSatan punches")
}

// The empty method that we want to avoid
func (m MRSatan) Transform() {
	// do nothing
}
