<script>
import SearchUserBoxVue from "../components/SearchUserBox.vue";
import Navbar from "../components/Navbar.vue";

export default {
    name: "Searchview",
    data() {
        return {

        }
    },
    components: {
        Navbar,
        SearchUserBoxVue,
    }
}

</script>

<template>
    <Navbar></Navbar>
    <p>SEARCH USER PAGE</p>
    <SearchUserBoxVue></SearchUserBoxVue>
</template>