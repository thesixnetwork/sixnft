import { Client, registry, MissingWalletError } from 'thesixnetwork-sixnft-client-ts'

import { ActionParam } from "thesixnetwork-sixnft-client-ts/thesixnetwork.sixnft.nftoracle/types"
import { ActionRequest } from "thesixnetwork-sixnft-client-ts/thesixnetwork.sixnft.nftoracle/types"
import { ActionSignature } from "thesixnetwork-sixnft-client-ts/thesixnetwork.sixnft.nftoracle/types"
import { TxOriginParam } from "thesixnetwork-sixnft-client-ts/thesixnetwork.sixnft.nftoracle/types"
import { CollectionOwnerRequest } from "thesixnetwork-sixnft-client-ts/thesixnetwork.sixnft.nftoracle/types"
import { CollectionOwnerSignature } from "thesixnetwork-sixnft-client-ts/thesixnetwork.sixnft.nftoracle/types"
import { MintRequest } from "thesixnetwork-sixnft-client-ts/thesixnetwork.sixnft.nftoracle/types"
import { Trait } from "thesixnetwork-sixnft-client-ts/thesixnetwork.sixnft.nftoracle/types"
import { Params } from "thesixnetwork-sixnft-client-ts/thesixnetwork.sixnft.nftoracle/types"
import { NftOriginData } from "thesixnetwork-sixnft-client-ts/thesixnetwork.sixnft.nftoracle/types"
import { TransactionOriginDataInfo } from "thesixnetwork-sixnft-client-ts/thesixnetwork.sixnft.nftoracle/types"
import { DataHash } from "thesixnetwork-sixnft-client-ts/thesixnetwork.sixnft.nftoracle/types"
import { OriginTxInfo } from "thesixnetwork-sixnft-client-ts/thesixnetwork.sixnft.nftoracle/types"


export { ActionParam, ActionRequest, ActionSignature, TxOriginParam, CollectionOwnerRequest, CollectionOwnerSignature, MintRequest, Trait, Params, NftOriginData, TransactionOriginDataInfo, DataHash, OriginTxInfo };

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
				CollectionOwnerRequest: {},
				CollectionOwnerRequestAll: {},
				
				_Structure: {
						ActionParam: getStructure(ActionParam.fromPartial({})),
						ActionRequest: getStructure(ActionRequest.fromPartial({})),
						ActionSignature: getStructure(ActionSignature.fromPartial({})),
						TxOriginParam: getStructure(TxOriginParam.fromPartial({})),
						CollectionOwnerRequest: getStructure(CollectionOwnerRequest.fromPartial({})),
						CollectionOwnerSignature: getStructure(CollectionOwnerSignature.fromPartial({})),
						MintRequest: getStructure(MintRequest.fromPartial({})),
						Trait: getStructure(Trait.fromPartial({})),
						Params: getStructure(Params.fromPartial({})),
						NftOriginData: getStructure(NftOriginData.fromPartial({})),
						TransactionOriginDataInfo: getStructure(TransactionOriginDataInfo.fromPartial({})),
						DataHash: getStructure(DataHash.fromPartial({})),
						OriginTxInfo: getStructure(OriginTxInfo.fromPartial({})),
						
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
				getCollectionOwnerRequest: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.CollectionOwnerRequest[JSON.stringify(params)] ?? {}
		},
				getCollectionOwnerRequestAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.CollectionOwnerRequestAll[JSON.stringify(params)] ?? {}
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
			console.log('Vuex module: thesixnetwork.sixnft.nftoracle initialized!')
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
				let value= (await client.ThesixnetworkSixnftNftoracle.query.queryParams()).data
				
					
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
				let value= (await client.ThesixnetworkSixnftNftoracle.query.queryMintRequest( key.id)).data
				
					
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
				let value= (await client.ThesixnetworkSixnftNftoracle.query.queryMintRequestAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.ThesixnetworkSixnftNftoracle.query.queryMintRequestAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
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
				let value= (await client.ThesixnetworkSixnftNftoracle.query.queryActionRequest( key.id)).data
				
					
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
				let value= (await client.ThesixnetworkSixnftNftoracle.query.queryActionRequestAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.ThesixnetworkSixnftNftoracle.query.queryActionRequestAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'ActionRequestAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryActionRequestAll', payload: { options: { all }, params: {...key},query }})
				return getters['getActionRequestAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryActionRequestAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryCollectionOwnerRequest({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ThesixnetworkSixnftNftoracle.query.queryCollectionOwnerRequest( key.id)).data
				
					
				commit('QUERY', { query: 'CollectionOwnerRequest', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryCollectionOwnerRequest', payload: { options: { all }, params: {...key},query }})
				return getters['getCollectionOwnerRequest']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryCollectionOwnerRequest API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryCollectionOwnerRequestAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ThesixnetworkSixnftNftoracle.query.queryCollectionOwnerRequestAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.ThesixnetworkSixnftNftoracle.query.queryCollectionOwnerRequestAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'CollectionOwnerRequestAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryCollectionOwnerRequestAll', payload: { options: { all }, params: {...key},query }})
				return getters['getCollectionOwnerRequestAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryCollectionOwnerRequestAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgCreateVerifyCollectionOwnerRequest({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.ThesixnetworkSixnftNftoracle.tx.sendMsgCreateVerifyCollectionOwnerRequest({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateVerifyCollectionOwnerRequest:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateVerifyCollectionOwnerRequest:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCreateMintRequest({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.ThesixnetworkSixnftNftoracle.tx.sendMsgCreateMintRequest({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateMintRequest:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateMintRequest:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCreateActionRequest({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.ThesixnetworkSixnftNftoracle.tx.sendMsgCreateActionRequest({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateActionRequest:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateActionRequest:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgSubmitVerifyCollectionOwner({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.ThesixnetworkSixnftNftoracle.tx.sendMsgSubmitVerifyCollectionOwner({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSubmitVerifyCollectionOwner:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgSubmitVerifyCollectionOwner:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgSubmitMintResponse({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.ThesixnetworkSixnftNftoracle.tx.sendMsgSubmitMintResponse({ value, fee: {amount: fee, gas: "200000"}, memo })
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
				const result = await client.ThesixnetworkSixnftNftoracle.tx.sendMsgSubmitActionResponse({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSubmitActionResponse:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgSubmitActionResponse:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgCreateVerifyCollectionOwnerRequest({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ThesixnetworkSixnftNftoracle.tx.msgCreateVerifyCollectionOwnerRequest({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateVerifyCollectionOwnerRequest:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateVerifyCollectionOwnerRequest:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCreateMintRequest({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ThesixnetworkSixnftNftoracle.tx.msgCreateMintRequest({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateMintRequest:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateMintRequest:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCreateActionRequest({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ThesixnetworkSixnftNftoracle.tx.msgCreateActionRequest({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateActionRequest:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateActionRequest:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgSubmitVerifyCollectionOwner({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ThesixnetworkSixnftNftoracle.tx.msgSubmitVerifyCollectionOwner({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSubmitVerifyCollectionOwner:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgSubmitVerifyCollectionOwner:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgSubmitMintResponse({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ThesixnetworkSixnftNftoracle.tx.msgSubmitMintResponse({value})
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
				const msg = await client.ThesixnetworkSixnftNftoracle.tx.msgSubmitActionResponse({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSubmitActionResponse:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgSubmitActionResponse:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}
