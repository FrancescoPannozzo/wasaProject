<script>
import Navbar from "../components/Navbar.vue";

export default {
    name: "Post",
    data() {
        return {
            photoObj: "",
            resp: [],
            idphoto: this.$route.params.idphoto,
            errormsg: "",
            likethis: false,
            feedbackMsg: "",
            loggedusername: "",
            comment: ""
        }
    },
    components: {
        Navbar
    },
    methods: {
        async loadPost() {
            try {
                let response = await this.$axios.get("/posts/" + this.idphoto);
                this.resp = response.data;
                this.likethis = response.data.likethis
                this.loggedusername = response.data.loggedusername
            } catch (e) {
                this.errormsg = e.toString();
                console.log(this.errormsg)
            }
            try {
                const photoData = await this.$axios.get("/photos/" + this.idphoto, { responseType: "blob" });
                this.photoObj = window.URL.createObjectURL(photoData.data);
            } catch (e) {
                this.errormsg = e.toString();
                console.log(this.errormsg)
            }
        },
        async likePhoto() {
            try {
                let response = await this.$axios.post("/photos/" + this.idphoto + "/likes");
                this.feedbackMsg = response.data.feedback;
                this.loadPost()
                window.alert(this.feedbackMsg)
            } catch (e) {
                this.errormsg = e.toString();
                console.log(this.errormsg)
            }
        },
        async removelikePhoto() {
            try {
                let response = await this.$axios.delete("/photos/" + this.idphoto + "/likes/" + this.loggedusername);
                this.feedbackMsg = response.data.feedback;
                this.loadPost()
                window.alert(this.feedbackMsg)
            } catch (e) {
                this.errormsg = e.toString();
                console.log(this.errormsg)
            }
        },
        async postComment() {
            try {
                let response = await this.$axios.post("/photos/" + this.idphoto + "/comments", { comment: this.comment });
                this.feedbackMsg = response.data.feedback;
                this.loadPost()
                window.alert(this.feedbackMsg)
            } catch (e) {
                this.errormsg = e.toString();
                console.log(this.errormsg)
            }
        },
        async cancelComment(idcomment) {
            try {
                let response = await this.$axios.delete("/photos/" + this.idphoto + "/comments/" + idcomment);
                this.feedbackMsg = response.data.feedback;
                this.loadPost()
                window.alert(this.feedbackMsg)
            } catch (e) {
                this.errormsg = e.toString();
                console.log(this.errormsg)
            }
        }
    },
    mounted() {
        this.loadPost()
    }
}
</script>

<template>
    <Navbar></Navbar>
    <div id="post">
        <img id="postpic" :src="photoObj" />
        <p>Photo by: <router-link :to="'/profiles/' + resp.username"> {{ resp.username }}</router-link></p>
        <p>Posted in data/time:{{ resp.datetime }}</p>
        <p>Likes received n.:{{ resp.nlikes }}</p>
        <div class="likebutton">
            <button v-if="likethis" @click="removelikePhoto()">YOU LIKE THIS</button>
            <button v-else @click="likePhoto">LIKE</button>

            <div>
                <input type="text" v-model="comment" placeholder="comment the photo!" minlength="1" maxlength="100" />
                <button @click="postComment()" :disabled="comment.length >= 1 ? false : true">COMMENT</button>
            </div>



        </div>
        <br />
        <li v-for=" item  in  resp.comments  ">
            <p>{{ item.name }} : {{ item.comment }}<button v-show="resp.loggedusername == item.name"
                    @click="cancelComment(item.commentid)">CANCEL</button>
            </p>
        </li>

    </div>
</template>

<style>
#post {
    color: blueviolet;
    border: 3px solid;
    border-radius: 16px;
    margin: auto;
    text-align: center;

}

#postpic {
    width: 30%;
    height: 30%;
    margin: auto;
    display: block;
}

#likebutton {
    text-align: center;
}
</style>