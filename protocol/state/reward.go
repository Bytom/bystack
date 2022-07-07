package state

import (
	"encoding/hex"

	"github.com/bytom/bystack/consensus"
	"github.com/bytom/bystack/protocol/bc/types"
)

func (c *Checkpoint) validatorReward() uint64 {
	if pledgeRate := c.pledgeRate(); pledgeRate <= consensus.RewardThreshold {
		return uint64((pledgeRate + consensus.RewardThreshold) * float64(consensus.GetCoinBaseReward(c.Height)))
	}

	return consensus.GetCoinBaseReward(c.Height)
}

// pledgeRate validator's pledge rate
func (c *Checkpoint) pledgeRate() float64 {
	var totalVotes uint64
	for _, vote := range c.Votes {
		totalVotes += vote
	}

	totalSupply := consensus.GetTotalSupply(c.Height)
	return float64(totalVotes) / float64(totalSupply)
}

// applyValidatorReward calculate the coinbase reward for validator
func (c *Checkpoint) applyValidatorReward(block *types.Block) {
	validatorScript := hex.EncodeToString(block.Transactions[0].Outputs[0].ControlProgram)
	for _, tx := range block.Transactions {
		c.Rewards[validatorScript] += tx.Fee()
	}

	c.Rewards[validatorScript] += c.validatorReward()
}
