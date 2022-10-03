import { Client, registry, MissingWalletError } from 'jackal-dao-canine-client-ts'

import { ActiveDeals } from "jackal-dao-canine-client-ts/jackaldao.canine.storage/types"
import { ClientUsage } from "jackal-dao-canine-client-ts/jackaldao.canine.storage/types"
import { Contracts } from "jackal-dao-canine-client-ts/jackaldao.canine.storage/types"
import { Miners } from "jackal-dao-canine-client-ts/jackaldao.canine.storage/types"
import { Params } from "jackal-dao-canine-client-ts/jackaldao.canine.storage/types"
import { PayBlocks } from "jackal-dao-canine-client-ts/jackaldao.canine.storage/types"
import { Proofs } from "jackal-dao-canine-client-ts/jackaldao.canine.storage/types"
import { Strays } from "jackal-dao-canine-client-ts/jackaldao.canine.storage/types"
import { MsgCreateContractsResponse } from "jackal-dao-canine-client-ts/jackaldao.canine.storage/types"
import { MsgUpdateContractsResponse } from "jackal-dao-canine-client-ts/jackaldao.canine.storage/types"
import { MsgDeleteContractsResponse } from "jackal-dao-canine-client-ts/jackaldao.canine.storage/types"
import { MsgCreateProofsResponse } from "jackal-dao-canine-client-ts/jackaldao.canine.storage/types"
import { MsgUpdateProofsResponse } from "jackal-dao-canine-client-ts/jackaldao.canine.storage/types"
import { MsgDeleteProofsResponse } from "jackal-dao-canine-client-ts/jackaldao.canine.storage/types"
import { MsgItemResponse } from "jackal-dao-canine-client-ts/jackaldao.canine.storage/types"
import { MsgCreateActiveDealsResponse } from "jackal-dao-canine-client-ts/jackaldao.canine.storage/types"
import { MsgUpdateActiveDealsResponse } from "jackal-dao-canine-client-ts/jackaldao.canine.storage/types"
import { MsgDeleteActiveDealsResponse } from "jackal-dao-canine-client-ts/jackaldao.canine.storage/types"
import { MsgCreateMinersResponse } from "jackal-dao-canine-client-ts/jackaldao.canine.storage/types"
import { MsgUpdateMinersResponse } from "jackal-dao-canine-client-ts/jackaldao.canine.storage/types"
import { MsgDeleteMinersResponse } from "jackal-dao-canine-client-ts/jackaldao.canine.storage/types"


export { ActiveDeals, ClientUsage, Contracts, Miners, Params, PayBlocks, Proofs, Strays, MsgCreateContractsResponse, MsgUpdateContractsResponse, MsgDeleteContractsResponse, MsgCreateProofsResponse, MsgUpdateProofsResponse, MsgDeleteProofsResponse, MsgItemResponse, MsgCreateActiveDealsResponse, MsgUpdateActiveDealsResponse, MsgDeleteActiveDealsResponse, MsgCreateMinersResponse, MsgUpdateMinersResponse, MsgDeleteMinersResponse };

function initClient(vuexGetters) {
	return new Client(vuexGetters['common/env/getEnv'], vuexGetters['common/wallet/signer'])
}

function mergeResults(value, next_values) {
	for (let prop of Object.keys(next_values)) {
		if (Array.isArray(next_values[prop])) {
			value[prop]=[...value[prop], ...next_values[prop]]
		}else{
			value[prop]=next_values[prop]
		}
	}
	return value
}

