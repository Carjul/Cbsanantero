"use strict";(self["webpackChunkfront"]=self["webpackChunkfront"]||[]).push([[376],{3376:function(a,e,r){r.r(e),r.d(e,{default:function(){return va}});var t=r(3396),i=r(9242),s=r(7139);const l=a=>((0,t.dD)("data-v-5b29a974"),a=a(),(0,t.Cn)(),a),n={id:"app"},o=l((()=>(0,t._)("nav",{class:"main-menu"},[(0,t._)("strong",null,"PANEL DE ADMINISTRACION"),(0,t._)("ul")],-1))),c={class:"content"},d={class:"left-content"},u={class:"card"},m={class:"card-body"},p=l((()=>(0,t._)("br",null,null,-1))),A=l((()=>(0,t._)("br",null,null,-1))),_=l((()=>(0,t._)("h5",{class:"card-title"},"Creación de Artesanía",-1))),f=l((()=>(0,t._)("p",{class:"card-text"},"En este espacio creamos nuevas artesanías para COSTA BRISA",-1))),h={ref:"modalCrear",class:"modal fade",id:"modalCrear",tabindex:"",role:"dialog","aria-labelledby":"modalCrearLabel","aria-hidden":"true"},b={class:"modal-dialog modal-dialog-centered",role:"document"},g={class:"modal-content"},v={class:"modal-header"},y=l((()=>(0,t._)("h5",{class:"modal-title",id:"modalCrearLabel"},"Crear Artesanía",-1))),z=l((()=>(0,t._)("span",{"aria-hidden":"true"},"×",-1))),C=[z],w={class:"modal-body"},M={class:"form-group"},k=l((()=>(0,t._)("label",{for:"nombreArtesania"},"Nombre:",-1))),x={class:"form-group"},D=l((()=>(0,t._)("label",{for:"direccionArtesania"},"Dirección:",-1))),I={class:"form-group"},L=l((()=>(0,t._)("label",{for:"imagenArtesania"},"Imagen:",-1))),q={class:"form-group"},E=l((()=>(0,t._)("label",{for:"telefonoArtesania"},"Teléfono:",-1))),U={class:"form-group"},S=l((()=>(0,t._)("label",{for:"descripcionArtesania"},"Descripción:",-1))),V=l((()=>(0,t._)("button",{type:"submit",class:"btn btn-primary"},"Crear Artesanía",-1))),N={class:"right-content"},T=l((()=>(0,t._)("strong",null,"ARTESANÍA ACTUAL",-1))),F={class:"row"},$={class:"card"},Z=["src"],j={class:"card-body"},R={class:"card-title"},O={class:"card-text"},B={class:"card-text"},G={class:"card-text"},H=["onClick"],K=["onClick"],P={ref:"modalActualizar",class:"modal fade",id:"modalActualizar",tabindex:"-1",role:"dialog","aria-labelledby":"modalActualizarLabel","aria-hidden":"true"},Y={class:"modal-dialog modal-dialog-centered",role:"document"},J={class:"modal-content"},Q={class:"modal-header"},W=l((()=>(0,t._)("h5",{class:"modal-title",id:"modalActualizarLabel"},"Actualizar Artesanía",-1))),X=l((()=>(0,t._)("span",{"aria-hidden":"true"},"×",-1))),aa=[X],ea={class:"modal-body"},ra={class:"form-group"},ta=l((()=>(0,t._)("label",{for:"nombreArtesaniaActualizar"},"Nombre:",-1))),ia={class:"form-group"},sa=l((()=>(0,t._)("label",{for:"direccionArtesaniaActualizar"},"Dirección:",-1))),la={class:"form-group"},na=l((()=>(0,t._)("label",{for:"imagenArtesaniaActualizar"},"Imagen:",-1))),oa={class:"form-group"},ca=l((()=>(0,t._)("label",{for:"telefonoArtesaniaActualizar"},"Teléfono:",-1))),da={class:"form-group"},ua=l((()=>(0,t._)("label",{for:"descripcionArtesaniaActualizar"},"Descripción:",-1))),ma=l((()=>(0,t._)("button",{type:"submit",class:"btn btn-primary"},"Guardar Cambios",-1))),pa={class:"modal-footer"},Aa={key:0,class:"text-success"};function _a(a,e,r,l,z,X){return(0,t.wg)(),(0,t.iD)("div",n,[(0,t._)("main",null,[o,(0,t._)("section",c,[(0,t._)("div",d,[(0,t._)("div",u,[(0,t._)("div",m,[p,A,_,f,(0,t._)("button",{class:"btn btn-success mb-4",onClick:e[0]||(e[0]=(...a)=>X.abrirModalCrear&&X.abrirModalCrear(...a))},"Crear")])]),(0,t._)("div",h,[(0,t._)("div",b,[(0,t._)("div",g,[(0,t._)("div",v,[y,(0,t._)("button",{type:"button",class:"close","data-dismiss":"modal","aria-label":"Close",onClick:e[1]||(e[1]=(...a)=>X.cerrarModalCrear&&X.cerrarModalCrear(...a))},C)]),(0,t._)("div",w,[(0,t._)("form",{onSubmit:e[7]||(e[7]=(0,i.iM)(((...a)=>X.crearArtesania&&X.crearArtesania(...a)),["prevent"]))},[(0,t._)("div",M,[k,(0,t.wy)((0,t._)("input",{type:"text","onUpdate:modelValue":e[2]||(e[2]=a=>z.nuevaArtesania.name=a),class:"form-control",id:"nombreArtesania",required:""},null,512),[[i.nr,z.nuevaArtesania.name]])]),(0,t._)("div",x,[D,(0,t.wy)((0,t._)("input",{type:"text","onUpdate:modelValue":e[3]||(e[3]=a=>z.nuevaArtesania.address=a),class:"form-control",id:"direccionArtesania",required:""},null,512),[[i.nr,z.nuevaArtesania.address]])]),(0,t._)("div",I,[L,(0,t._)("input",{name:"image",class:"form-control",accept:"image/jpeg, image/jpg, image/png",onChange:e[4]||(e[4]=(...e)=>a.uploadFile&&a.uploadFile(...e)),id:"imagenArtesania",ref:"file",type:"file",multiple:"",style:{display:"none"}},null,544)]),(0,t._)("div",q,[E,(0,t.wy)((0,t._)("input",{type:"text","onUpdate:modelValue":e[5]||(e[5]=a=>z.nuevaArtesania.phone=a),class:"form-control",id:"telefonoArtesania",required:""},null,512),[[i.nr,z.nuevaArtesania.phone]])]),(0,t._)("div",U,[S,(0,t.wy)((0,t._)("textarea",{"onUpdate:modelValue":e[6]||(e[6]=a=>z.nuevaArtesania.description=a),class:"form-control",id:"descripcionArtesania",required:""},null,512),[[i.nr,z.nuevaArtesania.description]])]),V],32)])])])],512)]),(0,t._)("div",N,[T,(0,t._)("div",F,[((0,t.wg)(!0),(0,t.iD)(t.HY,null,(0,t.Ko)(z.artesanias,(a=>((0,t.wg)(),(0,t.iD)("div",{key:a._id,class:"col-md-12 mb-6"},[(0,t._)("div",$,[(0,t._)("img",{src:a.image,class:"card-img-top",style:{width:"100%",height:"200px"},alt:"Artesanía Image"},null,8,Z),(0,t._)("div",j,[(0,t._)("h5",R,(0,s.zw)(a.name),1),(0,t._)("p",O,"Dirección: "+(0,s.zw)(a.address),1),(0,t._)("p",B,"Descripción: "+(0,s.zw)(a.description),1),(0,t._)("p",G,"Teléfono: "+(0,s.zw)(a.phone),1),(0,t._)("button",{class:"btn btn-danger",onClick:e=>X.eliminarArtesania(a._id)},"Eliminar",8,H),(0,t._)("button",{class:"btn btn-primary",onClick:e=>X.abrirModalActualizar(a)},"Actualizar",8,K)])])])))),128))])]),(0,t._)("div",P,[(0,t._)("div",Y,[(0,t._)("div",J,[(0,t._)("div",Q,[W,(0,t._)("button",{type:"button",class:"close","data-dismiss":"modal","aria-label":"Close",onClick:e[8]||(e[8]=(...a)=>X.cerrarModalActualizar&&X.cerrarModalActualizar(...a))},aa)]),(0,t._)("div",ea,[(0,t._)("form",{onSubmit:e[14]||(e[14]=(0,i.iM)(((...a)=>X.actualizarArtesania&&X.actualizarArtesania(...a)),["prevent"]))},[(0,t._)("div",ra,[ta,(0,t.wy)((0,t._)("input",{type:"text","onUpdate:modelValue":e[9]||(e[9]=a=>z.artesaniaActualizada.name=a),class:"form-control",id:"nombreArtesaniaActualizar",required:""},null,512),[[i.nr,z.artesaniaActualizada.name]])]),(0,t._)("div",ia,[sa,(0,t.wy)((0,t._)("input",{type:"text","onUpdate:modelValue":e[10]||(e[10]=a=>z.artesaniaActualizada.address=a),class:"form-control",id:"direccionArtesaniaActualizar",required:""},null,512),[[i.nr,z.artesaniaActualizada.address]])]),(0,t._)("div",la,[na,(0,t._)("input",{class:"form-control",accept:"image/jpeg, image/jpg, image/png",onChange:e[11]||(e[11]=e=>a.uploadFileActualizar()),id:"imagenArtesaniaActualizar",ref:"fileActualizar",type:"file"},null,544)]),(0,t._)("div",oa,[ca,(0,t.wy)((0,t._)("input",{type:"text","onUpdate:modelValue":e[12]||(e[12]=a=>z.artesaniaActualizada.phone=a),class:"form-control",id:"telefonoArtesaniaActualizar",required:""},null,512),[[i.nr,z.artesaniaActualizada.phone]])]),(0,t._)("div",da,[ua,(0,t.wy)((0,t._)("textarea",{"onUpdate:modelValue":e[13]||(e[13]=a=>z.artesaniaActualizada.description=a),class:"form-control",id:"descripcionArtesaniaActualizar",required:""},null,512),[[i.nr,z.artesaniaActualizada.description]])]),ma],32)]),(0,t._)("div",pa,[a.successMessage?((0,t.wg)(),(0,t.iD)("p",Aa,(0,s.zw)(a.successMessage),1)):(0,t.kq)("",!0)])])])],512)])])])}var fa=r(1076),ha={data(){return{artesanias:[],artesaniaActualizada:{_id:null,name:"",address:"",image:"",phone:"",description:""},file:null,nuevaArtesania:{name:"",address:"",image:"",phone:"",description:"",customer_id:""}}},mounted(){this.fetchArtesanias()},methods:{async fetchArtesanias(){let a=localStorage.getItem("customerId");try{const e=await fa.Z.get("http://localhost:3000/app/Artesania");let r=e.data.filter((e=>e._id===a));this.artesanias=r||[]}catch(e){console.error("Error fetching artesanias:",e)}},async eliminarArtesania(a){try{await fa.Z.delete(`http://localhost:3000/app/Artesania/${a}`),this.fetchArtesanias()}catch(e){console.error("Error deleting artesania:",e)}},abrirModalActualizar(a){this.artesaniaActualizada={...a};const e=this.$refs.modalActualizar;e&&(e.classList.add("show"),e.style.display="block")},cerrarModalActualizar(){const a=this.$refs.modalActualizar;a&&(a.classList.remove("show"),a.style.display="none")},abrirModalCrear(){const a=this.$refs.modalCrear;a&&(a.classList.add("show"),a.style.display="block")},cerrarModalCrear(){const a=this.$refs.modalCrear;a&&(a.classList.remove("show"),a.style.display="none")},async crearArtesania(){try{var a=new FormData;if(this.nuevaArtesania.customer_id=localStorage.getItem("customerId"),""!=this.nuevaArtesania.customer_id){a.append("image",this.file),a.append("name",this.nuevaArtesania.name),a.append("address",this.nuevaArtesania.address),a.append("phone",this.nuevaArtesania.phone),a.append("description",this.nuevaArtesania.description),a.append("customer_id",this.nuevaArtesania.customer_id);const e=await fa.Z.post("http://localhost:3000/app/Artesania",a);console.log(e),this.nuevaArtesania={name:"",address:"",image:"",phone:"",description:"",customer_id:""},this.cerrarModalCrear(),this.fetchArtesanias()}else alert("No tienes permisos para crear una artesanía")}catch(e){console.error("Error creating Artesania:",e)}},async actualizarArtesania(){try{const a=await fa.Z.put(`http://localhost:3000/app/Artesania/${this.artesaniaActualizada._id}`,this.artesaniaActualizada);console.log(a),this.fetchArtesanias(),this.cerrarModalActualizar()}catch(a){console.error("Error updating Artesania:",a)}}},uploadFile(a){this.file=a.target.files[0]},uploadFileActualizar(a){this.fileActualizar=a.target.files[0]}},ba=r(89);const ga=(0,ba.Z)(ha,[["render",_a],["__scopeId","data-v-5b29a974"]]);var va=ga}}]);
//# sourceMappingURL=376.601a4c1f.js.map