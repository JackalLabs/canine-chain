import { txClient, queryClient, MissingWalletError, registry } from './module';
import { Mined } from "./module/types/jklmining/mined";
import { MinerClaims } from "./module/types/jklmining/miner_claims";
import { Miners } from "./module/types/jklmining/miners";
import { Params } from "./module/types/jklmining/params";
import { SaveRequests } from "./module/types/jklmining/save_requests";
export { Mined, MinerClaims, Miners, Params, SaveRequests };
async function initTxClient(vuexGetters) {
    return await txClient(vuexGetters['common/wallet/signer'], {
        addr: vuexGetters['common/env/apiTendermint']
    });
}
async function initQueryClient(vuexGetters) {
    return await queryClient({
        addr: vuexGetters['common/env/apiCosmos']
    });
}
function mergeResults(value, next_values) {
    for (let prop of Object.keys(next_values)) {
        if (Array.isArray(next_values[prop])) {
            value[prop] = [...value[prop], ...next_values[prop]];
        }
        else {
            value[prop] = next_values[prop];
        }
    }
    return value;
}
function getStructure(template) {
    let structure = { fields: [] };
    for (const [key, value] of Object.entries(template)) {
        let field = {};
        field.name = key;
        field.type = typeof value;
        structure.fields.push(field);
    }
    return structure;
}
const getDefaultState = () => {
    return {
        Params: {},
        SaveRequests: {},
        SaveRequestsAll: {},
        Miners: {},
        MinersAll: {},
        Mined: {},
        MinedAll: {},
        CheckMinerIndex: {},
        GetMinerIndex: {},
        GetMinerStart: {},
        MinerClaims: {},
        MinerClaimsAll: {},
        _Structure: {
            Mined: getStructure(Mined.fromPartial({})),
            MinerClaims: getStructure(MinerClaims.fromPartial({})),
            Miners: getStructure(Miners.fromPartial({})),
            Params: getStructure(Params.fromPartial({})),
            SaveRequests: getStructure(SaveRequests.fromPartial({})),
        },
        _Registry: registry,
        _Subscriptions: new Set(),
    };
};
// initial state
const state = getDefaultState();
export default {
    namespaced: true,
    state,
    mutations: {
        RESET_STATE(state) {
            Object.assign(state, getDefaultState());
        },
        QUERY(state, { query, key, value }) {
            state[query][JSON.stringify(key)] = value;
        },
        SUBSCRIBE(state, subscription) {
            state._Subscriptions.add(JSON.stringify(subscription));
        },
        UNSUBSCRIBE(state, subscription) {
            state._Subscriptions.delete(JSON.stringify(subscription));
        }
    },
    getters: {
        getParams: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.Params[JSON.stringify(params)] ?? {};
        },
        getSaveRequests: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.SaveRequests[JSON.stringify(params)] ?? {};
        },
        getSaveRequestsAll: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.SaveRequestsAll[JSON.stringify(params)] ?? {};
        },
        getMiners: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.Miners[JSON.stringify(params)] ?? {};
        },
        getMinersAll: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.MinersAll[JSON.stringify(params)] ?? {};
        },
        getMined: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.Mined[JSON.stringify(params)] ?? {};
        },
        getMinedAll: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.MinedAll[JSON.stringify(params)] ?? {};
        },
        getCheckMinerIndex: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.CheckMinerIndex[JSON.stringify(params)] ?? {};
        },
        getGetMinerIndex: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.GetMinerIndex[JSON.stringify(params)] ?? {};
        },
        getGetMinerStart: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.GetMinerStart[JSON.stringify(params)] ?? {};
        },
        getMinerClaims: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.MinerClaims[JSON.stringify(params)] ?? {};
        },
        getMinerClaimsAll: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.MinerClaimsAll[JSON.stringify(params)] ?? {};
        },
        getTypeStructure: (state) => (type) => {
            return state._Structure[type].fields;
        },
        getRegistry: (state) => {
            return state._Registry;
        }
    },
    actions: {
        init({ dispatch, rootGetters }) {
            console.log('Vuex module: jackaldao.canine.jklmining initialized!');
            if (rootGetters['common/env/client']) {
                rootGetters['common/env/client'].on('newblock', () => {
                    dispatch('StoreUpdate');
                });
            }
        },
        resetState({ commit }) {
            commit('RESET_STATE');
        },
        unsubscribe({ commit }, subscription) {
            commit('UNSUBSCRIBE', subscription);
        },
        async StoreUpdate({ state, dispatch }) {
            state._Subscriptions.forEach(async (subscription) => {
                try {
                    const sub = JSON.parse(subscription);
                    await dispatch(sub.action, sub.payload);
                }
                catch (e) {
                    throw new Error('Subscriptions: ' + e.message);
                }
            });
        },
        async QueryParams({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null }) {
            try {
                const key = params ?? {};
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryParams()).data;
                commit('QUERY', { query: 'Params', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryParams', payload: { options: { all }, params: { ...key }, query } });
                return getters['getParams']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new Error('QueryClient:QueryParams API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QuerySaveRequests({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null }) {
            try {
                const key = params ?? {};
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.querySaveRequests(key.index)).data;
                commit('QUERY', { query: 'SaveRequests', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QuerySaveRequests', payload: { options: { all }, params: { ...key }, query } });
                return getters['getSaveRequests']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new Error('QueryClient:QuerySaveRequests API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QuerySaveRequestsAll({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null }) {
            try {
                const key = params ?? {};
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.querySaveRequestsAll(query)).data;
                while (all && value.pagination && value.pagination.next_key != null) {
                    let next_values = (await queryClient.querySaveRequestsAll({ ...query, 'pagination.key': value.pagination.next_key })).data;
                    value = mergeResults(value, next_values);
                }
                commit('QUERY', { query: 'SaveRequestsAll', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QuerySaveRequestsAll', payload: { options: { all }, params: { ...key }, query } });
                return getters['getSaveRequestsAll']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new Error('QueryClient:QuerySaveRequestsAll API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryMiners({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null }) {
            try {
                const key = params ?? {};
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryMiners(key.address)).data;
                commit('QUERY', { query: 'Miners', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryMiners', payload: { options: { all }, params: { ...key }, query } });
                return getters['getMiners']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new Error('QueryClient:QueryMiners API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryMinersAll({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null }) {
            try {
                const key = params ?? {};
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryMinersAll(query)).data;
                while (all && value.pagination && value.pagination.next_key != null) {
                    let next_values = (await queryClient.queryMinersAll({ ...query, 'pagination.key': value.pagination.next_key })).data;
                    value = mergeResults(value, next_values);
                }
                commit('QUERY', { query: 'MinersAll', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryMinersAll', payload: { options: { all }, params: { ...key }, query } });
                return getters['getMinersAll']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new Error('QueryClient:QueryMinersAll API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryMined({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null }) {
            try {
                const key = params ?? {};
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryMined(key.id)).data;
                commit('QUERY', { query: 'Mined', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryMined', payload: { options: { all }, params: { ...key }, query } });
                return getters['getMined']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new Error('QueryClient:QueryMined API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryMinedAll({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null }) {
            try {
                const key = params ?? {};
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryMinedAll(query)).data;
                while (all && value.pagination && value.pagination.next_key != null) {
                    let next_values = (await queryClient.queryMinedAll({ ...query, 'pagination.key': value.pagination.next_key })).data;
                    value = mergeResults(value, next_values);
                }
                commit('QUERY', { query: 'MinedAll', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryMinedAll', payload: { options: { all }, params: { ...key }, query } });
                return getters['getMinedAll']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new Error('QueryClient:QueryMinedAll API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryCheckMinerIndex({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null }) {
            try {
                const key = params ?? {};
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryCheckMinerIndex()).data;
                commit('QUERY', { query: 'CheckMinerIndex', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryCheckMinerIndex', payload: { options: { all }, params: { ...key }, query } });
                return getters['getCheckMinerIndex']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new Error('QueryClient:QueryCheckMinerIndex API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryGetMinerIndex({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null }) {
            try {
                const key = params ?? {};
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryGetMinerIndex(key.index)).data;
                commit('QUERY', { query: 'GetMinerIndex', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryGetMinerIndex', payload: { options: { all }, params: { ...key }, query } });
                return getters['getGetMinerIndex']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new Error('QueryClient:QueryGetMinerIndex API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryGetMinerStart({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null }) {
            try {
                const key = params ?? {};
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryGetMinerStart()).data;
                commit('QUERY', { query: 'GetMinerStart', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryGetMinerStart', payload: { options: { all }, params: { ...key }, query } });
                return getters['getGetMinerStart']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new Error('QueryClient:QueryGetMinerStart API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryMinerClaims({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null }) {
            try {
                const key = params ?? {};
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryMinerClaims(key.hash)).data;
                commit('QUERY', { query: 'MinerClaims', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryMinerClaims', payload: { options: { all }, params: { ...key }, query } });
                return getters['getMinerClaims']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new Error('QueryClient:QueryMinerClaims API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryMinerClaimsAll({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null }) {
            try {
                const key = params ?? {};
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryMinerClaimsAll(query)).data;
                while (all && value.pagination && value.pagination.next_key != null) {
                    let next_values = (await queryClient.queryMinerClaimsAll({ ...query, 'pagination.key': value.pagination.next_key })).data;
                    value = mergeResults(value, next_values);
                }
                commit('QUERY', { query: 'MinerClaimsAll', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryMinerClaimsAll', payload: { options: { all }, params: { ...key }, query } });
                return getters['getMinerClaimsAll']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new Error('QueryClient:QueryMinerClaimsAll API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async sendMsgUpdateMiners({ rootGetters }, { value, fee = [], memo = '' }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgUpdateMiners(value);
                const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee,
                        gas: "200000" }, memo });
                return result;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new Error('TxClient:MsgUpdateMiners:Init Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new Error('TxClient:MsgUpdateMiners:Send Could not broadcast Tx: ' + e.message);
                }
            }
        },
        async sendMsgCreateMinerClaims({ rootGetters }, { value, fee = [], memo = '' }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgCreateMinerClaims(value);
                const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee,
                        gas: "200000" }, memo });
                return result;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new Error('TxClient:MsgCreateMinerClaims:Init Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new Error('TxClient:MsgCreateMinerClaims:Send Could not broadcast Tx: ' + e.message);
                }
            }
        },
        async sendMsgDeleteMiners({ rootGetters }, { value, fee = [], memo = '' }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgDeleteMiners(value);
                const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee,
                        gas: "200000" }, memo });
                return result;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new Error('TxClient:MsgDeleteMiners:Init Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new Error('TxClient:MsgDeleteMiners:Send Could not broadcast Tx: ' + e.message);
                }
            }
        },
        async sendMsgDeleteMinerClaims({ rootGetters }, { value, fee = [], memo = '' }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgDeleteMinerClaims(value);
                const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee,
                        gas: "200000" }, memo });
                return result;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new Error('TxClient:MsgDeleteMinerClaims:Init Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new Error('TxClient:MsgDeleteMinerClaims:Send Could not broadcast Tx: ' + e.message);
                }
            }
        },
        async sendMsgAllowSave({ rootGetters }, { value, fee = [], memo = '' }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgAllowSave(value);
                const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee,
                        gas: "200000" }, memo });
                return result;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new Error('TxClient:MsgAllowSave:Init Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new Error('TxClient:MsgAllowSave:Send Could not broadcast Tx: ' + e.message);
                }
            }
        },
        async sendMsgUpdateMinerClaims({ rootGetters }, { value, fee = [], memo = '' }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgUpdateMinerClaims(value);
                const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee,
                        gas: "200000" }, memo });
                return result;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new Error('TxClient:MsgUpdateMinerClaims:Init Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new Error('TxClient:MsgUpdateMinerClaims:Send Could not broadcast Tx: ' + e.message);
                }
            }
        },
        async sendMsgCreateMiners({ rootGetters }, { value, fee = [], memo = '' }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgCreateMiners(value);
                const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee,
                        gas: "200000" }, memo });
                return result;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new Error('TxClient:MsgCreateMiners:Init Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new Error('TxClient:MsgCreateMiners:Send Could not broadcast Tx: ' + e.message);
                }
            }
        },
        async sendMsgClaimSave({ rootGetters }, { value, fee = [], memo = '' }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgClaimSave(value);
                const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee,
                        gas: "200000" }, memo });
                return result;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new Error('TxClient:MsgClaimSave:Init Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new Error('TxClient:MsgClaimSave:Send Could not broadcast Tx: ' + e.message);
                }
            }
        },
        async MsgUpdateMiners({ rootGetters }, { value }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgUpdateMiners(value);
                return msg;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new Error('TxClient:MsgUpdateMiners:Init  Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new Error('TxClient:MsgUpdateMiners:Create  Could not create message: ' + e.message);
                }
            }
        },
        async MsgCreateMinerClaims({ rootGetters }, { value }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgCreateMinerClaims(value);
                return msg;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new Error('TxClient:MsgCreateMinerClaims:Init  Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new Error('TxClient:MsgCreateMinerClaims:Create  Could not create message: ' + e.message);
                }
            }
        },
        async MsgDeleteMiners({ rootGetters }, { value }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgDeleteMiners(value);
                return msg;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new Error('TxClient:MsgDeleteMiners:Init  Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new Error('TxClient:MsgDeleteMiners:Create  Could not create message: ' + e.message);
                }
            }
        },
        async MsgDeleteMinerClaims({ rootGetters }, { value }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgDeleteMinerClaims(value);
                return msg;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new Error('TxClient:MsgDeleteMinerClaims:Init  Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new Error('TxClient:MsgDeleteMinerClaims:Create  Could not create message: ' + e.message);
                }
            }
        },
        async MsgAllowSave({ rootGetters }, { value }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgAllowSave(value);
                return msg;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new Error('TxClient:MsgAllowSave:Init  Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new Error('TxClient:MsgAllowSave:Create  Could not create message: ' + e.message);
                }
            }
        },
        async MsgUpdateMinerClaims({ rootGetters }, { value }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgUpdateMinerClaims(value);
                return msg;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new Error('TxClient:MsgUpdateMinerClaims:Init  Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new Error('TxClient:MsgUpdateMinerClaims:Create  Could not create message: ' + e.message);
                }
            }
        },
        async MsgCreateMiners({ rootGetters }, { value }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgCreateMiners(value);
                return msg;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new Error('TxClient:MsgCreateMiners:Init  Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new Error('TxClient:MsgCreateMiners:Create  Could not create message: ' + e.message);
                }
            }
        },
        async MsgClaimSave({ rootGetters }, { value }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgClaimSave(value);
                return msg;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new Error('TxClient:MsgClaimSave:Init  Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new Error('TxClient:MsgClaimSave:Create  Could not create message: ' + e.message);
                }
            }
        },
    }
};
