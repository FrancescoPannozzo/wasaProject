<script>
import Navbar from "../components/Navbar.vue";
import Thumbnail from "../components/Thumbnail.vue";

export default {
    name: "LoggedUserprofile",
    data() {
        return {
            errormsg: "",
            resp: [],
            loggedUsername: this.$route.params.username,
            uploadButtonDis: true,
            newName: "",
            selected: null,
            feedbackMsg: ""
        }
    },
    components: {
        Navbar,
        Thumbnail
    },
    methods: {
        async loadProfile() {
            try {
                let response = await this.$axios.get("/profiles/" + this.loggedUsername);
                this.resp = response.data;

            } catch (error) {
                if (error.response) {
                    this.errormsg = error.toString();
                    window.alert(error.data)
                    this.$router.push("/my-stream");
                } else {
                    console.log(error.toString())
                }

            }

        },
        async loadPhoto() {
            const selectedFile = this.$refs.userphoto.files[0]
            const reader = new FileReader();

            reader.addEventListener("load", async () => {
                try {
                    const prefix = "data:image/png;base64,"
                    const photodata = reader.result.substring(prefix.length)
                    await this.$axios.post("/photos", photodata, {
                        headers: "image/png"
                    })
                    window.alert("photo uploaded")
                    await this.loadProfile()

                } catch (error) {
                    if (error.response) {
                        this.errormsg = error.toString();
                        window.alert(error.response.data.feedback)
                        this.$router.push("/my-stream");
                    } else {
                        console.log(error.toString())
                    }
                }

            })


            try {
                reader.readAsDataURL(selectedFile)

            } catch (e) {
                console.log(e.toString())
            }

            this.uploadButtonDis = true

        },
        async deletePhoto(idPhoto) {
            try {
                let response = await this.$axios.delete("/photos/" + idPhoto);
                this.feedbackMsg = response.data.feedback;
                window.alert(this.feedbackMsg)
                await this.loadProfile()
            } catch (e) {
                this.errormsg = e.toString();
                console.log(this.errormsg)
            }

        },
        async setUsername() {
            try {
                let response = await this.$axios.put("/profiles/" + this.loggedUsername, { name: this.newName });
                this.feedbackMsg = response.data.feedback

                window.alert(this.feedbackMsg)

                this.$router.push("/my-stream");



            } catch (error) {
                if (error.response) {
                    this.errormsg = error.toString();
                    window.alert(error.data)
                    this.$router.push("/my-stream");
                } else {
                    console.log(error.toString())
                }

            }
        },
        visitProfile(item) {
            this.$router.push("/profiles/" + item);
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
    <Navbar></Navbar>
    <div v-show="isDataLoaded">
        <p>{{ loggedUsername }} PERSONAL AREA</p>

        <label for="followed">Followed:</label>
        <select name="followed" id="followed" v-model="selected">
            <option v-for="item in resp.followed" :value="item" @click="visitProfile(item)"> {{ item }} </option>
        </select>

        <label for="followers"> Followers:</label>
        <select name="followers" id="followers" v-model="selected">
            <option v-for="item in resp.followers" :value="item" @click="visitProfile(item)"> {{ item }} </option>
        </select>

        <div id="upload">
            <p>UPLOAD A NEW PHOTO (.png):</p><input type="file" accept="image/*" class="local" ref="userphoto"
                v-on:change="uploadButtonDis = false" />
            <button @click="loadPhoto" :disabled="uploadButtonDis">UPLOAD PHOTO</button>
        </div>
        <br /><br />
        <input type="text" v-model="newName" placeholder="set a new username" /><button @click="setUsername()">CHANGE
            USERNAME</button>
        <br />
        <p>-----------------</p>
        <br />
        <Thumbnail v-for="item in resp.thumbnails" :thumbnail="item" :loggedusername="resp.loggedUsername"
            :key="item.photoid" @deletephoto="deletePhoto(item.photoid)">
        </Thumbnail>
    </div>
</template>

<style>
#upload {
    width: 30%;
    border: 3px solid;
    border-radius: 16px;
    border-color: rgb(214, 43, 226);
    padding: 5%;
    text-align: center;
    margin: auto;
}
</style>