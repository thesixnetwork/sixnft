// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import ThesixnetworkSixnftEvmsupport from './thesixnetwork.sixnft.evmsupport'
import ThesixnetworkSixnftNftadmin from './thesixnetwork.sixnft.nftadmin'
import ThesixnetworkSixnftNftmngr from './thesixnetwork.sixnft.nftmngr'
import ThesixnetworkSixnftNftoracle from './thesixnetwork.sixnft.nftoracle'


export default { 
  ThesixnetworkSixnftEvmsupport: load(ThesixnetworkSixnftEvmsupport, 'thesixnetwork.sixnft.evmsupport'),
  ThesixnetworkSixnftNftadmin: load(ThesixnetworkSixnftNftadmin, 'thesixnetwork.sixnft.nftadmin'),
  ThesixnetworkSixnftNftmngr: load(ThesixnetworkSixnftNftmngr, 'thesixnetwork.sixnft.nftmngr'),
  ThesixnetworkSixnftNftoracle: load(ThesixnetworkSixnftNftoracle, 'thesixnetwork.sixnft.nftoracle'),
  
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