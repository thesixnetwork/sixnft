import { Client, registry, MissingWalletError } from 'sixnft-client-ts'

import { ActionParam } from "sixnft-client-ts/sixnft.nftoracle/types"
import { ActionRequest } from "sixnft-client-ts/sixnft.nftoracle/types"
import { ActionSignature } from "sixnft-client-ts/sixnft.nftoracle/types"
import { MintRequest } from "sixnft-client-ts/sixnft.nftoracle/types"
import { Trait } from "sixnft-client-ts/sixnft.nftoracle/types"
import { Params } from "sixnft-client-ts/sixnft.nftoracle/types"
import { NftOriginData } from "sixnft-client-ts/sixnft.nftoracle/types"
import { DataHash } from "sixnft-client-ts/sixnft.nftoracle/types"


export { ActionParam, ActionRequest, ActionSignature, MintRequest, Trait, Params, NftOriginData, DataHash };

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
				MintRequest: {},
				MintRequestAll: {},
				ActionRequest: {},
				ActionRequestAll: {},
				
				_Structure: {
						ActionParam: getStructure(ActionParam.fromPartial({})),
						ActionRequest: getStructure(ActionRequest.fromPartial({})),
						ActionSignature: getStructure(ActionSignature.fromPartial({})),
						MintRequest: getStructure(MintRequest.fromPartial({})),
						Trait: getStructure(Trait.fromPartial({})),
						Params: getStructure(Params.fromPartial({})),
						NftOriginData: getStructure(NftOriginData.fromPartial({})),
						DataHash: getStructure(DataHash.fromPartial({})),
						
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
				getMintRequest: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.MintRequest[JSON.stringify(params)] ?? {}
		},
				getMintRequestAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.MintRequestAll[JSON.stringify(params)] ?? {}
		},
				getActionRequest: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ActionRequest[JSON.stringify(params)] ?? {}
		},
				getActionRequestAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ActionRequestAll[JSON.stringify(params)] ?? {}
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
			console.log('Vuex module: sixnft.nftoracle initialized!')
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
				let value= (await client.SixnftNftoracle.query.queryParams()).data
				
					
				commit('QUERY', { query: 'Params', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryParams', payload: { options: { all }, params: {...key},query }})
				return getters['getParams']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryParams API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryMintRequest({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.SixnftNftoracle.query.queryMintRequest( key.id)).data
				
					
				commit('QUERY', { query: 'MintRequest', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryMintRequest', payload: { options: { all }, params: {...key},query }})
				return getters['getMintRequest']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryMintRequest API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryMintRequestAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.SixnftNftoracle.query.queryMintRequestAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.SixnftNftoracle.query.queryMintRequestAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'MintRequestAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryMintRequestAll', payload: { options: { all }, params: {...key},query }})
				return getters['getMintRequestAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryMintRequestAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryActionRequest({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.SixnftNftoracle.query.queryActionRequest( key.id)).data
				
					
				commit('QUERY', { query: 'ActionRequest', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryActionRequest', payload: { options: { all }, params: {...key},query }})
				return getters['getActionRequest']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryActionRequest API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryActionRequestAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.SixnftNftoracle.query.queryActionRequestAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.SixnftNftoracle.query.queryActionRequestAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'ActionRequestAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryActionRequestAll', payload: { options: { all }, params: {...key},query }})
				return getters['getActionRequestAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryActionRequestAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgCreateActionRequest({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.SixnftNftoracle.tx.sendMsgCreateActionRequest({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateActionRequest:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateActionRequest:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgSubmitMintResponse({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.SixnftNftoracle.tx.sendMsgSubmitMintResponse({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSubmitMintResponse:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgSubmitMintResponse:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgSubmitActionResponse({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.SixnftNftoracle.tx.sendMsgSubmitActionResponse({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSubmitActionResponse:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgSubmitActionResponse:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCreateMintRequest({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.SixnftNftoracle.tx.sendMsgCreateMintRequest({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateMintRequest:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateMintRequest:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgCreateActionRequest({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.SixnftNftoracle.tx.msgCreateActionRequest({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateActionRequest:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateActionRequest:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgSubmitMintResponse({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.SixnftNftoracle.tx.msgSubmitMintResponse({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSubmitMintResponse:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgSubmitMintResponse:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgSubmitActionResponse({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.SixnftNftoracle.tx.msgSubmitActionResponse({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSubmitActionResponse:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgSubmitActionResponse:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCreateMintRequest({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.SixnftNftoracle.tx.msgCreateMintRequest({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateMintRequest:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateMintRequest:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}
