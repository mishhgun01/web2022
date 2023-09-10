import axios from "axios";
import {url} from "@/main";


export function authUser(login, password) {
    const authToken = btoa(`${login}:${password}`)
    const authOpts = axios.create({
        headers: {
            Authorization : `Basic ${authToken}`
        }
    })
    return authOpts.get(url + "/api/v1/ping")
}
export function regUser(login, password) {
    return axios.post(url + "/api/v1/user", { Login: login, Password: password })
}

export function getUserInfo(login, password) {
    const authToken = btoa(`${login}:${password}`)
    const authOpts = axios.create({
        headers: {
            Authorization : `Basic ${authToken}`
        }
    })
    return authOpts.get(url + "/api/v1/user")
}