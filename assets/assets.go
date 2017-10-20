package assets

// MustAssetString returns the asset contents as a string (instead of a []byte).
func MustAssetString(name string) string {
	return string(MustAsset(name))
}
