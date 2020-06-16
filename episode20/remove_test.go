package episode20

func (t *TransformerTests) TestRemove() {
	r := t.Require()

	// if they key exists, it should be removed
	t.transformer.cache["a"] = "b"
	removed := t.transformer.remove("a")
	r.True(removed)
	_, found := t.transformer.cache["a"]
	t.False(found)

	// if the key doesn't exist, it should not be removed
	removed = t.transformer.remove("b")
	r.False(removed)
	_, found = t.transformer.cache["b"]
	t.False(found)
}