type Field = {
	name: string;
	type: unknown;
}
function getStructure(template) {
	let structure: {fields: Field[]} = { fields: [] }
	for (const [key, value] of Object.entries(template)) {
		let field = { name: key, type: typeof value }
		structure.fields.push(field)
	}
	return structure
}
const getDefaultState = () => {
	return {
				Params: {},
				Contracts: {},
				ContractsAll: {},
				Proofs: {},
				ProofsAll: {},
				ActiveDeals: {},
				ActiveDealsAll: {},
				Miners: {},
				MinersAll: {},
				Freespace: {},
				FindFile: {},
				PayBlocks: {},
				PayBlocksAll: {},
				ClientUsage: {},
				ClientUsageAll: {},
				Strays: {},
				StraysAll: {},
				GetClientFreeSpace: {},
				
				_Structure: {
						ActiveDeals: getStructure(ActiveDeals.fromPartial({})),
						ClientUsage: getStructure(ClientUsage.fromPartial({})),
						Contracts: getStructure(Contracts.fromPartial({})),
						Miners: getStructure(Miners.fromPartial({})),
						Params: getStructure(Params.fromPartial({})),
						PayBlocks: getStructure(PayBlocks.fromPartial({})),
						Proofs: getStructure(Proofs.fromPartial({})),
						Strays: getStructure(Strays.fromPartial({})),
						MsgCreateContractsResponse: getStructure(MsgCreateContractsResponse.fromPartial({})),
						MsgUpdateContractsResponse: getStructure(MsgUpdateContractsResponse.fromPartial({})),
						MsgDeleteContractsResponse: getStructure(MsgDeleteContractsResponse.fromPartial({})),
						MsgCreateProofsResponse: getStructure(MsgCreateProofsResponse.fromPartial({})),
						MsgUpdateProofsResponse: getStructure(MsgUpdateProofsResponse.fromPartial({})),
						MsgDeleteProofsResponse: getStructure(MsgDeleteProofsResponse.fromPartial({})),
						MsgItemResponse: getStructure(MsgItemResponse.fromPartial({})),
						MsgCreateActiveDealsResponse: getStructure(MsgCreateActiveDealsResponse.fromPartial({})),
						MsgUpdateActiveDealsResponse: getStructure(MsgUpdateActiveDealsResponse.fromPartial({})),
						MsgDeleteActiveDealsResponse: getStructure(MsgDeleteActiveDealsResponse.fromPartial({})),
						MsgCreateMinersResponse: getStructure(MsgCreateMinersResponse.fromPartial({})),
						MsgUpdateMinersResponse: getStructure(MsgUpdateMinersResponse.fromPartial({})),
						MsgDeleteMinersResponse: getStructure(MsgDeleteMinersResponse.fromPartial({})),
						
		},
		_Registry: registry,
		_Subscriptions: new Set(),
	}
}

// initial state
const state = getDefaultState()

