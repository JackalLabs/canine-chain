// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import JackalDaoCanineCosmwasmWasmV1 from './jackal-dao/canine/cosmwasm.wasm.v1'
import JackalDaoCanineJackaldaoCanineDsig from './jackal-dao/canine/jackaldao.canine.dsig'
import JackalDaoCanineJackaldaoCanineFiletree from './jackal-dao/canine/jackaldao.canine.filetree'
import JackalDaoCanineJackaldaoCanineJklmint from './jackal-dao/canine/jackaldao.canine.jklmint'
import JackalDaoCanineJackaldaoCanineLp from './jackal-dao/canine/jackaldao.canine.lp'
import JackalDaoCanineJackaldaoCanineNotifications from './jackal-dao/canine/jackaldao.canine.notifications'
import JackalDaoCanineJackaldaoCanineRns from './jackal-dao/canine/jackaldao.canine.rns'
import JackalDaoCanineJackaldaoCanineStorage from './jackal-dao/canine/jackaldao.canine.storage'


export default { 
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
