import{b as k,e as C,f as x,r as c,o as T,c as S,h as D,a as t,w as n,E as u,k as p,t as z}from"./index-546da2fe.js";import{_ as w}from"./_plugin-vue_export-helper-c27b6911.js";function U(s){return k({url:"/api/DbInsert",method:"post",data:s})}function B(s){return k({url:"/api/DbDelete",method:"post",data:s})}function E(s){return k({url:"/api/DbSearch",method:"post",data:s})}const I={name:"DbTest",setup(s,{emit:_}){const r=C({page:1,limit:10}),e=C([]),g=C(0),d=async(a=1)=>{r.value.page=a;const{data:o,code:i,msg:l}=await E(r.value);i!==0?u({type:"error",message:l}):(o.list==null&&(o.list=[]),e.value=o.list,g.value=o.count)},h=a=>{r.value.limit=a,d(1)},v=async()=>{const{code:a,msg:o}=await B();a!==0?u({type:"error",message:o}):(u({type:"success",message:o}),d())},y=()=>{console.log("list",e),e.value.unshift({key:"",value:""})},m=()=>{_("close",!1)};return x(()=>{d()}),{save:async a=>{const{code:o,msg:i}=await U({key:a.key,value:a.value});o!==0?u({type:"error",message:i}):u({type:"success",message:i})},input:r,list:e,total:g,handleSizeChange:h,clean:v,close:m,searchTest:d,addHeader:y}}},N={class:"app-container"},A={class:"search-container"},H={class:"pagination-container"};function M(s,_,r,e,g,d){const h=c("el-tag"),v=c("el-form-item"),y=c("el-form"),m=c("el-input"),f=c("el-table-column"),a=c("el-button"),o=c("el-table"),i=c("el-pagination");return T(),S("div",N,[D("div",A,[t(y,null,{default:n(()=>[t(v,null,{default:n(()=>[t(h,null,{default:n(()=>[p(" 数据库操作 demo ")]),_:1})]),_:1})]),_:1})]),t(o,{data:e.list,style:{width:"100%"},border:""},{default:n(()=>[t(f,{prop:"key",label:"键"},{default:n(({row:l,$index:V})=>[t(m,{modelValue:l.key,"onUpdate:modelValue":b=>l.key=b,placeholder:"键"},null,8,["modelValue","onUpdate:modelValue"])]),_:1}),t(f,{prop:"value",label:"值"},{default:n(({row:l,$index:V})=>[t(m,{modelValue:l.value,"onUpdate:modelValue":b=>l.value=b,placeholder:"值"},null,8,["modelValue","onUpdate:modelValue"])]),_:1}),t(f,{label:"操作",width:"300"},{header:n(()=>[t(a,{onClick:_[0]||(_[0]=l=>e.searchTest(1))},{default:n(()=>[p(z(s.$t("刷新")),1)]),_:1}),t(a,{type:"primary",onClick:e.addHeader},{default:n(()=>[p(" 新增 ")]),_:1},8,["onClick"]),t(a,{type:"danger",onClick:e.clean},{default:n(()=>[p(z(s.$t("清空")),1)]),_:1},8,["onClick"])]),default:n(({row:l})=>[t(a,{onClick:V=>e.save(l),type:"info"},{default:n(()=>[p("保存")]),_:2},1032,["onClick"])]),_:1})]),_:1},8,["data"]),D("div",H,[t(i,{"current-page":e.input.page,"page-sizes":[10,20,30,50,100,150,200],"page-size":e.input.limit,layout:"total, sizes, prev, pager, next",total:e.total,onSizeChange:e.handleSizeChange,onCurrentChange:e.searchTest},null,8,["current-page","page-size","total","onSizeChange","onCurrentChange"])])])}const j=w(I,[["render",M],["__scopeId","data-v-c5afde8b"]]);export{j as default};