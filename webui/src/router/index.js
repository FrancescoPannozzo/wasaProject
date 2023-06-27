import { createRouter, createWebHashHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import MyStream from '../views/MyStream.vue'
import SearchView from '../views/SearchView.vue'
import UserProfileView from '../views/UserProfileView.vue'
import PostView from '../views/PostView.vue'
import LoggedUserProfile from '../views/LoggedUserProfileView.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{ path: '/', component: HomeView },
		{ path: '/my-stream', component: MyStream },
		{ path: '/search', component: SearchView },
		{ path: '/profiles/:username', component: UserProfileView },
		{ path: '/posts/:idphoto', component: PostView },
		{ path: '/personal-area/:username', component: LoggedUserProfile },
		//{ path: '/link2', component: HomeView },
		//{ path: '/some/:id/link', component: HomeView },
	]
})

export default router
