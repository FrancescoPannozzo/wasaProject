<script>
export default {
    name: "LogoutButton",
    methods: {
        logout() {
            localStorage.setItem("token", "");
            window.alert("Logout done!")
        }
    }
}
</script>

<template>
    <router-link to="/" @click="logout">LOGOUT</router-link>
</template>