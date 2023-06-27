<script>
export default {
    name: 'Thumbnail',
    props: ['thumbnail', 'loggedusername'],
    emits: ['deletephoto'],
    data() {
        return {
            photoObj: "",
            errormsg: ""
        }
    },
    methods: {
        async loadThumbnail() {
            try {
                const localhost = "/http:/0.0.0.0:3000"
                const targetUrl = this.thumbnail.photourl.substring(localhost.length);

                const resp = await this.$axios.get(targetUrl, { responseType: "blob" });
                this.photoObj = window.URL.createObjectURL(resp.data);
            } catch (e) {
                this.errormsg = e.toString();
                console.log(this.errormsg);
            }
        },
    },
    mounted() {
        this.loadThumbnail()
    }
}
</script>

<template>
    <div class="thumbnail">
        <img :src="photoObj" />
        <p>By user: <router-link :to="'/profiles/' + thumbnail.username"> {{ thumbnail.username }}</router-link> </p>
        <p>datetime: {{ thumbnail.datetime }} </p>
        <p>Likes received n.:{{ thumbnail.nlikes }}</p>
        <p>Comments received n.: {{ thumbnail.ncomments }}</p>
        <router-link :to="`/posts/` + thumbnail.photoid">GO TO THE POST</router-link>
        <br />
        <button v-show="thumbnail.username == loggedusername" @click="$emit('deletephoto')">Delete photo</button>

    </div>
</template>

<style>
.thumbnail {
    width: 30%;
    border: 3px solid;
    border-radius: 16px;
    border-color: blueviolet;
    padding: 5%;
    text-align: center;
    margin: auto;
    background-color: rgba(182, 65, 137, 0.484);
}

img {
    width: 100px;
    height: auto;
}
</style>