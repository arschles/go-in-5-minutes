package episode20

func (t *TransformerTests) TestTransform() {
	r := t.Require()

	// key exists and the value is a string. should be transformed
	t.transformer.cache["a"] = "b"
	transformed := t.transformer.transform("a", func(s string) string {
		return s + "c"
	})
	r.True(transformed)
	r.Equal("bc", t.transformer.cache["a"])

	// key exists and the value is not a string. should not be transformed
	t.transformer.cache["b"] = 1
	transformed = t.transformer.transform("b", func(s string) string {
		return s + "d"
	})
	r.False(transformed)
	r.Equal(1, t.transformer.cache["b"])

	// key doesn't exist. should not be transformed
	transformed = t.transformer.transform("c", func(s string) string {
		return s + "e"
	})
	r.False(transformed)
	_, found := t.transformer.cache["c"]
	r.False(found)
}
