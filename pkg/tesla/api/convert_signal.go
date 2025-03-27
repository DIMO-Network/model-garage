package api

// func Decode(ce cloudevent.RawEvent) ([]vss.Signal, error) {
// 	did, err := cloudevent.DecodeNFTDID(ce.Subject)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to decode subject DID: %w", err)
// 	}

// 	tokenID := did.TokenID
// 	source := ce.Source

// 	baseSignal := vss.Signal{
// 		TokenID: tokenID,
// 		Source:  source,
// 	}

/*
	switch ce.DataVersion {
	case FleetTelemetryDataVersion:
		var td TelemetryData
		if err := json.Unmarshal(ce.Data, &td); err != nil {
			return nil, fmt.Errorf("failed to unmarshal telemetry wrapper: %w", err)
		}

		var batchedSigs []vss.Signal
		var batchedErrs []error
		for i, payload := range td.Payloads {
			var pl protos.Payload
			err := proto.Unmarshal(payload, &pl)
			if err != nil {
				batchedErrs = append(batchedErrs, fmt.Errorf("failed to unmarshal payload at index %d: %w", i, err))
				continue
			}
			sigs, errs := ftconv.ProcessPayload(&pl, tokenID, source)
			batchedSigs = append(batchedSigs, sigs...)
			batchedErrs = append(batchedErrs, errs...)
		}

		if len(batchedErrs) != 0 {
			return nil, convert.ConversionError{
				TokenID:        tokenID,
				Source:         source,
				DecodedSignals: batchedSigs,
				Errors:         batchedErrs,
			}
		}
		return batchedSigs, nil
	default:*/
// 	sigs, errs := SignalsFromTesla(baseSignal, ce.Data)
// 	if len(errs) != 0 {
// 		return nil, convert.ConversionError{
// 			TokenID:        tokenID,
// 			Source:         source,
// 			DecodedSignals: sigs,
// 			Errors:         errs,
// 		}
// 	}

// 	return sigs, nil
// 	// }
// }
