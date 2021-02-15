package types

func (c ChainInfo) ValidateBasic() error {
	return nil
}

func (c ChainInfo) Update(newInfo ChainInfo) {
	if len(newInfo.Seed) != 0 {
		c.Seed = newInfo.Seed
	}

	if newInfo.SourceCodeUrl != "" {
		c.SourceCodeUrl = newInfo.SourceCodeUrl
	}

	if newInfo.CanonicalIbcClient != "" {
		c.CanonicalIbcClient = newInfo.CanonicalIbcClient
	}
}
