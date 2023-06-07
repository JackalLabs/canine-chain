package wasmbinding

// STUB
// Remember to add in the sender
// Another dispatcher function that can be used to organise dispatching
// the storage module's messages in a different file

// func (m *CustomMessenger) DispatchStorageMsg(ctx sdk.Context, contractAddr sdk.AccAddress, contractIBCPortID string, msg wasmvmtypes.CosmosMsg, sender string) ([]sdk.Event, [][]byte, error) {
// 	if msg.Custom != nil {
// 		// only handle the happy path where this is really posting files
// 		// leave everything else for the wrapped version
// 		var contractMsg bindings.JackalMsg
// 		if err := json.Unmarshal(msg.Custom, &contractMsg); err != nil {
// 			return nil, nil, errorsmod.Wrap(err, "Jackal msg")
// 		}
// 		// TO DO
// 		// Replace with storage module's messages
// 		if contractMsg.MakeRoot != nil {
// 			return m.makeRoot(ctx, contractAddr, contractMsg.MakeRoot, sender) // put a storage msg here
// 		}

// 	}
// 	return m.wrapped.DispatchMsg(ctx, contractAddr, contractIBCPortID, msg, sender)
// }
