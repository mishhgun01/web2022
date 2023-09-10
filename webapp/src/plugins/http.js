import axios from "axios"

export default {
    install(Vue, options) {
        Vue.prototype.$http = axios.create({
            baseUrl: "https://localhost:8081",
            headers: options.headers || null
        })
    }
}