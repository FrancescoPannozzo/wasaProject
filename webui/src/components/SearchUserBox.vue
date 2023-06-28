<script>
import { watch } from 'vue';

export default {
    name: "SearchUserBox",
    data() {
        return {
            resp: [],
            username: ""
        }
    },
    computed: {
        usernameChanged() {
            console.log(this.username)
            return this.username
        }
    },
    methods: {
        async search() {
            try {
                let response = await this.$axios.get("/profiles", { params: { username: this.username } });
                this.resp = response.data;
            } catch (e) {
                this.errormsg = e.toString();
                window.alert(e.response.data.feedback)
            }
        }
    }
}
</script>

<template>
    <div id="search">
        SEARCH USERS:<input type="text" v-model="username" minlength="3" maxlength="13" required /><button
            @click="search()">GO</button>

        <ul id="cards" v-for="(item, index) in resp" :key="index">

            <router-link :to='"/profiles/" + item'> {{ item }}

            </router-link>
        </ul>
    </div>
</template>

<style>
#search {
    border: 3px solid;
    color: black;
    border-radius: 16px;
    margin: auto;
    padding: 5%;
}

#cards {
    align-content: center;
}

input:valid {
    background-color: palegreen;
}
</style>