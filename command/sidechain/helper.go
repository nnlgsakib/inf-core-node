package sidechain

import (
	"errors"
	"fmt"
	"os"

	"github.com/Infinity-Green/inf/command/polybftsecrets"
	rootHelper "github.com/Infinity-Green/inf/command/rootchain/helper"
	"github.com/Infinity-Green/inf/consensus/polybft"
	"github.com/Infinity-Green/inf/consensus/polybft/contractsapi"
	"github.com/Infinity-Green/inf/consensus/polybft/wallet"
	"github.com/Infinity-Green/inf/contracts"
	"github.com/Infinity-Green/inf/helper/common"
	"github.com/Infinity-Green/inf/txrelayer"
	"github.com/Infinity-Green/inf/types"
	"github.com/umbracle/ethgo"
)

const (
	AmountFlag = "amount"
)

func CheckIfDirectoryExist(dir string) error {
	if _, err := os.Stat(dir); errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("provided directory '%s' doesn't exist", dir)
	}

	return nil
}

func ValidateSecretFlags(dataDir, config string) error {
	if config == "" {
		if dataDir == "" {
			return polybftsecrets.ErrInvalidParams
		} else {
			return CheckIfDirectoryExist(dataDir)
		}
	}

	return nil
}

// GetAccount resolves secrets manager and returns an account object
func GetAccount(accountDir, accountConfig string) (*wallet.Account, error) {
	// resolve secrets manager instance and allow usage of insecure local secrets manager
	secretsManager, err := polybftsecrets.GetSecretsManager(accountDir, accountConfig, true)
	if err != nil {
		return nil, err
	}

	return wallet.NewAccountFromSecret(secretsManager)
}

// GetAccountFromDir returns an account object from local secrets manager
func GetAccountFromDir(accountDir string) (*wallet.Account, error) {
	return GetAccount(accountDir, "")
}

// GetValidatorInfo queries CustomSupernetManager, StakeManager and RewardPool smart contracts
// to retrieve validator info for given address
func GetValidatorInfo(validatorAddr ethgo.Address, supernetManager, stakeManager types.Address,
	chainID int64, rootRelayer, childRelayer txrelayer.TxRelayer) (*polybft.ValidatorInfo, error) {
	validatorInfo, err := rootHelper.GetValidatorInfo(validatorAddr, supernetManager, stakeManager,
		chainID, rootRelayer)
	if err != nil {
		return nil, err
	}

	withdrawableFn := contractsapi.RewardPool.Abi.GetMethod("pendingRewards")

	encode, err := withdrawableFn.Encode([]interface{}{validatorAddr})
	if err != nil {
		return nil, err
	}

	response, err := childRelayer.Call(ethgo.ZeroAddress, ethgo.Address(contracts.RewardPoolContract), encode)
	if err != nil {
		return nil, err
	}

	withdrawableRewards, err := common.ParseUint256orHex(&response)
	if err != nil {
		return nil, err
	}

	validatorInfo.WithdrawableRewards = withdrawableRewards

	return validatorInfo, nil
}
