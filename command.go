func (c *Command) ValidateRequiredFlags() error {
	// Check local flags
	for _, flag := range c.Flags().FlagUsages() {
		// ... existing logic ...
	}

	// Ensure we check persistent flags from parents
	for p := c.Parent(); p != nil; p = p.Parent() {
		for _, flag := range p.PersistentFlags().FlagUsages() {
			// ... existing logic ...
		}
	}
	return nil
}