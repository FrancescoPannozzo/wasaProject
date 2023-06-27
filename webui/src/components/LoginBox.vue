<script>
export default {
    name: "LoginBox",
    data() {
        return {
            errormsg: null,
            username: "",
        }
    },
    methods: {
        async login() {
            try {
                if (this.username.length < 3 || this.username.length > 13) {
                    throw new Error("Warning,  accepted usernames are min 3 up to 13 characters")
                }
                let response = await this.$axios.post("/session", { name: this.username });
                localStorage.setItem("token", response.data.identifier);
                this.$router.push("/my-stream");
            } catch (e) {
                this.errormsg = e.toString();
                window.alert(this.errormsg)
                this.username = "";
            }
        }
    }
}
</script>

<template>
    <div>
        <p>LOGGIN</p>
        <input type="text" v-model="username" /><button @click="login">GO</button>
    </div>
</template>

<style>
div {
    text-align: center;
}
</style>