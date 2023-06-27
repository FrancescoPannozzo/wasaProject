<script>
export default {
    name: "HomeButton",
    props: ['username']
}
</script>

<template>
    <router-link :to="`/profiles/` + username">{{ username }} PROFILE</router-link>
</template>