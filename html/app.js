Vue.component("greeting",{
	template:`
	<p>
	{{name}}: 大家好
	<button @click="changeName">改名</button>
	</p>`,
	data:function(){
		return{
			name:"鹿晗"
		}
	},
	methods:{
		changeName:function(){
		this.name="Henry";
	}
	}
})


//实例化vue对象
var one = new Vue({
	el: "#vue-app-one",
});


var two = new Vue({
	el: "#vue-app-two",
	
});

/*
 * el:element 
 *data:用于数据存储
 *methods:用于存储各种方法
 *data-binding:给属性绑定对应的值
 */