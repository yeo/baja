package baja

// Build execute template and content to generate our real static conent
func Build() int {
	config := DefaultConfig()

	tree := BuildNodeTree(config)
	tree.Compile()
	return 0
}
