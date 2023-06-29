<script>
import Navbar from "../components/Navbar.vue";
import Thumbnail from "../components/Thumbnail.vue";

export default {
    name: "Userprofile",
    data() {
        return {
            errormsg: "",
            resp: [],
            visitedUsername: this.$route.params.username,
            // Check if the visited user is banned from the logged user
            isBanned: false,
            isFollowing: false,
            selected: ""
        }
    },
    components: {
        Navbar,
        Thumbnail
    },
    methods: {
        async loadProfile() {
            try {
                this.resp = []
                let response = await this.$axios.get("/profiles/" + this.visitedUsername);
                this.resp = response.data;
                this.visitedUsername = this.resp.visitedUsername;
                this.isBanned = this.resp.visitedUserBanned;
                this.isFollowing = response.data.followers.includes(response.data.loggedUsername)
            } catch (error) {
                if (error.response) {
                    this.errormsg = error.toString();
                    window.alert(error.response.data.feedback)
                    this.$router.push("/my-stream");
                } else {
                    console.log(error.toString())
                }
            }
        },
        async banUser() {
            try {
                let response = await this.$axios.post("/bans", { name: this.visitedUsername });
                this.feedbackMsg = response.data.feedback;
                this.isBanned = response.data.isBanned
                this.loadProfile()
                window.alert(this.feedbackMsg)
            } catch (error) {
                if (error.response) {
                    this.errormsg = error.toString();
                    console.log(error.response);
                }
            }
        },
        async unBanUser() {
            try {
                console.log(this.resp.username)
                let response = await this.$axios.delete("/bans/" + this.visitedUsername);
                this.feedbackMsg = response.data.feedback;
                this.isBanned = response.data.isBanned
                this.loadProfile()
                window.alert(this.feedbackMsg)
            } catch (e) {
                this.errormsg = e.toString();
                console.log(this.errormsg)
            }
        },
        async follow() {
            try {
                let response = await this.$axios.post("/follows", { name: this.visitedUsername });
                this.feedbackMsg = response.data.feedback;
                this.loadProfile()
                window.alert(this.feedbackMsg)
            } catch (error) {
                if (error.response) {
                    this.errormsg = error.toString();
                    console.log(error.response);
                }

            }
        },
        async unfollow() {
            try {
                let response = await this.$axios.delete("/follows/" + this.visitedUsername);
                this.feedbackMsg = response.data.feedback;
                this.loadProfile()
                window.alert(this.feedbackMsg)
            } catch (error) {
                if (error.response) {
                    this.errormsg = error.toString();
                    console.log(error.response);
                }

            }
        },
        visitProfile(item) {
            this.$router.push("/profiles/" + item);
            this.visitedUsername = item
            this.loadProfile()
        }
    },
    mounted() {
        this.loadProfile()
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
        <p>{{ visitedUsername.toUpperCase() }} PROFILE</p>

        <div class="content">
            <button v-if="isFollowing" @click="unfollow">UNFOLLOW</button>
            <button v-else @click="follow">FOLLOW</button>
            <button v-if="isBanned" @click="unBanUser">UNBAN USER</button>
            <button v-else @click="banUser">BAN USER</button>

            <label for="followed"> Followed:</label>
            <select name="followed" id="followed" v-model="selected">
                <option v-for="(item, index) in resp.followed" :key="index" :value="item" @click="visitProfile(item)"> {{
                    item }} </option>
            </select>

            <label for="followers"> Followers:</label>
            <select name="followers" id="followers" v-model="selected">
                <option v-for="(item, index) in resp.followers" :key="index" :value="item" @click="visitProfile(item)"> {{
                    item }} </option>
            </select>


            <ul v-for="(item, index) in resp.thumbnails" :key="index">
                <Thumbnail :thumbnail="item"></Thumbnail>
            </ul>
        </div>

    </div>
</template>
