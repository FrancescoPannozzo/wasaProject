<script>

import HomeButton from '../components/HomeButton.vue';
import SearchButton from '../components/SearchButton.vue';
import LogoutButton from '../components/LogoutButton.vue';
import LoggedUserButton from '../components/LoggedUserButton.vue';

export default {
    name: 'Userprofile',
    props: ['loggeduser'],
    data() {
        return {
        }
    },
    components: {
        SearchButton,
        HomeButton,
        LogoutButton,
        LoggedUserButton
    }
}

</script>

<template>
    <div class="navbar">
        <HomeButton class="home"></HomeButton>
        <SearchButton class="search"></SearchButton>
        <LogoutButton class="logout"></LogoutButton>
    </div>
</template>

<style>
.home,
.search,
.logout,
.personal {
    padding: 2%;
    color: white;
}

.navbar {
    width: 100%;
    background-color: rgb(143, 29, 163);
    overflow: auto;
    text-align: center;
}
</style>