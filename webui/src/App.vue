<script setup>
import { RouterLink, RouterView } from 'vue-router'

</script>
<script>
export default {
	name: 'App',

}
</script>

<template>
	<div id="app">
		<div class="header">
			<h1>WASA PROJECT</h1>
		</div>
		<router-view>
		</router-view>
	</div>
</template>

<style>
.header {
	padding: 5px;
	text-align: left;
	background: #000000;
	color: white;
	font-size: 10px;
}

#app {
	background-color: rgb(244, 222, 250);
}
</style>
