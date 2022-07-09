// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import JackalDaoCanineCosmwasmWasmV1 from './jackal-dao/canine/cosmwasm.wasm.v1';
import JackalDaoCanineJackaldaoCanineRns from './jackal-dao/canine/jackaldao.canine.rns';
export default {
    JackalDaoCanineCosmwasmWasmV1: load(JackalDaoCanineCosmwasmWasmV1, 'cosmwasm.wasm.v1'),
    JackalDaoCanineJackaldaoCanineRns: load(JackalDaoCanineJackaldaoCanineRns, 'jackaldao.canine.rns'),
};
function load(mod, fullns) {
    return function init(store) {
        if (store.hasModule([fullns])) {
            throw new Error('Duplicate module name detected: ' + fullns);
        }
        else {
            store.registerModule([fullns], mod);
            store.subscribe((mutation) => {
                if (mutation.type == 'common/env/INITIALIZE_WS_COMPLETE') {
                    store.dispatch(fullns + '/init', null, {
                        root: true
                    });
                }
            });
        }
    };
}
