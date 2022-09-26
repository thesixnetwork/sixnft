// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import SixnftEvmsupport from './sixnft.evmsupport'
import SixnftNftadmin from './sixnft.nftadmin'
import SixnftNftmngr from './sixnft.nftmngr'
import SixnftNftoracle from './sixnft.nftoracle'


export default { 
  SixnftEvmsupport: load(SixnftEvmsupport, 'sixnft.evmsupport'),
  SixnftNftadmin: load(SixnftNftadmin, 'sixnft.nftadmin'),
  SixnftNftmngr: load(SixnftNftmngr, 'sixnft.nftmngr'),
  SixnftNftoracle: load(SixnftNftoracle, 'sixnft.nftoracle'),
  
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