import * as R from "ramda"

let histories = {};
function restore(vue,cb){
    let hash = location.hash;    
    let history =  histories[hash] || null;
    if(history!=null){
        R.forEachObjIndexed((value,key)=>{            
            vue[key] = value;
        },history);
    }else{
        if(cb) cb()
    }
}
function keep(vue){
    let hash = location.hash;
    histories[hash] = R.clone(vue.$data);
}
function clean(){
    let hash = location.hash;
    delete histories[hash];
}
function cleanAll(){
    histories = {};
}
export {restore,keep,clean,cleanAll}