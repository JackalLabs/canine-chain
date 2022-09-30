// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import CosmwasmWasmV1 from './cosmwasm.wasm.v1'
import JackaldaoCanineDsig from './jackaldao.canine.dsig'
import JackaldaoCanineFiletree from './jackaldao.canine.filetree'
import JackaldaoCanineJklmint from './jackaldao.canine.jklmint'
import JackaldaoCanineLp from './jackaldao.canine.lp'
import JackaldaoCanineNotifications from './jackaldao.canine.notifications'
import JackaldaoCanineRns from './jackaldao.canine.rns'
import JackaldaoCanineStorage from './jackaldao.canine.storage'


export default { 
  CosmwasmWasmV1: load(CosmwasmWasmV1, 'cosmwasm.wasm.v1'),
  JackaldaoCanineDsig: load(JackaldaoCanineDsig, 'jackaldao.canine.dsig'),
  JackaldaoCanineFiletree: load(JackaldaoCanineFiletree, 'jackaldao.canine.filetree'),
  JackaldaoCanineJklmint: load(JackaldaoCanineJklmint, 'jackaldao.canine.jklmint'),
  JackaldaoCanineLp: load(JackaldaoCanineLp, 'jackaldao.canine.lp'),
  JackaldaoCanineNotifications: load(JackaldaoCanineNotifications, 'jackaldao.canine.notifications'),
  JackaldaoCanineRns: load(JackaldaoCanineRns, 'jackaldao.canine.rns'),
  JackaldaoCanineStorage: load(JackaldaoCanineStorage, 'jackaldao.canine.storage'),
  
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