export default {
	namespaced: true,
	state,
	mutations: {
		RESET_STATE(state) {
			Object.assign(state, getDefaultState())
		},
		QUERY(state, { query, key, value }) {
			state[query][JSON.stringify(key)] = value
		},
		SUBSCRIBE(state, subscription) {
			state._Subscriptions.add(JSON.stringify(subscription))
		},
		UNSUBSCRIBE(state, subscription) {
			state._Subscriptions.delete(JSON.stringify(subscription))
		}
	},
	getters: {
				getParams: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Params[JSON.stringify(params)] ?? {}
		},
				getContracts: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Contracts[JSON.stringify(params)] ?? {}
		},
				getContractsAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ContractsAll[JSON.stringify(params)] ?? {}
		},
				getProofs: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Proofs[JSON.stringify(params)] ?? {}
		},
				getProofsAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ProofsAll[JSON.stringify(params)] ?? {}
		},
				getActiveDeals: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ActiveDeals[JSON.stringify(params)] ?? {}
		},
				getActiveDealsAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ActiveDealsAll[JSON.stringify(params)] ?? {}
		},
				getMiners: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Miners[JSON.stringify(params)] ?? {}
		},
				getMinersAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.MinersAll[JSON.stringify(params)] ?? {}
		},
				getFreespace: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Freespace[JSON.stringify(params)] ?? {}
		},
				getFindFile: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.FindFile[JSON.stringify(params)] ?? {}
		},
				getPayBlocks: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.PayBlocks[JSON.stringify(params)] ?? {}
		},
				getPayBlocksAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.PayBlocksAll[JSON.stringify(params)] ?? {}
		},
				getClientUsage: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ClientUsage[JSON.stringify(params)] ?? {}
		},
				getClientUsageAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ClientUsageAll[JSON.stringify(params)] ?? {}
		},
				getStrays: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Strays[JSON.stringify(params)] ?? {}
		},
				getStraysAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.StraysAll[JSON.stringify(params)] ?? {}
		},
				getGetClientFreeSpace: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.GetClientFreeSpace[JSON.stringify(params)] ?? {}
		},
				
		getTypeStructure: (state) => (type) => {
			return state._Structure[type].fields
		},
		getRegistry: (state) => {
			return state._Registry
		}
	},
	actions: {
		init({ dispatch, rootGetters }) {
			console.log('Vuex module: jackaldao.canine.storage initialized!')
			if (rootGetters['common/env/client']) {
				rootGetters['common/env/client'].on('newblock', () => {
					dispatch('StoreUpdate')
				})
			}
		},
		resetState({ commit }) {
			commit('RESET_STATE')
		},
		unsubscribe({ commit }, subscription) {
			commit('UNSUBSCRIBE', subscription)
		},
		async StoreUpdate({ state, dispatch }) {
			state._Subscriptions.forEach(async (subscription) => {
				try {
					const sub=JSON.parse(subscription)
					await dispatch(sub.action, sub.payload)
				}catch(e) {
					throw new Error('Subscriptions: ' + e.message)
				}
			})
		},
		
		
		
		 		
		
		
		async QueryParams({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JackaldaoCanineStorage.query.queryParams()).data
				
					
				commit('QUERY', { query: 'Params', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryParams', payload: { options: { all }, params: {...key},query }})
				return getters['getParams']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryParams API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryContracts({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JackaldaoCanineStorage.query.queryContracts( key.cid)).data
				
					
				commit('QUERY', { query: 'Contracts', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryContracts', payload: { options: { all }, params: {...key},query }})
				return getters['getContracts']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryContracts API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryContractsAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JackaldaoCanineStorage.query.queryContractsAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.JackaldaoCanineStorage.query.queryContractsAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'ContractsAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryContractsAll', payload: { options: { all }, params: {...key},query }})
				return getters['getContractsAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryContractsAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryProofs({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JackaldaoCanineStorage.query.queryProofs( key.cid)).data
				
					
				commit('QUERY', { query: 'Proofs', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryProofs', payload: { options: { all }, params: {...key},query }})
				return getters['getProofs']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryProofs API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryProofsAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JackaldaoCanineStorage.query.queryProofsAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.JackaldaoCanineStorage.query.queryProofsAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'ProofsAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryProofsAll', payload: { options: { all }, params: {...key},query }})
				return getters['getProofsAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryProofsAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryActiveDeals({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JackaldaoCanineStorage.query.queryActiveDeals( key.cid)).data
				
					
				commit('QUERY', { query: 'ActiveDeals', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryActiveDeals', payload: { options: { all }, params: {...key},query }})
				return getters['getActiveDeals']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryActiveDeals API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryActiveDealsAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JackaldaoCanineStorage.query.queryActiveDealsAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.JackaldaoCanineStorage.query.queryActiveDealsAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'ActiveDealsAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryActiveDealsAll', payload: { options: { all }, params: {...key},query }})
				return getters['getActiveDealsAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryActiveDealsAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryMiners({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JackaldaoCanineStorage.query.queryMiners( key.address)).data
				
					
				commit('QUERY', { query: 'Miners', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryMiners', payload: { options: { all }, params: {...key},query }})
				return getters['getMiners']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryMiners API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryMinersAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JackaldaoCanineStorage.query.queryMinersAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.JackaldaoCanineStorage.query.queryMinersAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'MinersAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryMinersAll', payload: { options: { all }, params: {...key},query }})
				return getters['getMinersAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryMinersAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryFreespace({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JackaldaoCanineStorage.query.queryFreespace( key.address)).data
				
					
				commit('QUERY', { query: 'Freespace', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryFreespace', payload: { options: { all }, params: {...key},query }})
				return getters['getFreespace']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryFreespace API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryFindFile({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JackaldaoCanineStorage.query.queryFindFile( key.fid)).data
				
					
				commit('QUERY', { query: 'FindFile', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryFindFile', payload: { options: { all }, params: {...key},query }})
				return getters['getFindFile']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryFindFile API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryPayBlocks({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JackaldaoCanineStorage.query.queryPayBlocks( key.blockid)).data
				
					
				commit('QUERY', { query: 'PayBlocks', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryPayBlocks', payload: { options: { all }, params: {...key},query }})
				return getters['getPayBlocks']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryPayBlocks API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryPayBlocksAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JackaldaoCanineStorage.query.queryPayBlocksAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.JackaldaoCanineStorage.query.queryPayBlocksAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'PayBlocksAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryPayBlocksAll', payload: { options: { all }, params: {...key},query }})
				return getters['getPayBlocksAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryPayBlocksAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryClientUsage({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JackaldaoCanineStorage.query.queryClientUsage( key.address)).data
				
					
				commit('QUERY', { query: 'ClientUsage', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryClientUsage', payload: { options: { all }, params: {...key},query }})
				return getters['getClientUsage']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryClientUsage API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryClientUsageAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JackaldaoCanineStorage.query.queryClientUsageAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.JackaldaoCanineStorage.query.queryClientUsageAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'ClientUsageAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryClientUsageAll', payload: { options: { all }, params: {...key},query }})
				return getters['getClientUsageAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryClientUsageAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryStrays({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JackaldaoCanineStorage.query.queryStrays( key.cid)).data
				
					
				commit('QUERY', { query: 'Strays', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryStrays', payload: { options: { all }, params: {...key},query }})
				return getters['getStrays']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryStrays API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryStraysAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JackaldaoCanineStorage.query.queryStraysAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.JackaldaoCanineStorage.query.queryStraysAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'StraysAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryStraysAll', payload: { options: { all }, params: {...key},query }})
				return getters['getStraysAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryStraysAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryGetClientFreeSpace({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JackaldaoCanineStorage.query.queryGetClientFreeSpace( key.address)).data
				
					
				commit('QUERY', { query: 'GetClientFreeSpace', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryGetClientFreeSpace', payload: { options: { all }, params: {...key},query }})
				return getters['getGetClientFreeSpace']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryGetClientFreeSpace API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgBuyStorage({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.JackaldaoCanineStorage.tx.sendMsgBuyStorage({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgBuyStorage:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgBuyStorage:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgUpdateContracts({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.JackaldaoCanineStorage.tx.sendMsgUpdateContracts({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateContracts:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgUpdateContracts:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgUpdateMiners({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.JackaldaoCanineStorage.tx.sendMsgUpdateMiners({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateMiners:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgUpdateMiners:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgInitMiner({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.JackaldaoCanineStorage.tx.sendMsgInitMiner({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgInitMiner:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgInitMiner:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgDeleteContracts({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.JackaldaoCanineStorage.tx.sendMsgDeleteContracts({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeleteContracts:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgDeleteContracts:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCreateActiveDeals({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.JackaldaoCanineStorage.tx.sendMsgCreateActiveDeals({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateActiveDeals:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateActiveDeals:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCreateMiners({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.JackaldaoCanineStorage.tx.sendMsgCreateMiners({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateMiners:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateMiners:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgPostproof({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.JackaldaoCanineStorage.tx.sendMsgPostproof({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgPostproof:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgPostproof:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCancelContract({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.JackaldaoCanineStorage.tx.sendMsgCancelContract({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCancelContract:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCancelContract:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgDeleteMiners({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.JackaldaoCanineStorage.tx.sendMsgDeleteMiners({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeleteMiners:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgDeleteMiners:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgDeleteActiveDeals({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.JackaldaoCanineStorage.tx.sendMsgDeleteActiveDeals({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeleteActiveDeals:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgDeleteActiveDeals:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgSetMinerTotalspace({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.JackaldaoCanineStorage.tx.sendMsgSetMinerTotalspace({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSetMinerTotalspace:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgSetMinerTotalspace:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgSignContract({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.JackaldaoCanineStorage.tx.sendMsgSignContract({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSignContract:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgSignContract:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgDeleteProofs({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.JackaldaoCanineStorage.tx.sendMsgDeleteProofs({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeleteProofs:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgDeleteProofs:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgUpdateProofs({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.JackaldaoCanineStorage.tx.sendMsgUpdateProofs({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateProofs:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgUpdateProofs:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgItem({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.JackaldaoCanineStorage.tx.sendMsgItem({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgItem:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgItem:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgPostContract({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.JackaldaoCanineStorage.tx.sendMsgPostContract({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgPostContract:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgPostContract:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgUpdateActiveDeals({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.JackaldaoCanineStorage.tx.sendMsgUpdateActiveDeals({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateActiveDeals:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgUpdateActiveDeals:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgSetMinerIp({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.JackaldaoCanineStorage.tx.sendMsgSetMinerIp({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSetMinerIp:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgSetMinerIp:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCreateProofs({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.JackaldaoCanineStorage.tx.sendMsgCreateProofs({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateProofs:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateProofs:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCreateContracts({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.JackaldaoCanineStorage.tx.sendMsgCreateContracts({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateContracts:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateContracts:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgBuyStorage({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.JackaldaoCanineStorage.tx.msgBuyStorage({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgBuyStorage:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgBuyStorage:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgUpdateContracts({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.JackaldaoCanineStorage.tx.msgUpdateContracts({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateContracts:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgUpdateContracts:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgUpdateMiners({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.JackaldaoCanineStorage.tx.msgUpdateMiners({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateMiners:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgUpdateMiners:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgInitMiner({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.JackaldaoCanineStorage.tx.msgInitMiner({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgInitMiner:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgInitMiner:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgDeleteContracts({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.JackaldaoCanineStorage.tx.msgDeleteContracts({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeleteContracts:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgDeleteContracts:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCreateActiveDeals({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.JackaldaoCanineStorage.tx.msgCreateActiveDeals({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateActiveDeals:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateActiveDeals:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCreateMiners({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.JackaldaoCanineStorage.tx.msgCreateMiners({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateMiners:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateMiners:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgPostproof({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.JackaldaoCanineStorage.tx.msgPostproof({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgPostproof:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgPostproof:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCancelContract({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.JackaldaoCanineStorage.tx.msgCancelContract({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCancelContract:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCancelContract:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgDeleteMiners({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.JackaldaoCanineStorage.tx.msgDeleteMiners({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeleteMiners:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgDeleteMiners:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgDeleteActiveDeals({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.JackaldaoCanineStorage.tx.msgDeleteActiveDeals({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeleteActiveDeals:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgDeleteActiveDeals:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgSetMinerTotalspace({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.JackaldaoCanineStorage.tx.msgSetMinerTotalspace({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSetMinerTotalspace:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgSetMinerTotalspace:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgSignContract({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.JackaldaoCanineStorage.tx.msgSignContract({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSignContract:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgSignContract:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgDeleteProofs({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.JackaldaoCanineStorage.tx.msgDeleteProofs({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeleteProofs:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgDeleteProofs:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgUpdateProofs({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.JackaldaoCanineStorage.tx.msgUpdateProofs({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateProofs:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgUpdateProofs:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgItem({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.JackaldaoCanineStorage.tx.msgItem({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgItem:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgItem:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgPostContract({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.JackaldaoCanineStorage.tx.msgPostContract({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgPostContract:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgPostContract:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgUpdateActiveDeals({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.JackaldaoCanineStorage.tx.msgUpdateActiveDeals({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateActiveDeals:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgUpdateActiveDeals:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgSetMinerIp({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.JackaldaoCanineStorage.tx.msgSetMinerIp({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSetMinerIp:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgSetMinerIp:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCreateProofs({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.JackaldaoCanineStorage.tx.msgCreateProofs({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateProofs:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateProofs:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCreateContracts({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.JackaldaoCanineStorage.tx.msgCreateContracts({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateContracts:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateContracts:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}
