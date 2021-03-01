package types

func (c ChainInfo) ValidateBasic() error {
	return nil
}

func (c *ChainInfo) Update(newInfo ChainInfo) {
	if len(newInfo.Seed) != 0 {
		c.Seed = newInfo.Seed
	}

	if newInfo.SourceCodeUrl != "" {
		c.SourceCodeUrl = newInfo.SourceCodeUrl
	}

	if newInfo.CanonicalIbcClient != "" {
		c.CanonicalIbcClient = newInfo.CanonicalIbcClient
	}

	if !newInfo.Version.Equal(VersionInfo{}) {
		if newInfo.Version.Version != 0 {
			c.Version.Version = newInfo.Version.Version
		}

		if newInfo.Version.GenesisHash != "" {
			c.Version.GenesisHash = newInfo.Version.GenesisHash
		}

		if newInfo.Version.SourceCommitHash != "" {
			c.Version.SourceCommitHash = newInfo.Version.SourceCommitHash
		}
	}
}
