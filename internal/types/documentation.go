package types

// Documentation is the interface for Action and Reusable
type Documentation interface {
	GetData() error
	RenderOutput() error
}