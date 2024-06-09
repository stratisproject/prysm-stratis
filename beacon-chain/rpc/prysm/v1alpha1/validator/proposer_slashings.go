package validator

import (
	"context"

	"github.com/stratisproject/prysm-stratis/beacon-chain/core/blocks"
	v "github.com/stratisproject/prysm-stratis/beacon-chain/core/validators"
	"github.com/stratisproject/prysm-stratis/beacon-chain/state"
	ethpb "github.com/stratisproject/prysm-stratis/proto/prysm/v1alpha1"
)

func (vs *Server) getSlashings(ctx context.Context, head state.BeaconState) ([]*ethpb.ProposerSlashing, []*ethpb.AttesterSlashing) {
	proposerSlashings := vs.SlashingsPool.PendingProposerSlashings(ctx, head, false /*noLimit*/)
	validProposerSlashings := make([]*ethpb.ProposerSlashing, 0, len(proposerSlashings))
	for _, slashing := range proposerSlashings {
		_, err := blocks.ProcessProposerSlashing(ctx, head, slashing, v.SlashValidator)
		if err != nil {
			log.WithError(err).Warn("Could not validate proposer slashing for block inclusion")
			continue
		}
		validProposerSlashings = append(validProposerSlashings, slashing)
	}
	attSlashings := vs.SlashingsPool.PendingAttesterSlashings(ctx, head, false /*noLimit*/)
	validAttSlashings := make([]*ethpb.AttesterSlashing, 0, len(attSlashings))
	for _, slashing := range attSlashings {
		_, err := blocks.ProcessAttesterSlashing(ctx, head, slashing, v.SlashValidator)
		if err != nil {
			log.WithError(err).Warn("Could not validate attester slashing for block inclusion")
			continue
		}
		validAttSlashings = append(validAttSlashings, slashing)
	}
	return validProposerSlashings, validAttSlashings
}
