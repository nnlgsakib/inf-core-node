package server

import (
	"github.com/Infinity-Green/inf/chain"
	"github.com/Infinity-Green/inf/consensus"
	consensusDev "github.com/Infinity-Green/inf/consensus/dev"
	consensusDummy "github.com/Infinity-Green/inf/consensus/dummy"
	consensusIBFT "github.com/Infinity-Green/inf/consensus/ibft"
	consensusPolyBFT "github.com/Infinity-Green/inf/consensus/polybft"
	"github.com/Infinity-Green/inf/forkmanager"
	"github.com/Infinity-Green/inf/secrets"
	"github.com/Infinity-Green/inf/secrets/awsssm"
	"github.com/Infinity-Green/inf/secrets/gcpssm"
	"github.com/Infinity-Green/inf/secrets/hashicorpvault"
	"github.com/Infinity-Green/inf/secrets/local"
	"github.com/Infinity-Green/inf/state"
)

type GenesisFactoryHook func(config *chain.Chain, engineName string) func(*state.Transition) error

type ConsensusType string

type ForkManagerFactory func(forks *chain.Forks) error

type ForkManagerInitialParamsFactory func(config *chain.Chain) (*forkmanager.ForkParams, error)

const (
	DevConsensus     ConsensusType = "dev"
	IBFTConsensus    ConsensusType = "ibft"
	PolyBFTConsensus ConsensusType = consensusPolyBFT.ConsensusName
	DummyConsensus   ConsensusType = "dummy"
)

var consensusBackends = map[ConsensusType]consensus.Factory{
	DevConsensus:     consensusDev.Factory,
	IBFTConsensus:    consensusIBFT.Factory,
	PolyBFTConsensus: consensusPolyBFT.Factory,
	DummyConsensus:   consensusDummy.Factory,
}

// secretsManagerBackends defines the SecretManager factories for different
// secret management solutions
var secretsManagerBackends = map[secrets.SecretsManagerType]secrets.SecretsManagerFactory{
	secrets.Local:          local.SecretsManagerFactory,
	secrets.HashicorpVault: hashicorpvault.SecretsManagerFactory,
	secrets.AWSSSM:         awsssm.SecretsManagerFactory,
	secrets.GCPSSM:         gcpssm.SecretsManagerFactory,
}

var genesisCreationFactory = map[ConsensusType]GenesisFactoryHook{
	PolyBFTConsensus: consensusPolyBFT.GenesisPostHookFactory,
}

var forkManagerFactory = map[ConsensusType]ForkManagerFactory{
	PolyBFTConsensus: consensusPolyBFT.ForkManagerFactory,
}

var forkManagerInitialParamsFactory = map[ConsensusType]ForkManagerInitialParamsFactory{
	PolyBFTConsensus: consensusPolyBFT.ForkManagerInitialParamsFactory,
}

func ConsensusSupported(value string) bool {
	_, ok := consensusBackends[ConsensusType(value)]

	return ok
}
