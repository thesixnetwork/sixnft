// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import SixnftAdmin from './sixnft.admin'
import SixnftEvmsupport from './sixnft.evmsupport'
import SixnftNftmngr from './sixnft.nftmngr'
import SixnftNftoracle from './sixnft.nftoracle'


export default { 
  SixnftAdmin: load(SixnftAdmin, 'sixnft.admin'),
  SixnftEvmsupport: load(SixnftEvmsupport, 'sixnft.evmsupport'),
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