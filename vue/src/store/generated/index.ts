// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import CosmosCosmosSdkCosmosAuthzV1Beta1 from './cosmos/cosmos-sdk/cosmos.authz.v1beta1'
import CosmosInterchainAccountsIntertx from './cosmos/interchain-accounts/intertx'
import IbcGoV3IbcApplicationsInterchainAccountsControllerV1 from './ibc-go/v3/ibc.applications.interchain_accounts.controller.v1'
import IbcGoV3IbcApplicationsInterchainAccountsHostV1 from './ibc-go/v3/ibc.applications.interchain_accounts.host.v1'
import IbcGoV3IbcApplicationsTransferV1 from './ibc-go/v3/ibc.applications.transfer.v1'
import IbcGoV3IbcCoreChannelV1 from './ibc-go/v3/ibc.core.channel.v1'
import IbcGoV3IbcCoreClientV1 from './ibc-go/v3/ibc.core.client.v1'
import IbcGoV3IbcCoreConnectionV1 from './ibc-go/v3/ibc.core.connection.v1'
import JackalDaoCanineCosmwasmWasmV1 from './jackal-dao/canine/cosmwasm.wasm.v1'
import JackalDaoCanineJackaldaoCanineDsig from './jackal-dao/canine/jackaldao.canine.dsig'
import JackalDaoCanineJackaldaoCanineFiletree from './jackal-dao/canine/jackaldao.canine.filetree'
import JackalDaoCanineJackaldaoCanineJklmint from './jackal-dao/canine/jackaldao.canine.jklmint'
import JackalDaoCanineJackaldaoCanineLp from './jackal-dao/canine/jackaldao.canine.lp'
import JackalDaoCanineJackaldaoCanineNotifications from './jackal-dao/canine/jackaldao.canine.notifications'
import JackalDaoCanineJackaldaoCanineRns from './jackal-dao/canine/jackaldao.canine.rns'
import JackalDaoCanineJackaldaoCanineStorage from './jackal-dao/canine/jackaldao.canine.storage'


export default { 
  CosmosCosmosSdkCosmosAuthzV1Beta1: load(CosmosCosmosSdkCosmosAuthzV1Beta1, 'cosmos.authz.v1beta1'),
  CosmosInterchainAccountsIntertx: load(CosmosInterchainAccountsIntertx, 'intertx'),
  IbcGoV3IbcApplicationsInterchainAccountsControllerV1: load(IbcGoV3IbcApplicationsInterchainAccountsControllerV1, 'ibc.applications.interchain_accounts.controller.v1'),
  IbcGoV3IbcApplicationsInterchainAccountsHostV1: load(IbcGoV3IbcApplicationsInterchainAccountsHostV1, 'ibc.applications.interchain_accounts.host.v1'),
  IbcGoV3IbcApplicationsTransferV1: load(IbcGoV3IbcApplicationsTransferV1, 'ibc.applications.transfer.v1'),
  IbcGoV3IbcCoreChannelV1: load(IbcGoV3IbcCoreChannelV1, 'ibc.core.channel.v1'),
  IbcGoV3IbcCoreClientV1: load(IbcGoV3IbcCoreClientV1, 'ibc.core.client.v1'),
  IbcGoV3IbcCoreConnectionV1: load(IbcGoV3IbcCoreConnectionV1, 'ibc.core.connection.v1'),
  JackalDaoCanineCosmwasmWasmV1: load(JackalDaoCanineCosmwasmWasmV1, 'cosmwasm.wasm.v1'),
  JackalDaoCanineJackaldaoCanineDsig: load(JackalDaoCanineJackaldaoCanineDsig, 'jackaldao.canine.dsig'),
  JackalDaoCanineJackaldaoCanineFiletree: load(JackalDaoCanineJackaldaoCanineFiletree, 'jackaldao.canine.filetree'),
  JackalDaoCanineJackaldaoCanineJklmint: load(JackalDaoCanineJackaldaoCanineJklmint, 'jackaldao.canine.jklmint'),
  JackalDaoCanineJackaldaoCanineLp: load(JackalDaoCanineJackaldaoCanineLp, 'jackaldao.canine.lp'),
  JackalDaoCanineJackaldaoCanineNotifications: load(JackalDaoCanineJackaldaoCanineNotifications, 'jackaldao.canine.notifications'),
  JackalDaoCanineJackaldaoCanineRns: load(JackalDaoCanineJackaldaoCanineRns, 'jackaldao.canine.rns'),
  JackalDaoCanineJackaldaoCanineStorage: load(JackalDaoCanineJackaldaoCanineStorage, 'jackaldao.canine.storage'),
  
}


function load(mod, fullns) {
    return function init(store) {        
        if (store.hasModule([fullns])) {
            throw new Error('Duplicate module name detected: '+ fullns)
        }else{
            store.registerModule([fullns], mod)
            store.subscribe((mutation) => {
                if (mutation.type == 'common/env/INITIALIZE_WS_COMPLETE') {
                    store.dispatch(fullns+ '/init', null, {
                        root: true
                    })
                }
            })
        }
    }
}
