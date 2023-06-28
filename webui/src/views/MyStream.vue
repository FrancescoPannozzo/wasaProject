<script>
import Navbar from "../components/Navbar.vue";
import Thumbnail from "../components/Thumbnail.vue";

export default {
    name: "Mystream",
    data() {
        return {
            resp: [],
            loggedUsername: "",
            errormsg: "",
            personalAreaURL: ""
        }
    },
    components: {
        Navbar,
        Thumbnail
    },
    methods: {
        async loadStream() {
            try {
                let response = await this.$axios.get("/my-stream");
                this.loggedUsername = response.data.loggedUsername
                this.personalAreaURL = `/personal-area/` + this.loggedUsername
                this.resp = response.data;

            } catch (error) {
                if (error.response) {
                    this.errormsg = error.toString();
                    console.log(error.response.data);
                } else {
                    console.log(this.errormsg)
                }

            }
        }
    },
    mounted() {
        this.loadStream();
    },
    computed: {
        isDataLoaded() {
            return this.resp.length != 0
        }
    }
}

</script>

<template>
    <div v-show="isDataLoaded">
        <Navbar></Navbar>
        <div class="personal"> <router-link :to="personalAreaURL">{{ loggedUsername.toUpperCase() }}
                PROFILE/PERSONAL
                AREA</router-link>
        </div>

        <p>USER STREAM PAGE, WELCOME {{ loggedUsername.toUpperCase() }} !</p>
        <p>User stream contents:</p>
        <ul v-for="(item, index) in resp.thumbnails" :key="index">
            <Thumbnail :thumbnail="item"></Thumbnail>
        </ul>
    </div>
</template>

<style>
.personal {
    width: 20%;
    border: 3px solid;
    border-radius: 16px;
    border-color: blueviolet;
    padding: 2%;
    text-align: center;
    margin: auto;
}
</style>