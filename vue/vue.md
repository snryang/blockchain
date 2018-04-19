### dataHistory.js保存vue页面数据状态
``` js
//main.js中增加代码
import * as H from "./common/dataHistory";
Vue.prototype.$H = H;

//.vue组件 created方法加入代码
created() {      
    //恢复当前组件状态
    this.$H.restore(this,()=>{
        this.queryList1(); //如果当前组件没有保存状态时，执行的默认代码
    });
}

//.vue组件中，需要保存当前页面状态的地方调用
this.$H.keep(this);  
```


https://stackblitz.com/github/  
如果你开发前端(Angular,React,ES6,Typescipt)，那么最推荐这种方式！直接把Chrome变成一个在线IDE。帮你把npm包都准备好，直接可以运行。