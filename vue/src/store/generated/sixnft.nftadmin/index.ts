import { Client, registry, MissingWalletError } from 'thesixnetwork-sixnft-client-ts'

import { Authorization } from "thesixnetwork-sixnft-client-ts/sixnft.nftadmin/types"
import { Params } from "thesixnetwork-sixnft-client-ts/sixnft.nftadmin/types"
import { AddressList } from "thesixnetwork-sixnft-client-ts/sixnft.nftadmin/types"
import { Permissions } from "thesixnetwork-sixnft-client-ts/sixnft.nftadmin/types"


export { Authorization, Params, AddressList, Permissions };

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
				Authorization: {},
				
				_Structure: {
						Authorization: getStructure(Authorization.fromPartial({})),
						Params: getStructure(Params.fromPartial({})),
						AddressList: getStructure(AddressList.fromPartial({})),
						Permissions: getStructure(Permissions.fromPartial({})),
						
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
				getAuthorization: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Authorization[JSON.stringify(params)] ?? {}
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
			console.log('Vuex module: sixnft.nftadmin initialized!')
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
				let value= (await client.SixnftNftadmin.query.queryParams()).data
				
					
				commit('QUERY', { query: 'Params', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryParams', payload: { options: { all }, params: {...key},query }})
				return getters['getParams']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryParams API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryAuthorization({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.SixnftNftadmin.query.queryAuthorization()).data
				
					
				commit('QUERY', { query: 'Authorization', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryAuthorization', payload: { options: { all }, params: {...key},query }})
				return getters['getAuthorization']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryAuthorization API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgGrantPermission({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.SixnftNftadmin.tx.sendMsgGrantPermission({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgGrantPermission:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgGrantPermission:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgBurn({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.SixnftNftadmin.tx.sendMsgBurn({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgBurn:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgBurn:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgRevokePermission({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.SixnftNftadmin.tx.sendMsgRevokePermission({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRevokePermission:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgRevokePermission:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgMint({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.SixnftNftadmin.tx.sendMsgMint({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgMint:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgMint:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgGrantPermission({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.SixnftNftadmin.tx.msgGrantPermission({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgGrantPermission:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgGrantPermission:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgBurn({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.SixnftNftadmin.tx.msgBurn({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgBurn:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgBurn:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgRevokePermission({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.SixnftNftadmin.tx.msgRevokePermission({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRevokePermission:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgRevokePermission:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgMint({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.SixnftNftadmin.tx.msgMint({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgMint:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgMint:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}
