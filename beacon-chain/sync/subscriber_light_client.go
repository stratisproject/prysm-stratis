package sync

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/stratisproject/prysm-stratis/beacon-chain/core/feed"
	statefeed "github.com/stratisproject/prysm-stratis/beacon-chain/core/feed/state"
	lightclientTypes "github.com/stratisproject/prysm-stratis/consensus-types/light-client"
	"google.golang.org/protobuf/proto"
)

func (s *Service) lightClientOptimisticUpdateSubscriber(_ context.Context, msg proto.Message) error {
	update, err := lightclientTypes.NewWrappedOptimisticUpdate(msg)
	if err != nil {
		return err
	}

	attestedHeaderRoot, err := update.AttestedHeader().Beacon().HashTreeRoot()
	if err != nil {
		return err
	}

	log.WithFields(logrus.Fields{
		"attestedSlot":       fmt.Sprintf("%d", update.AttestedHeader().Beacon().Slot),
		"signatureSlot":      fmt.Sprintf("%d", update.SignatureSlot()),
		"attestedHeaderRoot": fmt.Sprintf("%x", attestedHeaderRoot),
	}).Debug("Saving newly received light client optimistic update.")

	s.lcStore.SetLastOptimisticUpdate(update)

	s.cfg.stateNotifier.StateFeed().Send(&feed.Event{
		Type: statefeed.LightClientOptimisticUpdate,
		Data: update,
	})

	return nil
}

func (s *Service) lightClientFinalityUpdateSubscriber(_ context.Context, msg proto.Message) error {
	update, err := lightclientTypes.NewWrappedFinalityUpdate(msg)
	if err != nil {
		return err
	}

	attestedHeaderRoot, err := update.AttestedHeader().Beacon().HashTreeRoot()
	if err != nil {
		return err
	}

	log.WithFields(logrus.Fields{
		"attestedSlot":       fmt.Sprintf("%d", update.AttestedHeader().Beacon().Slot),
		"signatureSlot":      fmt.Sprintf("%d", update.SignatureSlot()),
		"attestedHeaderRoot": fmt.Sprintf("%x", attestedHeaderRoot),
	}).Debug("Saving newly received light client finality update.")

	s.lcStore.SetLastFinalityUpdate(update)

	s.cfg.stateNotifier.StateFeed().Send(&feed.Event{
		Type: statefeed.LightClientFinalityUpdate,
		Data: update,
	})

	return nil
}
