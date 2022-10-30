<!-- This file is auto-generated. Please do not modify it yourself. -->
# Protobuf Documentation
<a name="top"></a>

## Table of Contents

- [dsig/form.proto](#dsig/form.proto)
    - [Form](#jackaldao.canine.dsig.Form)
  
- [dsig/params.proto](#dsig/params.proto)
    - [Params](#jackaldao.canine.dsig.Params)
  
- [dsig/user_uploads.proto](#dsig/user_uploads.proto)
    - [UserUploads](#jackaldao.canine.dsig.UserUploads)
  
- [dsig/genesis.proto](#dsig/genesis.proto)
    - [GenesisState](#jackaldao.canine.dsig.GenesisState)
  
- [dsig/query.proto](#dsig/query.proto)
    - [QueryAllFormRequest](#jackaldao.canine.dsig.QueryAllFormRequest)
    - [QueryAllFormResponse](#jackaldao.canine.dsig.QueryAllFormResponse)
    - [QueryAllUserUploadsRequest](#jackaldao.canine.dsig.QueryAllUserUploadsRequest)
    - [QueryAllUserUploadsResponse](#jackaldao.canine.dsig.QueryAllUserUploadsResponse)
    - [QueryGetFormRequest](#jackaldao.canine.dsig.QueryGetFormRequest)
    - [QueryGetFormResponse](#jackaldao.canine.dsig.QueryGetFormResponse)
    - [QueryGetUserUploadsRequest](#jackaldao.canine.dsig.QueryGetUserUploadsRequest)
    - [QueryGetUserUploadsResponse](#jackaldao.canine.dsig.QueryGetUserUploadsResponse)
    - [QueryParamsRequest](#jackaldao.canine.dsig.QueryParamsRequest)
    - [QueryParamsResponse](#jackaldao.canine.dsig.QueryParamsResponse)
  
    - [Query](#jackaldao.canine.dsig.Query)
  
- [dsig/tx.proto](#dsig/tx.proto)
    - [MsgCreateform](#jackaldao.canine.dsig.MsgCreateform)
    - [MsgCreateformResponse](#jackaldao.canine.dsig.MsgCreateformResponse)
    - [MsgSignform](#jackaldao.canine.dsig.MsgSignform)
    - [MsgSignformResponse](#jackaldao.canine.dsig.MsgSignformResponse)
    - [MsgUploadfile](#jackaldao.canine.dsig.MsgUploadfile)
    - [MsgUploadfileResponse](#jackaldao.canine.dsig.MsgUploadfileResponse)
  
    - [Msg](#jackaldao.canine.dsig.Msg)
  
- [filetree/files.proto](#filetree/files.proto)
    - [Files](#jackaldao.canine.filetree.Files)
  
- [filetree/params.proto](#filetree/params.proto)
    - [Params](#jackaldao.canine.filetree.Params)
  
- [filetree/pubkey.proto](#filetree/pubkey.proto)
    - [Pubkey](#jackaldao.canine.filetree.Pubkey)
  
- [filetree/genesis.proto](#filetree/genesis.proto)
    - [GenesisState](#jackaldao.canine.filetree.GenesisState)
  
- [filetree/query.proto](#filetree/query.proto)
    - [QueryAllFilesRequest](#jackaldao.canine.filetree.QueryAllFilesRequest)
    - [QueryAllFilesResponse](#jackaldao.canine.filetree.QueryAllFilesResponse)
    - [QueryAllPubkeyRequest](#jackaldao.canine.filetree.QueryAllPubkeyRequest)
    - [QueryAllPubkeyResponse](#jackaldao.canine.filetree.QueryAllPubkeyResponse)
    - [QueryDecryptRequest](#jackaldao.canine.filetree.QueryDecryptRequest)
    - [QueryDecryptResponse](#jackaldao.canine.filetree.QueryDecryptResponse)
    - [QueryEncryptRequest](#jackaldao.canine.filetree.QueryEncryptRequest)
    - [QueryEncryptResponse](#jackaldao.canine.filetree.QueryEncryptResponse)
    - [QueryGetFilesRequest](#jackaldao.canine.filetree.QueryGetFilesRequest)
    - [QueryGetFilesResponse](#jackaldao.canine.filetree.QueryGetFilesResponse)
    - [QueryGetKeyRequest](#jackaldao.canine.filetree.QueryGetKeyRequest)
    - [QueryGetKeysRequest](#jackaldao.canine.filetree.QueryGetKeysRequest)
    - [QueryGetKeysResponse](#jackaldao.canine.filetree.QueryGetKeysResponse)
    - [QueryGetPubkeyRequest](#jackaldao.canine.filetree.QueryGetPubkeyRequest)
    - [QueryGetPubkeyResponse](#jackaldao.canine.filetree.QueryGetPubkeyResponse)
    - [QueryParamsRequest](#jackaldao.canine.filetree.QueryParamsRequest)
    - [QueryParamsResponse](#jackaldao.canine.filetree.QueryParamsResponse)
  
    - [Query](#jackaldao.canine.filetree.Query)
  
- [filetree/tx.proto](#filetree/tx.proto)
    - [MsgAddEditors](#jackaldao.canine.filetree.MsgAddEditors)
    - [MsgAddEditorsResponse](#jackaldao.canine.filetree.MsgAddEditorsResponse)
    - [MsgAddViewers](#jackaldao.canine.filetree.MsgAddViewers)
    - [MsgAddViewersResponse](#jackaldao.canine.filetree.MsgAddViewersResponse)
    - [MsgChangeOwner](#jackaldao.canine.filetree.MsgChangeOwner)
    - [MsgChangeOwnerResponse](#jackaldao.canine.filetree.MsgChangeOwnerResponse)
    - [MsgDeleteFile](#jackaldao.canine.filetree.MsgDeleteFile)
    - [MsgDeleteFileResponse](#jackaldao.canine.filetree.MsgDeleteFileResponse)
    - [MsgInitAll](#jackaldao.canine.filetree.MsgInitAll)
    - [MsgInitAllResponse](#jackaldao.canine.filetree.MsgInitAllResponse)
    - [MsgMakeRoot](#jackaldao.canine.filetree.MsgMakeRoot)
    - [MsgMakeRootResponse](#jackaldao.canine.filetree.MsgMakeRootResponse)
    - [MsgPostFile](#jackaldao.canine.filetree.MsgPostFile)
    - [MsgPostFileResponse](#jackaldao.canine.filetree.MsgPostFileResponse)
    - [MsgPostkey](#jackaldao.canine.filetree.MsgPostkey)
    - [MsgPostkeyResponse](#jackaldao.canine.filetree.MsgPostkeyResponse)
    - [MsgRemoveEditors](#jackaldao.canine.filetree.MsgRemoveEditors)
    - [MsgRemoveEditorsResponse](#jackaldao.canine.filetree.MsgRemoveEditorsResponse)
    - [MsgRemoveViewers](#jackaldao.canine.filetree.MsgRemoveViewers)
    - [MsgRemoveViewersResponse](#jackaldao.canine.filetree.MsgRemoveViewersResponse)
    - [MsgResetEditors](#jackaldao.canine.filetree.MsgResetEditors)
    - [MsgResetEditorsResponse](#jackaldao.canine.filetree.MsgResetEditorsResponse)
    - [MsgResetViewers](#jackaldao.canine.filetree.MsgResetViewers)
    - [MsgResetViewersResponse](#jackaldao.canine.filetree.MsgResetViewersResponse)
  
    - [Msg](#jackaldao.canine.filetree.Msg)
  
- [jklmint/params.proto](#jklmint/params.proto)
    - [Params](#jackaldao.canine.jklmint.Params)
  
- [jklmint/genesis.proto](#jklmint/genesis.proto)
    - [GenesisState](#jackaldao.canine.jklmint.GenesisState)
  
- [jklmint/query.proto](#jklmint/query.proto)
    - [QueryGetInflationRequest](#jackaldao.canine.jklmint.QueryGetInflationRequest)
    - [QueryGetInflationResponse](#jackaldao.canine.jklmint.QueryGetInflationResponse)
    - [QueryInflationRequest](#jackaldao.canine.jklmint.QueryInflationRequest)
    - [QueryInflationResponse](#jackaldao.canine.jklmint.QueryInflationResponse)
    - [QueryParamsRequest](#jackaldao.canine.jklmint.QueryParamsRequest)
    - [QueryParamsResponse](#jackaldao.canine.jklmint.QueryParamsResponse)
  
    - [Query](#jackaldao.canine.jklmint.Query)
  
- [jklmint/tx.proto](#jklmint/tx.proto)
    - [Msg](#jackaldao.canine.jklmint.Msg)
  
- [lp/params.proto](#lp/params.proto)
    - [Params](#jackaldao.canine.lp.Params)
  
- [lp/l_pool.proto](#lp/l_pool.proto)
    - [LPool](#jackaldao.canine.lp.LPool)
  
- [lp/l_provider_record.proto](#lp/l_provider_record.proto)
    - [LProviderRecord](#jackaldao.canine.lp.LProviderRecord)
  
- [lp/genesis.proto](#lp/genesis.proto)
    - [GenesisState](#jackaldao.canine.lp.GenesisState)
  
- [lp/query.proto](#lp/query.proto)
    - [QueryAllLPoolRequest](#jackaldao.canine.lp.QueryAllLPoolRequest)
    - [QueryAllLPoolResponse](#jackaldao.canine.lp.QueryAllLPoolResponse)
    - [QueryEstimateContributionRequest](#jackaldao.canine.lp.QueryEstimateContributionRequest)
    - [QueryEstimateContributionResponse](#jackaldao.canine.lp.QueryEstimateContributionResponse)
    - [QueryEstimatePoolRemoveRequest](#jackaldao.canine.lp.QueryEstimatePoolRemoveRequest)
    - [QueryEstimatePoolRemoveResponse](#jackaldao.canine.lp.QueryEstimatePoolRemoveResponse)
    - [QueryEstimateSwapInRequest](#jackaldao.canine.lp.QueryEstimateSwapInRequest)
    - [QueryEstimateSwapInResponse](#jackaldao.canine.lp.QueryEstimateSwapInResponse)
    - [QueryEstimateSwapOutRequest](#jackaldao.canine.lp.QueryEstimateSwapOutRequest)
    - [QueryEstimateSwapOutResponse](#jackaldao.canine.lp.QueryEstimateSwapOutResponse)
    - [QueryGetLPoolRequest](#jackaldao.canine.lp.QueryGetLPoolRequest)
    - [QueryGetLPoolResponse](#jackaldao.canine.lp.QueryGetLPoolResponse)
    - [QueryGetLProviderRecordRequest](#jackaldao.canine.lp.QueryGetLProviderRecordRequest)
    - [QueryGetLProviderRecordResponse](#jackaldao.canine.lp.QueryGetLProviderRecordResponse)
    - [QueryListRecordsFromPoolRequest](#jackaldao.canine.lp.QueryListRecordsFromPoolRequest)
    - [QueryListRecordsFromPoolResponse](#jackaldao.canine.lp.QueryListRecordsFromPoolResponse)
    - [QueryMakeValidPairRequest](#jackaldao.canine.lp.QueryMakeValidPairRequest)
    - [QueryMakeValidPairResponse](#jackaldao.canine.lp.QueryMakeValidPairResponse)
    - [QueryParamsRequest](#jackaldao.canine.lp.QueryParamsRequest)
    - [QueryParamsResponse](#jackaldao.canine.lp.QueryParamsResponse)
  
    - [Query](#jackaldao.canine.lp.Query)
  
- [lp/tx.proto](#lp/tx.proto)
    - [MsgCreateLPool](#jackaldao.canine.lp.MsgCreateLPool)
    - [MsgCreateLPoolResponse](#jackaldao.canine.lp.MsgCreateLPoolResponse)
    - [MsgExitPool](#jackaldao.canine.lp.MsgExitPool)
    - [MsgExitPoolResponse](#jackaldao.canine.lp.MsgExitPoolResponse)
    - [MsgJoinPool](#jackaldao.canine.lp.MsgJoinPool)
    - [MsgJoinPoolResponse](#jackaldao.canine.lp.MsgJoinPoolResponse)
    - [MsgSwap](#jackaldao.canine.lp.MsgSwap)
    - [MsgSwapResponse](#jackaldao.canine.lp.MsgSwapResponse)
  
    - [Msg](#jackaldao.canine.lp.Msg)
  
- [notifications/params.proto](#notifications/params.proto)
    - [Params](#jackaldao.canine.notifications.Params)
  
- [notifications/notifications.proto](#notifications/notifications.proto)
    - [Notifications](#jackaldao.canine.notifications.Notifications)
  
- [notifications/noti_counter.proto](#notifications/noti_counter.proto)
    - [NotiCounter](#jackaldao.canine.notifications.NotiCounter)
  
- [notifications/genesis.proto](#notifications/genesis.proto)
    - [GenesisState](#jackaldao.canine.notifications.GenesisState)
  
- [notifications/query.proto](#notifications/query.proto)
    - [QueryAllNotiCounterRequest](#jackaldao.canine.notifications.QueryAllNotiCounterRequest)
    - [QueryAllNotiCounterResponse](#jackaldao.canine.notifications.QueryAllNotiCounterResponse)
    - [QueryAllNotificationsRequest](#jackaldao.canine.notifications.QueryAllNotificationsRequest)
    - [QueryAllNotificationsResponse](#jackaldao.canine.notifications.QueryAllNotificationsResponse)
    - [QueryFilteredNotificationsRequest](#jackaldao.canine.notifications.QueryFilteredNotificationsRequest)
    - [QueryFilteredNotificationsResponse](#jackaldao.canine.notifications.QueryFilteredNotificationsResponse)
    - [QueryGetNotiCounterRequest](#jackaldao.canine.notifications.QueryGetNotiCounterRequest)
    - [QueryGetNotiCounterResponse](#jackaldao.canine.notifications.QueryGetNotiCounterResponse)
    - [QueryGetNotificationsRequest](#jackaldao.canine.notifications.QueryGetNotificationsRequest)
    - [QueryGetNotificationsResponse](#jackaldao.canine.notifications.QueryGetNotificationsResponse)
    - [QueryParamsRequest](#jackaldao.canine.notifications.QueryParamsRequest)
    - [QueryParamsResponse](#jackaldao.canine.notifications.QueryParamsResponse)
  
    - [Query](#jackaldao.canine.notifications.Query)
  
- [notifications/tx.proto](#notifications/tx.proto)
    - [MsgAddSenders](#jackaldao.canine.notifications.MsgAddSenders)
    - [MsgAddSendersResponse](#jackaldao.canine.notifications.MsgAddSendersResponse)
    - [MsgCreateNotifications](#jackaldao.canine.notifications.MsgCreateNotifications)
    - [MsgCreateNotificationsResponse](#jackaldao.canine.notifications.MsgCreateNotificationsResponse)
    - [MsgDeleteNotifications](#jackaldao.canine.notifications.MsgDeleteNotifications)
    - [MsgDeleteNotificationsResponse](#jackaldao.canine.notifications.MsgDeleteNotificationsResponse)
    - [MsgSetCounter](#jackaldao.canine.notifications.MsgSetCounter)
    - [MsgSetCounterResponse](#jackaldao.canine.notifications.MsgSetCounterResponse)
    - [MsgUpdateNotifications](#jackaldao.canine.notifications.MsgUpdateNotifications)
    - [MsgUpdateNotificationsResponse](#jackaldao.canine.notifications.MsgUpdateNotificationsResponse)
  
    - [Msg](#jackaldao.canine.notifications.Msg)
  
- [rns/bids.proto](#rns/bids.proto)
    - [Bids](#jackaldao.canine.rns.Bids)
  
- [rns/forsale.proto](#rns/forsale.proto)
    - [Forsale](#jackaldao.canine.rns.Forsale)
  
- [rns/params.proto](#rns/params.proto)
    - [Params](#jackaldao.canine.rns.Params)
  
- [rns/whois.proto](#rns/whois.proto)
    - [Whois](#jackaldao.canine.rns.Whois)
  
- [rns/names.proto](#rns/names.proto)
    - [Names](#jackaldao.canine.rns.Names)
  
- [rns/init.proto](#rns/init.proto)
    - [Init](#jackaldao.canine.rns.Init)
  
- [rns/genesis.proto](#rns/genesis.proto)
    - [GenesisState](#jackaldao.canine.rns.GenesisState)
  
- [rns/query.proto](#rns/query.proto)
    - [QueryAllBidsRequest](#jackaldao.canine.rns.QueryAllBidsRequest)
    - [QueryAllBidsResponse](#jackaldao.canine.rns.QueryAllBidsResponse)
    - [QueryAllForsaleRequest](#jackaldao.canine.rns.QueryAllForsaleRequest)
    - [QueryAllForsaleResponse](#jackaldao.canine.rns.QueryAllForsaleResponse)
    - [QueryAllInitRequest](#jackaldao.canine.rns.QueryAllInitRequest)
    - [QueryAllInitResponse](#jackaldao.canine.rns.QueryAllInitResponse)
    - [QueryAllNamesRequest](#jackaldao.canine.rns.QueryAllNamesRequest)
    - [QueryAllNamesResponse](#jackaldao.canine.rns.QueryAllNamesResponse)
    - [QueryAllWhoisRequest](#jackaldao.canine.rns.QueryAllWhoisRequest)
    - [QueryAllWhoisResponse](#jackaldao.canine.rns.QueryAllWhoisResponse)
    - [QueryGetBidsRequest](#jackaldao.canine.rns.QueryGetBidsRequest)
    - [QueryGetBidsResponse](#jackaldao.canine.rns.QueryGetBidsResponse)
    - [QueryGetForsaleRequest](#jackaldao.canine.rns.QueryGetForsaleRequest)
    - [QueryGetForsaleResponse](#jackaldao.canine.rns.QueryGetForsaleResponse)
    - [QueryGetInitRequest](#jackaldao.canine.rns.QueryGetInitRequest)
    - [QueryGetInitResponse](#jackaldao.canine.rns.QueryGetInitResponse)
    - [QueryGetNamesRequest](#jackaldao.canine.rns.QueryGetNamesRequest)
    - [QueryGetNamesResponse](#jackaldao.canine.rns.QueryGetNamesResponse)
    - [QueryGetWhoisRequest](#jackaldao.canine.rns.QueryGetWhoisRequest)
    - [QueryGetWhoisResponse](#jackaldao.canine.rns.QueryGetWhoisResponse)
    - [QueryListOwnedNamesRequest](#jackaldao.canine.rns.QueryListOwnedNamesRequest)
    - [QueryListOwnedNamesResponse](#jackaldao.canine.rns.QueryListOwnedNamesResponse)
    - [QueryParamsRequest](#jackaldao.canine.rns.QueryParamsRequest)
    - [QueryParamsResponse](#jackaldao.canine.rns.QueryParamsResponse)
  
    - [Query](#jackaldao.canine.rns.Query)
  
- [rns/tx.proto](#rns/tx.proto)
    - [MsgAcceptBid](#jackaldao.canine.rns.MsgAcceptBid)
    - [MsgAcceptBidResponse](#jackaldao.canine.rns.MsgAcceptBidResponse)
    - [MsgAddRecord](#jackaldao.canine.rns.MsgAddRecord)
    - [MsgAddRecordResponse](#jackaldao.canine.rns.MsgAddRecordResponse)
    - [MsgBid](#jackaldao.canine.rns.MsgBid)
    - [MsgBidResponse](#jackaldao.canine.rns.MsgBidResponse)
    - [MsgBuy](#jackaldao.canine.rns.MsgBuy)
    - [MsgBuyResponse](#jackaldao.canine.rns.MsgBuyResponse)
    - [MsgCancelBid](#jackaldao.canine.rns.MsgCancelBid)
    - [MsgCancelBidResponse](#jackaldao.canine.rns.MsgCancelBidResponse)
    - [MsgDelRecord](#jackaldao.canine.rns.MsgDelRecord)
    - [MsgDelRecordResponse](#jackaldao.canine.rns.MsgDelRecordResponse)
    - [MsgDelist](#jackaldao.canine.rns.MsgDelist)
    - [MsgDelistResponse](#jackaldao.canine.rns.MsgDelistResponse)
    - [MsgInit](#jackaldao.canine.rns.MsgInit)
    - [MsgInitResponse](#jackaldao.canine.rns.MsgInitResponse)
    - [MsgList](#jackaldao.canine.rns.MsgList)
    - [MsgListResponse](#jackaldao.canine.rns.MsgListResponse)
    - [MsgRegister](#jackaldao.canine.rns.MsgRegister)
    - [MsgRegisterResponse](#jackaldao.canine.rns.MsgRegisterResponse)
    - [MsgTransfer](#jackaldao.canine.rns.MsgTransfer)
    - [MsgTransferResponse](#jackaldao.canine.rns.MsgTransferResponse)
  
    - [Msg](#jackaldao.canine.rns.Msg)
  
- [storage/active_deals.proto](#storage/active_deals.proto)
    - [ActiveDeals](#jackaldao.canine.storage.ActiveDeals)
  
- [storage/client_usage.proto](#storage/client_usage.proto)
    - [ClientUsage](#jackaldao.canine.storage.ClientUsage)
  
- [storage/contracts.proto](#storage/contracts.proto)
    - [Contracts](#jackaldao.canine.storage.Contracts)
  
- [storage/fid_cid.proto](#storage/fid_cid.proto)
    - [FidCid](#jackaldao.canine.storage.FidCid)
  
- [storage/params.proto](#storage/params.proto)
    - [Params](#jackaldao.canine.storage.Params)
  
- [storage/proofs.proto](#storage/proofs.proto)
    - [Proofs](#jackaldao.canine.storage.Proofs)
  
- [storage/providers.proto](#storage/providers.proto)
    - [Providers](#jackaldao.canine.storage.Providers)
  
- [storage/pay_blocks.proto](#storage/pay_blocks.proto)
    - [PayBlocks](#jackaldao.canine.storage.PayBlocks)
  
- [storage/strays.proto](#storage/strays.proto)
    - [Strays](#jackaldao.canine.storage.Strays)
  
- [storage/genesis.proto](#storage/genesis.proto)
    - [GenesisState](#jackaldao.canine.storage.GenesisState)
  
- [storage/query.proto](#storage/query.proto)
    - [QueryAllActiveDealsRequest](#jackaldao.canine.storage.QueryAllActiveDealsRequest)
    - [QueryAllActiveDealsResponse](#jackaldao.canine.storage.QueryAllActiveDealsResponse)
    - [QueryAllClientUsageRequest](#jackaldao.canine.storage.QueryAllClientUsageRequest)
    - [QueryAllClientUsageResponse](#jackaldao.canine.storage.QueryAllClientUsageResponse)
    - [QueryAllContractsRequest](#jackaldao.canine.storage.QueryAllContractsRequest)
    - [QueryAllContractsResponse](#jackaldao.canine.storage.QueryAllContractsResponse)
    - [QueryAllFidCidRequest](#jackaldao.canine.storage.QueryAllFidCidRequest)
    - [QueryAllFidCidResponse](#jackaldao.canine.storage.QueryAllFidCidResponse)
    - [QueryAllPayBlocksRequest](#jackaldao.canine.storage.QueryAllPayBlocksRequest)
    - [QueryAllPayBlocksResponse](#jackaldao.canine.storage.QueryAllPayBlocksResponse)
    - [QueryAllProofsRequest](#jackaldao.canine.storage.QueryAllProofsRequest)
    - [QueryAllProofsResponse](#jackaldao.canine.storage.QueryAllProofsResponse)
    - [QueryAllProvidersRequest](#jackaldao.canine.storage.QueryAllProvidersRequest)
    - [QueryAllProvidersResponse](#jackaldao.canine.storage.QueryAllProvidersResponse)
    - [QueryAllStraysRequest](#jackaldao.canine.storage.QueryAllStraysRequest)
    - [QueryAllStraysResponse](#jackaldao.canine.storage.QueryAllStraysResponse)
    - [QueryFindFileRequest](#jackaldao.canine.storage.QueryFindFileRequest)
    - [QueryFindFileResponse](#jackaldao.canine.storage.QueryFindFileResponse)
    - [QueryFreespaceRequest](#jackaldao.canine.storage.QueryFreespaceRequest)
    - [QueryFreespaceResponse](#jackaldao.canine.storage.QueryFreespaceResponse)
    - [QueryGetActiveDealsRequest](#jackaldao.canine.storage.QueryGetActiveDealsRequest)
    - [QueryGetActiveDealsResponse](#jackaldao.canine.storage.QueryGetActiveDealsResponse)
    - [QueryGetClientFreeSpaceRequest](#jackaldao.canine.storage.QueryGetClientFreeSpaceRequest)
    - [QueryGetClientFreeSpaceResponse](#jackaldao.canine.storage.QueryGetClientFreeSpaceResponse)
    - [QueryGetClientUsageRequest](#jackaldao.canine.storage.QueryGetClientUsageRequest)
    - [QueryGetClientUsageResponse](#jackaldao.canine.storage.QueryGetClientUsageResponse)
    - [QueryGetContractsRequest](#jackaldao.canine.storage.QueryGetContractsRequest)
    - [QueryGetContractsResponse](#jackaldao.canine.storage.QueryGetContractsResponse)
    - [QueryGetFidCidRequest](#jackaldao.canine.storage.QueryGetFidCidRequest)
    - [QueryGetFidCidResponse](#jackaldao.canine.storage.QueryGetFidCidResponse)
    - [QueryGetPayBlocksRequest](#jackaldao.canine.storage.QueryGetPayBlocksRequest)
    - [QueryGetPayBlocksResponse](#jackaldao.canine.storage.QueryGetPayBlocksResponse)
    - [QueryGetPayDataRequest](#jackaldao.canine.storage.QueryGetPayDataRequest)
    - [QueryGetPayDataResponse](#jackaldao.canine.storage.QueryGetPayDataResponse)
    - [QueryGetProofsRequest](#jackaldao.canine.storage.QueryGetProofsRequest)
    - [QueryGetProofsResponse](#jackaldao.canine.storage.QueryGetProofsResponse)
    - [QueryGetProvidersRequest](#jackaldao.canine.storage.QueryGetProvidersRequest)
    - [QueryGetProvidersResponse](#jackaldao.canine.storage.QueryGetProvidersResponse)
    - [QueryGetStraysRequest](#jackaldao.canine.storage.QueryGetStraysRequest)
    - [QueryGetStraysResponse](#jackaldao.canine.storage.QueryGetStraysResponse)
    - [QueryParamsRequest](#jackaldao.canine.storage.QueryParamsRequest)
    - [QueryParamsResponse](#jackaldao.canine.storage.QueryParamsResponse)
  
    - [Query](#jackaldao.canine.storage.Query)
  
- [storage/tx.proto](#storage/tx.proto)
    - [MsgBuyStorage](#jackaldao.canine.storage.MsgBuyStorage)
    - [MsgBuyStorageResponse](#jackaldao.canine.storage.MsgBuyStorageResponse)
    - [MsgCancelContract](#jackaldao.canine.storage.MsgCancelContract)
    - [MsgCancelContractResponse](#jackaldao.canine.storage.MsgCancelContractResponse)
    - [MsgClaimStray](#jackaldao.canine.storage.MsgClaimStray)
    - [MsgClaimStrayResponse](#jackaldao.canine.storage.MsgClaimStrayResponse)
    - [MsgCreateActiveDeals](#jackaldao.canine.storage.MsgCreateActiveDeals)
    - [MsgCreateActiveDealsResponse](#jackaldao.canine.storage.MsgCreateActiveDealsResponse)
    - [MsgCreateContracts](#jackaldao.canine.storage.MsgCreateContracts)
    - [MsgCreateContractsResponse](#jackaldao.canine.storage.MsgCreateContractsResponse)
    - [MsgCreateProofs](#jackaldao.canine.storage.MsgCreateProofs)
    - [MsgCreateProofsResponse](#jackaldao.canine.storage.MsgCreateProofsResponse)
    - [MsgCreateProviders](#jackaldao.canine.storage.MsgCreateProviders)
    - [MsgCreateProvidersResponse](#jackaldao.canine.storage.MsgCreateProvidersResponse)
    - [MsgDeleteActiveDeals](#jackaldao.canine.storage.MsgDeleteActiveDeals)
    - [MsgDeleteActiveDealsResponse](#jackaldao.canine.storage.MsgDeleteActiveDealsResponse)
    - [MsgDeleteContracts](#jackaldao.canine.storage.MsgDeleteContracts)
    - [MsgDeleteContractsResponse](#jackaldao.canine.storage.MsgDeleteContractsResponse)
    - [MsgDeleteProofs](#jackaldao.canine.storage.MsgDeleteProofs)
    - [MsgDeleteProofsResponse](#jackaldao.canine.storage.MsgDeleteProofsResponse)
    - [MsgDeleteProviders](#jackaldao.canine.storage.MsgDeleteProviders)
    - [MsgDeleteProvidersResponse](#jackaldao.canine.storage.MsgDeleteProvidersResponse)
    - [MsgInitProvider](#jackaldao.canine.storage.MsgInitProvider)
    - [MsgInitProviderResponse](#jackaldao.canine.storage.MsgInitProviderResponse)
    - [MsgItem](#jackaldao.canine.storage.MsgItem)
    - [MsgItemResponse](#jackaldao.canine.storage.MsgItemResponse)
    - [MsgPostContract](#jackaldao.canine.storage.MsgPostContract)
    - [MsgPostContractResponse](#jackaldao.canine.storage.MsgPostContractResponse)
    - [MsgPostproof](#jackaldao.canine.storage.MsgPostproof)
    - [MsgPostproofResponse](#jackaldao.canine.storage.MsgPostproofResponse)
    - [MsgSetProviderIP](#jackaldao.canine.storage.MsgSetProviderIP)
    - [MsgSetProviderIPResponse](#jackaldao.canine.storage.MsgSetProviderIPResponse)
    - [MsgSetProviderTotalspace](#jackaldao.canine.storage.MsgSetProviderTotalspace)
    - [MsgSetProviderTotalspaceResponse](#jackaldao.canine.storage.MsgSetProviderTotalspaceResponse)
    - [MsgSignContract](#jackaldao.canine.storage.MsgSignContract)
    - [MsgSignContractResponse](#jackaldao.canine.storage.MsgSignContractResponse)
    - [MsgUpdateActiveDeals](#jackaldao.canine.storage.MsgUpdateActiveDeals)
    - [MsgUpdateActiveDealsResponse](#jackaldao.canine.storage.MsgUpdateActiveDealsResponse)
    - [MsgUpdateContracts](#jackaldao.canine.storage.MsgUpdateContracts)
    - [MsgUpdateContractsResponse](#jackaldao.canine.storage.MsgUpdateContractsResponse)
    - [MsgUpdateProofs](#jackaldao.canine.storage.MsgUpdateProofs)
    - [MsgUpdateProofsResponse](#jackaldao.canine.storage.MsgUpdateProofsResponse)
    - [MsgUpdateProviders](#jackaldao.canine.storage.MsgUpdateProviders)
    - [MsgUpdateProvidersResponse](#jackaldao.canine.storage.MsgUpdateProvidersResponse)
  
    - [Msg](#jackaldao.canine.storage.Msg)
  
- [Scalar Value Types](#scalar-value-types)



<a name="dsig/form.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## dsig/form.proto



<a name="jackaldao.canine.dsig.Form"></a>

### Form



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `ffid` | [string](#string) |  |  |
| `cid` | [string](#string) |  |  |
| `fid` | [string](#string) |  |  |
| `signees` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="dsig/params.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## dsig/params.proto



<a name="jackaldao.canine.dsig.Params"></a>

### Params
Params defines the parameters for the module.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="dsig/user_uploads.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## dsig/user_uploads.proto



<a name="jackaldao.canine.dsig.UserUploads"></a>

### UserUploads



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `fid` | [string](#string) |  |  |
| `cid` | [string](#string) |  |  |
| `createdAt` | [int64](#int64) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="dsig/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## dsig/genesis.proto



<a name="jackaldao.canine.dsig.GenesisState"></a>

### GenesisState
GenesisState defines the dsig module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#jackaldao.canine.dsig.Params) |  |  |
| `userUploadsList` | [UserUploads](#jackaldao.canine.dsig.UserUploads) | repeated |  |
| `formList` | [Form](#jackaldao.canine.dsig.Form) | repeated | this line is used by starport scaffolding # genesis/proto/state |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="dsig/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## dsig/query.proto



<a name="jackaldao.canine.dsig.QueryAllFormRequest"></a>

### QueryAllFormRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="jackaldao.canine.dsig.QueryAllFormResponse"></a>

### QueryAllFormResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `form` | [Form](#jackaldao.canine.dsig.Form) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="jackaldao.canine.dsig.QueryAllUserUploadsRequest"></a>

### QueryAllUserUploadsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="jackaldao.canine.dsig.QueryAllUserUploadsResponse"></a>

### QueryAllUserUploadsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `userUploads` | [UserUploads](#jackaldao.canine.dsig.UserUploads) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="jackaldao.canine.dsig.QueryGetFormRequest"></a>

### QueryGetFormRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `ffid` | [string](#string) |  |  |






<a name="jackaldao.canine.dsig.QueryGetFormResponse"></a>

### QueryGetFormResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `form` | [Form](#jackaldao.canine.dsig.Form) |  |  |






<a name="jackaldao.canine.dsig.QueryGetUserUploadsRequest"></a>

### QueryGetUserUploadsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `fid` | [string](#string) |  |  |






<a name="jackaldao.canine.dsig.QueryGetUserUploadsResponse"></a>

### QueryGetUserUploadsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `userUploads` | [UserUploads](#jackaldao.canine.dsig.UserUploads) |  |  |






<a name="jackaldao.canine.dsig.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is request type for the Query/Params RPC method.






<a name="jackaldao.canine.dsig.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is response type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#jackaldao.canine.dsig.Params) |  | params holds all the parameters of this module. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="jackaldao.canine.dsig.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#jackaldao.canine.dsig.QueryParamsRequest) | [QueryParamsResponse](#jackaldao.canine.dsig.QueryParamsResponse) | Parameters queries the parameters of the module. | GET|/dsig/dsig/params|
| `UserUploads` | [QueryGetUserUploadsRequest](#jackaldao.canine.dsig.QueryGetUserUploadsRequest) | [QueryGetUserUploadsResponse](#jackaldao.canine.dsig.QueryGetUserUploadsResponse) | Queries a UserUploads by index. | GET|/dsig/dsig/user_uploads/{fid}|
| `UserUploadsAll` | [QueryAllUserUploadsRequest](#jackaldao.canine.dsig.QueryAllUserUploadsRequest) | [QueryAllUserUploadsResponse](#jackaldao.canine.dsig.QueryAllUserUploadsResponse) | Queries a list of UserUploads items. | GET|/dsig/dsig/user_uploads|
| `Form` | [QueryGetFormRequest](#jackaldao.canine.dsig.QueryGetFormRequest) | [QueryGetFormResponse](#jackaldao.canine.dsig.QueryGetFormResponse) | Queries a Form by index. | GET|/dsig/dsig/form/{ffid}|
| `FormAll` | [QueryAllFormRequest](#jackaldao.canine.dsig.QueryAllFormRequest) | [QueryAllFormResponse](#jackaldao.canine.dsig.QueryAllFormResponse) | Queries a list of Form items. | GET|/dsig/dsig/form|

 <!-- end services -->



<a name="dsig/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## dsig/tx.proto



<a name="jackaldao.canine.dsig.MsgCreateform"></a>

### MsgCreateform



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `fid` | [string](#string) |  |  |
| `signees` | [string](#string) |  |  |






<a name="jackaldao.canine.dsig.MsgCreateformResponse"></a>

### MsgCreateformResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `ffid` | [string](#string) |  |  |






<a name="jackaldao.canine.dsig.MsgSignform"></a>

### MsgSignform



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `ffid` | [string](#string) |  |  |
| `vote` | [int32](#int32) |  |  |






<a name="jackaldao.canine.dsig.MsgSignformResponse"></a>

### MsgSignformResponse







<a name="jackaldao.canine.dsig.MsgUploadfile"></a>

### MsgUploadfile



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `fid` | [string](#string) |  |  |






<a name="jackaldao.canine.dsig.MsgUploadfileResponse"></a>

### MsgUploadfileResponse






 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="jackaldao.canine.dsig.Msg"></a>

### Msg
Msg defines the Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Uploadfile` | [MsgUploadfile](#jackaldao.canine.dsig.MsgUploadfile) | [MsgUploadfileResponse](#jackaldao.canine.dsig.MsgUploadfileResponse) |  | |
| `Createform` | [MsgCreateform](#jackaldao.canine.dsig.MsgCreateform) | [MsgCreateformResponse](#jackaldao.canine.dsig.MsgCreateformResponse) |  | |
| `Signform` | [MsgSignform](#jackaldao.canine.dsig.MsgSignform) | [MsgSignformResponse](#jackaldao.canine.dsig.MsgSignformResponse) | this line is used by starport scaffolding # proto/tx/rpc | |

 <!-- end services -->



<a name="filetree/files.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## filetree/files.proto



<a name="jackaldao.canine.filetree.Files"></a>

### Files



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |
| `contents` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |
| `viewingAccess` | [string](#string) |  |  |
| `editAccess` | [string](#string) |  |  |
| `trackingNumber` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="filetree/params.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## filetree/params.proto



<a name="jackaldao.canine.filetree.Params"></a>

### Params
Params defines the parameters for the module.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="filetree/pubkey.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## filetree/pubkey.proto



<a name="jackaldao.canine.filetree.Pubkey"></a>

### Pubkey



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |
| `key` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="filetree/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## filetree/genesis.proto



<a name="jackaldao.canine.filetree.GenesisState"></a>

### GenesisState
GenesisState defines the filetree module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#jackaldao.canine.filetree.Params) |  |  |
| `filesList` | [Files](#jackaldao.canine.filetree.Files) | repeated |  |
| `pubkeyList` | [Pubkey](#jackaldao.canine.filetree.Pubkey) | repeated | this line is used by starport scaffolding # genesis/proto/state |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="filetree/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## filetree/query.proto



<a name="jackaldao.canine.filetree.QueryAllFilesRequest"></a>

### QueryAllFilesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="jackaldao.canine.filetree.QueryAllFilesResponse"></a>

### QueryAllFilesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `files` | [Files](#jackaldao.canine.filetree.Files) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="jackaldao.canine.filetree.QueryAllPubkeyRequest"></a>

### QueryAllPubkeyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="jackaldao.canine.filetree.QueryAllPubkeyResponse"></a>

### QueryAllPubkeyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pubkey` | [Pubkey](#jackaldao.canine.filetree.Pubkey) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="jackaldao.canine.filetree.QueryDecryptRequest"></a>

### QueryDecryptRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `message` | [string](#string) |  |  |






<a name="jackaldao.canine.filetree.QueryDecryptResponse"></a>

### QueryDecryptResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `data` | [string](#string) |  |  |






<a name="jackaldao.canine.filetree.QueryEncryptRequest"></a>

### QueryEncryptRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |
| `message` | [string](#string) |  |  |






<a name="jackaldao.canine.filetree.QueryEncryptResponse"></a>

### QueryEncryptResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `encryptionData` | [string](#string) |  |  |






<a name="jackaldao.canine.filetree.QueryGetFilesRequest"></a>

### QueryGetFilesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |
| `ownerAddress` | [string](#string) |  |  |






<a name="jackaldao.canine.filetree.QueryGetFilesResponse"></a>

### QueryGetFilesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `files` | [Files](#jackaldao.canine.filetree.Files) |  |  |






<a name="jackaldao.canine.filetree.QueryGetKeyRequest"></a>

### QueryGetKeyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `filepath` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |






<a name="jackaldao.canine.filetree.QueryGetKeysRequest"></a>

### QueryGetKeysRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `hashpath` | [string](#string) |  |  |






<a name="jackaldao.canine.filetree.QueryGetKeysResponse"></a>

### QueryGetKeysResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `keys` | [string](#string) |  |  |






<a name="jackaldao.canine.filetree.QueryGetPubkeyRequest"></a>

### QueryGetPubkeyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |






<a name="jackaldao.canine.filetree.QueryGetPubkeyResponse"></a>

### QueryGetPubkeyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pubkey` | [Pubkey](#jackaldao.canine.filetree.Pubkey) |  |  |






<a name="jackaldao.canine.filetree.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is request type for the Query/Params RPC method.






<a name="jackaldao.canine.filetree.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is response type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#jackaldao.canine.filetree.Params) |  | params holds all the parameters of this module. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="jackaldao.canine.filetree.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#jackaldao.canine.filetree.QueryParamsRequest) | [QueryParamsResponse](#jackaldao.canine.filetree.QueryParamsResponse) | Parameters queries the parameters of the module. | GET|/jackaldao/canine/filetree/params|
| `Encrypt` | [QueryEncryptRequest](#jackaldao.canine.filetree.QueryEncryptRequest) | [QueryEncryptResponse](#jackaldao.canine.filetree.QueryEncryptResponse) | Queries a list of Encrypt items. | GET|/jackal-dao/canine/filetree/encrypt/{address}/{message}|
| `Decrypt` | [QueryDecryptRequest](#jackaldao.canine.filetree.QueryDecryptRequest) | [QueryDecryptResponse](#jackaldao.canine.filetree.QueryDecryptResponse) | Queries a list of Decrypt items. | GET|/jackal-dao/canine/filetree/decrypt/{message}|
| `Files` | [QueryGetFilesRequest](#jackaldao.canine.filetree.QueryGetFilesRequest) | [QueryGetFilesResponse](#jackaldao.canine.filetree.QueryGetFilesResponse) | Queries a Files by index. | GET|/jackal-dao/canine/filetree/files/{address}/{ownerAddress}|
| `FilesAll` | [QueryAllFilesRequest](#jackaldao.canine.filetree.QueryAllFilesRequest) | [QueryAllFilesResponse](#jackaldao.canine.filetree.QueryAllFilesResponse) | Queries a list of Files items. | GET|/jackal-dao/canine/filetree/files|
| `GetKeys` | [QueryGetKeysRequest](#jackaldao.canine.filetree.QueryGetKeysRequest) | [QueryGetKeysResponse](#jackaldao.canine.filetree.QueryGetKeysResponse) | Queries a list of GetKeys items. | GET|/jackal-dao/canine/filetree/get_keys/{hashpath}|
| `Pubkey` | [QueryGetPubkeyRequest](#jackaldao.canine.filetree.QueryGetPubkeyRequest) | [QueryGetPubkeyResponse](#jackaldao.canine.filetree.QueryGetPubkeyResponse) | Queries a Pubkey by index. | GET|/jackal-dao/canine/filetree/pubkey/{address}|
| `PubkeyAll` | [QueryAllPubkeyRequest](#jackaldao.canine.filetree.QueryAllPubkeyRequest) | [QueryAllPubkeyResponse](#jackaldao.canine.filetree.QueryAllPubkeyResponse) | Queries a list of Pubkey items. | GET|/jackal-dao/canine/filetree/pubkey|

 <!-- end services -->



<a name="filetree/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## filetree/tx.proto



<a name="jackaldao.canine.filetree.MsgAddEditors"></a>

### MsgAddEditors



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `editorIds` | [string](#string) |  |  |
| `editorKeys` | [string](#string) |  |  |
| `address` | [string](#string) |  |  |
| `fileowner` | [string](#string) |  |  |
| `notifyEditors` | [string](#string) |  |  |
| `notiForEditors` | [string](#string) |  |  |






<a name="jackaldao.canine.filetree.MsgAddEditorsResponse"></a>

### MsgAddEditorsResponse







<a name="jackaldao.canine.filetree.MsgAddViewers"></a>

### MsgAddViewers



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `viewerIds` | [string](#string) |  |  |
| `viewerKeys` | [string](#string) |  |  |
| `address` | [string](#string) |  |  |
| `fileowner` | [string](#string) |  |  |
| `notifyViewers` | [string](#string) |  |  |
| `notiForViewers` | [string](#string) |  |  |






<a name="jackaldao.canine.filetree.MsgAddViewersResponse"></a>

### MsgAddViewersResponse







<a name="jackaldao.canine.filetree.MsgChangeOwner"></a>

### MsgChangeOwner



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `address` | [string](#string) |  |  |
| `fileOwner` | [string](#string) |  |  |
| `newOwner` | [string](#string) |  |  |






<a name="jackaldao.canine.filetree.MsgChangeOwnerResponse"></a>

### MsgChangeOwnerResponse







<a name="jackaldao.canine.filetree.MsgDeleteFile"></a>

### MsgDeleteFile



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `hashPath` | [string](#string) |  |  |
| `account` | [string](#string) |  |  |






<a name="jackaldao.canine.filetree.MsgDeleteFileResponse"></a>

### MsgDeleteFileResponse







<a name="jackaldao.canine.filetree.MsgInitAll"></a>

### MsgInitAll



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `pubkey` | [string](#string) |  |  |






<a name="jackaldao.canine.filetree.MsgInitAllResponse"></a>

### MsgInitAllResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `name` | [string](#string) |  |  |






<a name="jackaldao.canine.filetree.MsgMakeRoot"></a>

### MsgMakeRoot



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `account` | [string](#string) |  |  |
| `rootHashPath` | [string](#string) |  |  |
| `contents` | [string](#string) |  |  |
| `editors` | [string](#string) |  |  |
| `viewers` | [string](#string) |  |  |
| `trackingNumber` | [string](#string) |  |  |






<a name="jackaldao.canine.filetree.MsgMakeRootResponse"></a>

### MsgMakeRootResponse







<a name="jackaldao.canine.filetree.MsgPostFile"></a>

### MsgPostFile



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `account` | [string](#string) |  |  |
| `hashParent` | [string](#string) |  |  |
| `hashChild` | [string](#string) |  |  |
| `contents` | [string](#string) |  |  |
| `viewers` | [string](#string) |  |  |
| `editors` | [string](#string) |  |  |
| `trackingNumber` | [string](#string) |  |  |
| `viewersToNotify` | [string](#string) |  |  |
| `editorsToNotify` | [string](#string) |  |  |
| `notiForViewers` | [string](#string) |  |  |
| `notiForEditors` | [string](#string) |  |  |






<a name="jackaldao.canine.filetree.MsgPostFileResponse"></a>

### MsgPostFileResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `path` | [string](#string) |  |  |






<a name="jackaldao.canine.filetree.MsgPostkey"></a>

### MsgPostkey



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `key` | [string](#string) |  |  |






<a name="jackaldao.canine.filetree.MsgPostkeyResponse"></a>

### MsgPostkeyResponse







<a name="jackaldao.canine.filetree.MsgRemoveEditors"></a>

### MsgRemoveEditors



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `editorIds` | [string](#string) |  |  |
| `address` | [string](#string) |  |  |
| `fileowner` | [string](#string) |  |  |
| `notifyEditors` | [string](#string) |  |  |
| `notiForEditors` | [string](#string) |  |  |






<a name="jackaldao.canine.filetree.MsgRemoveEditorsResponse"></a>

### MsgRemoveEditorsResponse







<a name="jackaldao.canine.filetree.MsgRemoveViewers"></a>

### MsgRemoveViewers



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `viewerIds` | [string](#string) |  |  |
| `address` | [string](#string) |  |  |
| `fileowner` | [string](#string) |  |  |
| `notifyviewers` | [string](#string) |  |  |
| `notiForViewers` | [string](#string) |  |  |






<a name="jackaldao.canine.filetree.MsgRemoveViewersResponse"></a>

### MsgRemoveViewersResponse







<a name="jackaldao.canine.filetree.MsgResetEditors"></a>

### MsgResetEditors



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `address` | [string](#string) |  |  |
| `fileowner` | [string](#string) |  |  |






<a name="jackaldao.canine.filetree.MsgResetEditorsResponse"></a>

### MsgResetEditorsResponse







<a name="jackaldao.canine.filetree.MsgResetViewers"></a>

### MsgResetViewers



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `address` | [string](#string) |  |  |
| `fileowner` | [string](#string) |  |  |






<a name="jackaldao.canine.filetree.MsgResetViewersResponse"></a>

### MsgResetViewersResponse






 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="jackaldao.canine.filetree.Msg"></a>

### Msg
Msg defines the Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `PostFile` | [MsgPostFile](#jackaldao.canine.filetree.MsgPostFile) | [MsgPostFileResponse](#jackaldao.canine.filetree.MsgPostFileResponse) |  | |
| `AddViewers` | [MsgAddViewers](#jackaldao.canine.filetree.MsgAddViewers) | [MsgAddViewersResponse](#jackaldao.canine.filetree.MsgAddViewersResponse) |  | |
| `Postkey` | [MsgPostkey](#jackaldao.canine.filetree.MsgPostkey) | [MsgPostkeyResponse](#jackaldao.canine.filetree.MsgPostkeyResponse) |  | |
| `DeleteFile` | [MsgDeleteFile](#jackaldao.canine.filetree.MsgDeleteFile) | [MsgDeleteFileResponse](#jackaldao.canine.filetree.MsgDeleteFileResponse) |  | |
| `InitAll` | [MsgInitAll](#jackaldao.canine.filetree.MsgInitAll) | [MsgInitAllResponse](#jackaldao.canine.filetree.MsgInitAllResponse) |  | |
| `RemoveViewers` | [MsgRemoveViewers](#jackaldao.canine.filetree.MsgRemoveViewers) | [MsgRemoveViewersResponse](#jackaldao.canine.filetree.MsgRemoveViewersResponse) |  | |
| `MakeRoot` | [MsgMakeRoot](#jackaldao.canine.filetree.MsgMakeRoot) | [MsgMakeRootResponse](#jackaldao.canine.filetree.MsgMakeRootResponse) |  | |
| `AddEditors` | [MsgAddEditors](#jackaldao.canine.filetree.MsgAddEditors) | [MsgAddEditorsResponse](#jackaldao.canine.filetree.MsgAddEditorsResponse) |  | |
| `RemoveEditors` | [MsgRemoveEditors](#jackaldao.canine.filetree.MsgRemoveEditors) | [MsgRemoveEditorsResponse](#jackaldao.canine.filetree.MsgRemoveEditorsResponse) |  | |
| `ResetEditors` | [MsgResetEditors](#jackaldao.canine.filetree.MsgResetEditors) | [MsgResetEditorsResponse](#jackaldao.canine.filetree.MsgResetEditorsResponse) |  | |
| `ResetViewers` | [MsgResetViewers](#jackaldao.canine.filetree.MsgResetViewers) | [MsgResetViewersResponse](#jackaldao.canine.filetree.MsgResetViewersResponse) |  | |
| `ChangeOwner` | [MsgChangeOwner](#jackaldao.canine.filetree.MsgChangeOwner) | [MsgChangeOwnerResponse](#jackaldao.canine.filetree.MsgChangeOwnerResponse) | this line is used by starport scaffolding # proto/tx/rpc | |

 <!-- end services -->



<a name="jklmint/params.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## jklmint/params.proto



<a name="jackaldao.canine.jklmint.Params"></a>

### Params
Params defines the parameters for the module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `mintDenom` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="jklmint/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## jklmint/genesis.proto



<a name="jackaldao.canine.jklmint.GenesisState"></a>

### GenesisState
GenesisState defines the jklmint module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#jackaldao.canine.jklmint.Params) |  | this line is used by starport scaffolding # genesis/proto/state |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="jklmint/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## jklmint/query.proto



<a name="jackaldao.canine.jklmint.QueryGetInflationRequest"></a>

### QueryGetInflationRequest







<a name="jackaldao.canine.jklmint.QueryGetInflationResponse"></a>

### QueryGetInflationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `inflation` | [string](#string) |  |  |






<a name="jackaldao.canine.jklmint.QueryInflationRequest"></a>

### QueryInflationRequest
QueryInflationRequest is the request type for the Query/Inflation RPC method.






<a name="jackaldao.canine.jklmint.QueryInflationResponse"></a>

### QueryInflationResponse
QueryInflationResponse is the response type for the Query/Inflation RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `inflation` | [bytes](#bytes) |  | inflation is the current minting inflation value. |






<a name="jackaldao.canine.jklmint.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is request type for the Query/Params RPC method.






<a name="jackaldao.canine.jklmint.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is response type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#jackaldao.canine.jklmint.Params) |  | params holds all the parameters of this module. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="jackaldao.canine.jklmint.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#jackaldao.canine.jklmint.QueryParamsRequest) | [QueryParamsResponse](#jackaldao.canine.jklmint.QueryParamsResponse) | Parameters queries the parameters of the module. | GET|/cosmos/mint/v1beta1/params|
| `Inflation` | [QueryInflationRequest](#jackaldao.canine.jklmint.QueryInflationRequest) | [QueryInflationResponse](#jackaldao.canine.jklmint.QueryInflationResponse) | Inflation returns the current minting inflation value. | GET|/cosmos/mint/v1beta1/inflation|
| `GetInflation` | [QueryGetInflationRequest](#jackaldao.canine.jklmint.QueryGetInflationRequest) | [QueryGetInflationResponse](#jackaldao.canine.jklmint.QueryGetInflationResponse) | Queries a list of GetInflation items. | GET|/jackal-dao/canine/jklmint/get_inflation|

 <!-- end services -->



<a name="jklmint/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## jklmint/tx.proto


 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="jackaldao.canine.jklmint.Msg"></a>

### Msg
Msg defines the Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |

 <!-- end services -->



<a name="lp/params.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## lp/params.proto



<a name="jackaldao.canine.lp.Params"></a>

### Params
Params defines the parameters for the module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `MinInitPoolDeposit` | [uint64](#uint64) |  |  |
| `MaxPoolDenomCount` | [uint32](#uint32) |  |  |
| `LPTokenUnit` | [uint32](#uint32) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="lp/l_pool.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## lp/l_pool.proto



<a name="jackaldao.canine.lp.LPool"></a>

### LPool



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `index` | [string](#string) |  |  |
| `name` | [string](#string) |  |  |
| `coins` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| `aMM_Id` | [uint32](#uint32) |  |  |
| `swap_fee_multi` | [string](#string) |  | sdk.Dec in string representation |
| `min_lock_duration` | [int64](#int64) |  |  |
| `penalty_multi` | [string](#string) |  | sdk.Dec in string representation |
| `lptoken_denom` | [string](#string) |  |  |
| `LPTokenBalance` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="lp/l_provider_record.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## lp/l_provider_record.proto



<a name="jackaldao.canine.lp.LProviderRecord"></a>

### LProviderRecord
LProviderRecord is a record of a liquidity provider depositing to a pool.
It is used to enforce withdraw panelty and calculate rewards collected. 
This record is created only once when provider contributes to a pool and
only updated after witdrawal or deposit.
It is deleted when the provider burns all of the liquidity pool token.
This is stored at KVStore with 
	{LProviderRecordKeyPrefix}{poolName}{provider} key.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `provider` | [string](#string) |  | Provider is the account address of the provider. |
| `poolName` | [string](#string) |  | A pool that the provider contributed to. |
| `unlockTime` | [string](#string) |  | Burning LP token is locked for certain duration the after provider deposits to the pool. Unlock time is updated every succeeding deposits. The provider can burn their LP token during lock time but has to take certain amount of panelty. Unlock time is blocktime + lockDuration at time of contribution. |
| `lockDuration` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="lp/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## lp/genesis.proto



<a name="jackaldao.canine.lp.GenesisState"></a>

### GenesisState
GenesisState defines the lp module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#jackaldao.canine.lp.Params) |  |  |
| `lPoolList` | [LPool](#jackaldao.canine.lp.LPool) | repeated |  |
| `lProviderRecordList` | [LProviderRecord](#jackaldao.canine.lp.LProviderRecord) | repeated | this line is used by starport scaffolding # genesis/proto/state |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="lp/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## lp/query.proto



<a name="jackaldao.canine.lp.QueryAllLPoolRequest"></a>

### QueryAllLPoolRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="jackaldao.canine.lp.QueryAllLPoolResponse"></a>

### QueryAllLPoolResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `lPool` | [LPool](#jackaldao.canine.lp.LPool) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="jackaldao.canine.lp.QueryEstimateContributionRequest"></a>

### QueryEstimateContributionRequest
Estimate amount of coins to deposit to get desired amount of LPToken


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `poolName` | [string](#string) |  |  |
| `desiredAmount` | [string](#string) |  |  |






<a name="jackaldao.canine.lp.QueryEstimateContributionResponse"></a>

### QueryEstimateContributionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `coins` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="jackaldao.canine.lp.QueryEstimatePoolRemoveRequest"></a>

### QueryEstimatePoolRemoveRequest
Estimate pool coins returned by burning LPToken


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `poolName` | [string](#string) |  |  |
| `amount` | [string](#string) |  |  |






<a name="jackaldao.canine.lp.QueryEstimatePoolRemoveResponse"></a>

### QueryEstimatePoolRemoveResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `coins` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="jackaldao.canine.lp.QueryEstimateSwapInRequest"></a>

### QueryEstimateSwapInRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `poolName` | [string](#string) |  |  |
| `outputCoins` | [string](#string) |  |  |






<a name="jackaldao.canine.lp.QueryEstimateSwapInResponse"></a>

### QueryEstimateSwapInResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `inputCoins` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="jackaldao.canine.lp.QueryEstimateSwapOutRequest"></a>

### QueryEstimateSwapOutRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `poolName` | [string](#string) |  |  |
| `inputCoin` | [string](#string) |  |  |






<a name="jackaldao.canine.lp.QueryEstimateSwapOutResponse"></a>

### QueryEstimateSwapOutResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `outputCoin` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="jackaldao.canine.lp.QueryGetLPoolRequest"></a>

### QueryGetLPoolRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `index` | [string](#string) |  |  |






<a name="jackaldao.canine.lp.QueryGetLPoolResponse"></a>

### QueryGetLPoolResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `lPool` | [LPool](#jackaldao.canine.lp.LPool) |  |  |






<a name="jackaldao.canine.lp.QueryGetLProviderRecordRequest"></a>

### QueryGetLProviderRecordRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `provider` | [string](#string) |  |  |
| `poolName` | [string](#string) |  |  |






<a name="jackaldao.canine.lp.QueryGetLProviderRecordResponse"></a>

### QueryGetLProviderRecordResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `lProviderRecord` | [LProviderRecord](#jackaldao.canine.lp.LProviderRecord) |  |  |






<a name="jackaldao.canine.lp.QueryListRecordsFromPoolRequest"></a>

### QueryListRecordsFromPoolRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `poolName` | [string](#string) |  |  |






<a name="jackaldao.canine.lp.QueryListRecordsFromPoolResponse"></a>

### QueryListRecordsFromPoolResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `records` | [LProviderRecord](#jackaldao.canine.lp.LProviderRecord) | repeated |  |






<a name="jackaldao.canine.lp.QueryMakeValidPairRequest"></a>

### QueryMakeValidPairRequest
Query amount of coins to deposit to make a valid liquidity pair


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `poolName` | [string](#string) |  |  |
| `coin` | [string](#string) |  |  |






<a name="jackaldao.canine.lp.QueryMakeValidPairResponse"></a>

### QueryMakeValidPairResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `coin` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="jackaldao.canine.lp.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is request type for the Query/Params RPC method.






<a name="jackaldao.canine.lp.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is response type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#jackaldao.canine.lp.Params) |  | params holds all the parameters of this module. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="jackaldao.canine.lp.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#jackaldao.canine.lp.QueryParamsRequest) | [QueryParamsResponse](#jackaldao.canine.lp.QueryParamsResponse) | Parameters queries the parameters of the module. | GET|/jackal-dao/canine/lp/params|
| `LPool` | [QueryGetLPoolRequest](#jackaldao.canine.lp.QueryGetLPoolRequest) | [QueryGetLPoolResponse](#jackaldao.canine.lp.QueryGetLPoolResponse) | Queries a LPool by index. | GET|/jackal-dao/canine/lp/l_pool/{index}|
| `LPoolAll` | [QueryAllLPoolRequest](#jackaldao.canine.lp.QueryAllLPoolRequest) | [QueryAllLPoolResponse](#jackaldao.canine.lp.QueryAllLPoolResponse) | Queries a list of LPool items. | GET|/jackal-dao/canine/lp/l_pool|
| `LProviderRecord` | [QueryGetLProviderRecordRequest](#jackaldao.canine.lp.QueryGetLProviderRecordRequest) | [QueryGetLProviderRecordResponse](#jackaldao.canine.lp.QueryGetLProviderRecordResponse) | Queries a LProviderRecord by pool name and provider address. | GET|/jackal-dao/canine/lp/l_provider_record/{poolName}/{provider}|
| `EstimateSwapOut` | [QueryEstimateSwapOutRequest](#jackaldao.canine.lp.QueryEstimateSwapOutRequest) | [QueryEstimateSwapOutResponse](#jackaldao.canine.lp.QueryEstimateSwapOutResponse) | Estimate coin output from a swap. | GET|/jackal-dao/canine/lp/estimate_out/{poolName}/{inputCoin}|
| `EstimateSwapIn` | [QueryEstimateSwapInRequest](#jackaldao.canine.lp.QueryEstimateSwapInRequest) | [QueryEstimateSwapInResponse](#jackaldao.canine.lp.QueryEstimateSwapInResponse) | Estimate coin input to get desired coin output from a swap. | GET|/jackal-dao/canine/lp/estimate_in/{poolName}/{outputCoins}|
| `EstimateContribution` | [QueryEstimateContributionRequest](#jackaldao.canine.lp.QueryEstimateContributionRequest) | [QueryEstimateContributionResponse](#jackaldao.canine.lp.QueryEstimateContributionResponse) | Estimate coin inputs to get desired amount of LPToken. | GET|/jackal-dao/canine/lp/estimate_contribution/{poolName}/{desiredAmount}|
| `MakeValidPair` | [QueryMakeValidPairRequest](#jackaldao.canine.lp.QueryMakeValidPairRequest) | [QueryMakeValidPairResponse](#jackaldao.canine.lp.QueryMakeValidPairResponse) | Query coins to deposit to make valid liquidity pair. Input one coin and get other coins to deposit to make liquidity pair. | GET|/jackal-dao/canine/lp/make_pair/{poolName}/{coin}|
| `EstimatePoolRemove` | [QueryEstimatePoolRemoveRequest](#jackaldao.canine.lp.QueryEstimatePoolRemoveRequest) | [QueryEstimatePoolRemoveResponse](#jackaldao.canine.lp.QueryEstimatePoolRemoveResponse) | Estimate amoutn of coins returned by burning a LPToken. | GET|/jackal-dao/canine/lp/estimate_pool_remove/{amount}|
| `ListRecordsFromPool` | [QueryListRecordsFromPoolRequest](#jackaldao.canine.lp.QueryListRecordsFromPoolRequest) | [QueryListRecordsFromPoolResponse](#jackaldao.canine.lp.QueryListRecordsFromPoolResponse) | Queries a list of ListRecordsFromPool items. | GET|/jackal-dao/canine/lp/list_records_from_pool/{poolName}|

 <!-- end services -->



<a name="lp/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## lp/tx.proto



<a name="jackaldao.canine.lp.MsgCreateLPool"></a>

### MsgCreateLPool



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `coins` | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) | repeated | Creator needs to deposit coins to create pool. Input format: "{amount0}{denomination},...,{amountN}{denominationN}" |
| `amm_Id` | [uint32](#uint32) |  | AMM model id used to balance the pool. |
| `swap_fee_multi` | [string](#string) |  | Swap fee charged to swap transaction (swap x swap_fee_multi). Is converted to type sdk.Dec |
| `min_lock_duration` | [int64](#int64) |  | Liquidity pool token (LPToken) lock duration in seconds. Liquidity provider's LPToken is locked when they contribute to a pool. Penalty is applied when LPToken is burned before the lock duration is over. |
| `penalty_multi` | [string](#string) |  | Penalty multiplier applied to LPToken when provider wishes to burn LPToken before lock duration is over (LPToken x penalty_multi) -> pool_tokens. Is converted to type sdk.Dec |






<a name="jackaldao.canine.lp.MsgCreateLPoolResponse"></a>

### MsgCreateLPoolResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [string](#string) |  | Pool id |






<a name="jackaldao.canine.lp.MsgExitPool"></a>

### MsgExitPool



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `poolName` | [string](#string) |  |  |
| `shares` | [int64](#int64) |  |  |






<a name="jackaldao.canine.lp.MsgExitPoolResponse"></a>

### MsgExitPoolResponse







<a name="jackaldao.canine.lp.MsgJoinPool"></a>

### MsgJoinPool



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `poolName` | [string](#string) |  |  |
| `coins` | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) | repeated | Input format: "{amount0}{denomination},...,{amountN}{denominationN}" |
| `lockDuration` | [int64](#int64) |  | The contributor can choose lock duration |






<a name="jackaldao.canine.lp.MsgJoinPoolResponse"></a>

### MsgJoinPoolResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `shares` | [uint64](#uint64) |  | Amount of shares given to pool contributor. |






<a name="jackaldao.canine.lp.MsgSwap"></a>

### MsgSwap
Swap a coin for other coin in a liquidity pool.
Swap fee is deducted from coin input before swap output is computed.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `poolName` | [string](#string) |  |  |
| `coin_input` | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) |  | Coin input to swap in return for the desired coin. Coin output from swap is determined by the input amount. |
| `min_coin_output` | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) |  | Minimum coin output from this swap. Swap will not proceed if computed swap output is less than this amount. |






<a name="jackaldao.canine.lp.MsgSwapResponse"></a>

### MsgSwapResponse






 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="jackaldao.canine.lp.Msg"></a>

### Msg
Msg defines the Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `CreateLPool` | [MsgCreateLPool](#jackaldao.canine.lp.MsgCreateLPool) | [MsgCreateLPoolResponse](#jackaldao.canine.lp.MsgCreateLPoolResponse) |  | |
| `JoinPool` | [MsgJoinPool](#jackaldao.canine.lp.MsgJoinPool) | [MsgJoinPoolResponse](#jackaldao.canine.lp.MsgJoinPoolResponse) |  | |
| `ExitPool` | [MsgExitPool](#jackaldao.canine.lp.MsgExitPool) | [MsgExitPoolResponse](#jackaldao.canine.lp.MsgExitPoolResponse) |  | |
| `Swap` | [MsgSwap](#jackaldao.canine.lp.MsgSwap) | [MsgSwapResponse](#jackaldao.canine.lp.MsgSwapResponse) | this line is used by starport scaffolding # proto/tx/rpc | |

 <!-- end services -->



<a name="notifications/params.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## notifications/params.proto



<a name="jackaldao.canine.notifications.Params"></a>

### Params
Params defines the parameters for the module.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="notifications/notifications.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## notifications/notifications.proto



<a name="jackaldao.canine.notifications.Notifications"></a>

### Notifications



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `count` | [uint64](#uint64) |  |  |
| `notification` | [string](#string) |  |  |
| `address` | [string](#string) |  |  |
| `sender` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="notifications/noti_counter.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## notifications/noti_counter.proto



<a name="jackaldao.canine.notifications.NotiCounter"></a>

### NotiCounter



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |
| `counter` | [uint64](#uint64) |  |  |
| `permittedSenders` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="notifications/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## notifications/genesis.proto



<a name="jackaldao.canine.notifications.GenesisState"></a>

### GenesisState
GenesisState defines the notifications module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#jackaldao.canine.notifications.Params) |  |  |
| `notificationsList` | [Notifications](#jackaldao.canine.notifications.Notifications) | repeated |  |
| `notiCounterList` | [NotiCounter](#jackaldao.canine.notifications.NotiCounter) | repeated | this line is used by starport scaffolding # genesis/proto/state |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="notifications/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## notifications/query.proto



<a name="jackaldao.canine.notifications.QueryAllNotiCounterRequest"></a>

### QueryAllNotiCounterRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="jackaldao.canine.notifications.QueryAllNotiCounterResponse"></a>

### QueryAllNotiCounterResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `notiCounter` | [NotiCounter](#jackaldao.canine.notifications.NotiCounter) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="jackaldao.canine.notifications.QueryAllNotificationsRequest"></a>

### QueryAllNotificationsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="jackaldao.canine.notifications.QueryAllNotificationsResponse"></a>

### QueryAllNotificationsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `notifications` | [Notifications](#jackaldao.canine.notifications.Notifications) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="jackaldao.canine.notifications.QueryFilteredNotificationsRequest"></a>

### QueryFilteredNotificationsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |






<a name="jackaldao.canine.notifications.QueryFilteredNotificationsResponse"></a>

### QueryFilteredNotificationsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `notifications` | [string](#string) |  | could turn it back to 'repeated Notifications notifications' if needed |






<a name="jackaldao.canine.notifications.QueryGetNotiCounterRequest"></a>

### QueryGetNotiCounterRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |






<a name="jackaldao.canine.notifications.QueryGetNotiCounterResponse"></a>

### QueryGetNotiCounterResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `notiCounter` | [NotiCounter](#jackaldao.canine.notifications.NotiCounter) |  |  |






<a name="jackaldao.canine.notifications.QueryGetNotificationsRequest"></a>

### QueryGetNotificationsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `count` | [uint64](#uint64) |  |  |
| `address` | [string](#string) |  |  |






<a name="jackaldao.canine.notifications.QueryGetNotificationsResponse"></a>

### QueryGetNotificationsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `notifications` | [Notifications](#jackaldao.canine.notifications.Notifications) |  |  |






<a name="jackaldao.canine.notifications.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is request type for the Query/Params RPC method.






<a name="jackaldao.canine.notifications.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is response type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#jackaldao.canine.notifications.Params) |  | params holds all the parameters of this module. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="jackaldao.canine.notifications.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#jackaldao.canine.notifications.QueryParamsRequest) | [QueryParamsResponse](#jackaldao.canine.notifications.QueryParamsResponse) | Parameters queries the parameters of the module. | GET|/jackal-dao/canine/notifications/params|
| `Notifications` | [QueryGetNotificationsRequest](#jackaldao.canine.notifications.QueryGetNotificationsRequest) | [QueryGetNotificationsResponse](#jackaldao.canine.notifications.QueryGetNotificationsResponse) | Queries a Notifications by index. | GET|/jackal-dao/canine/notifications/notifications/{count}/{address}|
| `NotificationsAll` | [QueryAllNotificationsRequest](#jackaldao.canine.notifications.QueryAllNotificationsRequest) | [QueryAllNotificationsResponse](#jackaldao.canine.notifications.QueryAllNotificationsResponse) | Queries a list of Notifications items. | GET|/jackal-dao/canine/notifications/notifications|
| `FilteredNotifications` | [QueryFilteredNotificationsRequest](#jackaldao.canine.notifications.QueryFilteredNotificationsRequest) | [QueryFilteredNotificationsResponse](#jackaldao.canine.notifications.QueryFilteredNotificationsResponse) | Queries a list of FilteredNotifications items. | GET|/jackal-dao/canine/notifications/filtered_notifications/{address}|
| `NotiCounter` | [QueryGetNotiCounterRequest](#jackaldao.canine.notifications.QueryGetNotiCounterRequest) | [QueryGetNotiCounterResponse](#jackaldao.canine.notifications.QueryGetNotiCounterResponse) | Queries a NotiCounter by index. | GET|/jackal-dao/canine/notifications/noti_counter/{address}|
| `NotiCounterAll` | [QueryAllNotiCounterRequest](#jackaldao.canine.notifications.QueryAllNotiCounterRequest) | [QueryAllNotiCounterResponse](#jackaldao.canine.notifications.QueryAllNotiCounterResponse) | Queries a list of NotiCounter items. | GET|/jackal-dao/canine/notifications/noti_counter|

 <!-- end services -->



<a name="notifications/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## notifications/tx.proto



<a name="jackaldao.canine.notifications.MsgAddSenders"></a>

### MsgAddSenders



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `senderIds` | [string](#string) |  |  |






<a name="jackaldao.canine.notifications.MsgAddSendersResponse"></a>

### MsgAddSendersResponse







<a name="jackaldao.canine.notifications.MsgCreateNotifications"></a>

### MsgCreateNotifications



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `notification` | [string](#string) |  |  |
| `address` | [string](#string) |  |  |






<a name="jackaldao.canine.notifications.MsgCreateNotificationsResponse"></a>

### MsgCreateNotificationsResponse







<a name="jackaldao.canine.notifications.MsgDeleteNotifications"></a>

### MsgDeleteNotifications



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `count` | [uint64](#uint64) |  |  |






<a name="jackaldao.canine.notifications.MsgDeleteNotificationsResponse"></a>

### MsgDeleteNotificationsResponse







<a name="jackaldao.canine.notifications.MsgSetCounter"></a>

### MsgSetCounter



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |






<a name="jackaldao.canine.notifications.MsgSetCounterResponse"></a>

### MsgSetCounterResponse







<a name="jackaldao.canine.notifications.MsgUpdateNotifications"></a>

### MsgUpdateNotifications



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `count` | [uint64](#uint64) |  |  |
| `notification` | [string](#string) |  |  |
| `address` | [string](#string) |  |  |






<a name="jackaldao.canine.notifications.MsgUpdateNotificationsResponse"></a>

### MsgUpdateNotificationsResponse






 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="jackaldao.canine.notifications.Msg"></a>

### Msg
Msg defines the Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `CreateNotifications` | [MsgCreateNotifications](#jackaldao.canine.notifications.MsgCreateNotifications) | [MsgCreateNotificationsResponse](#jackaldao.canine.notifications.MsgCreateNotificationsResponse) |  | |
| `UpdateNotifications` | [MsgUpdateNotifications](#jackaldao.canine.notifications.MsgUpdateNotifications) | [MsgUpdateNotificationsResponse](#jackaldao.canine.notifications.MsgUpdateNotificationsResponse) |  | |
| `DeleteNotifications` | [MsgDeleteNotifications](#jackaldao.canine.notifications.MsgDeleteNotifications) | [MsgDeleteNotificationsResponse](#jackaldao.canine.notifications.MsgDeleteNotificationsResponse) |  | |
| `SetCounter` | [MsgSetCounter](#jackaldao.canine.notifications.MsgSetCounter) | [MsgSetCounterResponse](#jackaldao.canine.notifications.MsgSetCounterResponse) |  | |
| `AddSenders` | [MsgAddSenders](#jackaldao.canine.notifications.MsgAddSenders) | [MsgAddSendersResponse](#jackaldao.canine.notifications.MsgAddSendersResponse) | this line is used by starport scaffolding # proto/tx/rpc | |

 <!-- end services -->



<a name="rns/bids.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## rns/bids.proto



<a name="jackaldao.canine.rns.Bids"></a>

### Bids



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `index` | [string](#string) |  |  |
| `name` | [string](#string) |  |  |
| `bidder` | [string](#string) |  |  |
| `price` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="rns/forsale.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## rns/forsale.proto



<a name="jackaldao.canine.rns.Forsale"></a>

### Forsale



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `name` | [string](#string) |  |  |
| `price` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="rns/params.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## rns/params.proto



<a name="jackaldao.canine.rns.Params"></a>

### Params
Params defines the parameters for the module.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="rns/whois.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## rns/whois.proto



<a name="jackaldao.canine.rns.Whois"></a>

### Whois



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `index` | [string](#string) |  |  |
| `name` | [string](#string) |  |  |
| `value` | [string](#string) |  |  |
| `data` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="rns/names.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## rns/names.proto



<a name="jackaldao.canine.rns.Names"></a>

### Names



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `name` | [string](#string) |  |  |
| `expires` | [int64](#int64) |  |  |
| `value` | [string](#string) |  |  |
| `data` | [string](#string) |  |  |
| `subdomains` | [Names](#jackaldao.canine.rns.Names) | repeated |  |
| `tld` | [string](#string) |  |  |
| `locked` | [int64](#int64) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="rns/init.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## rns/init.proto



<a name="jackaldao.canine.rns.Init"></a>

### Init



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |
| `complete` | [bool](#bool) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="rns/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## rns/genesis.proto



<a name="jackaldao.canine.rns.GenesisState"></a>

### GenesisState
GenesisState defines the rns module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#jackaldao.canine.rns.Params) |  |  |
| `whoisList` | [Whois](#jackaldao.canine.rns.Whois) | repeated |  |
| `namesList` | [Names](#jackaldao.canine.rns.Names) | repeated |  |
| `bidsList` | [Bids](#jackaldao.canine.rns.Bids) | repeated |  |
| `forsaleList` | [Forsale](#jackaldao.canine.rns.Forsale) | repeated |  |
| `initList` | [Init](#jackaldao.canine.rns.Init) | repeated | this line is used by starport scaffolding # genesis/proto/state |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="rns/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## rns/query.proto



<a name="jackaldao.canine.rns.QueryAllBidsRequest"></a>

### QueryAllBidsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="jackaldao.canine.rns.QueryAllBidsResponse"></a>

### QueryAllBidsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `bids` | [Bids](#jackaldao.canine.rns.Bids) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="jackaldao.canine.rns.QueryAllForsaleRequest"></a>

### QueryAllForsaleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="jackaldao.canine.rns.QueryAllForsaleResponse"></a>

### QueryAllForsaleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `forsale` | [Forsale](#jackaldao.canine.rns.Forsale) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="jackaldao.canine.rns.QueryAllInitRequest"></a>

### QueryAllInitRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="jackaldao.canine.rns.QueryAllInitResponse"></a>

### QueryAllInitResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `init` | [Init](#jackaldao.canine.rns.Init) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="jackaldao.canine.rns.QueryAllNamesRequest"></a>

### QueryAllNamesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="jackaldao.canine.rns.QueryAllNamesResponse"></a>

### QueryAllNamesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `names` | [Names](#jackaldao.canine.rns.Names) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="jackaldao.canine.rns.QueryAllWhoisRequest"></a>

### QueryAllWhoisRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="jackaldao.canine.rns.QueryAllWhoisResponse"></a>

### QueryAllWhoisResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `whois` | [Whois](#jackaldao.canine.rns.Whois) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="jackaldao.canine.rns.QueryGetBidsRequest"></a>

### QueryGetBidsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `index` | [string](#string) |  |  |






<a name="jackaldao.canine.rns.QueryGetBidsResponse"></a>

### QueryGetBidsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `bids` | [Bids](#jackaldao.canine.rns.Bids) |  |  |






<a name="jackaldao.canine.rns.QueryGetForsaleRequest"></a>

### QueryGetForsaleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `name` | [string](#string) |  |  |






<a name="jackaldao.canine.rns.QueryGetForsaleResponse"></a>

### QueryGetForsaleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `forsale` | [Forsale](#jackaldao.canine.rns.Forsale) |  |  |






<a name="jackaldao.canine.rns.QueryGetInitRequest"></a>

### QueryGetInitRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |






<a name="jackaldao.canine.rns.QueryGetInitResponse"></a>

### QueryGetInitResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `init` | [bool](#bool) |  |  |






<a name="jackaldao.canine.rns.QueryGetNamesRequest"></a>

### QueryGetNamesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `index` | [string](#string) |  |  |






<a name="jackaldao.canine.rns.QueryGetNamesResponse"></a>

### QueryGetNamesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `names` | [Names](#jackaldao.canine.rns.Names) |  |  |






<a name="jackaldao.canine.rns.QueryGetWhoisRequest"></a>

### QueryGetWhoisRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `index` | [string](#string) |  |  |






<a name="jackaldao.canine.rns.QueryGetWhoisResponse"></a>

### QueryGetWhoisResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `whois` | [Whois](#jackaldao.canine.rns.Whois) |  |  |






<a name="jackaldao.canine.rns.QueryListOwnedNamesRequest"></a>

### QueryListOwnedNamesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="jackaldao.canine.rns.QueryListOwnedNamesResponse"></a>

### QueryListOwnedNamesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `names` | [Names](#jackaldao.canine.rns.Names) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="jackaldao.canine.rns.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is request type for the Query/Params RPC method.






<a name="jackaldao.canine.rns.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is response type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#jackaldao.canine.rns.Params) |  | params holds all the parameters of this module. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="jackaldao.canine.rns.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#jackaldao.canine.rns.QueryParamsRequest) | [QueryParamsResponse](#jackaldao.canine.rns.QueryParamsResponse) | Parameters queries the parameters of the module. | GET|/jackal-dao/canine/rnsparams|
| `Names` | [QueryGetNamesRequest](#jackaldao.canine.rns.QueryGetNamesRequest) | [QueryGetNamesResponse](#jackaldao.canine.rns.QueryGetNamesResponse) | Queries a Name by index. | GET|/jackal-dao/canine/rnsnames/{index}|
| `NamesAll` | [QueryAllNamesRequest](#jackaldao.canine.rns.QueryAllNamesRequest) | [QueryAllNamesResponse](#jackaldao.canine.rns.QueryAllNamesResponse) | Queries a list of Names. | GET|/jackal-dao/canine/rnsnames|
| `Bids` | [QueryGetBidsRequest](#jackaldao.canine.rns.QueryGetBidsRequest) | [QueryGetBidsResponse](#jackaldao.canine.rns.QueryGetBidsResponse) | Queries a Bid by index. | GET|/jackal-dao/canine/rnsbids/{index}|
| `BidsAll` | [QueryAllBidsRequest](#jackaldao.canine.rns.QueryAllBidsRequest) | [QueryAllBidsResponse](#jackaldao.canine.rns.QueryAllBidsResponse) | Queries a list of Bids. | GET|/jackal-dao/canine/rnsbids|
| `Forsale` | [QueryGetForsaleRequest](#jackaldao.canine.rns.QueryGetForsaleRequest) | [QueryGetForsaleResponse](#jackaldao.canine.rns.QueryGetForsaleResponse) | Queries a Listing by index. | GET|/jackal-dao/canine/rnsforsale/{name}|
| `ForsaleAll` | [QueryAllForsaleRequest](#jackaldao.canine.rns.QueryAllForsaleRequest) | [QueryAllForsaleResponse](#jackaldao.canine.rns.QueryAllForsaleResponse) | Queries all Listings. | GET|/jackal-dao/canine/rnsforsale|
| `Init` | [QueryGetInitRequest](#jackaldao.canine.rns.QueryGetInitRequest) | [QueryGetInitResponse](#jackaldao.canine.rns.QueryGetInitResponse) | Queries a Init by index. | GET|/jackal-dao/canine/rns/init/{address}|
| `InitAll` | [QueryAllInitRequest](#jackaldao.canine.rns.QueryAllInitRequest) | [QueryAllInitResponse](#jackaldao.canine.rns.QueryAllInitResponse) | Queries a list of Init items. | GET|/jackal-dao/canine/rns/init|
| `ListOwnedNames` | [QueryListOwnedNamesRequest](#jackaldao.canine.rns.QueryListOwnedNamesRequest) | [QueryListOwnedNamesResponse](#jackaldao.canine.rns.QueryListOwnedNamesResponse) | Queries a list of ListOwnedNames items. | GET|/jackal-dao/canine/rns/list_owned_names/{address}|

 <!-- end services -->



<a name="rns/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## rns/tx.proto



<a name="jackaldao.canine.rns.MsgAcceptBid"></a>

### MsgAcceptBid



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `name` | [string](#string) |  |  |
| `from` | [string](#string) |  |  |






<a name="jackaldao.canine.rns.MsgAcceptBidResponse"></a>

### MsgAcceptBidResponse







<a name="jackaldao.canine.rns.MsgAddRecord"></a>

### MsgAddRecord



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `name` | [string](#string) |  |  |
| `value` | [string](#string) |  |  |
| `data` | [string](#string) |  |  |
| `record` | [string](#string) |  |  |






<a name="jackaldao.canine.rns.MsgAddRecordResponse"></a>

### MsgAddRecordResponse







<a name="jackaldao.canine.rns.MsgBid"></a>

### MsgBid



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `name` | [string](#string) |  |  |
| `bid` | [string](#string) |  |  |






<a name="jackaldao.canine.rns.MsgBidResponse"></a>

### MsgBidResponse







<a name="jackaldao.canine.rns.MsgBuy"></a>

### MsgBuy



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `name` | [string](#string) |  |  |






<a name="jackaldao.canine.rns.MsgBuyResponse"></a>

### MsgBuyResponse







<a name="jackaldao.canine.rns.MsgCancelBid"></a>

### MsgCancelBid



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `name` | [string](#string) |  |  |






<a name="jackaldao.canine.rns.MsgCancelBidResponse"></a>

### MsgCancelBidResponse







<a name="jackaldao.canine.rns.MsgDelRecord"></a>

### MsgDelRecord



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `name` | [string](#string) |  |  |






<a name="jackaldao.canine.rns.MsgDelRecordResponse"></a>

### MsgDelRecordResponse







<a name="jackaldao.canine.rns.MsgDelist"></a>

### MsgDelist



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `name` | [string](#string) |  |  |






<a name="jackaldao.canine.rns.MsgDelistResponse"></a>

### MsgDelistResponse







<a name="jackaldao.canine.rns.MsgInit"></a>

### MsgInit



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |






<a name="jackaldao.canine.rns.MsgInitResponse"></a>

### MsgInitResponse







<a name="jackaldao.canine.rns.MsgList"></a>

### MsgList



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `name` | [string](#string) |  |  |
| `price` | [string](#string) |  |  |






<a name="jackaldao.canine.rns.MsgListResponse"></a>

### MsgListResponse







<a name="jackaldao.canine.rns.MsgRegister"></a>

### MsgRegister



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `name` | [string](#string) |  |  |
| `years` | [string](#string) |  |  |
| `data` | [string](#string) |  |  |






<a name="jackaldao.canine.rns.MsgRegisterResponse"></a>

### MsgRegisterResponse







<a name="jackaldao.canine.rns.MsgTransfer"></a>

### MsgTransfer



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `name` | [string](#string) |  |  |
| `reciever` | [string](#string) |  |  |






<a name="jackaldao.canine.rns.MsgTransferResponse"></a>

### MsgTransferResponse






 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="jackaldao.canine.rns.Msg"></a>

### Msg
Msg defines the Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Register` | [MsgRegister](#jackaldao.canine.rns.MsgRegister) | [MsgRegisterResponse](#jackaldao.canine.rns.MsgRegisterResponse) |  | |
| `Bid` | [MsgBid](#jackaldao.canine.rns.MsgBid) | [MsgBidResponse](#jackaldao.canine.rns.MsgBidResponse) |  | |
| `AcceptBid` | [MsgAcceptBid](#jackaldao.canine.rns.MsgAcceptBid) | [MsgAcceptBidResponse](#jackaldao.canine.rns.MsgAcceptBidResponse) |  | |
| `CancelBid` | [MsgCancelBid](#jackaldao.canine.rns.MsgCancelBid) | [MsgCancelBidResponse](#jackaldao.canine.rns.MsgCancelBidResponse) |  | |
| `List` | [MsgList](#jackaldao.canine.rns.MsgList) | [MsgListResponse](#jackaldao.canine.rns.MsgListResponse) |  | |
| `Buy` | [MsgBuy](#jackaldao.canine.rns.MsgBuy) | [MsgBuyResponse](#jackaldao.canine.rns.MsgBuyResponse) |  | |
| `Delist` | [MsgDelist](#jackaldao.canine.rns.MsgDelist) | [MsgDelistResponse](#jackaldao.canine.rns.MsgDelistResponse) |  | |
| `Transfer` | [MsgTransfer](#jackaldao.canine.rns.MsgTransfer) | [MsgTransferResponse](#jackaldao.canine.rns.MsgTransferResponse) |  | |
| `AddRecord` | [MsgAddRecord](#jackaldao.canine.rns.MsgAddRecord) | [MsgAddRecordResponse](#jackaldao.canine.rns.MsgAddRecordResponse) |  | |
| `DelRecord` | [MsgDelRecord](#jackaldao.canine.rns.MsgDelRecord) | [MsgDelRecordResponse](#jackaldao.canine.rns.MsgDelRecordResponse) |  | |
| `Init` | [MsgInit](#jackaldao.canine.rns.MsgInit) | [MsgInitResponse](#jackaldao.canine.rns.MsgInitResponse) | this line is used by starport scaffolding # proto/tx/rpc | |

 <!-- end services -->



<a name="storage/active_deals.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## storage/active_deals.proto



<a name="jackaldao.canine.storage.ActiveDeals"></a>

### ActiveDeals



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `cid` | [string](#string) |  |  |
| `signee` | [string](#string) |  |  |
| `provider` | [string](#string) |  |  |
| `startblock` | [string](#string) |  |  |
| `endblock` | [string](#string) |  |  |
| `filesize` | [string](#string) |  |  |
| `proofverified` | [string](#string) |  |  |
| `proofsmissed` | [string](#string) |  |  |
| `blocktoprove` | [string](#string) |  |  |
| `creator` | [string](#string) |  |  |
| `merkle` | [string](#string) |  |  |
| `fid` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="storage/client_usage.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## storage/client_usage.proto



<a name="jackaldao.canine.storage.ClientUsage"></a>

### ClientUsage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |
| `usage` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="storage/contracts.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## storage/contracts.proto



<a name="jackaldao.canine.storage.Contracts"></a>

### Contracts



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `cid` | [string](#string) |  |  |
| `priceamt` | [string](#string) |  |  |
| `pricedenom` | [string](#string) |  |  |
| `merkle` | [string](#string) |  |  |
| `signee` | [string](#string) |  |  |
| `duration` | [string](#string) |  |  |
| `filesize` | [string](#string) |  |  |
| `fid` | [string](#string) |  |  |
| `creator` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="storage/fid_cid.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## storage/fid_cid.proto



<a name="jackaldao.canine.storage.FidCid"></a>

### FidCid



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `fid` | [string](#string) |  |  |
| `cids` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="storage/params.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## storage/params.proto



<a name="jackaldao.canine.storage.Params"></a>

### Params
Params defines the parameters for the module.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="storage/proofs.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## storage/proofs.proto



<a name="jackaldao.canine.storage.Proofs"></a>

### Proofs



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `cid` | [string](#string) |  |  |
| `item` | [string](#string) |  |  |
| `hashes` | [string](#string) |  |  |
| `creator` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="storage/providers.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## storage/providers.proto



<a name="jackaldao.canine.storage.Providers"></a>

### Providers



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |
| `ip` | [string](#string) |  |  |
| `totalspace` | [string](#string) |  |  |
| `burned_contracts` | [string](#string) |  |  |
| `creator` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="storage/pay_blocks.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## storage/pay_blocks.proto



<a name="jackaldao.canine.storage.PayBlocks"></a>

### PayBlocks



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `blockid` | [string](#string) |  |  |
| `bytes` | [string](#string) |  |  |
| `blocktype` | [string](#string) |  |  |
| `blocknum` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="storage/strays.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## storage/strays.proto



<a name="jackaldao.canine.storage.Strays"></a>

### Strays



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `cid` | [string](#string) |  |  |
| `fid` | [string](#string) |  |  |
| `signee` | [string](#string) |  |  |
| `filesize` | [string](#string) |  |  |
| `merkle` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="storage/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## storage/genesis.proto



<a name="jackaldao.canine.storage.GenesisState"></a>

### GenesisState
GenesisState defines the storage module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#jackaldao.canine.storage.Params) |  |  |
| `contractsList` | [Contracts](#jackaldao.canine.storage.Contracts) | repeated |  |
| `proofsList` | [Proofs](#jackaldao.canine.storage.Proofs) | repeated |  |
| `activeDealsList` | [ActiveDeals](#jackaldao.canine.storage.ActiveDeals) | repeated |  |
| `providersList` | [Providers](#jackaldao.canine.storage.Providers) | repeated |  |
| `payBlocksList` | [PayBlocks](#jackaldao.canine.storage.PayBlocks) | repeated |  |
| `clientUsageList` | [ClientUsage](#jackaldao.canine.storage.ClientUsage) | repeated |  |
| `straysList` | [Strays](#jackaldao.canine.storage.Strays) | repeated |  |
| `fidCidList` | [FidCid](#jackaldao.canine.storage.FidCid) | repeated | this line is used by starport scaffolding # genesis/proto/state |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="storage/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## storage/query.proto



<a name="jackaldao.canine.storage.QueryAllActiveDealsRequest"></a>

### QueryAllActiveDealsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="jackaldao.canine.storage.QueryAllActiveDealsResponse"></a>

### QueryAllActiveDealsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `activeDeals` | [ActiveDeals](#jackaldao.canine.storage.ActiveDeals) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="jackaldao.canine.storage.QueryAllClientUsageRequest"></a>

### QueryAllClientUsageRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="jackaldao.canine.storage.QueryAllClientUsageResponse"></a>

### QueryAllClientUsageResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `clientUsage` | [ClientUsage](#jackaldao.canine.storage.ClientUsage) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="jackaldao.canine.storage.QueryAllContractsRequest"></a>

### QueryAllContractsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="jackaldao.canine.storage.QueryAllContractsResponse"></a>

### QueryAllContractsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `contracts` | [Contracts](#jackaldao.canine.storage.Contracts) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="jackaldao.canine.storage.QueryAllFidCidRequest"></a>

### QueryAllFidCidRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="jackaldao.canine.storage.QueryAllFidCidResponse"></a>

### QueryAllFidCidResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `fidCid` | [FidCid](#jackaldao.canine.storage.FidCid) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="jackaldao.canine.storage.QueryAllPayBlocksRequest"></a>

### QueryAllPayBlocksRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="jackaldao.canine.storage.QueryAllPayBlocksResponse"></a>

### QueryAllPayBlocksResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `payBlocks` | [PayBlocks](#jackaldao.canine.storage.PayBlocks) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="jackaldao.canine.storage.QueryAllProofsRequest"></a>

### QueryAllProofsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="jackaldao.canine.storage.QueryAllProofsResponse"></a>

### QueryAllProofsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proofs` | [Proofs](#jackaldao.canine.storage.Proofs) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="jackaldao.canine.storage.QueryAllProvidersRequest"></a>

### QueryAllProvidersRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="jackaldao.canine.storage.QueryAllProvidersResponse"></a>

### QueryAllProvidersResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `providers` | [Providers](#jackaldao.canine.storage.Providers) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="jackaldao.canine.storage.QueryAllStraysRequest"></a>

### QueryAllStraysRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="jackaldao.canine.storage.QueryAllStraysResponse"></a>

### QueryAllStraysResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `strays` | [Strays](#jackaldao.canine.storage.Strays) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="jackaldao.canine.storage.QueryFindFileRequest"></a>

### QueryFindFileRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `fid` | [string](#string) |  |  |






<a name="jackaldao.canine.storage.QueryFindFileResponse"></a>

### QueryFindFileResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `providerIps` | [string](#string) |  |  |






<a name="jackaldao.canine.storage.QueryFreespaceRequest"></a>

### QueryFreespaceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |






<a name="jackaldao.canine.storage.QueryFreespaceResponse"></a>

### QueryFreespaceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `space` | [string](#string) |  |  |






<a name="jackaldao.canine.storage.QueryGetActiveDealsRequest"></a>

### QueryGetActiveDealsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `cid` | [string](#string) |  |  |






<a name="jackaldao.canine.storage.QueryGetActiveDealsResponse"></a>

### QueryGetActiveDealsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `activeDeals` | [ActiveDeals](#jackaldao.canine.storage.ActiveDeals) |  |  |






<a name="jackaldao.canine.storage.QueryGetClientFreeSpaceRequest"></a>

### QueryGetClientFreeSpaceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |






<a name="jackaldao.canine.storage.QueryGetClientFreeSpaceResponse"></a>

### QueryGetClientFreeSpaceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `bytesfree` | [string](#string) |  |  |






<a name="jackaldao.canine.storage.QueryGetClientUsageRequest"></a>

### QueryGetClientUsageRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |






<a name="jackaldao.canine.storage.QueryGetClientUsageResponse"></a>

### QueryGetClientUsageResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `clientUsage` | [ClientUsage](#jackaldao.canine.storage.ClientUsage) |  |  |






<a name="jackaldao.canine.storage.QueryGetContractsRequest"></a>

### QueryGetContractsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `cid` | [string](#string) |  |  |






<a name="jackaldao.canine.storage.QueryGetContractsResponse"></a>

### QueryGetContractsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `contracts` | [Contracts](#jackaldao.canine.storage.Contracts) |  |  |






<a name="jackaldao.canine.storage.QueryGetFidCidRequest"></a>

### QueryGetFidCidRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `fid` | [string](#string) |  |  |






<a name="jackaldao.canine.storage.QueryGetFidCidResponse"></a>

### QueryGetFidCidResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `fidCid` | [FidCid](#jackaldao.canine.storage.FidCid) |  |  |






<a name="jackaldao.canine.storage.QueryGetPayBlocksRequest"></a>

### QueryGetPayBlocksRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `blockid` | [string](#string) |  |  |






<a name="jackaldao.canine.storage.QueryGetPayBlocksResponse"></a>

### QueryGetPayBlocksResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `payBlocks` | [PayBlocks](#jackaldao.canine.storage.PayBlocks) |  |  |






<a name="jackaldao.canine.storage.QueryGetPayDataRequest"></a>

### QueryGetPayDataRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |






<a name="jackaldao.canine.storage.QueryGetPayDataResponse"></a>

### QueryGetPayDataResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `blocks_remaining` | [int64](#int64) |  |  |
| `bytes` | [int64](#int64) |  |  |






<a name="jackaldao.canine.storage.QueryGetProofsRequest"></a>

### QueryGetProofsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `cid` | [string](#string) |  |  |






<a name="jackaldao.canine.storage.QueryGetProofsResponse"></a>

### QueryGetProofsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proofs` | [Proofs](#jackaldao.canine.storage.Proofs) |  |  |






<a name="jackaldao.canine.storage.QueryGetProvidersRequest"></a>

### QueryGetProvidersRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |






<a name="jackaldao.canine.storage.QueryGetProvidersResponse"></a>

### QueryGetProvidersResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `providers` | [Providers](#jackaldao.canine.storage.Providers) |  |  |






<a name="jackaldao.canine.storage.QueryGetStraysRequest"></a>

### QueryGetStraysRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `cid` | [string](#string) |  |  |






<a name="jackaldao.canine.storage.QueryGetStraysResponse"></a>

### QueryGetStraysResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `strays` | [Strays](#jackaldao.canine.storage.Strays) |  |  |






<a name="jackaldao.canine.storage.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is request type for the Query/Params RPC method.






<a name="jackaldao.canine.storage.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is response type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#jackaldao.canine.storage.Params) |  | params holds all the parameters of this module. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="jackaldao.canine.storage.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#jackaldao.canine.storage.QueryParamsRequest) | [QueryParamsResponse](#jackaldao.canine.storage.QueryParamsResponse) | Parameters queries the parameters of the module. | GET|/jackal-dao/canine/storage/params|
| `Contracts` | [QueryGetContractsRequest](#jackaldao.canine.storage.QueryGetContractsRequest) | [QueryGetContractsResponse](#jackaldao.canine.storage.QueryGetContractsResponse) | Queries a Contracts by index. | GET|/jackal-dao/canine/storage/contracts/{cid}|
| `ContractsAll` | [QueryAllContractsRequest](#jackaldao.canine.storage.QueryAllContractsRequest) | [QueryAllContractsResponse](#jackaldao.canine.storage.QueryAllContractsResponse) | Queries a list of Contracts items. | GET|/jackal-dao/canine/storage/contracts|
| `Proofs` | [QueryGetProofsRequest](#jackaldao.canine.storage.QueryGetProofsRequest) | [QueryGetProofsResponse](#jackaldao.canine.storage.QueryGetProofsResponse) | Queries a Proofs by index. | GET|/jackal-dao/canine/storage/proofs/{cid}|
| `ProofsAll` | [QueryAllProofsRequest](#jackaldao.canine.storage.QueryAllProofsRequest) | [QueryAllProofsResponse](#jackaldao.canine.storage.QueryAllProofsResponse) | Queries a list of Proofs items. | GET|/jackal-dao/canine/storage/proofs|
| `ActiveDeals` | [QueryGetActiveDealsRequest](#jackaldao.canine.storage.QueryGetActiveDealsRequest) | [QueryGetActiveDealsResponse](#jackaldao.canine.storage.QueryGetActiveDealsResponse) | Queries a ActiveDeals by index. | GET|/jackal-dao/canine/storage/active_deals/{cid}|
| `ActiveDealsAll` | [QueryAllActiveDealsRequest](#jackaldao.canine.storage.QueryAllActiveDealsRequest) | [QueryAllActiveDealsResponse](#jackaldao.canine.storage.QueryAllActiveDealsResponse) | Queries a list of ActiveDeals items. | GET|/jackal-dao/canine/storage/active_deals|
| `Providers` | [QueryGetProvidersRequest](#jackaldao.canine.storage.QueryGetProvidersRequest) | [QueryGetProvidersResponse](#jackaldao.canine.storage.QueryGetProvidersResponse) | Queries a Providers by index. | GET|/jackal-dao/canine/storage/providers/{address}|
| `ProvidersAll` | [QueryAllProvidersRequest](#jackaldao.canine.storage.QueryAllProvidersRequest) | [QueryAllProvidersResponse](#jackaldao.canine.storage.QueryAllProvidersResponse) | Queries a list of Providers items. | GET|/jackal-dao/canine/storage/providers|
| `Freespace` | [QueryFreespaceRequest](#jackaldao.canine.storage.QueryFreespaceRequest) | [QueryFreespaceResponse](#jackaldao.canine.storage.QueryFreespaceResponse) | Queries a list of Freespace items. | GET|/jackal-dao/canine/storage/freespace/{address}|
| `FindFile` | [QueryFindFileRequest](#jackaldao.canine.storage.QueryFindFileRequest) | [QueryFindFileResponse](#jackaldao.canine.storage.QueryFindFileResponse) | Queries a list of FindFile items. | GET|/jackal-dao/canine/storage/find_file/{fid}|
| `PayBlocks` | [QueryGetPayBlocksRequest](#jackaldao.canine.storage.QueryGetPayBlocksRequest) | [QueryGetPayBlocksResponse](#jackaldao.canine.storage.QueryGetPayBlocksResponse) | Queries a PayBlocks by index. | GET|/jackal-dao/canine/storage/pay_blocks/{blockid}|
| `PayBlocksAll` | [QueryAllPayBlocksRequest](#jackaldao.canine.storage.QueryAllPayBlocksRequest) | [QueryAllPayBlocksResponse](#jackaldao.canine.storage.QueryAllPayBlocksResponse) | Queries a list of PayBlocks items. | GET|/jackal-dao/canine/storage/pay_blocks|
| `ClientUsage` | [QueryGetClientUsageRequest](#jackaldao.canine.storage.QueryGetClientUsageRequest) | [QueryGetClientUsageResponse](#jackaldao.canine.storage.QueryGetClientUsageResponse) | Queries a ClientUsage by index. | GET|/jackal-dao/canine/storage/client_usage/{address}|
| `ClientUsageAll` | [QueryAllClientUsageRequest](#jackaldao.canine.storage.QueryAllClientUsageRequest) | [QueryAllClientUsageResponse](#jackaldao.canine.storage.QueryAllClientUsageResponse) | Queries a list of ClientUsage items. | GET|/jackal-dao/canine/storage/client_usage|
| `Strays` | [QueryGetStraysRequest](#jackaldao.canine.storage.QueryGetStraysRequest) | [QueryGetStraysResponse](#jackaldao.canine.storage.QueryGetStraysResponse) | Queries a Strays by index. | GET|/jackal-dao/canine/storage/strays/{cid}|
| `StraysAll` | [QueryAllStraysRequest](#jackaldao.canine.storage.QueryAllStraysRequest) | [QueryAllStraysResponse](#jackaldao.canine.storage.QueryAllStraysResponse) | Queries a list of Strays items. | GET|/jackal-dao/canine/storage/strays|
| `GetClientFreeSpace` | [QueryGetClientFreeSpaceRequest](#jackaldao.canine.storage.QueryGetClientFreeSpaceRequest) | [QueryGetClientFreeSpaceResponse](#jackaldao.canine.storage.QueryGetClientFreeSpaceResponse) | Queries a list of GetClientFreeSpace items. | GET|/jackal-dao/canine/storage/get_client_free_space/{address}|
| `FidCid` | [QueryGetFidCidRequest](#jackaldao.canine.storage.QueryGetFidCidRequest) | [QueryGetFidCidResponse](#jackaldao.canine.storage.QueryGetFidCidResponse) | Queries a FidCid by index. | GET|/jackal-dao/canine/storage/fid_cid/{fid}|
| `FidCidAll` | [QueryAllFidCidRequest](#jackaldao.canine.storage.QueryAllFidCidRequest) | [QueryAllFidCidResponse](#jackaldao.canine.storage.QueryAllFidCidResponse) | Queries a list of FidCid items. | GET|/jackal-dao/canine/storage/fid_cid|
| `GetPayData` | [QueryGetPayDataRequest](#jackaldao.canine.storage.QueryGetPayDataRequest) | [QueryGetPayDataResponse](#jackaldao.canine.storage.QueryGetPayDataResponse) | Queries a list of GetPayData items. | GET|/jackal-dao/canine/storage/get_pay_data/{address}|

 <!-- end services -->



<a name="storage/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## storage/tx.proto



<a name="jackaldao.canine.storage.MsgBuyStorage"></a>

### MsgBuyStorage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `forAddress` | [string](#string) |  |  |
| `duration` | [string](#string) |  |  |
| `bytes` | [string](#string) |  |  |
| `paymentDenom` | [string](#string) |  |  |






<a name="jackaldao.canine.storage.MsgBuyStorageResponse"></a>

### MsgBuyStorageResponse







<a name="jackaldao.canine.storage.MsgCancelContract"></a>

### MsgCancelContract



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `cid` | [string](#string) |  |  |






<a name="jackaldao.canine.storage.MsgCancelContractResponse"></a>

### MsgCancelContractResponse







<a name="jackaldao.canine.storage.MsgClaimStray"></a>

### MsgClaimStray



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `cid` | [string](#string) |  |  |






<a name="jackaldao.canine.storage.MsgClaimStrayResponse"></a>

### MsgClaimStrayResponse







<a name="jackaldao.canine.storage.MsgCreateActiveDeals"></a>

### MsgCreateActiveDeals



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `cid` | [string](#string) |  |  |
| `signee` | [string](#string) |  |  |
| `provider` | [string](#string) |  |  |
| `startblock` | [string](#string) |  |  |
| `endblock` | [string](#string) |  |  |
| `filesize` | [string](#string) |  |  |
| `proofverified` | [string](#string) |  |  |
| `proofsmissed` | [string](#string) |  |  |
| `blocktoprove` | [string](#string) |  |  |






<a name="jackaldao.canine.storage.MsgCreateActiveDealsResponse"></a>

### MsgCreateActiveDealsResponse







<a name="jackaldao.canine.storage.MsgCreateContracts"></a>

### MsgCreateContracts



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `cid` | [string](#string) |  |  |
| `priceamt` | [string](#string) |  |  |
| `pricedenom` | [string](#string) |  |  |
| `chunks` | [string](#string) |  |  |
| `merkle` | [string](#string) |  |  |
| `signee` | [string](#string) |  |  |
| `duration` | [string](#string) |  |  |
| `filesize` | [string](#string) |  |  |
| `fid` | [string](#string) |  |  |






<a name="jackaldao.canine.storage.MsgCreateContractsResponse"></a>

### MsgCreateContractsResponse







<a name="jackaldao.canine.storage.MsgCreateProofs"></a>

### MsgCreateProofs



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `cid` | [string](#string) |  |  |
| `item` | [string](#string) |  |  |
| `hashes` | [string](#string) |  |  |






<a name="jackaldao.canine.storage.MsgCreateProofsResponse"></a>

### MsgCreateProofsResponse







<a name="jackaldao.canine.storage.MsgCreateProviders"></a>

### MsgCreateProviders



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `address` | [string](#string) |  |  |
| `ip` | [string](#string) |  |  |
| `totalspace` | [string](#string) |  |  |






<a name="jackaldao.canine.storage.MsgCreateProvidersResponse"></a>

### MsgCreateProvidersResponse







<a name="jackaldao.canine.storage.MsgDeleteActiveDeals"></a>

### MsgDeleteActiveDeals



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `cid` | [string](#string) |  |  |






<a name="jackaldao.canine.storage.MsgDeleteActiveDealsResponse"></a>

### MsgDeleteActiveDealsResponse







<a name="jackaldao.canine.storage.MsgDeleteContracts"></a>

### MsgDeleteContracts



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `cid` | [string](#string) |  |  |






<a name="jackaldao.canine.storage.MsgDeleteContractsResponse"></a>

### MsgDeleteContractsResponse







<a name="jackaldao.canine.storage.MsgDeleteProofs"></a>

### MsgDeleteProofs



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `cid` | [string](#string) |  |  |






<a name="jackaldao.canine.storage.MsgDeleteProofsResponse"></a>

### MsgDeleteProofsResponse







<a name="jackaldao.canine.storage.MsgDeleteProviders"></a>

### MsgDeleteProviders



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `address` | [string](#string) |  |  |






<a name="jackaldao.canine.storage.MsgDeleteProvidersResponse"></a>

### MsgDeleteProvidersResponse







<a name="jackaldao.canine.storage.MsgInitProvider"></a>

### MsgInitProvider



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `ip` | [string](#string) |  |  |
| `totalspace` | [string](#string) |  |  |






<a name="jackaldao.canine.storage.MsgInitProviderResponse"></a>

### MsgInitProviderResponse







<a name="jackaldao.canine.storage.MsgItem"></a>

### MsgItem



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `hashlist` | [string](#string) |  |  |






<a name="jackaldao.canine.storage.MsgItemResponse"></a>

### MsgItemResponse







<a name="jackaldao.canine.storage.MsgPostContract"></a>

### MsgPostContract



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `priceamt` | [string](#string) |  |  |
| `pricedenom` | [string](#string) |  |  |
| `merkle` | [string](#string) |  |  |
| `signee` | [string](#string) |  |  |
| `duration` | [string](#string) |  |  |
| `filesize` | [string](#string) |  |  |
| `fid` | [string](#string) |  |  |






<a name="jackaldao.canine.storage.MsgPostContractResponse"></a>

### MsgPostContractResponse







<a name="jackaldao.canine.storage.MsgPostproof"></a>

### MsgPostproof



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `item` | [string](#string) |  |  |
| `hashlist` | [string](#string) |  |  |
| `cid` | [string](#string) |  |  |






<a name="jackaldao.canine.storage.MsgPostproofResponse"></a>

### MsgPostproofResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `merkle` | [string](#string) |  |  |






<a name="jackaldao.canine.storage.MsgSetProviderIP"></a>

### MsgSetProviderIP



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `ip` | [string](#string) |  |  |






<a name="jackaldao.canine.storage.MsgSetProviderIPResponse"></a>

### MsgSetProviderIPResponse







<a name="jackaldao.canine.storage.MsgSetProviderTotalspace"></a>

### MsgSetProviderTotalspace



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `space` | [string](#string) |  |  |






<a name="jackaldao.canine.storage.MsgSetProviderTotalspaceResponse"></a>

### MsgSetProviderTotalspaceResponse







<a name="jackaldao.canine.storage.MsgSignContract"></a>

### MsgSignContract



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `cid` | [string](#string) |  |  |






<a name="jackaldao.canine.storage.MsgSignContractResponse"></a>

### MsgSignContractResponse







<a name="jackaldao.canine.storage.MsgUpdateActiveDeals"></a>

### MsgUpdateActiveDeals



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `cid` | [string](#string) |  |  |
| `signee` | [string](#string) |  |  |
| `provider` | [string](#string) |  |  |
| `startblock` | [string](#string) |  |  |
| `endblock` | [string](#string) |  |  |
| `filesize` | [string](#string) |  |  |
| `proofverified` | [string](#string) |  |  |
| `proofsmissed` | [string](#string) |  |  |
| `blocktoprove` | [string](#string) |  |  |






<a name="jackaldao.canine.storage.MsgUpdateActiveDealsResponse"></a>

### MsgUpdateActiveDealsResponse







<a name="jackaldao.canine.storage.MsgUpdateContracts"></a>

### MsgUpdateContracts



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `cid` | [string](#string) |  |  |
| `chunks` | [string](#string) |  |  |
| `merkle` | [string](#string) |  |  |
| `signee` | [string](#string) |  |  |
| `duration` | [string](#string) |  |  |
| `filesize` | [string](#string) |  |  |
| `fid` | [string](#string) |  |  |






<a name="jackaldao.canine.storage.MsgUpdateContractsResponse"></a>

### MsgUpdateContractsResponse







<a name="jackaldao.canine.storage.MsgUpdateProofs"></a>

### MsgUpdateProofs



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `cid` | [string](#string) |  |  |
| `item` | [string](#string) |  |  |
| `hashes` | [string](#string) |  |  |






<a name="jackaldao.canine.storage.MsgUpdateProofsResponse"></a>

### MsgUpdateProofsResponse







<a name="jackaldao.canine.storage.MsgUpdateProviders"></a>

### MsgUpdateProviders



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |
| `address` | [string](#string) |  |  |
| `ip` | [string](#string) |  |  |
| `totalspace` | [string](#string) |  |  |






<a name="jackaldao.canine.storage.MsgUpdateProvidersResponse"></a>

### MsgUpdateProvidersResponse






 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="jackaldao.canine.storage.Msg"></a>

### Msg
Msg defines the Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `PostContract` | [MsgPostContract](#jackaldao.canine.storage.MsgPostContract) | [MsgPostContractResponse](#jackaldao.canine.storage.MsgPostContractResponse) |  | |
| `Postproof` | [MsgPostproof](#jackaldao.canine.storage.MsgPostproof) | [MsgPostproofResponse](#jackaldao.canine.storage.MsgPostproofResponse) |  | |
| `SignContract` | [MsgSignContract](#jackaldao.canine.storage.MsgSignContract) | [MsgSignContractResponse](#jackaldao.canine.storage.MsgSignContractResponse) |  | |
| `SetProviderIP` | [MsgSetProviderIP](#jackaldao.canine.storage.MsgSetProviderIP) | [MsgSetProviderIPResponse](#jackaldao.canine.storage.MsgSetProviderIPResponse) |  | |
| `SetProviderTotalspace` | [MsgSetProviderTotalspace](#jackaldao.canine.storage.MsgSetProviderTotalspace) | [MsgSetProviderTotalspaceResponse](#jackaldao.canine.storage.MsgSetProviderTotalspaceResponse) |  | |
| `InitProvider` | [MsgInitProvider](#jackaldao.canine.storage.MsgInitProvider) | [MsgInitProviderResponse](#jackaldao.canine.storage.MsgInitProviderResponse) |  | |
| `CancelContract` | [MsgCancelContract](#jackaldao.canine.storage.MsgCancelContract) | [MsgCancelContractResponse](#jackaldao.canine.storage.MsgCancelContractResponse) |  | |
| `BuyStorage` | [MsgBuyStorage](#jackaldao.canine.storage.MsgBuyStorage) | [MsgBuyStorageResponse](#jackaldao.canine.storage.MsgBuyStorageResponse) |  | |
| `ClaimStray` | [MsgClaimStray](#jackaldao.canine.storage.MsgClaimStray) | [MsgClaimStrayResponse](#jackaldao.canine.storage.MsgClaimStrayResponse) | this line is used by starport scaffolding # proto/tx/rpc | |

 <!-- end services -->



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

