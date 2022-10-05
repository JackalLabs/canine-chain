import { txClient, queryClient, MissingWalletError , registry} from './module'

import { Bids } from "./module/types/rns/bids"
import { Forsale } from "./module/types/rns/forsale"
import { Init } from "./module/types/rns/init"
import { Names } from "./module/types/rns/names"
import { Params } from "./module/types/rns/params"
import { QueryGetWhoisRequest } from "./module/types/rns/query"
import { QueryGetWhoisResponse } from "./module/types/rns/query"
import { QueryAllWhoisRequest } from "./module/types/rns/query"
import { QueryAllWhoisResponse } from "./module/types/rns/query"
import { Whois } from "./module/types/rns/whois"


export { Bids, Forsale, Init, Names, Params, QueryGetWhoisRequest, QueryGetWhoisResponse, QueryAllWhoisRequest, QueryAllWhoisResponse, Whois };

async function initTxClient(vuexGetters) {
	return await txClient(vuexGetters['common/wallet/signer'], {
		addr: vuexGetters['common/env/apiTendermint']
	})
}

async function initQueryClient(vuexGetters) {
	return await queryClient({
		addr: vuexGetters['common/env/apiCosmos']
	})
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

function getStructure(template) {
	let structure = { fields: [] }
	for (const [key, value] of Object.entries(template)) {
		let field: any = {}
		field.name = key
		field.type = typeof value
		structure.fields.push(field)
	}
	return structure
}

const getDefaultState = () => {
	return {
				Params: {},
				Names: {},
				NamesAll: {},
				Bids: {},
				BidsAll: {},
				Forsale: {},
				ForsaleAll: {},
				Init: {},
				InitAll: {},
				ListOwnedNames: {},
				
				_Structure: {
						Bids: getStructure(Bids.fromPartial({})),
						Forsale: getStructure(Forsale.fromPartial({})),
						Init: getStructure(Init.fromPartial({})),
						Names: getStructure(Names.fromPartial({})),
						Params: getStructure(Params.fromPartial({})),
						QueryGetWhoisRequest: getStructure(QueryGetWhoisRequest.fromPartial({})),
						QueryGetWhoisResponse: getStructure(QueryGetWhoisResponse.fromPartial({})),
						QueryAllWhoisRequest: getStructure(QueryAllWhoisRequest.fromPartial({})),
						QueryAllWhoisResponse: getStructure(QueryAllWhoisResponse.fromPartial({})),
						Whois: getStructure(Whois.fromPartial({})),
						
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
				getNames: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Names[JSON.stringify(params)] ?? {}
		},
				getNamesAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.NamesAll[JSON.stringify(params)] ?? {}
		},
				getBids: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Bids[JSON.stringify(params)] ?? {}
		},
				getBidsAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.BidsAll[JSON.stringify(params)] ?? {}
		},
				getForsale: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Forsale[JSON.stringify(params)] ?? {}
		},
				getForsaleAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ForsaleAll[JSON.stringify(params)] ?? {}
		},
				getInit: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Init[JSON.stringify(params)] ?? {}
		},
				getInitAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.InitAll[JSON.stringify(params)] ?? {}
		},
				getListOwnedNames: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ListOwnedNames[JSON.stringify(params)] ?? {}
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
			console.log('Vuex module: jackaldao.canine.rns initialized!')
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
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryParams()).data
				
					
				commit('QUERY', { query: 'Params', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryParams', payload: { options: { all }, params: {...key},query }})
				return getters['getParams']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryParams API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryNames({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryNames( key.index)).data
				
					
				commit('QUERY', { query: 'Names', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryNames', payload: { options: { all }, params: {...key},query }})
				return getters['getNames']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryNames API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryNamesAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryNamesAll(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryNamesAll({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'NamesAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryNamesAll', payload: { options: { all }, params: {...key},query }})
				return getters['getNamesAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryNamesAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryBids({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryBids( key.index)).data
				
					
				commit('QUERY', { query: 'Bids', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryBids', payload: { options: { all }, params: {...key},query }})
				return getters['getBids']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryBids API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryBidsAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryBidsAll(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryBidsAll({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'BidsAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryBidsAll', payload: { options: { all }, params: {...key},query }})
				return getters['getBidsAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryBidsAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryForsale({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryForsale( key.name)).data
				
					
				commit('QUERY', { query: 'Forsale', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryForsale', payload: { options: { all }, params: {...key},query }})
				return getters['getForsale']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryForsale API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryForsaleAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryForsaleAll(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryForsaleAll({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'ForsaleAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryForsaleAll', payload: { options: { all }, params: {...key},query }})
				return getters['getForsaleAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryForsaleAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryInit({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryInit( key.address)).data
				
					
				commit('QUERY', { query: 'Init', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryInit', payload: { options: { all }, params: {...key},query }})
				return getters['getInit']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryInit API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryInitAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryInitAll(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryInitAll({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'InitAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryInitAll', payload: { options: { all }, params: {...key},query }})
				return getters['getInitAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryInitAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryListOwnedNames({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryListOwnedNames( key.address, query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryListOwnedNames( key.address, {...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'ListOwnedNames', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryListOwnedNames', payload: { options: { all }, params: {...key},query }})
				return getters['getListOwnedNames']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryListOwnedNames API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgAcceptBid({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgAcceptBid(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgAcceptBid:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgAcceptBid:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgList({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgList(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgList:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgList:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCancelBid({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgCancelBid(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCancelBid:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCancelBid:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgBuy({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgBuy(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgBuy:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgBuy:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgDelist({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgDelist(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDelist:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgDelist:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgDelRecord({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgDelRecord(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDelRecord:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgDelRecord:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgInit({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgInit(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgInit:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgInit:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgRegister({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgRegister(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRegister:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgRegister:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgTransfer({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgTransfer(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgTransfer:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgTransfer:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgAddRecord({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgAddRecord(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgAddRecord:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgAddRecord:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgBid({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgBid(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgBid:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgBid:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgAcceptBid({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgAcceptBid(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgAcceptBid:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgAcceptBid:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgList({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgList(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgList:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgList:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCancelBid({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgCancelBid(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCancelBid:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCancelBid:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgBuy({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgBuy(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgBuy:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgBuy:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgDelist({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgDelist(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDelist:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgDelist:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgDelRecord({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgDelRecord(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDelRecord:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgDelRecord:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgInit({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgInit(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgInit:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgInit:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgRegister({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgRegister(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRegister:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgRegister:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgTransfer({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgTransfer(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgTransfer:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgTransfer:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgAddRecord({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgAddRecord(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgAddRecord:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgAddRecord:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgBid({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgBid(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgBid:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgBid:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}
