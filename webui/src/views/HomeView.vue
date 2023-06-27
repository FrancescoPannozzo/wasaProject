<script>
import LoginBox from "../components/LoginBox.vue";

export default {
	// configuration object
	name: "HomeView",
	data() {
		return {
		}
	},
	components: {
		LoginBox
	}
}
</script>

<template>
	<LoginBox></LoginBox>
</template>

<style>
.test-area {
	border-style: solid;
	color: blueviolet;
}
</style>
